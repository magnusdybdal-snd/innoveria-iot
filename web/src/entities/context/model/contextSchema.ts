export type BucketUnit = "minutes" | "hours" | "days" | "weeks" | "months";

export interface OrderReporting {
  reportingId: string;
  timestamp: string;
  quantity: number;
  status: string;
}

export interface Order {
  orderId: string;
  productName: string;
  status: string;
  startTime: string;
  endTime: string;
  workcenterId: string;
  workcenterName: string;
  reportings: OrderReporting[];
}

export interface CreateRuleRequest {
  companyId: string;
  name: string;
  contextType: string;
  measurementType: string;
  aggregationMethod: string;
  timeBucketMinutes: number;
  isActive: boolean;
}

export interface AggregationRule {
  id: string;
  companyId: string;
  name: string;
  contextType: string;
  measurementType: string;
  aggregationMethod: string;
  timeBucketMinutes: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

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
