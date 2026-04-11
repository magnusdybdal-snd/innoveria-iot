import { LineChart } from "@mui/x-charts/LineChart";

import type { BucketResponse } from "@entities/context/model/contextSchema";

interface BucketLineChartProps {
  buckets: BucketResponse[];
  unit?: string;
  area?: boolean;
}

/**
 * Renders a line chart of aggregated time bucket values.
 * @param props - Component props
 * @param props.buckets - Array of time buckets to plot
 * @param props.unit - Optional unit label shown on the value axis
 * @param props.area - If true, fills the area under the line with color
 * @returns Line chart with period start labels on the x-axis and bucket values on the y-axis
 */
export function BucketLineChart({
  buckets,
  unit,
  area = false,
}: BucketLineChartProps) {
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
    <LineChart
      xAxis={[
        { scaleType: "band", data: labels, tickLabelStyle: { fontSize: 11 } },
      ]}
      yAxis={[{ label: unit }]}
      series={[{ data: values, color: "#fe8019", area: area, showMark: false }]}
      height={300}
    />
  );
}
