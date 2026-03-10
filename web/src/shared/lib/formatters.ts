/**
 * Pure formatting utilities for visual presentation only.
 * These functions do not alter underlying data — they return formatted strings for display purposes.
 */

/**
 * Formats an ISO 8601 or RFC1123 timestamp into a human-readable local date and time string.
 * Returns "Never" for zero-value timestamps (0001-01-01), and "Unknown" if the date is invalid.
 * @param timestamp - The raw timestamp string from the API
 * @returns A formatted date/time string, e.g. "09.03.2026, 15:42"
 */
export function formatTimestamp(timestamp: string): string {
  const date = new Date(timestamp);

  if (isNaN(date.getTime())) return "Unknown";

  // Zero-value timestamp from Go backends
  if (date.getFullYear() === 1) return "Never";

  return date.toLocaleString(undefined, {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

/**
 * Formats a numeric sensor reading to one decimal place for display.
 * Returns "N/A" if the value is null, undefined, or not a finite number.
 * @param value - The raw numeric reading from the sensor payload
 * @param unit - Optional unit suffix to append, e.g. "°C", "%", "V"
 * @returns A formatted string, e.g. "23.4°C"
 */
export function formatReading(value: unknown, unit?: string): string {
  if (value === null || value === undefined) return "N/A";
  const num = Number(value);
  if (!isFinite(num)) return "N/A";
  const formatted = num.toFixed(1);
  return unit ? `${formatted}${unit}` : formatted;
}
