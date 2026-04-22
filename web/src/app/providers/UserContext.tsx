import { createContext } from "react";

import type { UserApiResponse } from "@entities/user";

export type UserContextValue = {
  user: UserApiResponse | null;
  isLoading: boolean;
};

export const UserContext = createContext<UserContextValue>({
  user: null,
  isLoading: true,
});
