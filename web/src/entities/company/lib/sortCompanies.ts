import type { CompanyApiResponse } from "@entities/company";

export type SortDirection = "asc" | "desc";
export type CompanySortKey = "Name" | "Address" | "Created at" | "Updated at";

// Parses an RFC1123 date string (from the API) into a Unix timestamp for comparison.
// Returns 0 if the string is not a valid date.
function parseDateToMs(dateAt: string): number {
  const ms = Date.parse(dateAt);
  return isNaN(ms) ? 0 : ms;
}

// Sorting companys based on Key (CompanySortKey)
/**
 * Returns a sorted copy of the company array based on the given column and direction.
 * @param companies - Array of CompanyApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted CompanyApiResponse array (does not mutate the input)
 */
export function sortCompanies(
  companies: CompanyApiResponse[],
  key: CompanySortKey | null,
  direction: SortDirection,
): CompanyApiResponse[] {
  if (!key) return companies;
  return [...companies].sort((a, b) => {
    let cmp = 0;
    if (key === "Name") {
      cmp = a.name.localeCompare(b.name);
    } else if (key === "Address") {
      cmp = a.address.localeCompare(b.address);
    } else if (key === "Created at") {
      cmp = parseDateToMs(a.createdAt) - parseDateToMs(b.createdAt);
    } else if (key === "Updated at") {
      cmp = parseDateToMs(a.updatedAt) - parseDateToMs(b.updatedAt);
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
