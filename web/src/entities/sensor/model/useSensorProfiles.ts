import { useCallback, useEffect, useState } from "react";

import { getSensorProfiles } from "@entities/sensor/api";
import type { SensorProfileApiResponse } from "@entities/sensor/model/sensorSchema";

export interface UseSensorProfilesResult {
  sensorProfiles: SensorProfileApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches all sensor profiles and exposes loading/error state.
 * @returns Sensor profiles array, loading flag, error state, and a stable refetch callback.
 */
export function useSensorProfiles(): UseSensorProfilesResult {
  const [sensorProfiles, setSensorProfiles] = useState<
    SensorProfileApiResponse[]
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
    getSensorProfiles()
      .then((data) => {
        if (!cancelled) {
          setSensorProfiles(data);
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

  return { sensorProfiles, isLoading, error, refetch };
}
