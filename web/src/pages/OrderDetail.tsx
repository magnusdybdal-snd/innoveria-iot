import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { useParams } from "react-router";

import { useOrderContext } from "@entities/context";
import { MachineEnergyCard } from "@entities/context/ui/MachineEnergyCard";
import { LoadingIndicator } from "@shared/ui/LoadingIndicator";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { OrderHeader } from "@widgets/orderHeader";

/**
 * Page displaying a single ERP order with its operations and sensor data.
 * Reads the order ID from the URL parameter `:id`.
 * @returns The rendered order detail page
 */
export default function OrderDetail() {
  const { id } = useParams<{ id: string }>();
  const orderId = id ? Number(id) : null;

  const { orderContext, isLoading, error } = useOrderContext(orderId);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader />

        {isLoading && <LoadingIndicator message="Loading order…" />}

        {!isLoading && error && (
          <Typography variant="body2" sx={{ mt: 3, color: "error.main" }}>
            Failed to load order. Please try again.
          </Typography>
        )}

        {!isLoading && !error && orderContext && (
          <>
            <OrderHeader order={orderContext.order} />

            <Typography variant="h6" sx={{ mb: 2 }}>
              Work Centers
            </Typography>

            <Box sx={{ display: "flex", flexWrap: "wrap", gap: 2 }}>
              {orderContext.operations.map((opCtx) => (
                <MachineEnergyCard
                  key={opCtx.operation.id}
                  operationContext={opCtx}
                />
              ))}
            </Box>
          </>
        )}
      </PageContent>
    </div>
  );
}
