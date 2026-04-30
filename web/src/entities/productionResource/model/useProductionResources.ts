import { useCallback, useEffect, useState } from "react";

import { getProductionResources } from "@entities/productionResource/api/getProductionResources";
import type { ProductionResourceApiResponse } from "@entities/productionResource/model/productionResourceSchema";

export interface UseProductionResourcesResult {
  productionResources: ProductionResourceApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches all ERP production resources and exposes loading/error state.
 * @returns Production resources array, loading flag, error state, and a stable refetch callback.
 */
export function useProductionResources(): UseProductionResourcesResult {
  const [productionResources, setProductionResources] = useState<
    ProductionResourceApiResponse[]
  >([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getProductionResources()
      .then((data) => {
        if (!cancelled) {
          setProductionResources(data);
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

  return { productionResources, isLoading, error, refetch };
}
