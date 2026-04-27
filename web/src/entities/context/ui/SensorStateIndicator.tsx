import ErrorIcon from "@mui/icons-material/Error";
import WarningAmberIcon from "@mui/icons-material/WarningAmber";
import Tooltip from "@mui/material/Tooltip";

import type { SensorState } from "./sensorState";

/** Props for the `SensorStateIndicator` component. */
interface SensorStateIndicatorProps {
  /** The sensor state to visualise. Renders nothing when `"normal"`. */
  state: SensorState;
}

/**
 * Icon badge that communicates sensor mapping and fetch status for a work center.
 * Renders nothing when state is `"normal"`.
 * @param props - Component props
 * @param props.state - One of `"no_sensor"`, `"degraded"`, or `"normal"`
 * @returns A tooltipped icon, or null when state is normal
 */
export function SensorStateIndicator({ state }: SensorStateIndicatorProps) {
  if (state === "normal") return null;

  if (state === "no_sensor") {
    return (
      <Tooltip title="No sensor mapped to this work center">
        <WarningAmberIcon sx={{ color: "warning.main", fontSize: 20 }} />
      </Tooltip>
    );
  }

  return (
    <Tooltip title="Sensor data unavailable — fetch failed">
      <ErrorIcon sx={{ color: "error.main", fontSize: 20 }} />
    </Tooltip>
  );
}
