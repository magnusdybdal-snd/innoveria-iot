/**
 * This file contains the TypeScript interfaces for the API responses related to payload schemas and their readings.
 * It defines the structure of the data returned by the API when fetching payload schema information.
 */
export interface PayloadSchemaListApiResponse {
  totalCount: number;
  schemas: PayloadSchemaApiResponse[];
}

export interface PayloadSchemaApiResponse {
  chirpstackProfileId: string;
  id: string;
  measurementType: string;
  payloadKey: string;
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
