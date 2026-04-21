import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Chip from "@mui/material/Chip";
import Typography from "@mui/material/Typography";

import type { Order } from "@entities/context/model/contextSchema";
import { formatStatus, formatTimestamp } from "@shared/lib";

/** Maps ERP order status strings to MUI Chip color variants. */
const STATUS_COLOR: Record<string, "default" | "warning" | "info" | "success"> =
  {
    pending: "warning",
    in_progress: "info",
    completed: "success",
  };

/** Props for the `OrderHeader` component. */
interface OrderHeaderProps {
  /** The order whose metadata is displayed in the header. */
  order: Order;
}

/**
 * Card displaying the key metadata for an ERP order: order number, part description,
 * status badge, and planned start/finish dates.
 * @param props - Component props
 * @param props.order - The order to render
 * @returns The rendered order header card
 */
export function OrderHeader({ order }: OrderHeaderProps) {
  return (
    <Card
      sx={{
        backgroundColor: "secondary.light",
        borderRadius: 3,
        p: 2,
        color: "primary.main",
        mt: 2,
        mb: 3,
      }}
    >
      <Box
        sx={{
          display: "flex",
          alignItems: "flex-start",
          justifyContent: "space-between",
          flexWrap: "wrap",
          gap: 1,
        }}
      >
        <Box>
          <Typography variant="h5" fontWeight={600}>
            {order.orderNumber}
          </Typography>
          <Typography variant="body1" sx={{ opacity: 0.7, mt: 0.5 }}>
            {order.partDescription}
          </Typography>
        </Box>
        <Chip
          label={formatStatus(order.status)}
          color={STATUS_COLOR[order.status] ?? "default"}
        />
      </Box>

      <Box sx={{ display: "flex", gap: 4, mt: 2, flexWrap: "wrap" }}>
        <Box>
          <Typography variant="caption" sx={{ opacity: 0.6 }}>
            Planned start
          </Typography>
          <Typography variant="body2">
            {formatTimestamp(order.plannedStartDate)}
          </Typography>
        </Box>
        <Box>
          <Typography variant="caption" sx={{ opacity: 0.6 }}>
            Planned finish
          </Typography>
          <Typography variant="body2">
            {formatTimestamp(order.plannedFinishDate)}
          </Typography>
        </Box>
      </Box>
    </Card>
  );
}
