import type { CreateSensorRequest } from "@entities/sensor";

import { API_ROUTES } from "@/shared/api/routes";

/**
 * Posts a new sensor to the API. This function sends a POST request to the /api/v1/device/sensors endpoint with the provided sensor data.
 * @param sensorData - An object containing the companyId, deviceEui, deviceProfileId, and name of the sensor to be created.
 */
export const postSensor = async (
  sensorData: CreateSensorRequest,
): Promise<void> => {
  const response = await fetch(API_ROUTES.sensors, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    // The API expects snake_case keys.
    body: JSON.stringify({
      company_id: sensorData.companyId,
      device_eui: sensorData.deviceEui,
      device_profile_id: sensorData.sensorProfileId,
      name: sensorData.name,
    }),
  });

  if (!response.ok) {
    throw new Error(`Error posting sensor: ${response.statusText}`);
  }
};
