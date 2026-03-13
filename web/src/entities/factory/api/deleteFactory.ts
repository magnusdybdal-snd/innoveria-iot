import { apiRequest, serviceClient } from "@/shared/api";
import { API_ROUTES } from "@/shared/api/routes";

/**
 * Deletes a factory via the API gateway.
 * @param factoryId - The ID of the factory to delete
 */
export const deleteFactory = async (factoryId: string): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.factories}/${encodeURIComponent(factoryId)}`,
    "DELETE",
  );
};
