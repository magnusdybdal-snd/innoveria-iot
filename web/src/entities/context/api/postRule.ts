import type { CreateRuleRequest } from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Sends a POST request to create a new aggregation rule.
 * @param payload - The rule data to create.
 * @returns The ID of the newly created rule.
 */
export const postRule = async (payload: CreateRuleRequest): Promise<string> => {
  return apiRequest<string>(serviceClient, API_ROUTES.contextRules, "POST", {
    company_id: payload.companyId,
    name: payload.name,
    context_type: payload.contextType,
    measurement_type: payload.measurementType,
    aggregation_method: payload.aggregationMethod,
    time_bucket_minutes: payload.timeBucketMinutes,
    is_active: payload.isActive,
  });
};
