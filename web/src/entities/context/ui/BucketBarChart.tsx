import { BarChart } from "@mui/x-charts/BarChart";

import type { BucketResponse } from "@entities/context/model/contextSchema";
import { buildChartAxisConfig, downsample } from "@shared/lib";

interface BucketBarChartProps {
  buckets: BucketResponse[];
  unit?: string;
}

/**
 * Renders a bar chart of aggregated time bucket values.
 * Downsamples to at most 500 bars when the bucket array exceeds that limit.
 * @param props - Component props
 * @param props.buckets - Array of time buckets to plot
 * @param props.unit - Optional unit label shown on the value axis
 * @returns Bar chart with period start labels on the x-axis and bucket values on the y-axis
 */
export function BucketBarChart({ buckets, unit }: BucketBarChartProps) {
  const sampled = downsample(buckets);

  const timestamps = sampled.map((b) => new Date(b.periodStart));
  const { data, valueFormatter, tickLabelInterval } =
    buildChartAxisConfig(timestamps);
  const values = sampled.map((b) => b.value);

  return (
    <BarChart
      xAxis={[
        {
          scaleType: "band",
          data,
          valueFormatter,
          tickLabelStyle: { fontSize: 11 },
          tickLabelInterval,
        },
      ]}
      yAxis={[{ label: unit }]}
      series={[
        {
          data: values,
          color: "#fe8019",
          // TODO: add valueFormatter with unit derived from sensor field type (e.g. temperature → °C, humidity → %)
        },
      ]}
      height={300}
    />
  );
}
