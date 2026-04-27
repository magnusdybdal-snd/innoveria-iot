import { useCallback, useEffect, useState } from "react";

import { fetchSensorReading } from "@entities/sensor/api";
import type { SensorReadingApiResponse } from "@entities/sensor/model/sensorSchema";

export interface UseSensorReadingResult {
  reading: SensorReadingApiResponse | null;
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches the latest sensor reading for a given device EUI and exposes loading/error state.
 * @param deviceEUI - LoRaWAN Device EUI used to identify the sensor
 * @returns Reading, loading flag, error state, and a stable refetch callback.
 */
export function useSensorReading(deviceEUI: string): UseSensorReadingResult {
  const [reading, setReading] = useState<SensorReadingApiResponse | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [refetchIndex, setRefetchIndex] = useState(0);

  const refetch = useCallback(() => {
    setIsLoading(true);
    setRefetchIndex((i) => i + 1);
  }, []);

  useEffect(() => {
    let cancelled = false;
    fetchSensorReading(deviceEUI)
      .then((data) => {
        if (!cancelled) {
          setReading(data);
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
  }, [deviceEUI, refetchIndex]);

  return { reading, isLoading, error, refetch };
}
