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
  description: string | null;
  factory: string;
  factoryAreaId: string;
  status: number;
  productionResource: number | null;
  electricitySensor: boolean;
  voltage: number | null;
  lastReading: string;
  sensorProfileId: string;
}

export interface UpdateSensorRequest {
  name?: string;
  description?: string;
  factoryId?: string;
  factoryAreaId?: string;
  sensorProfileId?: string;
  electricitySensor?: boolean;
  voltage?: number | null;
  productionResource?: number | null;
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
  sensorProfileId: string;
  name: string;
  description?: string;
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
  mac_version: string;
  vendor_id: string;
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
