import type { ReactNode } from "react";

import Box from "@mui/material/Box";

interface PageContentProps {
  children: ReactNode;
}

export function PageContent({ children }: PageContentProps) {
  return (
    <Box
      sx={{
        backgroundColor: "primary.dark",
        color: "primary.main",
        paddingTop: 2,
        paddingBottom: 5,
        paddingLeft: 2,
        paddingRight: 2,
      }}
      className="flex-1 overflow-auto"
    >
      {children}
    </Box>
  );
}
