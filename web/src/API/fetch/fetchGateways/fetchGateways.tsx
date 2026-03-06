import {
  apiRequest,
  serviceClient,
  type GatewayListApiResponse,
} from "@/API/apiClient";
import type { Gateway } from "@/mocks/gateways";

/**
 * Fetches all gateways from the collection-service via the API gateway and maps them to the local Gateway shape.
 * @returns Array of Gateway objects, or an empty array if the request fails
 */
export const fetchGateways = async (): Promise<Gateway[]> => {
  try {
    const data = await apiRequest<GatewayListApiResponse>(
      serviceClient,
      "/v1/device/gateways",
      "GET",
    );

    return (data.gateways ?? []).map((gw) => ({
      id: gw.id,
      company_id: gw.company_id,
      name: gw.name,
      status: gw.status,
      euid: gw.device_eui,
      lastSeen: gw.lastSeenAt,
    }));
  } catch (error) {
    console.error("Failed to fetch gateways:", error);
    return [];
  }
};
