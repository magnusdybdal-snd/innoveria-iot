import type { CreateGatewayRequest } from "@entities/gateway/model/gatewaySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new gateway to the API. This function sends a POST request to the /api/v1/device/gateways endpoint with the provided gateway data.
 * @param gatewayData - An object containing the data of the gateway to be created.
 * @param gatewayData.gatewayEui - The EUI of the gateway
 * @param gatewayData.name - Display name of the gateway
 * @param gatewayData.factoryId - ID of the factory the gateway belongs to
 * @param gatewayData.factoryAreaId - ID of the factory area the gateway is located in
 */
export const postGateway = async (
  gatewayData: CreateGatewayRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.gateways, "POST", {
    gateway_eui: gatewayData.gatewayEui,
    name: gatewayData.name,
    factory_id: gatewayData.factoryId,
    factory_area_id: gatewayData.factoryAreaId,
  });
};
