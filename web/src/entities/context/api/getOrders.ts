import type { OrderSummary } from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/** Raw slim order shape returned by GET /orders. */
type RawOrderSummary = {
  id: number;
  name: string;
};

/**
 * Fetches the slim ERP order list from the context service.
 * Each entry contains only the order ID and display name — use `getOrderContext`
 * to fetch full detail for a specific order.
 * @returns Array of order summaries for the current company
 */
export const getOrders = async (): Promise<OrderSummary[]> => {
  const data = await apiRequest<RawOrderSummary[]>(
    serviceClient,
    API_ROUTES.contextOrders,
    "GET",
  );

  return data.map((o) => ({ id: o.id, name: o.name }));
};
