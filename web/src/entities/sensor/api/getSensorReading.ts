import type { SensorReadingApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";

/**
 * Fetches the latest sensor reading for a given device from the collection-service.
 * @param deviceEUI - LoRaWAN Device EUI used to identify the sensor in ChirpStack
 * @returns The latest SensorReadingApiResponse, or null if the request fails
 */
export const fetchSensorReading = async (
  deviceEUI: string,
): Promise<SensorReadingApiResponse | null> => {
  try {
    const data = await apiRequest<SensorReadingApiResponse>(
      serviceClient,
      `v1/collection/latest?device_eui=${encodeURIComponent(deviceEUI)}`,
      "GET",
    );
    return data;
  } catch (error) {
    console.error("Failed to fetch sensor reading:", error);
    return null;
  }
};
