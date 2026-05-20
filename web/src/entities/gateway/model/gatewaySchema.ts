/**
 * This file contains the types for the Gateway entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching gateway information.
 */
export interface GatewayApiResponse {
  id: string;
  companyId: string;
  gatewayEui: string;
  name: string;
  description: string | null;
  factoryId: string;
  factoryAreaId: string;
  status: number;
  lastSeenAt: string; // RFC1123 - directly parsable in JS.
}

export interface GatewayListApiResponse {
  totalCount: number;
  gateways: GatewayApiResponse[];
}

export interface CreateGatewayRequest {
  gatewayEui: string;
  name: string;
  description?: string;
  factoryId: string;
  factoryAreaId: string;
}

export interface UpdateGatewayRequest {
  name?: string;
  description?: string;
  factoryId?: string;
  factoryAreaId?: string;
}
