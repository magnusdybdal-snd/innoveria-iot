import { LineChart } from "@mui/x-charts/LineChart";

import type { BucketResponse } from "@entities/context/model/contextSchema";
import { buildChartAxisConfig, downsample } from "@shared/lib";

interface BucketLineChartProps {
  buckets: BucketResponse[];
  unit?: string;
}

/**
 * Renders a line chart of aggregated time bucket values.
 * Downsamples to at most 500 points when the bucket array exceeds that limit.
 * @param props - Component props
 * @param props.buckets - Array of time buckets to plot
 * @param props.unit - Optional unit label shown on the value axis
 * @returns Line chart with period start labels on the x-axis and bucket values on the y-axis
 */
export function BucketLineChart({ buckets, unit }: BucketLineChartProps) {
  const sampled = downsample(buckets);

  const timestamps = sampled.map((b) => new Date(b.periodStart));
  const { data, valueFormatter, tickLabelInterval } =
    buildChartAxisConfig(timestamps);
  const values = sampled.map((b) => b.value);

  return (
    <LineChart
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
          area: true,
          showMark: false,
          // TODO: add valueFormatter with unit derived from sensor field type (e.g. temperature → °C, humidity → %)
        },
      ]}
      height={300}
    />
  );
}
