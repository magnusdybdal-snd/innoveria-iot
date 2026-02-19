import type { Gateway } from "@/mocks/gateways";

export type SortDirection = "asc" | "desc";
export type GatewaySortKey = "Name" | "Status" | "Last seen";

// Tuple with units to seconds.
// Used to unify unit that is being used to sort in @sortGateway.
const unitToSeconds: Record<string, number> = {
  sec: 1,
  min: 60,
  hour: 3600,
  hours: 3600,
  day: 86400,
  days: 86400,
};

//
function parseLastSeenToSeconds(lastSeen: string): number {
  const [amount, unit] = lastSeen.split(" ");
  return (parseInt(amount, 10) || 0) * (unitToSeconds[unit] ?? 0);
}

// Sorting gateways based on Key (GatewaySortkeys)
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
      cmp = (a.online ? 1 : 0) - (b.online ? 1 : 0);
    } else if (key === "Last seen") {
      cmp =
        parseLastSeenToSeconds(a.lastSeen) - parseLastSeenToSeconds(b.lastSeen);
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
