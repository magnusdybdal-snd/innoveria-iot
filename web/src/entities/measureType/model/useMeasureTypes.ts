import { useEffect, useState } from "react";

import { getMeasureTypes } from "@entities/measureType";
import type { MeasureTypeApiResponse } from "@entities/measureType/model/measureTypeSchema.ts";

export interface UseMeasureTypesResult {
  measureTypes: MeasureTypeApiResponse[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
}

/**
 * Fetches and manages the list of measure types from the collection-service.
 * @returns measure types array, loading state, any fetch error, and a refetch function
 */
export function useMeasureTypes(): UseMeasureTypesResult {
  const [measureTypes, setMeasureTypes] = useState<MeasureTypeApiResponse[]>(
    [],
  );
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  const fetchMeasureTypes = () => {
    getMeasureTypes()
      .then(setMeasureTypes)
      .catch((err: unknown) => {
        setError(err instanceof Error ? err : new Error(String(err)));
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    fetchMeasureTypes();
  }, []);

  return { measureTypes, isLoading, error, refetch: fetchMeasureTypes };
}
