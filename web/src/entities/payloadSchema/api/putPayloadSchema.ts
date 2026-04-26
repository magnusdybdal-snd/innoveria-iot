import type { SavePayloadSchemaListRequest } from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Puts payload schema data to the API gateway.
 * @param payloadSchemaDataList - A list of payload schemas containing the measurementType, payloadKey, and unit of the payloadSchema to be saved.
 * @param chirpstackProfileId - Chirpstack ID used to identify the profile
 */
export const putPayloadSchema = async (
  payloadSchemaDataList: SavePayloadSchemaListRequest,
  chirpstackProfileId: string,
): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.payloadSchema}${encodeURIComponent(chirpstackProfileId)}`,
    "PUT",
    {
      labels: payloadSchemaDataList.labels.map((payloadSchemaData) => ({
        measurement_type: payloadSchemaData.measurementType,
        payload_key: payloadSchemaData.payloadKey,
        unit: payloadSchemaData.unit,
      })),
    },
  );
};
