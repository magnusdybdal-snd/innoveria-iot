import { serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Downloads measurements for a device within a time range as a CSV file.
 * Uses axios (blob) so the Bearer token is included in the request.
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
  });

  const response = await serviceClient.get(
    `${API_ROUTES.collectionExport}?${params.toString()}`,
    { responseType: "blob" },
  );

  const url = window.URL.createObjectURL(response.data as Blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = `${deviceEUI}.csv`;
  a.click();
  window.URL.revokeObjectURL(url);
};
