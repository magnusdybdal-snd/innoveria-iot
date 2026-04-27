import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

export interface UpdateGatewayRequest {
  description?: string;
  factory_area_id?: string;
  factory_id?: string;
  name?: string;
}

/**
 * Updates a gateway's fields by ID. Only provided fields are updated.
 * @param gatewayId - The ID of the gateway to update
 * @param payload - The fields to update
 */
export const patchGateway = async (
  gatewayId: string,
  payload: UpdateGatewayRequest,
): Promise<void> => {
  await apiRequest<void>(
    serviceClient,
    `${API_ROUTES.gateways}/${encodeURIComponent(gatewayId)}`,
    "PATCH",
    payload,
  );
};
