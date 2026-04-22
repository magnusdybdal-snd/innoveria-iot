import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Typography from "@mui/material/Typography";

import type {
  OperationContext,
  SensorContext,
} from "@entities/context/model/contextSchema";

/**
 * Derives the latest display value and unit from the first sensor's first metric.
 * Returns null if no measurements exist.
 * @param sensor - Sensor context to extract the latest reading from
 * @returns Object with `value` and `unit` strings, or null if no data is available
 */
function getLatestReading(
  sensor: SensorContext,
): { value: string; unit: string } | null {
  const metric = sensor.metrics[0];
  if (!metric || sensor.measurements.length === 0) return null;

  const sorted = [...sensor.measurements].sort(
    (a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime(),
  );
  const raw = sorted[0].payload[metric.payloadKey];
  if (raw === undefined || raw === null) return null;

  const value =
    typeof raw === "number" ? String(Math.round(raw * 10) / 10) : String(raw);
  return { value, unit: metric.unit ?? "" };
}

type CardState = "normal" | "warning" | "missing";

/**
 * Determines the visual state of a machine card based on its sensor data.
 * - `missing`: no sensor is mapped to this operation at all
 * - `warning`: a sensor is mapped but has no measurements, or data fetch failed
 * - `normal`: sensor data is available
 * @param sensors - List of sensor contexts for the operation
 * @param degraded - True if the sensor data fetch failed with a technical error
 * @returns The resolved card state
 */
function resolveCardState(
  sensors: OperationContext["sensors"],
  degraded: boolean,
): CardState {
  if (sensors.length === 0 && !degraded) return "missing";
  if (degraded || sensors.every((s) => s.measurements.length === 0))
    return "warning";
  return "normal";
}

const OVERLAY: Record<"warning" | "missing", string> = {
  missing: "rgba(211, 47, 47, 0.25)",
  warning: "rgba(245, 124, 0, 0.25)",
};

/** Props for the `MachineCard` component. */
interface MachineCardProps {
  /** The operation context containing sensor data for this work center. */
  operationContext: OperationContext;
}

/**
 * Card displaying a work center and its live sensor reading.
 * Visual states:
 * - Normal: shows metric value and unit
 * - Yellow overlay: sensor mapped but no data (config issue / machine unused)
 * - Red overlay: no sensor mapped at all
 * @param props - Component props
 * @param props.operationContext - The operation context to render
 * @returns The rendered machine card
 */
export function MachineCard({ operationContext }: MachineCardProps) {
  const { operation, sensors, degraded } = operationContext;
  const cardState = resolveCardState(sensors, degraded);

  const firstSensor = sensors[0] ?? null;
  const reading = firstSensor ? getLatestReading(firstSensor) : null;

  let displayValue: string;
  let displayUnit: string;

  if (cardState === "missing") {
    displayValue = "No sensor";
    displayUnit = "";
  } else if (cardState === "warning") {
    displayValue = "No data";
    displayUnit = "";
  } else {
    displayValue = reading?.value ?? "—";
    displayUnit = reading?.unit ?? "";
  }

  return (
    <Card
      sx={{
        backgroundColor: "secondary.light",
        borderRadius: 3,
        p: 2,
        color: "primary.main",
        aspectRatio: "1",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        gap: 1,
        width: 200,
        position: "relative",
        overflow: "hidden",
        containerType: "inline-size",
      }}
    >
      {cardState !== "normal" && (
        <Box
          sx={{
            position: "absolute",
            inset: 0,
            backgroundColor: OVERLAY[cardState],
            pointerEvents: "none",
          }}
        />
      )}
      <Typography variant="subtitle2" sx={{ opacity: 0.7 }}>
        {operation.productionResource.number}
      </Typography>
      <Typography
        variant="h2"
        sx={{
          fontWeight: 600,
          textAlign: "center",
          whiteSpace: "nowrap",
          fontSize: "clamp(0.90rem, 90cqi, 3rem)",
        }}
      >
        {displayValue}
      </Typography>
      <Typography variant="subtitle2" sx={{ opacity: 0.7 }}>
        {displayUnit}
      </Typography>
    </Card>
  );
}
