import type {
  PayloadSchemaApiResponse,
  PayloadSchemaListApiResponse,
} from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all payload tags from the collection-service via the API gateway.
 * Maps snake_case API response keys to camelCase.
 * @param chirpstackProfileId - Device EUI used to identify the tags
 * @returns Array of PayloadTagsApiResponse objects, or an empty array if the request fails
 */
export const getPayloadSchema = async (
  chirpstackProfileId: string,
): Promise<PayloadSchemaApiResponse[]> => {
  try {
    const data = await apiRequest<PayloadSchemaListApiResponse>(
      serviceClient,
      `${API_ROUTES.payloadSchema}${encodeURIComponent(chirpstackProfileId)}`,
      "GET",
    );
    return data.schemas ?? [];
  } catch (error) {
    console.error("Failed to get payload tags:", error);
    return [];
  }
};
