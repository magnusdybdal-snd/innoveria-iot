import Box from "@mui/material/Box";

import { Path } from "@shared/ui/Path";
import { AddBox } from "@widgets/dashboard";

/**
 * Dashboard page that displays the current breadcrumb path and a placeholder to add new dashboard components.
 * @returns The rendered Dashboard page
 */
export default function Home() {
  return (
    <div className="flex h-screen">
      <Box
        sx={{
          backgroundColor: "primary.dark",
          color: "primary.main",
        }}
        className="flex-1 overflow-auto"
      >
        <Box
          sx={{
            paddingLeft: 2,
            paddingRight: 2,
            paddingTop: 2,
          }}
          className="flex-1 overflow-auto"
        >
          <Path />
          <AddBox />
        </Box>
      </Box>
    </div>
  );
}
