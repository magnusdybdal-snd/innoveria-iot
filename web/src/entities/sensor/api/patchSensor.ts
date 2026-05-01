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
      name: data.name,
      description: data.description,
      factory_id: data.factoryId,
      factory_area_id: data.factoryAreaId,
      device_profile_id: data.sensorProfileId,
      electricity_sensor: data.electricitySensor,
      voltage: data.voltage,
      production_resource: data.productionResource,
    },
  );
};
