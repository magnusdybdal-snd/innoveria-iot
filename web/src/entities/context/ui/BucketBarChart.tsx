import type { BucketResponse } from "@entities/context/model/contextSchema";
import WarningAmber from "@mui/icons-material/WarningAmber";
import { BarChart } from "@mui/x-charts/BarChart";

interface BucketBarChartProps {
  buckets: BucketResponse[];
  unit?: string;
}

const MAX_BARS = 500;

/**
 * Renders a bar chart of aggregated time bucket values.
 * @param props - Component props
 * @param props.buckets - Array of time buckets to plot
 * @param props.unit - Optional unit label shown on the value axis
 * @returns Bar chart with period start labels on the x-axis and bucket values on the y-axis
 */
export function BucketBarChart({ buckets, unit }: BucketBarChartProps) {
  const isDownsampled = buckets.length > MAX_BARS;

  const sampled = isDownsampled
    ? Array.from(
        { length: MAX_BARS },
        (_, i) =>
          buckets[Math.round((i * (buckets.length - 1)) / (MAX_BARS - 1))],
      )
    : buckets;

  const labels = sampled.map((b) =>
    new Date(b.periodStart).toLocaleString(undefined, {
      month: "short",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    }),
  );
  const values = sampled.map((b) => b.value);

  return (
    <div>
      {isDownsampled && (
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: 6,
            color: "#f59e0b",
            fontSize: 13,
            marginBottom: 4,
          }}
        >
          <WarningAmber fontSize="small" />
          <span>
            Data downsampled from {buckets.length} to {MAX_BARS} points —
            rendering may not reflect all values accurately.
          </span>
        </div>
      )}
      <BarChart
        xAxis={[
          { scaleType: "band", data: labels, tickLabelStyle: { fontSize: 11 } },
        ]}
        yAxis={[{ label: unit }]}
        series={[{ data: values, color: "#90caf9" }]}
        height={300}
      />
    </div>
  );
}
