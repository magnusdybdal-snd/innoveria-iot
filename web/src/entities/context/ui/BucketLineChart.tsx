import { LineChart } from "@mui/x-charts/LineChart";

import type { BucketResponse } from "@entities/context/model/contextSchema";
import { buildChartAxisConfig, downsample } from "@shared/lib";

interface BucketLineChartProps {
  buckets: BucketResponse[];
  unit?: string;
  area?: boolean;
  /** Left edge of the x-axis — anchors the chart to the operation start even when data starts later. */
  from?: Date;
  /** Right edge of the x-axis — anchors the chart to the operation end (or now) even when data ends earlier. */
  to?: Date;
}

/**
 * Renders a line chart of aggregated time bucket values.
 * Downsamples to at most 500 points when the bucket array exceeds that limit.
 * When `from`/`to` are provided the x-axis spans the full operation time window,
 * padding with null boundary points where no data exists.
 * @param props - Component props
 * @param props.buckets - Array of time buckets to plot
 * @param props.unit - Optional unit label shown on the value axis
 * @param props.area - Whether to render a filled area under the line (default false)
 * @param props.from - Operation actual start date — anchors the left x-axis edge
 * @param props.to - Operation actual end date (or now if still running) — anchors the right x-axis edge
 * @returns Line chart with period start labels on the x-axis and bucket values on the y-axis
 */
export function BucketLineChart({
  buckets,
  unit,
  area = false,
  from,
  to,
}: BucketLineChartProps) {
  const sampled = downsample(buckets);

  const timestamps = sampled.map((b) => new Date(b.periodStart));
  const values: (number | null)[] = sampled.map((b) => b.value);

  const firstTs = timestamps[0]?.getTime() ?? Infinity;
  const lastTs = timestamps[timestamps.length - 1]?.getTime() ?? -Infinity;
  if (from && from.getTime() < firstTs) {
    timestamps.unshift(from);
    values.unshift(null);
  }
  if (to && to.getTime() > lastTs) {
    timestamps.push(to);
    values.push(null);
  }

  const { data, valueFormatter, tickLabelInterval } =
    buildChartAxisConfig(timestamps);

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
          area,
          showMark: false,
          // TODO: add valueFormatter with unit derived from sensor field type (e.g. temperature → °C, humidity → %)
        },
      ]}
      height={300}
    />
  );
}
