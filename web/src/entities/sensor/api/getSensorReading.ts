import type { SensorReadingApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches the latest sensor reading for a given device from the collection-service.
 * @param deviceEUI - LoRaWAN Device EUI used to identify the sensor in ChirpStack
 * @returns The latest SensorReadingApiResponse
 */
export const fetchSensorReading = async (
  deviceEUI: string,
): Promise<SensorReadingApiResponse> => {
  return apiRequest<SensorReadingApiResponse>(
    serviceClient,
    `${API_ROUTES.sensorLatest}${encodeURIComponent(deviceEUI)}`,
    "GET",
  );
};
