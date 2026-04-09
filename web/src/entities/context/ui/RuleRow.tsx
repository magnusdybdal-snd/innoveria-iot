import type { AggregationRule } from "@entities/context/model/contextSchema";
import Typography from "@mui/material/Typography";

interface RuleRowProps {
  rule: AggregationRule;
  onDelete?: (id: string) => void;
}

/**
 * Displays a single aggregation rule's fields as a row in the rules list.
 * @param props - Component props.
 * @param props.rule - The aggregation rule to display.
 * @returns The rendered rule row.
 */
export function RuleRow({ rule }: RuleRowProps) {
  return (
    <>
      <Typography>{rule.name}</Typography>
      <Typography>{rule.measurementType}</Typography>
      <Typography>{rule.aggregationMethod}</Typography>
      <Typography>{rule.isActive ? "Yes" : "No"}</Typography>
    </>
  );
}
