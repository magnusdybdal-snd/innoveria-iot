import type {
  Order,
  OrderReporting,
} from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/** Raw order reporting shape as returned by the API (snake_case). */
type RawOrderReporting = {
  reporting_id: string;
  timestamp: string;
  quantity: number;
  status: string;
};

/** Raw order shape as returned by the API (snake_case). */
type RawOrder = {
  order_id: string;
  product_name: string;
  status: string;
  start_time: string;
  end_time: string;
  workcenter_id: string;
  workcenter_name: string;
  reportings: RawOrderReporting[];
};

/**
 * Maps a raw API order reporting to the domain `OrderReporting` type.
 * @param r - Raw reporting object from the API response
 * @returns Camel-cased `OrderReporting`
 */
const toOrderReporting = (r: RawOrderReporting): OrderReporting => ({
  reportingId: r.reporting_id,
  timestamp: r.timestamp,
  quantity: r.quantity,
  status: r.status,
});

/**
 * Maps a raw API order to the domain `Order` type.
 * @param o - Raw order object from the API response
 * @returns Camel-cased `Order`
 */
const toOrder = (o: RawOrder): Order => ({
  orderId: o.order_id,
  productName: o.product_name,
  status: o.status,
  startTime: o.start_time,
  endTime: o.end_time,
  workcenterId: o.workcenter_id,
  workcenterName: o.workcenter_name,
  reportings: o.reportings.map(toOrderReporting),
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
