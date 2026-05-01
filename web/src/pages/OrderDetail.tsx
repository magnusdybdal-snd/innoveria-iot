import { useState } from "react";

import Box from "@mui/material/Box";
import ToggleButton from "@mui/material/ToggleButton";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";
import Typography from "@mui/material/Typography";
import { useParams } from "react-router";

import { useOrderContext, type AggregationMethod } from "@entities/context";
import { LoadingIndicator } from "@shared/ui/LoadingIndicator";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { OrderHeader } from "@widgets/orderHeader";
import { OrderOverview } from "@widgets/orderOverview";
import { WorkCenterGrid } from "@widgets/workCenterGrid";

const AGGREGATION_OPTIONS: { value: AggregationMethod; label: string }[] = [
  { value: "latest", label: "Latest" },
  { value: "min", label: "Min" },
  { value: "max", label: "Max" },
  { value: "avg", label: "Avg" },
  { value: "sum", label: "Sum" },
];

/**
 * Page displaying a single ERP order with its operations and sensor data.
 * Reads the order ID from the URL parameter `:id`.
 * @returns The rendered order detail page
 */
export default function OrderDetail() {
  const { id } = useParams<{ id: string }>();
  const parsed = Number(id);
  const orderId = Number.isInteger(parsed) && parsed > 0 ? parsed : null;

  const { orderContext, isLoading, error } = useOrderContext(orderId);
  const [aggregationMethod, setAggregationMethod] =
    useState<AggregationMethod>("latest");

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader />

        {orderId === null && (
          <Typography variant="body2" sx={{ mt: 3, color: "error.main" }}>
            Invalid order ID. Please navigate back to the order list.
          </Typography>
        )}

        {orderId !== null && isLoading && (
          <LoadingIndicator message="Loading order…" />
        )}

        {orderId !== null && !isLoading && error && (
          <Typography variant="body2" sx={{ mt: 3, color: "error.main" }}>
            Failed to load order. Please try again.
          </Typography>
        )}

        {orderId !== null && !isLoading && !error && orderContext && (
          <>
            <OrderHeader order={orderContext.order} />
            <OrderOverview operations={orderContext.operations} />
            <Box
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "flex-end",
                mt: 2,
                mb: 1,
                px: 2,
              }}
            >
              <ToggleButtonGroup
                value={aggregationMethod}
                exclusive
                size="small"
                onChange={(_, v: AggregationMethod | null) => {
                  if (v !== null) setAggregationMethod(v);
                }}
              >
                {AGGREGATION_OPTIONS.map((opt) => (
                  <ToggleButton key={opt.value} value={opt.value}>
                    {opt.label}
                  </ToggleButton>
                ))}
              </ToggleButtonGroup>
            </Box>
            <WorkCenterGrid
              operations={orderContext.operations}
              aggregationMethod={aggregationMethod}
            />
          </>
        )}
      </PageContent>
    </div>
  );
}
