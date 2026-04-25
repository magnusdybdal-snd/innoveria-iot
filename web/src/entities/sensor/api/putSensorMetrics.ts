import type { SaveSensorMetricListRequest } from "@entities/sensor/model/sensorSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Puts new payload schema data to the API.
 * @param sensorMetricsDataList - A list of metrics containing the measurementType, payloadKey, and unit of the sensor metrics to be saved.
 * @param deviceEui -
 */
export const putSensorMetrics = async (
  sensorMetricsDataList: SaveSensorMetricListRequest,
  deviceEui: string,
): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.sensors}/${encodeURIComponent(deviceEui)}/metrics`,
    "PUT",
    {
      metrics: sensorMetricsDataList.metrics.map((payloadSchemaData) => ({
        measurement_type: payloadSchemaData.measurementType,
        payload_key: payloadSchemaData.payloadKey,
        unit: payloadSchemaData.unit,
      })),
    },
  );
};
