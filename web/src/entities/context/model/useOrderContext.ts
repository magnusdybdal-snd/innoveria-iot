import { useCallback, useEffect, useState } from "react";

import { getOrderContext } from "@entities/context/api/getOrderContext";
import type { OrderContext } from "@entities/context/model/contextSchema";

/** Return type of the `useOrderContext` hook. */
export interface UseOrderContextResult {
  /** The fetched order context, null until the request resolves or when orderId is null. */
  orderContext: OrderContext | null;
  /** True while a fetch is in-flight for the current orderId. */
  isLoading: boolean;
  /** Set to an `Error` if the request failed, otherwise null. */
  error: Error | null;
  /** Clears current data and triggers a fresh fetch. */
  refetch: () => void;
}

/**
 * Fetches sensor-enriched context for a single order.
 * `isLoading` is derived by comparing `orderId` against the id inside the
 * resolved `orderContext`, so no synchronous setState is needed in the effect.
 * @param orderId - Numeric order ID to fetch context for, or null to skip fetching.
 * @returns Order context data, loading flag, error state, and a stable refetch callback.
 */
export function useOrderContext(orderId: number | null): UseOrderContextResult {
  const [orderContext, setOrderContext] = useState<OrderContext | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const isLoading =
    orderId !== null && error === null && orderContext?.order.id !== orderId;

  // Clears stale data so isLoading becomes true for same-order refetches.
  const refetch = useCallback(() => {
    setOrderContext(null);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    if (orderId === null) return;

    let cancelled = false;
    getOrderContext(orderId)
      .then((data) => {
        if (!cancelled) setOrderContext(data);
      })
      .catch((err: unknown) => {
        if (!cancelled)
          setError(err instanceof Error ? err : new Error(String(err)));
      });
    return () => {
      cancelled = true;
    };
  }, [orderId, refetchIndex]);

  return { orderContext, isLoading, error, refetch };
}
