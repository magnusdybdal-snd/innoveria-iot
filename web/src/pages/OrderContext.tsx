import { useState } from "react";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import { DropDownSelect } from "@shared/ui/DropDownSelect";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

// TODO: replace with data from GET /api/v1/context/orders once the hook is wired up
const ORDER_OPTIONS: { id: string; name: string }[] = [];

/**
 * OrderContext page for displaying ERP order data enriched with sensor context.
 * @returns The rendered OrderContext page
 */
export default function OrderContext() {
  const [selectedOrderId, setSelectedOrderId] = useState<string>("");

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
            <DropDownSelect
              options={ORDER_OPTIONS}
              value={selectedOrderId}
              onChange={setSelectedOrderId}
            />
          </Box>
        </Box>
      </PageContent>
    </div>
  );
}
