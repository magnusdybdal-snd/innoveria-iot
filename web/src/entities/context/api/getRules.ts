import type { AggregationRule } from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawAggregationRule = {
  id: string;
  company_id: string;
  name: string;
  context_type: string;
  measurement_type: string;
  aggregation_method: string;
  time_bucket_minutes: number;
  is_active: boolean;
  created_at: string;
  updated_at: string;
};

/**
 * Fetches all aggregation rules for a given company from the context service.
 * @param companyId - The company UUID to scope the rules query
 * @returns Array of aggregation rules belonging to the company, or an empty array on failure
 */
export const getRules = async (
  companyId: string,
): Promise<AggregationRule[]> => {
  try {
    const data = await apiRequest<RawAggregationRule[]>(
      serviceClient,
      `${API_ROUTES.contextRules}?company_id=${companyId}`,
      "GET",
    );

    return data.map((r) => ({
      id: r.id,
      companyId: r.company_id,
      name: r.name,
      contextType: r.context_type,
      measurementType: r.measurement_type,
      aggregationMethod: r.aggregation_method,
      timeBucketMinutes: r.time_bucket_minutes,
      isActive: r.is_active,
      createdAt: r.created_at,
      updatedAt: r.updated_at,
    }));
  } catch (error) {
    console.error("Failed to fetch aggregation rules:", error);
    return [];
  }
};
