import { useCallback, useEffect, useState } from "react";

import { getSensors } from "@entities/sensor/api";
import type { SensorApiResponse } from "@entities/sensor/model/sensorSchema";

export interface UseSensorsResult {
  sensors: SensorApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches and manages the list of sensors from the collection-service.
 * @returns Sensors array, loading flag, error state, and a stable refetch callback.
 */
export function useSensors(): UseSensorsResult {
  const [sensors, setSensors] = useState<SensorApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    getSensors()
      .then((data) => {
        if (!cancelled) {
          setSensors(data);
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

  return { sensors, isLoading, error, refetch };
}
