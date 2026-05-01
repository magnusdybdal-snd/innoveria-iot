import type { SensorApiResponse } from "@entities/sensor/model/sensorSchema";

export type SortDirection = "asc" | "desc";
export type SensorSortKey =
  | "Name"
  | "Status"
  | "Factory"
  | "Machine"
  | "Last reading"
  | "Sensor profile";

// Sorting sensors based on Key (SensorSortKey)
/**
 * Returns a sorted copy of the sensor array based on the given column and direction.
 * @param sensors - Array of SensorApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted SensorApiResponse array (does not mutate the input)
 */
export function sortSensors(
  sensors: SensorApiResponse[],
  key: SensorSortKey | null,
  direction: SortDirection,
): SensorApiResponse[] {
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
        cmp = (a.productionResource ?? 0) - (b.productionResource ?? 0);
        break;
      case "Last reading":
        cmp = a.lastReading.localeCompare(b.lastReading);
        break;
      case "Sensor profile":
        cmp = a.sensorProfileId.localeCompare(b.sensorProfileId);
        break;
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
