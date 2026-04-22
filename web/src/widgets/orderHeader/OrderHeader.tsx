import { PageDivider } from "@/shared/ui/PageDivider";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Typography from "@mui/material/Typography";

import type { Order } from "@entities/context/model/contextSchema";

import { OrderDates } from "./OrderDates";
import { OrderStatusChip } from "./OrderStatusChip";

/** Props for the `OrderHeader` component. */
interface OrderHeaderProps {
  /** The order whose metadata is displayed in the header. */
  order: Order;
}

/**
 * Card displaying the key metadata for an ERP order: order number, part description,
 * status badge, and planned/actual start and finish dates.
 * @param props - Component props
 * @param props.order - The order to render
 * @returns The rendered order header card
 */
export function OrderHeader({ order }: OrderHeaderProps) {
  return (
    <>
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
            mb: 2,
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
          <OrderStatusChip status={order.status} />
        </Box>

        <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
          <OrderDates
            plannedLabel="Planned start"
            plannedDate={order.plannedStartDate}
            actualLabel="Actual start"
            actualDate={order.actualStartDate}
          />
          <OrderDates
            plannedLabel="Planned finish"
            plannedDate={order.plannedFinishDate}
            actualLabel="Actual finish"
            actualDate={order.actualFinishDate}
          />
        </Box>
      </Card>
      <PageDivider />
    </>
  );
}
