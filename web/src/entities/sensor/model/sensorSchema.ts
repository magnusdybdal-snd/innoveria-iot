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
  status: number;
  machine: string;
  lastReading: string;
  appKey: string;
  senProf: string;
}

export interface SensorReadingApiResponse {
  deviceEui: string;
  timestamp: string;
  payload: Record<string, unknown>;
  companyId: string;
}

export interface CreateSensorRequest {
  companyId: string;
  deviceEui: string;
  deviceProfileId: string;
  name: string;
}
