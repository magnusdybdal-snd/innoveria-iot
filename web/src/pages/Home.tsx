import Box from "@mui/material/Box";
import { Path } from "@/components/path";
import Menu from "@/Menu.tsx";
import { JokeViewer } from "@/API/getJoke";
import { ApiViewer } from "@/API/getData/";

export default function Home() {
  return (
    <div className="flex h-screen">
      <Menu />
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
          <JokeViewer />
          <ApiViewer />
        </Box>
      </Box>
    </div>
  );
}
