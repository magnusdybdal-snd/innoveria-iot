import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawDeviceEUIResponse = {
  device_euis: string[];
};

/**
 * Fetches all device EUIs registered on a given Chirpstack profile, ordered oldest-first.
 * The caller should try each EUI against collection-service /payload-tags and use the first that returns data.
 * @param profileId - Chirpstack profile ID
 * @returns Array of device EUIs, empty if none found or request fails
 */
export const getDeviceEUIs = async (profileId: string): Promise<string[]> => {
  try {
    const data = await apiRequest<RawDeviceEUIResponse>(
      serviceClient,
      `${API_ROUTES.deviceEUI}${encodeURIComponent(profileId)}`,
      "GET",
    );
    return data.device_euis ?? [];
  } catch (error) {
    console.error("Failed to fetch device EUIs:", error);
    return [];
  }
};
