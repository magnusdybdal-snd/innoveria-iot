import { useEffect, useState } from "react";

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
 * @returns sensors array, loading state, any fetch error, and a refetch function
 */
export function useSensors(): UseSensorsResult {
  const [sensors, setSensors] = useState<SensorApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  const fetchSensors = () => {
    getSensors()
      .then(setSensors)
      .catch((err: unknown) => {
        setError(err instanceof Error ? err : new Error(String(err)));
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    fetchSensors();
  }, []);

  return { sensors, isLoading, error, refetch: fetchSensors };
}
