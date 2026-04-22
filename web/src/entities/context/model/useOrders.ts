import { useCallback, useEffect, useState } from "react";

import { getOrders } from "@entities/context/api";
import type { Order } from "@entities/context/model/contextSchema";

/** Return type of the `useOrders` hook. */
export interface UseOrdersResult {
  /** The fetched orders, empty until the request resolves. */
  orders: Order[];
  /** True while the request is in-flight. */
  isLoading: boolean;
  /** Set to an `Error` if the request failed, otherwise `null`. */
  error: Error | null;
  /** Triggers a fresh fetch of the orders list. */
  refetch: () => void;
}

/**
 * Fetches the ERP orders list from the context service and exposes loading/error state.
 * @returns Orders array, loading flag, error state, and a stable refetch callback.
 */
export function useOrders(): UseOrdersResult {
  const [orders, setOrders] = useState<Order[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setError(null);
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getOrders()
      .then((data) => {
        if (!cancelled) {
          setOrders(data);
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

  return { orders, isLoading, error, refetch };
}
