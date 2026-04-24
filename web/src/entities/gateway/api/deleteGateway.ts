import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Deletes a gateway by ID.
 * @param gatewayId - The ID of the gateway to delete
 * @returns A promise that resolves when the deletion succeeds, or rejects on failure
 */
export const deleteGateway = (gatewayId: string): Promise<void> =>
  apiRequest(
    serviceClient,
    `${API_ROUTES.gateways}/${encodeURIComponent(gatewayId)}`,
    "DELETE",
  );
