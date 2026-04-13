import type { BucketUnit } from "@entities/context/model/contextSchema";

const UNIT_TO_MINUTES: Record<BucketUnit, number> = {
  minutes: 1,
  hours: 60,
  days: 60 * 24,
  weeks: 60 * 24 * 7,
  months: 60 * 24 * 30,
};

/**
 * Converts a bucket size value and unit into an equivalent number of minutes.
 * @param value - The numeric bucket size
 * @param unit - The time unit to convert from
 * @returns The equivalent number of minutes
 */
export function toMinutes(value: number, unit: BucketUnit): number {
  return value * UNIT_TO_MINUTES[unit];
}

export const BUCKET_UNIT_OPTIONS: { id: BucketUnit; name: string }[] = [
  { id: "minutes", name: "Minutes" },
  { id: "hours", name: "Hours" },
  { id: "days", name: "Days" },
  { id: "weeks", name: "Weeks" },
  { id: "months", name: "Months" },
];

const HOUR_MS = 60 * 60 * 1000;
const DAY_MS = 24 * HOUR_MS;

/**
 * Suggests a human-friendly bucket interval for the given time range.
 * Keeps the number of buckets roughly between 50 and 500.
 * @param from - Range start as an ISO/datetime-local string
 * @param to - Range end as an ISO/datetime-local string
 * @returns Suggested bucket value and unit
 */
export function suggestBucketInterval(
  from: string,
  to: string,
): { value: number; unit: BucketUnit } {
  const rangeMs = new Date(to).getTime() - new Date(from).getTime();

  if (rangeMs > 90 * DAY_MS) return { value: 1, unit: "days" };
  if (rangeMs > 30 * DAY_MS) return { value: 12, unit: "hours" };
  if (rangeMs > 7 * DAY_MS) return { value: 6, unit: "hours" };
  if (rangeMs > DAY_MS) return { value: 1, unit: "hours" };
  if (rangeMs > 2 * HOUR_MS) return { value: 15, unit: "minutes" };
  return { value: 1, unit: "minutes" };
}
