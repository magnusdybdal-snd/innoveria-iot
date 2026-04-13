type Tier =
  | { maxMinutes: number; type: "intraday"; intervalMinutes: number }
  | { maxMinutes: number; type: "daily"; intervalDays: number };

// Intraday tiers filter by minute-of-day divisibility.
// Daily tiers filter by midnight ticks spaced intervalDays apart.
const INTERVAL_TIERS: Tier[] = [
  { maxMinutes: 30, type: "intraday", intervalMinutes: 5 },
  { maxMinutes: 60, type: "intraday", intervalMinutes: 15 },
  { maxMinutes: 4 * 60, type: "intraday", intervalMinutes: 30 },
  { maxMinutes: 12 * 60, type: "intraday", intervalMinutes: 60 },
  { maxMinutes: 2 * 24 * 60, type: "intraday", intervalMinutes: 6 * 60 },
  { maxMinutes: 7 * 24 * 60, type: "daily", intervalDays: 1 },
  { maxMinutes: 21 * 24 * 60, type: "daily", intervalDays: 3 },
  { maxMinutes: Infinity, type: "daily", intervalDays: 7 },
];

function pickTier(totalMinutes: number): Tier {
  return (
    INTERVAL_TIERS.find((t) => totalMinutes <= t.maxMinutes) ??
    INTERVAL_TIERS[INTERVAL_TIERS.length - 1]
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
 * ("13 Apr 2026, 14:00").
 *
 * Tick spacing adapts to the selected time range:
 * - ≤ 30 min   → every 5 min
 * - ≤ 1 h      → every 15 min
 * - ≤ 4 h      → every 30 min
 * - ≤ 12 h     → every 1 h
 * - ≤ 2 days   → every 6 h
 * - ≤ 7 days   → every day (midnight)
 * - ≤ 21 days  → every 3 days
 * - > 21 days  → every 7 days
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
  const tier = pickTier(totalMinutes);

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

  if (tier.type === "intraday") {
    timestamps.forEach((d, i) => {
      const iso = data[i];
      if (isFirstOfDay(d, i)) firstOfDayIsos.add(iso);
      const minutesInDay = d.getHours() * 60 + d.getMinutes();
      if (isFirstOfDay(d, i) || minutesInDay % tier.intervalMinutes === 0) {
        visibleIsos.add(iso);
      }
    });
  } else {
    // Daily tier: show the first bucket of every Nth day, counting day
    // boundaries as they appear in the data. This is robust to timezone
    // offsets and non-midnight bucket starts — no clock-time check needed.
    let dayCount = 0;
    timestamps.forEach((d, i) => {
      const iso = data[i];
      if (isFirstOfDay(d, i)) {
        firstOfDayIsos.add(iso);
        if (dayCount % tier.intervalDays === 0) visibleIsos.add(iso);
        dayCount++;
      }
    });
  }

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
