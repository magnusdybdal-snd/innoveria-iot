import type { UserApiResponse } from "@entities/user/model/userSchema.ts";

export type SortDirection = "asc" | "desc";
export type UserSortKey = "Name" | "Email" | "Role" | "Created at";

/**
 * Returns a sorted copy of the user array based on the given column and direction.
 * @param users - Array of UserApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted UserApiResponse array (does not mutate the input)
 */
export function sortUsers(
  users: UserApiResponse[],
  key: UserSortKey | null,
  direction: SortDirection,
): UserApiResponse[] {
  if (!key) return users;
  return [...users].sort((a, b) => {
    let cmp = 0;
    if (key === "Name") {
      cmp = a.name.localeCompare(b.name);
    } else if (key === "Email") {
      cmp = a.email.localeCompare(b.email);
    } else if (key === "Role") {
      cmp = a.role.localeCompare(b.role);
    } else if (key === "Created at") {
      cmp = a.createdAt.localeCompare(b.createdAt);
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
