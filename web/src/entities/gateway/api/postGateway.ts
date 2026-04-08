import type { CreateGatewayRequest } from "@entities/gateway/model/gatewaySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new gateway to the API. This function sends a POST request to the /api/v1/device/gateways endpoint with the provided gateway data.
 * @param gatewayData - An object containing the companyId, deviceEui, and name of the gateway to be created.
 */
export const postGateway = async (
  gatewayData: CreateGatewayRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.gateways, "POST", {
    company_id: gatewayData.companyId,
    gateway_eui: gatewayData.gatewayEui,
    name: gatewayData.name,
    factory_id: gatewayData.factoryId,
    factory_area_id: gatewayData.factoryAreaId,
  });
};
