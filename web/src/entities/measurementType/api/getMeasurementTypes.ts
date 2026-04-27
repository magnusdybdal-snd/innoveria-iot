import type {
  MeasurementTypeApiResponse,
  RawMeasurementTypeListApiResponse,
} from "@entities/measurementType/model/measurementTypeSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all non-deprecated measure types from the collection-service via the API gateway.
 * Maps snake_case API response keys to camelCase.
 * @returns Array of MeasurementTypeApiResponse objects, or an empty array if the request fails
 */
export const getMeasurementTypes = async (): Promise<
  MeasurementTypeApiResponse[]
> => {
  const data = await apiRequest<RawMeasurementTypeListApiResponse>(
    serviceClient,
    API_ROUTES.measurementTypes,
    "GET",
  );

  return (data.measurement_types ?? []).map((s) => ({
    defaultUnit: s.default_unit,
    deprecated: s.deprecated,
    description: s.description,
    displayName: s.display_name,
    slug: s.slug,
  }));
};
