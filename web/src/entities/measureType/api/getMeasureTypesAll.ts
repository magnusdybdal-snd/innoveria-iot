import type { MeasureTypeApiResponse } from "@entities/measureType/model/measureTypeSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawMeasureType = {
  default_unit: string;
  deprecated: boolean;
  description: string;
  display_name: string;
  slug: string;
};

type RawMeasureTypeListApiResponse = {
  total_count: number;
  measureTypes: RawMeasureType[];
};

/**
 * Fetches all measure types from the collection-service via the API gateway.
 * Maps snake_case API response keys to camelCase.
 * @returns Array of MeasureTypeApiResponse objects, or an empty array if the request fails
 */
export const getMeasureTypesAll = async (): Promise<
  MeasureTypeApiResponse[]
> => {
  try {
    const data = await apiRequest<RawMeasureTypeListApiResponse>(
      serviceClient,
      API_ROUTES.measureTypesAll,
      "GET",
    );

    return (data.measureTypes ?? []).map((s) => ({
      defaultUnit: s.default_unit,
      deprecated: s.deprecated,
      description: s.description,
      displayName: s.display_name,
      slug: s.slug,
    }));
  } catch (error) {
    console.error("Failed to fetch measurement types:", error);
    return [];
  }
};
