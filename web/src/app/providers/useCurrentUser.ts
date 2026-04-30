import { useContext } from "react";

import { UserContext, type UserContextValue } from "./UserContext";

/**
 * Returns the current logged-in user and loading state from UserContext.
 * @returns The current user and isLoading flag from UserContext
 */
export function useCurrentUser(): UserContextValue {
  return useContext(UserContext);
}
