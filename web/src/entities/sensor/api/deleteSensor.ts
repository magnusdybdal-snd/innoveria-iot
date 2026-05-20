import { apiRequest, serviceClient } from "@/shared/api";
import { API_ROUTES } from "@/shared/api/routes";

/**
 * Deletes a sensor by ID.
 * @param sensorId - The ID of the sensor to delete
 * @returns A promise that resolves when the deletion succeeds, or rejects on failure
 */
export const deleteSensor = (sensorId: string): Promise<void> =>
  apiRequest(
    serviceClient,
    `${API_ROUTES.sensors}/${encodeURIComponent(sensorId)}`,
    "DELETE",
  );
