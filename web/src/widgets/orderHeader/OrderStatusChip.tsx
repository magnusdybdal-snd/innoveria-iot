import Chip from "@mui/material/Chip";

import { ORDER_STATUS_COLOR } from "@entities/context/model/statusColors";
import { formatStatus } from "@shared/lib";

/** Props for the `OrderStatusChip` component. */
interface OrderStatusChipProps {
  /** ERP order status string, e.g. `"started"` or `"finished"`. */
  status: string;
}

/**
 * MUI Chip that displays a formatted ERP order status with the appropriate color.
 * @param props - Component props
 * @param props.status - The raw ERP order status string
 * @returns A colored Chip with a human-readable status label
 */
export function OrderStatusChip({ status }: OrderStatusChipProps) {
  return (
    <Chip
      label={formatStatus(status)}
      color={ORDER_STATUS_COLOR[status] ?? "default"}
    />
  );
}
