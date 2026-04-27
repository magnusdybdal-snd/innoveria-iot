import { useEffect, useState } from "react";

import { getMeasurementTypes } from "@entities/measurementType";
import type { MeasurementTypeApiResponse } from "@entities/measurementType/model/measurementTypeSchema.ts";

export interface UseMeasurementTypesResult {
  measurementTypes: MeasurementTypeApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches and manages the list of measure types from the collection-service.
 * @returns measure types array, loading state, any fetch error, and a refetch function
 */
export function useMeasurementTypes(): UseMeasurementTypesResult {
  const [measurementTypes, setMeasurementTypes] = useState<
    MeasurementTypeApiResponse[]
  >([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  const fetchMeasurementTypes = () => {
    getMeasurementTypes()
      .then(setMeasurementTypes)
      .catch((err: unknown) => {
        setError(err instanceof Error ? err : new Error(String(err)));
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    fetchMeasurementTypes();
  }, []);

  return {
    measurementTypes: measurementTypes,
    isLoading,
    error,
    refetch: fetchMeasurementTypes,
  };
}
