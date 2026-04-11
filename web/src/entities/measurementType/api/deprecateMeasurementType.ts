import { serviceClient } from "@/shared/api";
import { API_ROUTES } from "@/shared/api/routes";

/**
 * Deprecates a measure type from the device-service via the API gateway.
 * @param slug - The slug of the measure type to delete
 * @returns True if the deletion was successful, false otherwise
 */
export const deprecateMeasurementType = async (
  slug: string,
): Promise<boolean> => {
  try {
    const response = await serviceClient.patch(
      `${API_ROUTES.measurementTypes}/${encodeURIComponent(slug)}/deprecate`,
    );
    return response.status === 204;
  } catch (error) {
    console.error(`Failed to delete measure type with slug ${slug}:`, error);
    return false;
  }
};
