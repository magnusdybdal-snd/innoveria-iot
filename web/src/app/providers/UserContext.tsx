import { createContext } from "react";

import type { CurrentUserApiResponse } from "@entities/user";

export type UserContextValue = {
  user: CurrentUserApiResponse | null;
  isLoading: boolean;
  error: Error | null;
};

export const UserContext = createContext<UserContextValue>({
  user: null,
  isLoading: true,
  error: null,
});
