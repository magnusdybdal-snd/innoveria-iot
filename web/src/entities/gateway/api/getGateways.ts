import type {
  GatewayApiResponse,
  GatewayListApiResponse,
} from "@entities/gateway/model/gatewaySchema";
import { apiRequest, serviceClient } from "@shared/api";

/**
 * Fetches all gateways from the collection-service via the API gateway.
 * @returns Array of GatewayApiResponse objects, or an empty array if the request fails
 */
export const getGateways = async (): Promise<GatewayApiResponse[]> => {
  try {
    const data = await apiRequest<GatewayListApiResponse>(
      serviceClient,
      "/v1/device/gateways",
      "GET",
    );

    return data.gateways ?? [];
  } catch (error) {
    console.error("Failed to fetch gateways:", error);
    return [];
  }
};
