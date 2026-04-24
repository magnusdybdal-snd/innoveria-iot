import { useCallback, useEffect, useState } from "react";

import { getFactoryAreas } from "@entities/factoryArea/api/getFactoryAreas";
import type { FactoryAreaApiResponse } from "@entities/factoryArea/model/factoryAreaSchema";

export interface UseFactoryAreasResult {
  factoryAreas: FactoryAreaApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches factory areas for the given factory and exposes loading/error state.
 * Skips the fetch when factoryId is empty.
 * @param factoryId - The factory to fetch areas for
 * @returns Factory areas array, loading flag, error state, and a stable refetch callback.
 */
export function useFactoryAreas(factoryId?: string): UseFactoryAreasResult {
  const [factoryAreas, setFactoryAreas] = useState<FactoryAreaApiResponse[]>(
    [],
  );
  const [fetchedFactoryId, setFetchedFactoryId] = useState<
    string | undefined
  >();
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setFetchedFactoryId(undefined);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    if (!factoryId) return;

    let cancelled = false;
    getFactoryAreas(factoryId)
      .then((data) => {
        if (!cancelled) {
          setFactoryAreas(data);
          setFetchedFactoryId(factoryId);
        }
      })
      .catch((err: unknown) => {
        if (!cancelled) {
          setError(err instanceof Error ? err : new Error(String(err)));
          setFetchedFactoryId(factoryId);
        }
      });
    return () => {
      cancelled = true;
    };
  }, [factoryId, refetchIndex]);

  const isLoading = !!factoryId && fetchedFactoryId !== factoryId;

  return {
    factoryAreas:
      factoryId && fetchedFactoryId === factoryId ? factoryAreas : [],
    isLoading,
    error,
    refetch,
  };
}
