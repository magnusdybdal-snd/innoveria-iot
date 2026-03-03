import {
  apiRequest,
  serviceClient,
  type GatewayListApiResponse,
} from "@/API/apiClient";
import type { Gateway } from "@/mocks/gateways";

/**
 *
 */
export const fetchGateways = async (): Promise<Gateway[]> => {
  try {
    const data = await apiRequest<GatewayListApiResponse>(
      serviceClient,
      "/api/v1/device/gateways",
      "GET",
    );

    return data.gateways.map((gw) => ({
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
