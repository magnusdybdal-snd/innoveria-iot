import type { OperationContext } from "@entities/context/model/contextSchema";

/** Visual state of a work center's sensor mapping. */
export type SensorState = "no_sensor" | "degraded" | "normal";

/**
 * Derives the sensor state for an operation from its sensors array and degraded flag.
 * @param sensors - Sensors mapped to this operation
 * @param degraded - True when the sensor data fetch failed with a technical error
 * @returns The resolved sensor state
 */
export function resolveSensorState(
  sensors: OperationContext["sensors"],
  degraded: boolean,
): SensorState {
  if (sensors.length === 0 && !degraded) return "no_sensor";
  if (degraded) return "degraded";
  return "normal";
}
