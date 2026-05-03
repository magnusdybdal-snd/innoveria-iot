import type { CreateSensorRequest } from "@entities/sensor";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new sensor to the API.
 * @param sensorData - An object containing the data of the sensor to be created.
 * @param sensorData.deviceEui - The EUI of the sensor device
 * @param sensorData.sensorProfileId - ID of the sensor profile to apply
 * @param sensorData.appKey - Application key for LoRaWAN activation
 * @param sensorData.name - Display name of the sensor
 * @param sensorData.factoryId - ID of the factory the sensor belongs to
 * @param sensorData.factoryAreaId - ID of the factory area the sensor is located in
 */
export const postSensor = async (
  sensorData: CreateSensorRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.sensors, "POST", {
    electricity_sensor: sensorData.electricitySensor,
    factory_id: sensorData.factoryId,
    factory_area_id: sensorData.factoryAreaId,
    device_eui: sensorData.deviceEui,
    device_profile_id: sensorData.sensorProfileId,
    production_resource: sensorData.productionResource,
    app_key: sensorData.appKey,
    name: sensorData.name,
    voltage: sensorData.voltage,
    description: sensorData.description,
  });
};
