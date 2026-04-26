import type {
  RawSensorProfileConfigApiResponse,
  SensorProfileConfigApiResponse,
} from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches saved payload schema labels for a Chirpstack profile from device-service.
 * @param chirpstackProfileId - Chirpstack profile ID to fetch labels for
 * @returns Array of labeled schema rows, or an empty array if the request fails
 */
export const getSensorProfileConfig = async (
  chirpstackProfileId: string,
): Promise<SensorProfileConfigApiResponse> => {
  const data = await apiRequest<RawSensorProfileConfigApiResponse>(
    serviceClient,
    `${API_ROUTES.sensorProfileConfig}${encodeURIComponent(chirpstackProfileId)}`,
    "GET",
  );
  return {
    chirpstackProfileId: data.chirpstack_profile_id,
    configurableSchema: data.configurable_schema,
  };
};
