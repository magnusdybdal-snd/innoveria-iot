import { useState } from "react";

import Typography from "@mui/material/Typography";

import type { AggregationRule } from "@entities/context/model/contextSchema";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

interface RuleRowProps {
  rule: AggregationRule;
  onDelete?: (id: string) => Promise<void>;
}

/**
 * Displays a single aggregation rule's fields as a row in the rules list.
 * Includes an action menu with a delete option; delegates the actual delete
 * API call and snackbar feedback to the parent via {@link RuleRowProps.onDelete}.
 * @param props - Component props.
 * @param props.rule - The aggregation rule to display.
 * @param props.onDelete - Called with the rule ID when the user confirms deletion.
 * @returns The rendered rule row.
 */
export function RuleRow({ rule, onDelete }: RuleRowProps) {
  const [deleteOpen, setDeleteOpen] = useState(false);

  const handleDeleteConfirm = async () => {
    await onDelete?.(rule.id);
    setDeleteOpen(false);
  };

  const menuItems = [
    {
      label: "Delete",
      onClick: () => {
        (document.activeElement as HTMLElement)?.blur();
        setDeleteOpen(true);
      },
    },
  ];

  return (
    <>
      <Typography>{rule.name}</Typography>
      <Typography>{rule.measurementType}</Typography>
      <Typography>{rule.aggregationMethod}</Typography>
      <Typography>{rule.isActive ? "Yes" : "No"}</Typography>
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
    </>
  );
}
