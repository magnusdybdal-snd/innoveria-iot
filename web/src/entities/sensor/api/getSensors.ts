import type { SensorApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawSensor = {
  id: string;
  device_eui: string;
  name: string;
  factory_id: string;
  factory_area_id: string;
  status: number;
  production_resource: string | null;
  last_seen_at: string;
  device_profile_id: string;
  app_key: string;
  voltage: number | null;
  electricity_sensor: boolean;
  updated_at: string;
  description: string;
  company_id: string;
  state: string;
  created_at: string;
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
    factory: s.factory_id,
    factoryArea: s.factory_area_id,
    status: s.status,
    productionResource: s.production_resource ?? "",
    lastReading: s.last_seen_at,
    sensorProfile: s.device_profile_id,
    appKey: s.app_key,
    voltage: s.voltage,
    electricitySensor: s.electricity_sensor,
    updatedAt: s.updated_at,
    description: s.description,
    companyId: s.company_id,
    state: s.state,
    createdAt: s.created_at,
  }));
};
