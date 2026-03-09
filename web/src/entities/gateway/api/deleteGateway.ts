import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Deletes a gateway from the collection-service via the API gateway.
 * @param gatewayId - The ID of the gateway to delete
 * @returns True if the deletion was successful, false otherwise
 */
export const deleteGateway = async (gatewayId: string): Promise<boolean> => {
  try {
    await apiRequest(
      serviceClient,
      `${API_ROUTES.gateways}/${encodeURIComponent(gatewayId)}`,
      "DELETE",
    );
    return true;
  } catch (error) {
    console.error(`Failed to delete gateway with ID ${gatewayId}:`, error);
    return false;
  }
};
