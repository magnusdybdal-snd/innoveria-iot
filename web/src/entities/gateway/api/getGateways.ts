import type { GatewayApiResponse } from "@entities/gateway/model/gatewaySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawGatewayListApiResponse = {
  total_count: number;
  gateways: RawGatewayApiResponse[];
};

type RawGatewayApiResponse = {
  id: string;
  company_id: string;
  gateway_eui: string;
  name: string;
  description: string | null;
  factory_id: string;
  factory_area_id: string;
  status: number;
  last_seen_at: string;
};
/**
 * Fetches all gateways from the collection-service via the API gateway.
 * @returns Array of GatewayApiResponse objects
 */
export const getGateways = async (): Promise<GatewayApiResponse[]> => {
  const data = await apiRequest<RawGatewayListApiResponse>(
    serviceClient,
    API_ROUTES.gateways,
    "GET",
  );

  return (data.gateways ?? []).map((a) => ({
    id: a.id,
    companyId: a.company_id,
    gatewayEui: a.gateway_eui,
    name: a.name,
    description: a.description,
    factoryId: a.factory_id,
    factoryAreaId: a.factory_area_id,
    status: a.status,
    lastSeenAt: a.last_seen_at,
  }));
};
