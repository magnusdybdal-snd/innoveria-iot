import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

export interface UpdateSensorRequest {
  description?: string;
  device_profile_id?: string;
  electricity_sensor?: boolean;
  factory_area_id?: string;
  factory_id?: string;
  name?: string;
  production_resource?: string;
  voltage?: number;
}

/**
 * Updates a sensor's fields by ID. Only provided fields are updated.
 * @param sensorId - The ID of the sensor to update
 * @param payload - The fields to update
 */
export const patchSensor = async (
  sensorId: string,
  payload: UpdateSensorRequest,
): Promise<void> => {
  await apiRequest<void>(
    serviceClient,
    `${API_ROUTES.sensors}/${encodeURIComponent(sensorId)}`,
    "PATCH",
    payload,
  );
};
