import type { DeviceEUIApiResponse } from "@entities/sensor/model/sensorSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all device EUIs registered on a given Chirpstack profile, ordered oldest-first.
 * The caller should try each EUI against collection-service /payload-tags and use the first that returns data.
 * @param profileId - Chirpstack profile ID
 * @returns Array of device EUIs, empty if none found or request fails
 */
export const getDeviceEUI = async (profileId: string): Promise<string[]> => {
  try {
    const data = await apiRequest<DeviceEUIApiResponse>(
      serviceClient,
      `${API_ROUTES.deviceEUI}${encodeURIComponent(profileId)}`,
      "GET",
    );
    return data.deviceEuis ?? [];
  } catch (error) {
    console.error("Failed to fetch device EUIs:", error);
    return [];
  }
};
