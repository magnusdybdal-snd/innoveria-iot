import type { ReactNode } from "react";

import Box from "@mui/material/Box";

interface PageContentProps {
  children: ReactNode;
}
/**
 * Wrapper for the main scrollable area of a page.
 * Provides consistent background color and padding across all pages.
 * @param props - Component props
 * @param props.children - Page content to render inside the wrapper
 * @returns A scrollable Box with consistent page-level styling
 */
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

        // Hide scrollbar
        "&::-webkit-scrollbar": {
          display: "none", // Chrome, Safari, Edge
        },
        scrollbarWidth: "none", // Firefox
      }}
      className="flex-1 overflow-auto"
    >
      {children}
    </Box>
  );
}
