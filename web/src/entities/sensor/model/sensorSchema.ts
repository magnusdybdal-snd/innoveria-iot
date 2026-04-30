/**
 * This file contains the TypeScript interfaces for the API responses related to sensors and their readings.
 * It defines the structure of the data returned by the API when fetching sensor information and their latest readings.
 */
export interface SensorListApiResponse {
  totalCount: number;
  sensors: SensorApiResponse[];
}

export interface SensorApiResponse {
  id: string;
  deviceEui: string;
  name: string;
  factory: string;
  factoryArea: string;
  status: number;
  lastReading: string;
  updatedAt: string;
  sensorProfile: string;
  productionResource: string | null;
  appKey: string;
  voltage: number | null;
  electricitySensor: boolean;
  description: string;
  companyId: string;
  state: string;
  createdAt: string;
}

export interface SensorReadingApiResponse {
  deviceEui: string;
  timestamp: string;
  payload: Record<string, unknown> | null;
  companyId: string;
}

export interface DeviceEUIApiResponse {
  deviceEui: string;
}

export interface CreateSensorRequest {
  electricitySensor: boolean;
  factoryId: string;
  factoryAreaId: string;
  deviceEui: string;
  productionResource: number | null;
  appKey: string;
  sensorProfile: string;
  name: string;
  voltage: number | null;
}

// Sensor profile API responses
export interface SensorProfileListApiResponse {
  total_count: number;
  sensor_profiles: SensorProfileApiResponse[];
}

export interface SensorProfileApiResponse {
  id: string;
  name: string;
  region: string;
  vendorId: string;
  vendor: string;
}

// Sensor profile config
export interface RawSensorProfileConfigApiResponse {
  chirpstack_profile_id: string;
  configurable_schema: boolean;
}

export interface SensorProfileConfigApiResponse {
  chirpstackProfileId: string;
  configurableSchema: boolean;
}

export interface SaveSensorProfileConfigListRequest {
  configurableSchema: boolean;
}

export interface SaveSensorMetricRequest {
  measurementType: string;
  payloadKey: string;
  unit: string;
}

export interface SaveSensorMetricListRequest {
  metrics: SaveSensorMetricRequest[];
}

export interface SensorMetricApiResponse {
  measurementType: string;
  payloadKey: string;
  unit: string;
}

export interface UpdateSensorRequest {
  description?: string;
  device_profile_id?: string;
  electricity_sensor?: boolean;
  factory_area_id?: string;
  factory_id?: string;
  name?: string;
  app_key?: string;
  production_resource?: string;
  voltage?: number;
}
