import { useEffect, useState } from "react";

import { getSensors } from "@entities/sensor/api";
import type { SensorApiResponse } from "@entities/sensor/model/sensorSchema";

export interface UseSensorsResult {
  sensors: SensorApiResponse[];
  isLoading: boolean;
  error: Error | null;
}

/**
 * Fetches and manages the list of sensors from the collection-service.
 * @returns sensors array, loading state, and any fetch error
 */
export function useSensors(): UseSensorsResult {
  const [sensors, setSensors] = useState<SensorApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    getSensors()
      .then(setSensors)
      .catch((err: unknown) => {
        setError(err instanceof Error ? err : new Error(String(err)));
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  return { sensors, isLoading, error };
}
