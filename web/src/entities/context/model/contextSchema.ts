export type BucketUnit = "minutes" | "hours" | "days" | "weeks" | "months";

export interface ProductionResource {
  id: number;
  number: string;
  description: string | null;
  type: string;
}

export interface OrderOperation {
  id: number;
  productionResource: ProductionResource;
  plannedStartDate: string;
  plannedFinishDate: string;
  actualStartDate: string | null;
  actualFinishDate: string | null;
  status: string;
  productionResourceStatus: string;
}

export interface Order {
  id: number;
  orderNumber: string;
  partDescription: string;
  plannedStartDate: string;
  plannedFinishDate: string;
  actualStartDate: string | null;
  actualFinishDate: string | null;
  status: string;
  priority: number;
  operations: OrderOperation[];
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
