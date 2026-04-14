import { useState } from "react";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import { useOrders } from "@entities/context";
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
    id: o.orderId,
    name: o.productName,
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
        </Box>
      </PageContent>
    </div>
  );
}
