import type { UpdateSensorRequest } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Sends a PATCH request to update a sensor's editable fields.
 * All fields are optional — only provided fields will be updated.
 * @param sensorId - ID of the sensor to update
 * @param data - Fields to update
 */
export const patchSensor = async (
  sensorId: string,
  data: UpdateSensorRequest,
): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.sensors}/${sensorId}`,
    "PATCH",
    {
      ...(data.name !== undefined && { name: data.name }),
      ...(data.description !== undefined && { description: data.description }),
      ...(data.factoryId !== undefined && { factory_id: data.factoryId }),
      ...(data.factoryAreaId !== undefined && {
        factory_area_id: data.factoryAreaId,
      }),
      ...(data.sensorProfileId !== undefined && {
        device_profile_id: data.sensorProfileId,
      }),
      ...(data.electricitySensor !== undefined && {
        electricity_sensor: data.electricitySensor,
      }),
      ...(data.voltage !== undefined && { voltage: data.voltage }),
      ...(data.productionResource !== undefined && {
        production_resource: data.productionResource,
      }),
    },
  );
};
