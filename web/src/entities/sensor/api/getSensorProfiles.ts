import type {
  SensorProfileApiResponse,
  SensorProfileListApiResponse,
} from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all sensor profiles from the collection-service via the API gateway.
 * @returns Array of SensorProfileApiResponse objects
 */
export const getSensorProfiles = async (): Promise<
  SensorProfileApiResponse[]
> => {
  const data = await apiRequest<SensorProfileListApiResponse>(
    serviceClient,
    `${API_ROUTES.sensorProfile}?limit=1000`, // TODO: remove hardcoded limit once pagination is implemented
    "GET",
  );

  return data.sensor_profiles ?? [];
};
