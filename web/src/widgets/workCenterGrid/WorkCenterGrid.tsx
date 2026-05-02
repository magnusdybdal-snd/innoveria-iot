import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import type { AggregationMethod } from "@entities/context";
import type { OperationContext } from "@entities/context/model/contextSchema";
import { MachineEnergyCard } from "@entities/context/ui/MachineEnergyCard";
import { useMeasurementTypes } from "@entities/measurementType";

/** Props for the `WorkCenterGrid` component. */
interface WorkCenterGridProps {
  /** List of operation contexts to render as machine energy cards. */
  operations: OperationContext[];
  /** How to aggregate sensor readings across the measurements window. Defaults to latest. */
  aggregationMethod?: AggregationMethod;
}

/**
 * Renders the machines (work centers) section for an order: a heading and a responsive
 * card grid where each card represents one operation context.
 * @param props - Component props
 * @param props.operations - Operation contexts to display
 * @param props.aggregationMethod - How to aggregate sensor readings. Defaults to "latest".
 * @returns The rendered machine grid, or null when the list is empty
 */
export function WorkCenterGrid({
  operations,
  aggregationMethod = "latest",
}: WorkCenterGridProps) {
  const { measurementTypes, error, refetch } = useMeasurementTypes();

  if (error) {
    return (
      <Box sx={{ mt: 3 }}>
        <Typography variant="body2" sx={{ color: "error.main", mb: 1 }}>
          Failed to load measurement types. {error.message}
        </Typography>
        <Button variant="outlined" size="small" onClick={refetch}>
          Retry
        </Button>
      </Box>
    );
  }

  if (operations.length === 0) {
    return (
      <Typography variant="body2" sx={{ mt: 3, opacity: 0.6 }}>
        No machines found for this order.
      </Typography>
    );
  }

  return (
    <Box>
      <Typography variant="h6" sx={{ mb: 2 }}>
        Machines
      </Typography>
      <Box sx={{ display: "flex", flexWrap: "wrap", gap: 2 }}>
        {operations.map((opCtx) => (
          <MachineEnergyCard
            key={opCtx.operation.id}
            operationContext={opCtx}
            measurementTypes={measurementTypes}
            aggregationMethod={aggregationMethod}
          />
        ))}
      </Box>
    </Box>
  );
}
