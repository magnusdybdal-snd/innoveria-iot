import type { DeviceEUIApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches a sample EUI based on a given chirpstack profile from the device-service.
 * @param profileId - Chirpstack profile ID used to identify the EUI
 * @returns The given EUI for the profile, or null if the request fails
 */
export const getDeviceEUI = async (
  profileId: string,
): Promise<string | null> => {
  try {
    const data = await apiRequest<DeviceEUIApiResponse>(
      serviceClient,
      `${API_ROUTES.deviceEUI}${encodeURIComponent(profileId)}`,
      "GET",
    );
    return data.deviceEui ?? null;
  } catch (error) {
    console.error("Failed to fetch device EUI:", error);
    return null;
  }
};
