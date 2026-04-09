import type { MeasureTypeApiResponse } from "@entities/measureType/model/measureTypeSchema.ts";

export type SortDirection = "asc" | "desc";
export type MeasureTypeSortKey =
  | "Default unit"
  | "Description"
  | "Display name"
  | "Slug";

// Sorting measureTypes based on Key (MeasureTypeSortKey)
/**
 * Returns a sorted copy of the measure type array based on the given column and direction.
 * @param measureTypes - Array of MeasureTypeApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted MeasureTypeApiResponse array (does not mutate the input)
 */
export function sortMeasureTypes(
  measureTypes: MeasureTypeApiResponse[],
  key: MeasureTypeSortKey | null,
  direction: SortDirection,
): MeasureTypeApiResponse[] {
  if (!key) return measureTypes;
  return [...measureTypes].sort((a, b) => {
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
