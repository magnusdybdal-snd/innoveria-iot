import type { Gateway } from "@/mocks/gateways";

export type SortDirection = "asc" | "desc";
export type GatewaySortKey = "Name" | "Status" | "Last seen";

// Parses an RFC1123 date string (from the API) into a Unix timestamp for comparison.
// Returns 0 if the string is not a valid date.
function parseLastSeenToMs(lastSeen: string): number {
  const ms = Date.parse(lastSeen);
  return isNaN(ms) ? 0 : ms;
}

// Sorting gateways based on Key (GatewaySortkeys)
/**
 * Returns a sorted copy of the gateway array based on the given column and direction.
 * @param gateways - Array of Gateway objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted Gateway array (does not mutate the input)
 */
export function sortGateways(
  gateways: Gateway[],
  key: GatewaySortKey | null,
  direction: SortDirection,
): Gateway[] {
  if (!key) return gateways;
  return [...gateways].sort((a, b) => {
    let cmp = 0;
    if (key === "Name") {
      cmp = a.name.localeCompare(b.name);
    } else if (key === "Status") {
      // asc = Online first (0), neverconnected second (1), offline third (2)
      cmp = a.status - b.status;
    } else if (key === "Last seen") {
      cmp = parseLastSeenToMs(a.lastSeen) - parseLastSeenToMs(b.lastSeen);
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
