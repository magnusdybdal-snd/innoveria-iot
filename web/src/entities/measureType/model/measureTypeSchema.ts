/**
 * This file contains the TypeScript interfaces for the API responses related to measurement types and their readings.
 * It defines the structure of the data returned by the API when fetching measurement type information and their latest readings.
 */
export interface MeasureTypeListApiResponse {
  totalCount: number;
  sensors: MeasureTypeApiResponse[];
}

export interface MeasureTypeApiResponse {
  defaultUnit: string;
  deprecated: boolean;
  description: string;
  displayName: string;
  slug: string;
}

export interface CreateMeasureTypeRequest {
  defaultUnit: string;
  description: string;
  displayName: string;
  slug: string;
}
