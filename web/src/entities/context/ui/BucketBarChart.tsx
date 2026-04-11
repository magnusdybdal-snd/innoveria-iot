import { BarChart } from "@mui/x-charts/BarChart";

import type { BucketResponse } from "@entities/context/model/contextSchema";

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
    <BarChart
      xAxis={[
        { scaleType: "band", data: labels, tickLabelStyle: { fontSize: 11 } },
      ]}
      yAxis={[{ label: unit }]}
      series={[{ data: values, color: "#fe8019" }]}
      height={300}
    />
  );
}
