import { useCallback, useEffect, useState } from "react";

import { getRules } from "@entities/context/api";
import type { AggregationRule } from "@entities/context/model/contextSchema";

export interface UseRulesResult {
  rules: AggregationRule[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches aggregation rules for the given company and exposes loading/error state.
 * @param companyId - UUID of the company to scope the query.
 * @returns Rules array, loading flag, error state, and a stable refetch callback.
 */
export function useRules(companyId: string): UseRulesResult {
  const [rules, setRules] = useState<AggregationRule[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getRules(companyId)
      .then((data) => {
        if (!cancelled) {
          setRules(data);
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
  }, [companyId, refetchIndex]);

  return { rules, isLoading, error, refetch };
}
