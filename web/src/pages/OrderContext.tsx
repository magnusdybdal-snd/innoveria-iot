import { useState } from "react";

import { InfoWidget } from "@/widgets/infoWidget";
import Box from "@mui/material/Box";
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
  const { orders, isLoading } = useOrders();
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
            {!isLoading && (
              <DropDownSelect
                options={orderOptions}
                value={selectedOrderId}
                onChange={setSelectedOrderId}
              />
            )}
          </Box>
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
        <InfoWidget label="Test-label" value="25" unit="ppm" />
      </PageContent>
    </div>
  );
}
