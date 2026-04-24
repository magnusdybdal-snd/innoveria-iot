import type {
  RawSensorProfileConfigApiResponse,
  SensorProfileConfigApiResponse,
} from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches one EUI based on a given chirpstack profile from the collection-service.
 * @param profileId - Chirpstack profile ID used to identify the EUI
 * @returns The given EUI for the profile, or null if the request fails
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
