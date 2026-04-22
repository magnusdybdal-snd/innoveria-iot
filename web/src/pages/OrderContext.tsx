import { useState } from "react";

import { InfoWidget } from "@/widgets/infoWidget";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import { useOrders, type Order } from "@entities/context";
import { DropDownSelect } from "@shared/ui/DropDownSelect";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

/**
 * OrderContext page for displaying ERP order data enriched with sensor context.
 * @returns The rendered OrderContext page
 */
export default function OrderContext() {
  const { orders, isLoading, error, refetch } = useOrders();
  const [selectedOrderId, setSelectedOrderId] = useState<string>("");

  const orderOptions = orders.map((o) => ({
    id: String(o.id),
    name: `${o.orderNumber} — ${o.partDescription}`,
  }));

  const selectedOrder: Order | undefined = orders.find(
    (o) => String(o.id) === selectedOrderId,
  );

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
            {error ? (
              <Box>
                <Typography variant="body2" sx={{ color: "error.main", mb: 1 }}>
                  Failed to load orders. {error.message}
                </Typography>
                <Button variant="outlined" size="small" onClick={refetch}>
                  Retry
                </Button>
              </Box>
            ) : (
              !isLoading && (
                <DropDownSelect
                  options={orderOptions}
                  value={selectedOrderId}
                  onChange={setSelectedOrderId}
                />
              )
            )}
          </Box>

          {/* Displays the operations of the selected order as InfoWidgets */}
          {selectedOrder && (
            <>
              {" "}
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
                {selectedOrder.operations.map((op) => (
                  <InfoWidget
                    key={op.id}
                    label={op.productionResource.number} // TODO display workcenter name instead of number
                    value={op.productionResourceStatus} // TODO: display sensor value
                    unit={op.productionResource.description ?? undefined} // TODO: display sensor unit instead of workcenter description
                  />
                ))}
              </Box>
            </>
          )}
          <PageDivider />
          {/* DEBUG: Display selected order details in a formatted JSON block */}
          {selectedOrder && (
            <Box
              component="pre"
              sx={{
                mt: 3,
                p: 2,
                backgroundColor: "primary.dark",
                color: "primary.main",
                borderRadius: 1,
                fontSize: "0.8rem",
                overflowX: "auto",
              }}
            >
              {JSON.stringify(selectedOrder, null, 2)}
            </Box>
          )}
        </Box>
      </PageContent>
    </div>
  );
}
