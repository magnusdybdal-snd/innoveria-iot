import type { BucketUnit } from "@entities/context";

export const CHART_TYPE = {
  bar: "bar",
  line: "line",
} as const;

export type ChartType = (typeof CHART_TYPE)[keyof typeof CHART_TYPE];

export interface GraphWidgetConfig {
  title: string;
  deviceEui: string;
  ruleId: string;
  from: string; // datetime-local string e.g. "2026-03-30T10:00"
  to: string;
  useCurrentTime: boolean; // when true, "to" is resolved to now() at fetch time
  bucketValue: string; // controlled string from number TextField
  bucketUnit: BucketUnit;
}
