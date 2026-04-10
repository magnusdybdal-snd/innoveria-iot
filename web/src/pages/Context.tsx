import { useRef, useState } from "react";

import AddIcon from "@mui/icons-material/Add";
import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { GraphWidget, GraphWidgetPlaceholder } from "@widgets/graphWidget";

/**
 * Context page for visualizing sensor data through configurable graph widgets.
 * @returns The rendered Context page
 */
export default function Context() {
  const [widgetIds, setWidgetIds] = useState<number[]>([0]);
  const nextWidgetId = useRef(1);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Context" />
        <PageDivider />
        <Box sx={{ mt: 4, ml: 4, mr: 4, mb: 2 }}>
          <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 2 }}>
            <Typography variant="h5" sx={{ fontWeight: 600 }}>
              Graph Widgets
            </Typography>
            <IconButton
              size="small"
              onClick={() => {
                setWidgetIds((ids) => [...ids, nextWidgetId.current++]);
              }}
            >
              <AddIcon fontSize="small" />
            </IconButton>
          </Box>
          <Box
            sx={{
              display: "grid",
              gridTemplateColumns: "repeat(2, 1fr)",
              gap: 2,
            }}
          >
            {widgetIds.length === 0 ? (
              <GraphWidgetPlaceholder
                onAdd={() => {
                  setWidgetIds((ids) => [...ids, nextWidgetId.current++]);
                }}
              />
            ) : (
              widgetIds.map((id) => (
                <GraphWidget
                  key={id}
                  onDelete={() =>
                    setWidgetIds((ids) => ids.filter((w) => w !== id))
                  }
                />
              ))
            )}
          </Box>
        </Box>
      </PageContent>
    </div>
  );
}
