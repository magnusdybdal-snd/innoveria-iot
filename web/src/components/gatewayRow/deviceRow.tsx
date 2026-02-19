import type { ReactNode } from "react";

import Box from "@mui/material/Box";

export function DeviceRow({ children }: { children: ReactNode }) {
  return (
    <Box
      sx={{
        // Span all parent grid columns and inherit their sizing so children align with headers
        gridColumn: "1 / -1",
        display: "grid",
        gridTemplateColumns: "subgrid",
        backgroundColor: "secondary.light",
        padding: 2,
        borderRadius: 2,
        marginTop: 1,
        marginBottom: 1,
      }}
    >
      {children}
    </Box>
  );
}
