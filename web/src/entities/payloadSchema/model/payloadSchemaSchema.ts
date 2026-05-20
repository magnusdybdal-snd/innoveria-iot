/**
 * This file contains the TypeScript interfaces for the API responses related to payload schemas and their readings.
 * It defines the structure of the data returned by the API when fetching payload schema information.
 */
export interface PayloadSchemaListApiResponse {
  total_count: number;
  schemas: PayloadSchemaApiResponse[];
}

export interface PayloadSchemaApiResponse {
  chirpstack_profile_id: string;
  measurement_type: string;
  payload_key: string;
  unit: string;
}

export interface SavePayloadSchemaListRequest {
  labels: SavePayloadSchemaRequest[];
}

export interface SavePayloadSchemaRequest {
  measurementType: string;
  payloadKey: string;
  unit: string;
}

export interface RawPayloadTagsApiResponse {
  keys: string[];
}
