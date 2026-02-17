import Box from "@mui/material/Box";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu";

export default function Gateways() {
  return (
    <div className="flex h-screen">
      <Menu />
      <Box
        sx={{ backgroundColor: "primary.dark", color: "primary.main" }}
        className="flex-1 overflow-auto"
      >
        <Box sx={{ paddingLeft: 2, paddingRight: 2 }}>
          <SubPageHeader title="Gateways" />
        </Box>
      </Box>
    </div>
  );
}
