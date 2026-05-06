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
  const params = new URLSearchParams({
    device_eui: deviceEUI,
    from,
    to,
    timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
  });

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
      let message = "Export failed";
      try {
        const body = JSON.parse(text) as { error?: { message?: string } };
        message = body.error?.message ?? message;
      } catch {
        /* not JSON, use default message */
      }
      throw new Error(message);
    }
    throw err;
  }

  // Use the filename from Content-Disposition so it matches what the backend generated.
  const disposition = response.headers["content-disposition"] as
    | string
    | undefined;
  const filename =
    disposition?.match(/filename="([^"]+)"/)?.[1] ?? `${deviceEUI}.csv`;

  const url = window.URL.createObjectURL(response.data as Blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = filename;
  a.click();
  window.URL.revokeObjectURL(url);
};
