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
  status: number;
  machine: string;
  lastReading: string;
  sensorProfileId: string;
}

export interface SensorReadingApiResponse {
  deviceEui: string;
  timestamp: string;
  payload: Record<string, unknown>;
  companyId: string;
}

export interface DeviceEUIApiResponse {
  deviceEui: string;
}

export interface CreateSensorRequest {
  companyId: string;
  electricitySensor: boolean;
  factoryId: string;
  factoryAreaId: string;
  deviceEui: string;
  appKey: string;
  sensorProfileId: string;
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
