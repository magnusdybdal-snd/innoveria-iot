/**
 * This file contains the types for the Gateway entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching gateway information.
 */
export interface GatewayApiResponse {
  id: string;
  company_id: string;
  device_eui: string;
  name: string;
  status: number;
  lastSeenAt: string; // RFC1123 - directly parsable in JS.
}

export interface GatewayListApiResponse {
  totalCount: number; // TODO: check id backens uses batching of max fetched in one fetch - if so update logic to fetch again. (pagination)
  gateways: GatewayApiResponse[];
}
