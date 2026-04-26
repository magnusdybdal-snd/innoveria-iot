import type {
  PayloadSchemaApiResponse,
  PayloadSchemaListApiResponse,
} from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all payload schemas  for a profile from the device-service via the API gateway.
 * @param chirpstackProfileId - ID used to identify the profile
 * @returns Array of PayloadSchemaApiResponse objects, or an empty array if the request fails
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
    console.error("Failed to get payload schemas:", error);
    return [];
  }
};
