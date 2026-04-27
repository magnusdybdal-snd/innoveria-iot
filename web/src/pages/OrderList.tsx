import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { useNavigate } from "react-router";

import { useOrders, type OrderSummary } from "@entities/context";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { LoadingIndicator } from "@shared/ui/LoadingIndicator";
import { NotFoundCard } from "@shared/ui/NotFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

/** Column labels for the orders table. */
const COLUMNS = ["Order", ""];

/**
 * Renders a single clickable row for an order summary.
 * @param props - Component props
 * @param props.order - The order summary to display
 * @param props.onClick - Callback invoked when the row is clicked
 * @returns The rendered order row
 */
function OrderRow({
  order,
  onClick,
}: {
  order: OrderSummary;
  onClick: () => void;
}) {
  return (
    <DeviceRow onClick={onClick}>
      <Typography variant="body1">{order.name}</Typography>
      <span />
    </DeviceRow>
  );
}

/**
 * Page listing all ERP orders. Each row navigates to the order detail view.
 * @returns The rendered order list page
 */
export default function OrderList() {
  const navigate = useNavigate();
  const { orders, isLoading, error, refetch } = useOrders();

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Orders" />

        {isLoading && <LoadingIndicator message="Loading orders…" />}

        {!isLoading && error && (
          <Box sx={{ mt: 3 }}>
            <Typography variant="body2" sx={{ color: "error.main", mb: 1 }}>
              Failed to load orders. Please try again.
            </Typography>
            <Button variant="outlined" size="small" onClick={refetch}>
              Retry
            </Button>
          </Box>
        )}

        {!isLoading && !error && orders.length === 0 && (
          <NotFoundCard page="orders" />
        )}

        {!isLoading && !error && orders.length > 0 && (
          <CategoryHeader categories={COLUMNS} columns={COLUMNS.length}>
            {orders.map((order) => (
              <OrderRow
                key={order.id}
                order={order}
                onClick={() =>
                  void navigate(`/Dashboard/Context/OrderList/${order.id}`)
                }
              />
            ))}
          </CategoryHeader>
        )}
      </PageContent>
    </div>
  );
}
