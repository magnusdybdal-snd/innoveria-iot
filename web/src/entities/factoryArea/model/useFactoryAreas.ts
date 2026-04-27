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
 * Fetches all factory areas and exposes loading/error state.
 * @returns Factory areas array, loading flag, error state, and a stable refetch callback.
 */
export function useFactoryAreas(): UseFactoryAreasResult {
  const [factoryAreas, setFactoryAreas] = useState<FactoryAreaApiResponse[]>(
    [],
  );
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getFactoryAreas()
      .then((data) => {
        if (!cancelled) {
          setFactoryAreas(data);
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

  return { factoryAreas, isLoading, error, refetch };
}
