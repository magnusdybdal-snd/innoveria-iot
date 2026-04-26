import type {
  RawSensorProfileConfigApiResponse,
  SensorProfileConfigApiResponse,
} from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches one sensor profile config based on a given profile from the collection-service.
 * @param profileId - Profile ID used to identify the sensor profile
 * @returns The given SensorProfileConfigApiResponse for the profile
 */
export const getSensorProfileConfig = async (
  profileId: string,
): Promise<SensorProfileConfigApiResponse> => {
  const data = await apiRequest<RawSensorProfileConfigApiResponse>(
    serviceClient,
    `${API_ROUTES.sensorProfileConfig}${encodeURIComponent(profileId)}`,
    "GET",
  );
  return {
    chirpstackProfileId: data.chirpstack_profile_id,
    configurableSchema: data.configurable_schema,
  };
};
