import type { CreateSensorRequest } from "@entities/sensor";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new sensor to the API.
 * @param sensorData - An object containing the companyId, deviceEui, sensorProfileId, and name of the sensor to be created.
 */
export const postSensor = async (
  sensorData: CreateSensorRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.sensors, "POST", {
    company_id: sensorData.companyId,
    factory_id: sensorData.factoryId,
    factory_area_id: sensorData.factoryAreaId,
    device_eui: sensorData.deviceEui,
    device_profile_id: sensorData.sensorProfileId,
    production_resource: sensorData.productionResource,
    app_key: sensorData.appKey,
    name: sensorData.name,
    description: "hardcoded description", //TODO: remove hardcoded when description is added to form"
  });
};
