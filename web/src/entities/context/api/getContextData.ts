import type { ContextDataResponse } from "@entities/context/model/contextSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawBucket = {
  period_start: string;
  period_end: string;
  value: number;
};

type RawContextDataResponse = {
  device_eui: string;
  company_id: string;
  context_type: string;
  unit: string;
  period_start: string;
  period_end: string;
  total_value: number;
  buckets: RawBucket[];
  calculated_at: string;
};

/**
 * Fetches computed context data for one or more devices over a time window.
 * @param companyId - The company UUID to scope the request
 * @param deviceEuis - One or more device EUIs to include in the query
 * @param ruleId - The aggregation rule UUID to apply
 * @param from - Period start as an RFC3339 timestamp string
 * @param to - Period end as an RFC3339 timestamp string
 * @param bucketMinutes - Optional bucket size in minutes; overrides the rule default
 * @returns Array of context data responses, one per device EUI
 */
export const getContextData = async (
  companyId: string,
  deviceEuis: string[],
  ruleId: string,
  from: string,
  to: string,
  bucketMinutes?: number,
): Promise<ContextDataResponse[]> => {
  const params = new URLSearchParams({
    company_id: companyId,
    rule_id: ruleId,
    from,
    to,
  });
  deviceEuis.forEach((eui) => params.append("device_eui", eui));
  if (bucketMinutes && bucketMinutes > 0) {
    params.set("bucket_minutes", String(bucketMinutes));
  }

  const data = await apiRequest<RawContextDataResponse[]>(
    serviceClient,
    `${API_ROUTES.contextData}?${params.toString()}`,
    "GET",
  );

  return data.map((item) => ({
    deviceEui: item.device_eui,
    companyId: item.company_id,
    contextType: item.context_type,
    unit: item.unit,
    periodStart: item.period_start,
    periodEnd: item.period_end,
    totalValue: item.total_value,
    buckets: item.buckets.map((b) => ({
      periodStart: b.period_start,
      periodEnd: b.period_end,
      value: b.value,
    })),
    calculatedAt: item.calculated_at,
  }));
};
