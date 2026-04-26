import type {
  PayloadSchemaApiResponse,
  PayloadSchemaListApiResponse,
} from "@entities/payloadSchema/model/payloadSchemaSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches saved payload schema labels for a Chirpstack profile from device-service.
 * @param chirpstackProfileId - Chirpstack profile ID to fetch labels for
 * @returns Array of labeled schema rows, or an empty array if the request fails
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
