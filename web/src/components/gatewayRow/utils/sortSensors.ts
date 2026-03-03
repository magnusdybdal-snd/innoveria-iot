import type { Sensor } from "@/mocks/sensors";

export type SortDirection = "asc" | "desc";
export type SensorSortKey =
  | "Name"
  | "Status"
  | "Machine"
  | "Last reading"
  | "Application key"
  | "Sensor profile";

// Sorting sensors based on Key (SensorSortkeys)
/**
 *
 * @param sensors
 * @param key
 * @param direction
 */
export function sortSensors(
  sensors: Sensor[],
  key: SensorSortKey | null,
  direction: SortDirection,
): Sensor[] {
  if (!key) return sensors;
  return [...sensors].sort((a, b) => {
    let cmp = 0;
    switch (key) {
      case "Name":
        cmp = a.name.localeCompare(b.name);
        break;
      case "Status":
        cmp = a.status - b.status;
        break;
      case "Machine":
        cmp = a.machine.localeCompare(b.machine);
        break;
      case "Last reading":
        cmp = a.lastReading.localeCompare(b.lastReading);
        break;
      case "Application key":
        cmp = a.appKey.localeCompare(b.appKey);
        break;
      case "Sensor profile":
        cmp = a.senProf.localeCompare(b.senProf);
        break;
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
