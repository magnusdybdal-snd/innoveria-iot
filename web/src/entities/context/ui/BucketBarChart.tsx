import { BarChart } from "@mui/x-charts/BarChart";

import type { BucketResponse } from "@entities/context/model/contextSchema";

interface BucketBarChartProps {
  buckets: BucketResponse[];
  unit?: string;
}

/**
 * Renders a bar chart of aggregated time bucket values.
 * @param props - Component props
 * @param props.buckets - Array of time buckets to plot
 * @param props.unit - Optional unit label shown on the value axis
 * @returns Bar chart with period start labels on the x-axis and bucket values on the y-axis
 */
export function BucketBarChart({ buckets, unit }: BucketBarChartProps) {
  const labels = buckets.map((b) =>
    new Date(b.periodStart).toLocaleString(undefined, {
      month: "short",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    }),
  );
  const values = buckets.map((b) => b.value);

  return (
    <BarChart
      xAxis={[
        { scaleType: "band", data: labels, tickLabelStyle: { fontSize: 11 } },
      ]}
      yAxis={[{ label: unit }]}
      series={[{ data: values, color: "#90caf9" }]}
      height={300}
    />
  );
}
