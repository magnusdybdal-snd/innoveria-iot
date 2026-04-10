import type { SensorReadingApiResponse } from "@entities/sensor/model/sensorSchema";

/**
 * Type guard that narrows a sensor reading to one with a non-null, non-empty payload.
 *
 * The backend cannot guarantee a payload for newly registered or misconfigured
 * devices, so this guard must be applied before accessing payload keys.
 * When it returns true, TypeScript narrows the type so that both `reading`
 * and `reading.payload` are guaranteed non-null.
 * @param reading - The latest sensor reading, or null if none has been received
 * @returns True if the reading has at least one payload field, false otherwise
 */
export function hasPayloadData(
  reading: SensorReadingApiResponse | null,
): reading is SensorReadingApiResponse & { payload: Record<string, unknown> } {
  return (
    reading !== null &&
    reading.payload !== null &&
    Object.keys(reading.payload).length > 0
  );
}
