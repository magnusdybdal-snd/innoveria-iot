import type { SensorReadingApiResponse } from "@/entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@/shared/api";

export const fetchSensorReading = async (
  deviceEUI: string,
): Promise<SensorReadingApiResponse | null> => {
  try {
    const data = await apiRequest<SensorReadingApiResponse>(
      serviceClient,
      `/api/v1/collection/latest?device_eui=${encodeURIComponent(deviceEUI)}`,
      "GET",
    );
    return data;
  } catch (error) {
    console.error("Failed to fetch sensor reading:", error);
    return null;
  }
};
