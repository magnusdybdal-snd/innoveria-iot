import type { RawPayloadTagsApiResponse } from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all payload tags from the collection-service via the API gateway.
 * @param deviceEUI - Device EUI used to identify the tags
 * @returns Array of RawPayloadTagsApiResponse objects, or an empty array if the request fails
 */
export const getPayloadTags = async (deviceEUI: string): Promise<string[]> => {
  try {
    const data = await apiRequest<RawPayloadTagsApiResponse>(
      serviceClient,
      `${API_ROUTES.payloadTags}${encodeURIComponent(deviceEUI)}`,
      "GET",
    );
    return data.keys ?? [];
  } catch (error) {
    console.error("Failed to get payload tags:", error);
    return [];
  }
};
