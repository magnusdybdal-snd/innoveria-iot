import type { MeasurementTypeApiResponse } from "@entities/measurementType/model/measurementTypeSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawMeasurementType = {
  default_unit: string;
  deprecated: boolean;
  description: string;
  display_name: string;
  slug: string;
};

type RawMeasurementTypeListApiResponse = {
  total_count: number;
  measurement_types: RawMeasurementType[];
};

/**
 * Fetches all measure types from the collection-service via the API gateway.
 * Maps snake_case API response keys to camelCase.
 * @returns Array of MeasurementTypeApiResponse objects, or an empty array if the request fails
 */
export const getMeasurementTypesAll = async (): Promise<
  MeasurementTypeApiResponse[]
> => {
  try {
    const data = await apiRequest<RawMeasurementTypeListApiResponse>(
      serviceClient,
      API_ROUTES.measurementTypesAll,
      "GET",
    );

    return (data.measurement_types ?? []).map((s) => ({
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
