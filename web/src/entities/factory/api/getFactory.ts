import type {
  FactoryApiResponse,
  FactoryListApiResponse,
} from "@entities/factory";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all factories from the collection-service via the API gateway.
 * @returns Array of FactoryApiResponse objects, or an empty array if the request fails
 */
export const getFactories = async (): Promise<FactoryApiResponse[]> => {
  try {
    const data = await apiRequest<FactoryListApiResponse>(
      serviceClient,
      API_ROUTES.factories,
      "GET",
    );

    return data.factories ?? [];
  } catch (error) {
    console.error("Failed to fetch factories:", error);
    return [];
  }
};
