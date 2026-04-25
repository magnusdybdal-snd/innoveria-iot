import type { CreateFactoryAreaRequest } from "@entities/factoryArea/model/factoryAreaSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new factory area to the auth-service via the API gateway.
 * @param data - Factory area payload including factoryId, name, and optional description
 */
export const postFactoryArea = async (
  data: CreateFactoryAreaRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.factoryAreas, "POST", {
    factory_id: data.factoryId,
    name: data.name,
    description: data.description ?? null,
  });
};
