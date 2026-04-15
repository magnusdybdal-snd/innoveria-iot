import type {
  Order,
  OrderOperation,
  ProductionResource,
} from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/** Raw production resource shape as returned by the API (snake_case). */
type RawProductionResource = {
  id: number;
  number: string;
  description: string | null;
  type: string;
};

/** Raw order operation shape as returned by the API (snake_case). */
type RawOrderOperation = {
  id: number;
  production_resource: RawProductionResource;
  planned_start_date: string;
  planned_finish_date: string;
  actual_start_date: string | null;
  actual_finish_date: string | null;
  status: string;
  production_resource_status: string;
};

/** Raw order shape as returned by the API (snake_case). */
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

/**
 * Maps a raw API production resource to the domain `ProductionResource` type.
 * @param r - Raw production resource object from the API response
 * @returns Camel-cased `ProductionResource`
 */
const toProductionResource = (
  r: RawProductionResource,
): ProductionResource => ({
  id: r.id,
  number: r.number,
  description: r.description,
  type: r.type,
});

/**
 * Maps a raw API order operation to the domain `OrderOperation` type.
 * @param op - Raw operation object from the API response
 * @returns Camel-cased `OrderOperation`
 */
const toOrderOperation = (op: RawOrderOperation): OrderOperation => ({
  id: op.id,
  productionResource: toProductionResource(op.production_resource),
  plannedStartDate: op.planned_start_date,
  plannedFinishDate: op.planned_finish_date,
  actualStartDate: op.actual_start_date,
  actualFinishDate: op.actual_finish_date,
  status: op.status,
  productionResourceStatus: op.production_resource_status,
});

/**
 * Maps a raw API order to the domain `Order` type.
 * @param o - Raw order object from the API response
 * @returns Camel-cased `Order`
 */
const toOrder = (o: RawOrder): Order => ({
  id: o.id,
  orderNumber: o.order_number,
  partDescription: o.part_description,
  plannedStartDate: o.planned_start_date,
  plannedFinishDate: o.planned_finish_date,
  actualStartDate: o.actual_start_date,
  actualFinishDate: o.actual_finish_date,
  status: o.status,
  priority: o.priority,
  operations: o.operations.map(toOrderOperation),
});

/**
 * Fetches all ERP orders enriched with reportings from the context service.
 * @returns Array of orders belonging to the current company
 */
export const getOrders = async (): Promise<Order[]> => {
  const data = await apiRequest<RawOrder[]>(
    serviceClient,
    API_ROUTES.contextOrders,
    "GET",
  );

  return data.map(toOrder);
};
