import type { SensorApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawSensor = {
  id: string;
  device_eui: string;
  name: string;
  description: string | null;
  factory_id: string;
  factory_area_id: string;
  status: number;
  production_resource: number | null;
  electricity_sensor: boolean;
  voltage: number | null;
  last_seen_at: string;
  device_profile_id: string;
  global_device_profile_id: string;
};

type RawSensorListApiResponse = {
  total_count: number;
  sensors: RawSensor[];
};

/**
 * Fetches all sensors from the collection-service via the API gateway.
 * Maps snake_case API response keys to camelCase.
 * @returns Array of SensorApiResponse objects
 */
export const getSensors = async (): Promise<SensorApiResponse[]> => {
  const data = await apiRequest<RawSensorListApiResponse>(
    serviceClient,
    API_ROUTES.sensors,
    "GET",
  );

  return (data.sensors ?? []).map((s) => ({
    id: s.id,
    deviceEui: s.device_eui,
    name: s.name,
    description: s.description,
    factory: s.factory_id,
    factoryAreaId: s.factory_area_id,
    status: s.status,
    productionResource: s.production_resource,
    electricitySensor: s.electricity_sensor,
    voltage: s.voltage,
    lastReading: s.last_seen_at,
    sensorProfileId: s.global_device_profile_id,
  }));
};
