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
