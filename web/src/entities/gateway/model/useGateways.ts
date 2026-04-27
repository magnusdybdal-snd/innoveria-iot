import { useCallback, useEffect, useState } from "react";

import { getGateways } from "@entities/gateway/api";
import type { GatewayApiResponse } from "@entities/gateway/model/gatewaySchema";

export interface UseGatewaysResult {
  gateways: GatewayApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches all gateways and exposes loading/error state.
 * @returns Gateways array, loading flag, error state, and a stable refetch callback.
 */
export function useGateways(): UseGatewaysResult {
  const [gateways, setGateways] = useState<GatewayApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getGateways()
      .then((data) => {
        if (!cancelled) {
          setGateways(data);
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

  return { gateways, isLoading, error, refetch };
}
