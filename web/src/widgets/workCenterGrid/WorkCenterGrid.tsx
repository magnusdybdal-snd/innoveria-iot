import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import type { OperationContext } from "@entities/context/model/contextSchema";
import { MachineEnergyCard } from "@entities/context/ui/MachineEnergyCard";

/** Props for the `WorkCenterGrid` component. */
interface WorkCenterGridProps {
  /** List of operation contexts to render as machine energy cards. */
  operations: OperationContext[];
}

/**
 * Renders the work centers section for an order: a heading and a responsive
 * card grid where each card represents one operation context.
 * @param props - Component props
 * @param props.operations - Operation contexts to display
 * @returns The rendered work center grid, or null when the list is empty
 */
export function WorkCenterGrid({ operations }: WorkCenterGridProps) {
  if (operations.length === 0) return null;

  return (
    <Box>
      <Typography variant="h6" sx={{ mb: 2 }}>
        Work Centers
      </Typography>
      <Box sx={{ display: "flex", flexWrap: "wrap", gap: 2 }}>
        {operations.map((opCtx) => (
          <MachineEnergyCard
            key={opCtx.operation.id}
            operationContext={opCtx}
          />
        ))}
      </Box>
    </Box>
  );
}
