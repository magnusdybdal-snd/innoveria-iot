/**
 * This file contains the TypeScript interfaces for the API responses related to measurement types and their readings.
 * It defines the structure of the data returned by the API when fetching measurement type information and their latest readings.
 */
export interface MeasurementTypeListApiResponse {
  totalCount: number;
  sensors: MeasurementTypeApiResponse[];
}

export interface MeasurementTypeApiResponse {
  defaultUnit: string;
  deprecated: boolean;
  description: string;
  displayName: string;
  slug: string;
}

export interface CreateMeasurementTypeRequest {
  defaultUnit: string;
  description: string;
  displayName: string;
  slug: string;
}
