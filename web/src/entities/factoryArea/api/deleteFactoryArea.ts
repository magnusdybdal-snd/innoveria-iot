import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Deletes a factory area via the API gateway.
 * @param areaId - The ID of the factory area to delete
 */
export const deleteFactoryArea = async (areaId: string): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.factoryAreas}/${encodeURIComponent(areaId)}`,
    "DELETE",
  );
};
