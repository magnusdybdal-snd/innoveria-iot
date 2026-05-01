import { useState } from "react";

import ErrorOutlineIcon from "@mui/icons-material/ErrorOutline";
import ExpandLessIcon from "@mui/icons-material/ExpandLess";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Collapse from "@mui/material/Collapse";
import IconButton from "@mui/material/IconButton";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

import type {
  BucketResponse,
  Measurement,
  OperationContext,
  SensorContext,
  SensorMetric,
} from "@entities/context/model/contextSchema";
import { BucketLineChart } from "@entities/context/ui/BucketLineChart";
import { resolveSensorState } from "@entities/context/ui/sensorState";
import { SensorStateIndicator } from "@entities/context/ui/SensorStateIndicator";
import type { MeasurementTypeApiResponse } from "@entities/measurementType";
import { formatStatus } from "@shared/lib";

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
    .filter((m) => Number.isFinite(Number(m.payload[payloadKey])))
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
  id: string;
  payloadKey: string;
  unit: string | null;
  label: string;
  buckets: BucketResponse[];
}

/** Chart data grouped by sensor. */
interface SensorChartGroup {
  sensorId: string;
  deviceEui: string;
  powerCharts: MetricChartData[];
  secondaryCharts: MetricChartData[];
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
 * Returns the display name for a measurement type slug, falling back to a formatted slug.
 * @param slug - The measurement type slug from a sensor metric
 * @param measurementTypes - All known measurement types
 * @returns Display label string
 */
function resolveLabel(
  slug: string,
  measurementTypes: MeasurementTypeApiResponse[],
): string {
  return (
    measurementTypes.find((m) => m.slug === slug)?.displayName ??
    formatStatus(slug)
  );
}

/**
 * Derives the primary metric display value from the first sensor that has both a
 * defined schema and a latest measurement.
 * @param sensors - Sensor contexts to search
 * @param measurementTypes - All known measurement types for label resolution
 * @returns Display object or null if no displayable data exists
 */
function getPrimaryMetricDisplay(
  sensors: SensorContext[],
  measurementTypes: MeasurementTypeApiResponse[],
): { label: string; value: string; unit: string } | null {
  for (const sensor of sensors) {
    const metric = sensor.metrics[0];
    if (!metric) continue; // no schema defined for this sensor, skip to next
    if (sensor.measurements.length === 0) continue; // no data points at all, noting to display

    // Get the most recent measurement for this metric and extract the raw value
    const sorted = [...sensor.measurements].sort(
      (a, b) =>
        new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime(),
    );
    const raw = sorted[0].payload[metric.payloadKey];
    if (raw === undefined || raw === null) continue;

    // Resolve the display label and unit, preferring measurement type info but falling back to metric info
    const mt = measurementTypes.find((m) => m.slug === metric.measurementType);
    const value =
      typeof raw === "number" ? String(Math.round(raw * 10) / 10) : String(raw); // round numbers to 1 decimal place for cleaner display, leave non-numeric values as-is
    return {
      label: mt?.displayName ?? formatStatus(metric.measurementType),
      value,
      unit: metric.unit ?? mt?.defaultUnit ?? "",
    };
  }
  return null;
}

/**
 * Builds chart data grouped per sensor for an operation context.
 * Each sensor produces its own group with power and secondary chart arrays.
 * Groups with no chart data are omitted.
 * @param operationContext - The operation context containing sensors and measurements
 * @param measurementTypes - All known measurement types for label resolution
 * @returns Array of per-sensor chart groups
 */
function buildChartData(
  operationContext: OperationContext,
  measurementTypes: MeasurementTypeApiResponse[],
): SensorChartGroup[] {
  return operationContext.sensors
    .map((sensor) => {
      const powerCharts: MetricChartData[] = [];
      const secondaryCharts: MetricChartData[] = [];

      for (const metric of sensor.metrics) {
        const buckets = measurementsToBuckets(
          sensor.measurements,
          metric.payloadKey,
        );
        if (buckets.length === 0) continue;

        const entry: MetricChartData = {
          id: `${sensor.id}:${metric.payloadKey}`,
          payloadKey: metric.payloadKey,
          unit: metric.unit,
          label: resolveLabel(metric.measurementType, measurementTypes),
          buckets,
        };

        if (isPowerMetric(metric)) {
          powerCharts.push(entry);
        } else {
          secondaryCharts.push(entry);
        }
      }

      return {
        sensorId: sensor.id,
        deviceEui: sensor.deviceEui,
        powerCharts,
        secondaryCharts,
      };
    })
    .filter((g) => g.powerCharts.length > 0 || g.secondaryCharts.length > 0);
}

/** Props for the `SensorChartSection` component. */
interface SensorChartSectionProps {
  /** Chart group for a single sensor. */
  group: SensorChartGroup;
  /** Whether to show the sensor EUI label (only needed when multiple sensors exist). */
  showLabel: boolean;
}

/**
 * Collapsible section showing all charts for a single sensor.
 * Uses the same expand/collapse toggle pattern as `MachineEnergyCard`.
 * @param props - Component props
 * @param props.group - The sensor chart group to render
 * @param props.showLabel - Whether to display the sensor EUI as a section label
 * @returns The rendered sensor chart section
 */
function SensorChartSection({ group, showLabel }: SensorChartSectionProps) {
  const [expanded, setExpanded] = useState(false);

  return (
    <Box>
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
        <IconButton size="small" sx={{ color: "primary.main", p: 0, mr: 0.5 }}>
          {expanded ? <ExpandLessIcon /> : <ExpandMoreIcon />}
        </IconButton>
        <Typography variant="body2">
          {expanded ? "Hide graphs" : "Show graphs"}
          {showLabel && (
            <Typography
              component="span"
              variant="caption"
              sx={{ ml: 1, opacity: 0.6 }}
            >
              {group.deviceEui}
            </Typography>
          )}
        </Typography>
      </Box>

      <Collapse in={expanded} unmountOnExit>
        <Box sx={{ mt: 2 }}>
          {group.powerCharts.map((chart) => (
            <Box key={chart.id} sx={{ mb: 2 }}>
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
          {group.secondaryCharts.map((chart) => (
            <Box key={chart.id} sx={{ mb: 2 }}>
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
    </Box>
  );
}

/** Props for the `MachineEnergyCard` component. */
interface MachineEnergyCardProps {
  /** The operation context containing sensor data for this work center. */
  operationContext: OperationContext;
  /** All known measurement types, used to resolve display labels and units. */
  measurementTypes: MeasurementTypeApiResponse[];
}

/**
 * Card displaying a work center's operational status, sensor state, and on-demand
 * sensor graphs. Charts are not mounted until the user expands the card, keeping
 * initial render cost low when many cards are shown simultaneously.
 * @param props - Component props
 * @param props.operationContext - The operation context to render
 * @param props.measurementTypes
 * @returns The rendered machine energy card
 */
export function MachineEnergyCard({
  operationContext,
  measurementTypes,
}: MachineEnergyCardProps) {
  const { operation, sensors, degraded } = operationContext;
  const hasAnySchema = sensors.some((s) => s.metrics.length > 0);
  const primaryDisplay = getPrimaryMetricDisplay(sensors, measurementTypes);
  const sensorState = resolveSensorState(sensors, degraded);
  const sensorGroups = buildChartData(operationContext, measurementTypes);
  const hasCharts = sensorGroups.length > 0;

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
      <Box sx={{ mb: 1.5 }}></Box>

      {sensors.length > 0 && !hasAnySchema ? (
        <Box sx={{ display: "flex", alignItems: "center", gap: 0.5, mb: 1.5 }}>
          <Tooltip title="No payload schema defined for this sensor">
            <ErrorOutlineIcon fontSize="small" color="error" />
          </Tooltip>
          <Typography variant="body2" sx={{ color: "error.main" }}>
            No schema defined
          </Typography>
        </Box>
      ) : hasAnySchema ? (
        <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 1.5 }}>
          <Typography variant="body2" sx={{ opacity: 0.7 }}>
            {primaryDisplay?.label ?? "—"}:
          </Typography>
          <Typography variant="body1" fontWeight={600}>
            {primaryDisplay?.value ?? "—"}
          </Typography>
          {primaryDisplay?.unit && (
            <Typography variant="body2" sx={{ opacity: 0.5 }}>
              {primaryDisplay.unit}
            </Typography>
          )}
        </Box>
      ) : null}

      {/* One collapsible section per sensor — label shown only when multiple sensors exist */}
      {hasCharts &&
        sensorGroups.map((group) => (
          <SensorChartSection
            key={group.sensorId}
            group={group}
            showLabel={sensorGroups.length > 1}
          />
        ))}
    </Card>
  );
}
