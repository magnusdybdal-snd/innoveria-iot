import Box from "@mui/material/Box";

import { AddBox } from "@/components/addCompBox";
import { Path } from "@/components/path";
import Menu from "@/Menu.tsx";

export default function Home() {
  return (
    <Menu>
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
    </Menu>
  );
}
