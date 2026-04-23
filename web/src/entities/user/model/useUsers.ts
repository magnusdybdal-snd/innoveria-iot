import { useCallback, useEffect, useReducer, useState } from "react";

import { getUsers } from "@entities/user/api";
import type { UserApiResponse } from "@entities/user/model/userSchema";

export interface UseUsersResult {
  users: UserApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

type State = {
  users: UserApiResponse[];
  isLoading: boolean;
  error: Error | null;
};

type Action =
  | { type: "FETCH_START" }
  | { type: "FETCH_SUCCESS"; users: UserApiResponse[] }
  | { type: "FETCH_ERROR"; error: Error };

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case "FETCH_START":
      return { users: [], isLoading: true, error: null };
    case "FETCH_SUCCESS":
      return { users: action.users, isLoading: false, error: null };
    case "FETCH_ERROR":
      return { ...state, isLoading: false, error: action.error };
  }
}

/**
 * Fetches users and exposes loading/error state.
 * @param companyId - If provided, only users belonging to this company are returned
 * @returns Users array, loading flag, error state, and a stable refetch callback
 */
export function useUsers(companyId?: string): UseUsersResult {
  const [state, dispatch] = useReducer(reducer, {
    users: [],
    isLoading: true,
    error: null,
  });
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => setRefetchIndex((i) => i + 1), []);

  useEffect(() => {
    let cancelled = false;
    dispatch({ type: "FETCH_START" });
    getUsers(companyId)
      .then((data) => {
        if (!cancelled) dispatch({ type: "FETCH_SUCCESS", users: data });
      })
      .catch((err: unknown) => {
        if (!cancelled)
          dispatch({
            type: "FETCH_ERROR",
            error: err instanceof Error ? err : new Error(String(err)),
          });
      });
    return () => {
      cancelled = true;
    };
  }, [refetchIndex, companyId]);

  return { ...state, refetch };
}
