import { useCallback, useEffect, useState } from "react";

import { getCompanies } from "@entities/company/api";
import type { CompanyApiResponse } from "@entities/company/model/companySchema";

export interface UseCompaniesResult {
  companies: CompanyApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches all companies and exposes loading/error state.
 * @returns Companies array, loading flag, error state, and a stable refetch callback.
 */
export function useCompanies(): UseCompaniesResult {
  const [companies, setCompanies] = useState<CompanyApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getCompanies()
      .then((data) => {
        if (!cancelled) {
          setCompanies(data);
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

  return { companies, isLoading, error, refetch };
}
