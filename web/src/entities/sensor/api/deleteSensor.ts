import { serviceClient } from "@/shared/api";
import { API_ROUTES } from "@/shared/api/routes";

/**
 * Deletes a sensor from the collection-service via the API gateway.
 * @param sensorId - The ID of the sensor to delete
 * @returns True if the deletion was successful, false otherwise
 */
export const deleteSensor = async (sensorId: string): Promise<boolean> => {
  try {
    const response = await serviceClient.delete(
      `${API_ROUTES.sensors}/${encodeURIComponent(sensorId)}`,
    );
    return response.status === 204;
  } catch (error) {
    console.error(`Failed to delete sensor with ID ${sensorId}:`, error);
    return false;
  }
};
