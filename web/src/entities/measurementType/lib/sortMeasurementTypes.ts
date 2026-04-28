import type { MeasurementTypeApiResponse } from "@entities/measurementType/model/measurementTypeSchema.ts";

export type SortDirection = "asc" | "desc";
export type MeasurementTypeSortKey =
  | "Default unit"
  | "Description"
  | "Display name"
  | "Slug";

// Sorting measurementTypes based on Key (MeasurementTypeSortKey)
/**
 * Returns a sorted copy of the measure type array based on the given column and direction.
 * @param measurementTypes - Array of MeasurementTypeApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted MeasurementTypeApiResponse array (does not mutate the input)
 */
export function sortMeasurementTypes(
  measurementTypes: MeasurementTypeApiResponse[],
  key: MeasurementTypeSortKey | null,
  direction: SortDirection,
): MeasurementTypeApiResponse[] {
  if (!key) return measurementTypes;
  return [...measurementTypes].sort((a, b) => {
    let cmp = 0;
    switch (key) {
      case "Default unit":
        cmp = a.defaultUnit.localeCompare(b.defaultUnit);
        break;
      case "Description":
        cmp = a.description.localeCompare(b.description);
        break;
      case "Display name":
        cmp = a.displayName.localeCompare(b.displayName);
        break;
      case "Slug":
        cmp = a.slug.localeCompare(b.slug);
        break;
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
