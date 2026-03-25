export type BucketUnit = "minutes" | "hours" | "days" | "weeks" | "months";

export interface ContextQueryParams {
  companyId: string;
  deviceEui: string;
  ruleId: string;
  from: string;
  to: string;
  bucketMinutes?: number;
}

export interface BucketResponse {
  periodStart: string;
  periodEnd: string;
  value: number;
}

export interface ContextDataResponse {
  deviceEui: string;
  companyId: string;
  contextType: string;
  unit: string;
  periodStart: string;
  periodEnd: string;
  totalValue: number;
  buckets: BucketResponse[];
  calculatedAt: string;
}
