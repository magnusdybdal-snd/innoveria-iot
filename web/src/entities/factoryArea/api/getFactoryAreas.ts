import type { FactoryAreaApiResponse } from "@entities/factoryArea";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawFactoryAreaListApiResponse = {
  total_count: number;
  factory_areas: RawFactoryAreaApiResponse[];
};

type RawFactoryAreaApiResponse = {
  id: string;
  factory_id: string;
  name: string;
  description: string | null;
  created_at: string;
  updated_at: string;
};

/**
 * Fetches all factory areas from the auth-service via the API gateway.
 * @returns Array of FactoryAreaApiResponse objects
 */
export const getFactoryAreas = async (): Promise<FactoryAreaApiResponse[]> => {
  const data = await apiRequest<RawFactoryAreaListApiResponse>(
    serviceClient,
    API_ROUTES.factoryAreas,
    "GET",
  );

  return (data.factory_areas ?? []).map((a) => ({
    id: a.id,
    factoryId: a.factory_id,
    name: a.name,
    description: a.description,
    createdAt: new Date(a.created_at),
    updatedAt: new Date(a.updated_at),
  }));
};
