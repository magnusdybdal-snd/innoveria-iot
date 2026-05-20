import type { SensorContext } from "@entities/context/model/contextSchema";
import type { MeasurementTypeApiResponse } from "@entities/measurementType";
import { formatStatus } from "@shared/lib";

/** How to reduce multiple measurements for a single metric into one display value. */
export type AggregationMethod = "latest" | "min" | "max" | "avg" | "sum";

/** A single resolved metric reading ready to display. */
export interface MetricReading {
  /** Stable key from the payload schema — use this as the React list key. */
  payloadKey: string;
  label: string;
  value: string;
  unit: string;
  /** Computed watts value (A × V). Present only when unit is "A" and a voltage metric exists in the same payload. */
  wattsValue?: string;
  /** Label to show when displaying watts instead of amps. */
  wattsLabel?: string;
}

/** All readable metric values for one sensor. */
export interface SensorReadingGroup {
  sensorId: string;
  sensorName: string;
  readings: MetricReading[];
}

/**
 * Reduces an array of numeric values to a single number using the chosen method.
 * Returns 0 for an empty array.
 * @param values - Array of finite numbers
 * @param method - Aggregation method to apply
 * @returns The aggregated value
 */
export function aggregate(values: number[], method: AggregationMethod): number {
  if (values.length === 0) return 0;
  switch (method) {
    case "min":
      return values.reduce((a, b) => (b < a ? b : a));
    case "max":
      return values.reduce((a, b) => (b > a ? b : a));
    case "avg":
      return values.reduce((s, v) => s + v, 0) / values.length;
    case "sum":
      return values.reduce((s, v) => s + v, 0);
    case "latest":
    default:
      return values[0];
  }
}

/**
 * Collects a reading for every schema-defined metric across all sensors,
 * aggregating values across the full measurements window.
 * Non-numeric payload values always fall back to the latest raw value.
 * Groups with no readable values are omitted.
 * @param sensors - Sensor contexts for the operation
 * @param measurementTypes - All known measurement types for label resolution
 * @param method - How to aggregate numeric values across measurements
 * @returns Per-sensor reading groups, in sensor order
 */
export function getAllSensorReadings(
  sensors: SensorContext[],
  measurementTypes: MeasurementTypeApiResponse[],
  method: AggregationMethod,
): SensorReadingGroup[] {
  const groups: SensorReadingGroup[] = [];

  for (const sensor of sensors) {
    if (sensor.metrics.length === 0) continue;

    const sorted = [...sensor.measurements]
      .filter((m) => Number.isFinite(new Date(m.timestamp).getTime()))
      .sort(
        (a, b) =>
          new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime(),
      );

    const readings: MetricReading[] = [];
    for (const metric of sensor.metrics) {
      const latestRaw = sorted[0]?.payload[metric.payloadKey];
      if (latestRaw === undefined || latestRaw === null) continue;

      const mt = measurementTypes.find(
        (m) => m.slug === metric.measurementType,
      );

      let displayValue: string;
      let wattsValue: string | undefined;
      let wattsLabel: string | undefined;

      if (typeof latestRaw === "number") {
        const numericValues = sorted
          .map((m) => m.payload[metric.payloadKey])
          .filter((v): v is number => typeof v === "number");
        const result = aggregate(numericValues, method);
        displayValue = String(Math.round(result * 10) / 10);

        // Special case: if the metric is a current (A), attempt to find a voltage metric in the same payload to compute watts
        if (metric.unit === "A" && sensor.voltage !== null) {
          wattsValue = String(Math.round(result * sensor.voltage * 10) / 10);
          const baseLabel =
            mt?.displayName ?? formatStatus(metric.measurementType);
          wattsLabel = baseLabel.replace(/\bcurrent\b/i, "Power");
        }
      } else {
        displayValue = String(latestRaw);
      }

      readings.push({
        payloadKey: metric.payloadKey,
        label: mt?.displayName ?? formatStatus(metric.measurementType),
        value: displayValue,
        unit: metric.unit ?? mt?.defaultUnit ?? "",
        wattsValue,
        wattsLabel,
      });
    }

    if (readings.length > 0) {
      groups.push({ sensorId: sensor.id, sensorName: sensor.name, readings });
    }
  }

  return groups;
}
