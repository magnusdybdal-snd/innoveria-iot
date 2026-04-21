import { useState } from "react";

import ExpandLessIcon from "@mui/icons-material/ExpandLess";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Chip from "@mui/material/Chip";
import Collapse from "@mui/material/Collapse";
import IconButton from "@mui/material/IconButton";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

import type {
  BucketResponse,
  Measurement,
  OperationContext,
  SensorMetric,
} from "@entities/context/model/contextSchema";
import { BucketLineChart } from "@entities/context/ui/BucketLineChart";
import {
  resolveSensorState,
  SensorStateIndicator,
} from "@entities/context/ui/SensorStateIndicator";
import { formatStatus } from "@shared/lib";

/** Maps ERP operation status strings to MUI Chip color variants. */
const STATUS_COLOR: Record<string, "default" | "warning" | "info" | "success"> =
  {
    pending: "warning",
    in_progress: "info",
    completed: "success",
  };

/**
 * Converts raw measurements for a single payload key into the `BucketResponse`
 * shape expected by `BucketLineChart`.
 * @param measurements - All measurements for a sensor
 * @param payloadKey - The specific key to extract from each measurement's payload
 * @returns Sorted array of bucket-shaped data points
 */
function measurementsToBuckets(
  measurements: Measurement[],
  payloadKey: string,
): BucketResponse[] {
  return measurements
    .filter((m) => m.payload[payloadKey] !== undefined)
    .sort(
      (a, b) =>
        new Date(a.timestamp).getTime() - new Date(b.timestamp).getTime(),
    )
    .map((m) => ({
      periodStart: m.timestamp,
      periodEnd: m.timestamp,
      value: Number(m.payload[payloadKey]),
    }));
}

/** A resolved metric with its chart data ready to render. */
interface MetricChartData {
  payloadKey: string;
  unit: string | null;
  label: string;
  buckets: BucketResponse[];
}

/**
 * Returns true for metrics representing power or electrical energy.
 * Separates the primary power trend from secondary metrics (temperature, nitrogen, etc.).
 * @param metric - Sensor metric to evaluate
 * @returns Whether the metric is power-related
 */
function isPowerMetric(metric: SensorMetric): boolean {
  const t = metric.measurementType.toLowerCase();
  return t.includes("power") || t.includes("electric");
}

/**
 * Builds chart data arrays for an operation context, split into
 * power-related metrics and all other secondary metrics.
 * @param operationContext - The operation context containing sensors and measurements
 * @returns Object with `powerCharts` and `secondaryCharts` arrays
 */
function buildChartData(operationContext: OperationContext): {
  powerCharts: MetricChartData[];
  secondaryCharts: MetricChartData[];
} {
  const powerCharts: MetricChartData[] = [];
  const secondaryCharts: MetricChartData[] = [];

  for (const sensor of operationContext.sensors) {
    for (const metric of sensor.metrics) {
      const buckets = measurementsToBuckets(
        sensor.measurements,
        metric.payloadKey,
      );
      if (buckets.length === 0) continue;

      const entry: MetricChartData = {
        payloadKey: metric.payloadKey,
        unit: metric.unit,
        label: formatStatus(metric.payloadKey),
        buckets,
      };

      if (isPowerMetric(metric)) {
        powerCharts.push(entry);
      } else {
        secondaryCharts.push(entry);
      }
    }
  }

  return { powerCharts, secondaryCharts };
}

/** Props for the `MachineEnergyCard` component. */
interface MachineEnergyCardProps {
  /** The operation context containing sensor data for this work center. */
  operationContext: OperationContext;
}

/**
 * Card displaying a work center's operational status, sensor state, and on-demand
 * sensor graphs. Charts are not mounted until the user expands the card, keeping
 * initial render cost low when many cards are shown simultaneously.
 *
 * The energy (kWh) field is a placeholder pending backend support.
 * @param props - Component props
 * @param props.operationContext - The operation context to render
 * @returns The rendered machine energy card
 */
export function MachineEnergyCard({
  operationContext,
}: MachineEnergyCardProps) {
  const { operation, sensors, degraded } = operationContext;
  const sensorState = resolveSensorState(sensors, degraded);
  const { powerCharts, secondaryCharts } = buildChartData(operationContext);
  const hasCharts = powerCharts.length > 0 || secondaryCharts.length > 0;

  const [expanded, setExpanded] = useState(false);

  return (
    <Card
      sx={{
        backgroundColor: "secondary.light",
        borderRadius: 3,
        p: 2,
        color: "primary.main",
        minWidth: 320,
        flex: "1 1 320px",
      }}
    >
      {/* Header row: machine name + sensor state */}
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          mb: 1,
        }}
      >
        <Typography variant="h6" fontWeight={600}>
          {operation.productionResource.number}
        </Typography>
        <SensorStateIndicator state={sensorState} />
      </Box>

      {/* Production resource description */}
      {operation.productionResource.description && (
        <Typography variant="body2" sx={{ opacity: 0.7, mb: 1 }}>
          {operation.productionResource.description}
        </Typography>
      )}

      {/* Operation status chip */}
      <Box sx={{ mb: 1.5 }}>
        <Chip
          label={formatStatus(operation.status)}
          color={STATUS_COLOR[operation.status] ?? "default"}
          size="small"
        />
      </Box>

      {/* Energy placeholder — TODO: wire up kWh once backend returns energy totals (see issue #___) */}
      <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 1.5 }}>
        <Typography variant="body2" sx={{ opacity: 0.7 }}>
          Energy:
        </Typography>
        <Tooltip title="Energy totals will be available in a future update">
          <Typography variant="body1" fontWeight={600}>
            —
          </Typography>
        </Tooltip>
        <Typography variant="body2" sx={{ opacity: 0.5 }}>
          kWh
        </Typography>
      </Box>

      {/* Graph toggle — only shown when there is chart data to display */}
      {hasCharts && (
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            cursor: "pointer",
            userSelect: "none",
            mt: 1,
          }}
          onClick={() => setExpanded((v) => !v)}
        >
          <IconButton
            size="small"
            sx={{ color: "primary.main", p: 0, mr: 0.5 }}
          >
            {expanded ? <ExpandLessIcon /> : <ExpandMoreIcon />}
          </IconButton>
          <Typography variant="body2">
            {expanded ? "Hide graphs" : "Show graphs"}
          </Typography>
        </Box>
      )}

      {/* Charts — only mounted when expanded to avoid rendering cost upfront */}
      <Collapse in={expanded} unmountOnExit>
        <Box sx={{ mt: 2 }}>
          {powerCharts.map((chart) => (
            <Box key={chart.payloadKey} sx={{ mb: 2 }}>
              <Typography variant="caption" sx={{ opacity: 0.7 }}>
                {chart.label}
                {chart.unit ? ` (${chart.unit})` : ""}
              </Typography>
              <BucketLineChart
                buckets={chart.buckets}
                unit={chart.unit ?? undefined}
              />
            </Box>
          ))}
          {secondaryCharts.map((chart) => (
            <Box key={chart.payloadKey} sx={{ mb: 2 }}>
              <Typography variant="caption" sx={{ opacity: 0.7 }}>
                {chart.label}
                {chart.unit ? ` (${chart.unit})` : ""}
              </Typography>
              <BucketLineChart
                buckets={chart.buckets}
                unit={chart.unit ?? undefined}
              />
            </Box>
          ))}
        </Box>
      </Collapse>
    </Card>
  );
}
