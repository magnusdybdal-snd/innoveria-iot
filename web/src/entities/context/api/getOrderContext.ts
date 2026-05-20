import type {
  Measurement,
  OperationContext,
  OrderContext,
  OrderOperation,
  OrderReport,
  ProductionResource,
  SensorContext,
  SensorMetric,
} from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawProductionResource = {
  id: number;
  number: string;
  description: string | null;
  type: string;
};

type RawOrderReport = {
  id: number;
  quantity: number;
  rest_quantity: number;
  type: string;
  reporting_timestamp: string;
  actual_reported_date: string | null;
};

type RawOrderOperation = {
  id: number;
  production_resource: RawProductionResource;
  planned_start_date: string;
  planned_finish_date: string;
  actual_start_date: string | null;
  actual_finish_date: string | null;
  status: string;
  production_resource_status: string;
  reports: RawOrderReport[];
};

type RawOrder = {
  id: number;
  order_number: string;
  part_description: string;
  planned_start_date: string;
  planned_finish_date: string;
  actual_start_date: string | null;
  actual_finish_date: string | null;
  status: string;
  priority: number;
  operations: RawOrderOperation[];
};

type RawSensorMetric = {
  payload_key: string;
  measurement_type: string;
  unit: string | null;
};

type RawMeasurement = {
  device_eui: string;
  timestamp: string;
  payload: Record<string, unknown>;
};

type RawSensorContext = {
  id: string;
  name: string;
  device_eui: string;
  metrics: RawSensorMetric[];
  measurements: RawMeasurement[];
  total_power_wh: number | null;
  voltage: number | null;
};

type RawOperationContext = {
  operation: RawOrderOperation;
  sensors: RawSensorContext[];
  degraded: boolean;
};

type RawOrderContext = {
  order: RawOrder;
  operations: RawOperationContext[];
};

const toProductionResource = (
  r: RawProductionResource,
): ProductionResource => ({
  id: r.id,
  number: r.number,
  description: r.description,
  type: r.type,
});

const toOrderReport = (r: RawOrderReport): OrderReport => ({
  id: r.id,
  quantity: r.quantity,
  restQuantity: r.rest_quantity,
  type: r.type,
  reportingTimestamp: r.reporting_timestamp,
  actualReportedDate: r.actual_reported_date,
});

const toOrderOperation = (op: RawOrderOperation): OrderOperation => ({
  id: op.id,
  productionResource: toProductionResource(op.production_resource),
  plannedStartDate: op.planned_start_date,
  plannedFinishDate: op.planned_finish_date,
  actualStartDate: op.actual_start_date,
  actualFinishDate: op.actual_finish_date,
  status: op.status,
  productionResourceStatus: op.production_resource_status,
  reports: (op.reports ?? []).map(toOrderReport),
});

const toSensorMetric = (m: RawSensorMetric): SensorMetric => ({
  payloadKey: m.payload_key,
  measurementType: m.measurement_type,
  unit: m.unit,
});

const toMeasurement = (m: RawMeasurement): Measurement => ({
  deviceEui: m.device_eui,
  timestamp: m.timestamp,
  payload: m.payload,
});

const toSensorContext = (s: RawSensorContext): SensorContext => ({
  id: s.id,
  name: s.name,
  deviceEui: s.device_eui,
  metrics: s.metrics.map(toSensorMetric),
  measurements: s.measurements.map(toMeasurement),
  totalPowerWh: s.total_power_wh,
  voltage: s.voltage,
});

const toOperationContext = (op: RawOperationContext): OperationContext => ({
  operation: toOrderOperation(op.operation),
  sensors: op.sensors.map(toSensorContext),
  degraded: op.degraded,
});

/**
 * Fetches a single order enriched with sensor context for each operation.
 * @param orderId - Numeric ID of the order to fetch context for
 * @returns The order context with per-operation sensor data
 */
export const getOrderContext = async (
  orderId: number,
): Promise<OrderContext> => {
  const url = `${API_ROUTES.contextOrders}/${orderId}/context`;
  const data = await apiRequest<RawOrderContext>(serviceClient, url, "GET");

  return {
    order: {
      id: data.order.id,
      orderNumber: data.order.order_number,
      partDescription: data.order.part_description,
      plannedStartDate: data.order.planned_start_date,
      plannedFinishDate: data.order.planned_finish_date,
      actualStartDate: data.order.actual_start_date,
      actualFinishDate: data.order.actual_finish_date,
      status: data.order.status,
      priority: data.order.priority,
      operations: data.order.operations.map(toOrderOperation),
    },
    operations: data.operations.map(toOperationContext),
  };
};
