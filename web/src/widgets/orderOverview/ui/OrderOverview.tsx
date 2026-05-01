import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

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
  const totalQuantity = operations
    .flatMap((op) => op.operation.reports ?? [])
    .reduce((sum, r) => sum + r.quantity, 0);

  const totalWh = operations
    .flatMap((op) => op.sensors)
    .map((s) => s.totalPowerWh)
    .filter((wh): wh is number => wh !== null)
    .reduce((sum, wh) => sum + wh, 0);
  const hasEnergyData = operations.some((op) =>
    op.sensors.some((s) => s.totalPowerWh !== null),
  );
  const totalKwh = hasEnergyData ? (totalWh / 1000).toFixed(2) : "—";
  const perPartKwh =
    hasEnergyData && totalQuantity > 0
      ? (totalWh / 1000 / totalQuantity).toFixed(2)
      : "—";

  return (
    <>
      <Typography variant="h5" fontWeight={500} sx={{ mb: 2 }}>
        Order overview
      </Typography>
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
        <OrderOverviewStat
          label="Parts produced"
          value={String(totalQuantity)}
        />
        <OrderOverviewStat
          label="Total energy"
          value={totalKwh}
          unit="kWh"
          tooltip={
            !hasEnergyData
              ? "No electricity sensors with a configured voltage found on this order"
              : undefined
          }
        />
        <OrderOverviewStat
          label="Per part"
          value={perPartKwh}
          unit="kWh/part"
          tooltip={
            !hasEnergyData
              ? "No electricity sensors with a configured voltage found on this order"
              : totalQuantity === 0
                ? "No parts reported for this order yet"
                : undefined
          }
        />
      </Box>
    </>
  );
}
