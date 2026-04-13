import type { RawPayloadTagsApiResponse } from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all payload tags from the collection-service via the API gateway.
 * Maps snake_case API response keys to camelCase.
 * @returns Array of PayloadTagsApiResponse objects, or an empty array if the request fails
 */
export const getPayloadTags = async (): Promise<string[]> => {
  const data = await apiRequest<RawPayloadTagsApiResponse>(
    serviceClient,
    API_ROUTES.payloadTags,
    "GET",
  );

  return data.keys ?? [];
};
