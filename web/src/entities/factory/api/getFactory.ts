import type { FactoryApiResponse } from "@entities/factory";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawFactoryListApiResponse = {
  total_count: number;
  factories: RawFactoryApiResponse[];
};
type RawFactoryApiResponse = {
  id: string;
  factory_id: string;
  name: string;
  address: string;
  created_at: Date;
  updated_at: Date;
};
/**
 * Fetches all factories from the collection-service via the API gateway.
 * @returns Array of FactoryApiResponse objects, or an empty array if the request fails
 */
export const getFactories = async (): Promise<FactoryApiResponse[]> => {
  try {
    const data = await apiRequest<RawFactoryListApiResponse>(
      serviceClient,
      API_ROUTES.factories,
      "GET",
    );

    return (data.factories ?? []).map((a) => ({
      id: a.id,
      factoryId: a.factory_id,
      name: a.name,
      address: a.address,
      createdAt: a.created_at,
      updatedAt: a.updated_at,
    }));
  } catch (error) {
    console.error("Failed to fetch factories:", error);
    return [];
  }
};
