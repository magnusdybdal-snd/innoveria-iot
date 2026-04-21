import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import { formatTimestamp } from "@shared/lib";

/** Props for the `OrderDates` component. */
interface OrderDatesProps {
  /** Label shown above the planned date column. */
  plannedLabel: string;
  /** ISO timestamp string for the planned date. */
  plannedDate: string;
  /** Label shown above the actual date column. */
  actualLabel: string;
  /** ISO timestamp string for the actual date, or null/undefined if not yet set. */
  actualDate?: string | null;
}

/**
 * Renders a pair of labelled date fields side by side: one planned, one actual.
 * Displays an em-dash when the actual date is not yet available.
 * @param props - Component props
 * @param props.plannedLabel - Column header for the planned date
 * @param props.plannedDate - ISO timestamp for the planned date
 * @param props.actualLabel - Column header for the actual date
 * @param props.actualDate - ISO timestamp for the actual date, or null/undefined
 * @returns A flex row with two labelled date values
 */
export function OrderDates({
  plannedLabel,
  plannedDate,
  actualLabel,
  actualDate,
}: OrderDatesProps) {
  return (
    <Box sx={{ display: "flex", gap: 4, flexWrap: "wrap" }}>
      <Box>
        <Typography variant="caption" sx={{ opacity: 0.6 }}>
          {plannedLabel}
        </Typography>
        <Typography variant="body2">{formatTimestamp(plannedDate)}</Typography>
      </Box>
      <Box>
        <Typography variant="caption" sx={{ opacity: 0.6 }}>
          {actualLabel}
        </Typography>
        <Typography variant="body2">
          {actualDate ? formatTimestamp(actualDate) : "—"}
        </Typography>
      </Box>
    </Box>
  );
}
