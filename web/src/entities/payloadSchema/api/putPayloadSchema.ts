import type { SavePayloadSchemaListRequest } from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Puts new payload schema data to the API.
 * @param payloadSchemaDataList - A list of labels containing the measurementType, payloadKey, and unit of the payloadSchema to be saved.
 * @param chirpstackProfileId -
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
