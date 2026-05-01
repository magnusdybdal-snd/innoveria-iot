import type { UpdateGatewayRequest } from "@entities/gateway/model/gatewaySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Sends a PATCH request to update a gateway's editable fields.
 * All fields are optional — only provided fields will be updated.
 * @param gatewayId - ID of the gateway to update
 * @param gatewayData - Fields to update
 */
export const patchGateway = async (
  gatewayId: string,
  gatewayData: UpdateGatewayRequest,
): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.gateways}/${gatewayId}`,
    "PATCH",
    {
      name: gatewayData.name,
      description: gatewayData.description,
      factory_id: gatewayData.factoryId,
      factory_area_id: gatewayData.factoryAreaId,
    },
  );
};
