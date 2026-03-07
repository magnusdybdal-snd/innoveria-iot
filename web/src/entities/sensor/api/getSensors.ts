import type {
  SensorApiResponse,
  SensorListApiResponse,
} from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";

/**
 * Fetches all sensors from the collection-service via the API gateway.
 * @returns Array of SensorApiResponse objects, or an empty array if the request fails
 */
export const getSensors = async (): Promise<SensorApiResponse[]> => {
  try {
    const data = await apiRequest<SensorListApiResponse>(
      serviceClient,
      "/v1/device/sensors",
      "GET",
    );

    return data.sensors ?? [];
  } catch (error) {
    console.error("Failed to fetch sensors:", error);
    return [];
  }
};
