import type { SensorMetricApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawSensorMetric = {
  measurement_type: string;
  payload_key: string;
  unit: string;
};

export interface RawSensorMetricListApiResponse {
  total_count: number;
  metrics: RawSensorMetric[];
}

/**
 * Fetches all sensor metrics of a given sensor from the device-service via the API gateway.
 * @param deviceEui - Device EUI used to identify the sensor
 * Maps snake_case API response keys to camelCase.
 * @returns Array of SensorMetricApiResponse objects, or an empty array if the request fails
 */
export const getSensorMetrics = async (
  deviceEui: string,
): Promise<SensorMetricApiResponse[]> => {
  try {
    const data = await apiRequest<RawSensorMetricListApiResponse>(
      serviceClient,
      `${API_ROUTES.sensors}/${encodeURIComponent(deviceEui)}/metrics`,
      "GET",
    );

    return (data.metrics ?? []).map((m) => ({
      measurementType: m.measurement_type,
      payloadKey: m.payload_key,
      unit: m.unit,
    }));
  } catch (error) {
    console.error("Failed to fetch sensor metrics:", error);
    return [];
  }
};
