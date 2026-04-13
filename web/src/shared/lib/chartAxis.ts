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

function formatFull(d: Date): string {
  return d.toLocaleString(undefined, {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

/**
 * Computes x-axis data, a value formatter, and a tick visibility filter for
 * time-series bucket charts.
 *
 * X-axis data is stored as ISO timestamp strings so the full date is always
 * available. The `valueFormatter` renders an abbreviated label for axis ticks
 * ("Apr 13", "14:00") and a full date+time string in the hover tooltip
 * ("13 Apr 2026, 14:00"). Ticks are shown at logical time boundaries (5 min,
 * 15 min, 30 min, 1 h, 6 h, midnight) based on the total range.
 * @param timestamps - Ordered array of bucket start times
 * @returns `data` (ISO strings), `valueFormatter`, and a `tickLabelInterval`
 *   predicate for MUI x-charts
 */
export function buildChartAxisConfig(timestamps: Date[]): {
  data: string[];
  valueFormatter: (value: string, context: { location: string }) => string;
  tickLabelInterval: (value: unknown) => boolean;
} {
  if (timestamps.length === 0) {
    return {
      data: [],
      valueFormatter: () => "",
      tickLabelInterval: () => false,
    };
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

  const data = timestamps.map((d) => d.toISOString());

  // Precompute which ISOs are first-of-day (drives tick label format) and
  // which should have visible tick labels.
  const firstOfDayIsos = new Set<string>();
  const visibleIsos = new Set<string>();

  timestamps.forEach((d, i) => {
    const iso = data[i];
    if (isFirstOfDay(d, i)) {
      firstOfDayIsos.add(iso);
      visibleIsos.add(iso);
    }
    const minutesInDay = d.getHours() * 60 + d.getMinutes();
    if (minutesInDay % intervalMinutes === 0) {
      visibleIsos.add(iso);
    }
  });

  const valueFormatter = (
    value: string,
    context: { location: string },
  ): string => {
    const d = new Date(value);
    if (context.location === "tooltip") return formatFull(d);
    return firstOfDayIsos.has(value) ? formatDate(d) : formatTime(d);
  };

  const tickLabelInterval = (value: unknown): boolean =>
    visibleIsos.has(String(value));

  return { data, valueFormatter, tickLabelInterval };
}
