import type { ReactNode } from "react";

import Box from "@mui/material/Box";

type DeviceRowProps = {
  children: ReactNode;
  onClick?: () => void;
};

export function DeviceRow({ children, onClick }: DeviceRowProps) {
  return (
    <Box
      onClick={onClick}
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
        whiteSpace: "nowrap", // Prevent wrapping
        overflow: "hidden", // Hide overflow text

        "&:hover": onClick ? { backgroundColor: "primary.dark" } : undefined,
      }}
    >
      {children}
    </Box>
  );
}
