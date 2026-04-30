import type { UpdateSensorRequest } from "@entities/sensor";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

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
