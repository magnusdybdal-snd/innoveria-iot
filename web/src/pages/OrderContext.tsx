import { useState } from "react";

import { InfoWidget } from "@/widgets/infoWidget";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import {
  MachineCard,
  useOrderContext,
  useOrders,
  type OrderSummary,
} from "@entities/context";
import { DropDownSelect } from "@shared/ui/DropDownSelect";
import { LoadingIndicator } from "@shared/ui/LoadingIndicator";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

/**
 * OrderContext page for displaying ERP order data enriched with sensor context.
 * @returns The rendered OrderContext page
 */
export default function OrderContext() {
  const { orders, isLoading } = useOrders();
  const [selectedOrderId, setSelectedOrderId] = useState<string>("");

  const selectedOrder: OrderSummary | undefined = orders.find(
    (o) => String(o.id) === selectedOrderId,
  );

  const {
    orderContext,
    isLoading: isContextLoading,
    error: contextError,
  } = useOrderContext(selectedOrder ? selectedOrder.id : null);

  const orderOptions = orders.map((o) => ({
    id: String(o.id),
    name: o.name,
  }));

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Order Context" />
        <PageDivider />
        <Box sx={{ mt: 4, ml: 4, mr: 4, mb: 2 }}>
          <Typography variant="body2" sx={{ mb: 0.5, color: "primary.main" }}>
            Orders
          </Typography>
          <Box sx={{ maxWidth: 400 }}>
            {!isLoading && (
              <DropDownSelect
                options={orderOptions}
                value={selectedOrderId}
                onChange={setSelectedOrderId}
              />
            )}
          </Box>

          {selectedOrder && (
            <>
              <PageDivider />
              <Typography
                variant="h3"
                sx={{ color: "primary.main", fontWeight: 500 }}
              >
                Machines
              </Typography>
              <Typography
                variant="subtitle1"
                sx={{ mb: 2, color: "primary.main" }}
              >
                Work centers and their current status for the selected order
              </Typography>
              <Box sx={{ mt: 3, display: "flex", flexWrap: "wrap", gap: 2 }}>
                {orderContext?.order.operations.map((op) => (
                  <InfoWidget
                    key={op.id}
                    label={op.productionResource.number}
                    value={op.productionResourceStatus}
                    unit={op.productionResource.description ?? undefined}
                  />
                ))}
              </Box>

              <PageDivider />
              <Typography
                variant="h3"
                sx={{ color: "primary.main", fontWeight: 500 }}
              >
                Sensor Data
              </Typography>
              <Typography
                variant="subtitle1"
                sx={{ mb: 2, color: "primary.main" }}
              >
                Live sensor readings per work center
              </Typography>

              {isContextLoading ? (
                <LoadingIndicator message="Loading sensor data…" />
              ) : contextError ? (
                <Typography variant="body2" sx={{ mt: 3, color: "error.main" }}>
                  Failed to load sensor data. Please try again.
                </Typography>
              ) : (
                <Box sx={{ mt: 3, display: "flex", flexWrap: "wrap", gap: 2 }}>
                  {orderContext?.operations.map((opCtx) => (
                    <MachineCard
                      key={opCtx.operation.id}
                      operationContext={opCtx}
                    />
                  ))}
                </Box>
              )}
            </>
          )}
        </Box>
      </PageContent>
    </div>
  );
}
