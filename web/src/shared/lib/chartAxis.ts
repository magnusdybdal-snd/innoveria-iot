const INTERVAL_TIERS = [
  { maxMinutes: 30, intervalMinutes: 5 },
  { maxMinutes: 60, intervalMinutes: 15 },
  { maxMinutes: 4 * 60, intervalMinutes: 30 },
  { maxMinutes: 12 * 60, intervalMinutes: 60 },
  { maxMinutes: 2 * 24 * 60, intervalMinutes: 6 * 60 },
  { maxMinutes: Infinity, intervalMinutes: 24 * 60 },
];

function pickInterval(totalMinutes: number): number {
  return (
    INTERVAL_TIERS.find((t) => totalMinutes <= t.maxMinutes)?.intervalMinutes ??
    24 * 60
  );
}

function formatDate(d: Date): string {
  return d.toLocaleString(undefined, { month: "short", day: "numeric" });
}

function formatTime(d: Date): string {
  return d.toLocaleString(undefined, { hour: "2-digit", minute: "2-digit" });
}

/**
 * Computes x-axis labels and a tick visibility filter for time-series bucket charts.
 *
 * Ticks are shown at logical time boundaries (5 min, 15 min, 30 min, 1 h, 6 h, midnight)
 * based on the total range of the timestamps. Midnight ticks show as date-only ("Mar 5").
 * If no midnight falls in the range, the first tick shows the date instead of the time
 * so the axis always has at least one date anchor.
 * @param timestamps - Ordered array of bucket start times
 * @returns labels for every point and a tickLabelInterval predicate for MUI x-charts
 */
export function buildChartAxisConfig(timestamps: Date[]): {
  labels: string[];
  tickLabelInterval: (value: unknown) => boolean;
} {
  if (timestamps.length === 0) {
    return { labels: [], tickLabelInterval: () => false };
  }

  const first = timestamps[0].getTime();
  const last = timestamps[timestamps.length - 1].getTime();
  const totalMinutes = (last - first) / 60_000;
  const intervalMinutes = pickInterval(totalMinutes);

  const isFirstOfDay = (d: Date, i: number): boolean => {
    if (i === 0) return true;
    const prev = timestamps[i - 1];
    return d.getDate() !== prev.getDate() || d.getMonth() !== prev.getMonth();
  };

  const labels = timestamps.map((d, i) => {
    if (isFirstOfDay(d, i)) return formatDate(d);
    return formatTime(d);
  });

  // Pre-compute which labels are visible so the predicate uses value lookup
  // instead of index — MUI x-charts may call tickLabelInterval with indices
  // outside our timestamps array bounds.
  const visibleLabels = new Set<string>();
  timestamps.forEach((d, i) => {
    const minutesInDay = d.getHours() * 60 + d.getMinutes();
    if (isFirstOfDay(d, i) || minutesInDay % intervalMinutes === 0) {
      visibleLabels.add(labels[i]);
    }
  });

  const tickLabelInterval = (value: unknown): boolean =>
    visibleLabels.has(String(value));

  return { labels, tickLabelInterval };
}
