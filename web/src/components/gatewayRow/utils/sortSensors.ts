import type { Sensor } from "@/mocks/sensors";

export type SortDirection = "asc" | "desc";
export type SensorSortKey =
  | "Name"
  | "Status"
  | "Machine"
  | "Last reading"
  | "Application key"
  | "Device profile";

// Sorting sensors based on Key (SensorSortkeys)
export function sortSensors(
  sensors: Sensor[],
  key: SensorSortKey | null,
  direction: SortDirection,
): Sensor[] {
  if (!key) return sensors;
  return [...sensors].sort((a, b) => {
    let cmp = 0;
    if (key === "Name") {
      cmp = a.name.localeCompare(b.name);
    } else if (key === "Status") {
      cmp = a.status - b.status;
    } else if (key === "Machine") {
      cmp = a.machine.localeCompare(b.machine);
    } else if (key === "Last reading") {
      cmp = a.lastReading.localeCompare(b.lastReading);
    } else if (key === "Application key") {
      cmp = a.appKey.localeCompare(b.appKey);
    } else if (key === "Device profile") {
      cmp = a.devProf.localeCompare(b.devProf);
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
