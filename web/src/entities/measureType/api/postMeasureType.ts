import type { CreateMeasureTypeRequest } from "@entities/measureType/model/measureTypeSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new measure type to the API.
 * @param measureTypeData - An object containing the defaultUnit, description, displayName, and slug of the measureType to be created.
 */
export const postMeasureType = async (
  measureTypeData: CreateMeasureTypeRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.measureTypes, "POST", {
    default_unit: measureTypeData.defaultUnit,
    description: measureTypeData.description,
    display_name: measureTypeData.displayName,
    slug: measureTypeData.slug,
  });
};
