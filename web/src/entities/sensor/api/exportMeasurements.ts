import axios from "axios";

import { serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Downloads measurements for a device within a time range as a CSV file.
 * Uses axios (blob) so the Bearer token is included in the request.
 * Throws an Error with a human-readable message on failure.
 * @param deviceEUI - LoRaWAN Device EUI
 * @param from - Start of range (RFC3339, e.g. "2026-01-01T00:00:00Z")
 * @param to - End of range (RFC3339, e.g. "2026-05-01T00:00:00Z")
 */
export const exportMeasurements = async (
  deviceEUI: string,
  from: string,
  to: string,
): Promise<void> => {
  const params = new URLSearchParams({ device_eui: deviceEUI, from, to });

  let response;
  try {
    response = await serviceClient.get(
      `${API_ROUTES.collectionExport}?${params.toString()}`,
      { responseType: "blob" },
    );
  } catch (err) {
    if (axios.isAxiosError(err) && err.response?.data instanceof Blob) {
      // Blob responses need to be read back as text to extract the error message.
      const text = await err.response.data.text();
      try {
        const body = JSON.parse(text) as { error?: { message?: string } };
        throw new Error(body.error?.message ?? "Export failed");
      } catch {
        throw new Error("Export failed");
      }
    }
    throw err;
  }

  const url = window.URL.createObjectURL(response.data as Blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = `${deviceEUI}.csv`;
  a.click();
  window.URL.revokeObjectURL(url);
};
