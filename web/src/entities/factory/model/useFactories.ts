import { useCallback, useEffect, useState } from "react";

import { getFactories } from "@entities/factory/api/getFactory";
import type { FactoryApiResponse } from "@entities/factory/model/factorySchema";

export interface UseFactoriesResult {
  factories: FactoryApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches all factories and exposes loading/error state.
 * @returns Factories array, loading flag, error state, and a stable refetch callback.
 */
export function useFactories(): UseFactoriesResult {
  const [factories, setFactories] = useState<FactoryApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getFactories()
      .then((data) => {
        if (!cancelled) {
          setFactories(data);
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
  }, [refetchIndex]);

  return { factories, isLoading, error, refetch };
}
