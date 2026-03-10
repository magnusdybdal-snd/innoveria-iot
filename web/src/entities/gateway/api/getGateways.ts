import type {
  GatewayApiResponse,
  GatewayListApiResponse,
} from "@entities/gateway/model/gatewaySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all gateways from the collection-service via the API gateway.
 * @returns Array of GatewayApiResponse objects, or an empty array if the request fails
 */
export const getGateways = async (): Promise<GatewayApiResponse[]> => {
  console.log("api-route", process.env.API_PROXY_TARGET);
  try {
    const data = await apiRequest<GatewayListApiResponse>(
      serviceClient,
      API_ROUTES.gateways,
      "GET",
    );

    return data.gateways ?? [];
  } catch (error) {
    console.error("Failed to fetch gateways:", error);
    return [];
  }
};
