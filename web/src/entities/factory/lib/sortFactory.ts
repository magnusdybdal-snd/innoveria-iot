import type { FactoryApiResponse } from "@entities/factory/model/factorySchema";

export type SortDirection = "asc" | "desc";
export type FactorySortKey = "Name" | "Address" | "Created at" | "Updated at";

type SortComparator<T> = (a: T, b: T) => number;

/**
 * Generic sort utility that can be used for any entity.
 * Takes a comparators map of column label → comparator function,
 * picks the one matching the given key, and applies direction.
 * Returns the original array unsorted if key is null or has no comparator.
 * @param items - Array of objects to sort
 * @param key - Column label to sort by, or null to return unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @param comparators - Map of column label to comparator function for that entity
 * @returns A new sorted array (does not mutate the input)
 */
export function sortItems<T>(
  items: T[],
  key: string | null,
  direction: SortDirection,
  comparators: Partial<Record<string, SortComparator<T>>>,
): T[] {
  if (!key) return items;
  const comparator = comparators[key];
  if (!comparator) return items;
  return [...items].sort((a, b) => {
    const cmp = comparator(a, b);
    return direction === "asc" ? cmp : -cmp;
  });
}

// Comparators for each sortable factory column.
const factoryComparators: Record<
  FactorySortKey,
  SortComparator<FactoryApiResponse>
> = {
  Name: (a, b) => a.name.localeCompare(b.name),
  Address: (a, b) => a.address.localeCompare(b.address),
  "Created at": (a, b) =>
    new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime(),
  "Updated at": (a, b) =>
    new Date(a.updatedAt).getTime() - new Date(b.updatedAt).getTime(),
};

/**
 * Returns a sorted copy of the factory array based on the given column and direction.
 * @param factories - Array of FactoryApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted FactoryApiResponse array (does not mutate the input)
 */
export function sortFactories(
  factories: FactoryApiResponse[],
  key: FactorySortKey | null,
  direction: SortDirection,
): FactoryApiResponse[] {
  return sortItems(factories, key, direction, factoryComparators);
}
