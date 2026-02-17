import Box from "@mui/material/Box";
import { Path } from "@/components/path";
import Typography from "@mui/material/Typography";

export function SubPageHeader({ title }: { title: string }) {
  return (
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
        }}
        className="flex-1 overflow-auto"
      >
        <Path />
        <div className="flex justify-between flex-wrap">
          <Typography variant="h4">{title}</Typography>
        </div>
      </Box>
    </Box>
  );
}
