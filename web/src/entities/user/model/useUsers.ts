import { useCallback, useEffect, useState } from "react";

import { getUsers } from "@entities/user/api";
import type { UserApiResponse } from "@entities/user/model/userSchema";

export interface UseUsersResult {
  users: UserApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches users and exposes loading/error state.
 * @param companyId - If provided, only users belonging to this company are returned
 * @returns Users array, loading flag, error state, and a stable refetch callback
 */
export function useUsers(companyId?: string): UseUsersResult {
  const [users, setUsers] = useState<UserApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getUsers(companyId)
      .then((data) => {
        if (!cancelled) {
          setUsers(data);
          setIsLoading(false);
        }
      })
      .catch((err: unknown) => {
        if (!cancelled) {
          setError(err instanceof Error ? err : new Error(String(err)));
          setIsLoading(false);
        }
      });
    return () => {
      cancelled = true;
    };
  }, [refetchIndex, companyId]);

  return { users, isLoading, error, refetch };
}
