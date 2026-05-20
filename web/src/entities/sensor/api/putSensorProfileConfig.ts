import type { SaveSensorProfileConfigListRequest } from "@entities/sensor/model/sensorSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Puts sensor profile metrics data to the API.
 * @param sensorProfileConfigData - A list of metrics containing the measurementType, payloadKey, and unit of the sensor metrics to be saved.
 * @param profileId - Profile ID used to identify the sensor profile
 */
export const putSensorProfileConfig = async (
  sensorProfileConfigData: SaveSensorProfileConfigListRequest,
  profileId: string,
): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.sensorProfileConfig}${encodeURIComponent(profileId)}`,
    "PUT",
    {
      configurable_schema: sensorProfileConfigData.configurableSchema,
    },
  );
};
