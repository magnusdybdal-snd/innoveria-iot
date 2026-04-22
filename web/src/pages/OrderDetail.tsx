import Typography from "@mui/material/Typography";
import { useParams } from "react-router";

import { useOrderContext } from "@entities/context";
import { LoadingIndicator } from "@shared/ui/LoadingIndicator";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { OrderHeader } from "@widgets/orderHeader";
import { OrderOverview } from "@widgets/orderOverview";
import { WorkCenterGrid } from "@widgets/workCenterGrid";

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
            <OrderOverview operations={orderContext.operations} />
            <WorkCenterGrid operations={orderContext.operations} />
          </>
        )}
      </PageContent>
    </div>
  );
}
