import type {
  SensorApiResponse,
  SensorListApiResponse,
} from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all sensors from the collection-service via the API gateway.
 * @returns Array of SensorApiResponse objects, or an empty array if the request fails
 */
export const getSensors = async (): Promise<SensorApiResponse[]> => {
  try {
    const data = await apiRequest<SensorListApiResponse>(
      serviceClient,
      API_ROUTES.sensors,
      "GET",
    );

    return data.sensors ?? [];
  } catch (error) {
    console.error("Failed to fetch sensors:", error);
    return [];
  }
};
