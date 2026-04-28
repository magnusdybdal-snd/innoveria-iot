import type { CreateMeasurementTypeRequest } from "@entities/measurementType/model/measurementTypeSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new measure type to the API.
 * @param measurementTypeData - An object containing the defaultUnit, description, displayName, and slug of the measurementType to be created.
 */
export const postMeasurementType = async (
  measurementTypeData: CreateMeasurementTypeRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.measurementTypes, "POST", {
    default_unit: measurementTypeData.defaultUnit,
    description: measurementTypeData.description,
    display_name: measurementTypeData.displayName,
    slug: measurementTypeData.slug,
  });
};
