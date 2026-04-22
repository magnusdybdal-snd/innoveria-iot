import Box from "@mui/material/Box";

import type { OperationContext } from "@entities/context/model/contextSchema";

import { OrderOverviewStat } from "./OrderOverviewStat";

/** Props for the `OrderOverview` component. */
interface OrderOverviewProps {
  /** Operation contexts used to derive all summary stats. */
  operations: OperationContext[];
}

/**
 * Horizontal row of stat tiles summarising key metrics for an order.
 * Shown between the order header and the machine cards.
 * Power and per-part stats are placeholders pending backend aggregation support.
 * @param props - Component props
 * @param props.operations - The operation contexts for the order
 * @returns The rendered overview stat row
 */
export function OrderOverview({ operations }: OrderOverviewProps) {
  const totalOperations = operations.length;
  const machinesWithSensors = operations.filter(
    (op) => op.sensors.length > 0,
  ).length;
  const degradedCount = operations.filter((op) => op.degraded).length;
  const totalSensors = operations.reduce(
    (sum, op) => sum + op.sensors.length,
    0,
  );

  const allMachinesHaveSensors = machinesWithSensors === totalOperations;

  return (
    <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap", mb: 3 }}>
      <OrderOverviewStat label="Operations" value={String(totalOperations)} />
      <OrderOverviewStat
        label="Machines with sensors"
        value={`${machinesWithSensors}/${totalOperations}`}
        warningTooltip={
          !allMachinesHaveSensors
            ? "Not all machines on this order have sensors mapped"
            : undefined
        }
      />
      <OrderOverviewStat
        label="Degraded sensors"
        value={String(degradedCount)}
        valueColor={degradedCount > 0 ? "warning.main" : undefined}
      />
      <OrderOverviewStat label="Total sensors" value={String(totalSensors)} />
      {/* Power totals placeholder — TODO: wire up once backend aggregation is available */}
      <OrderOverviewStat
        label="Total power"
        value="—"
        unit="kWh"
        tooltip="Power totals will be available in a future update"
      />
      {/* Per-part consumption placeholder — TODO: wire up once context-service exposes parts produced */}
      <OrderOverviewStat
        label="Per part"
        value="—"
        unit="kWh/part"
        tooltip="Per-part consumption will be available in a future update"
      />
    </Box>
  );
}
