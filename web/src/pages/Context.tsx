import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

/**
 * Context page that is currently under construction and serves as a placeholder for future context-related features.
 * @returns The rendered Context page with a message indicating it's under construction
 */
export default function Context() {
  return (
    <div className="flex h-screen">
      <Box
        sx={{
          backgroundColor: "primary.dark",
          color: "primary.main",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          textAlign: "center",
        }}
        className="flex-1 overflow-auto"
      >
        <Typography className="font-normal text-3xl">
          Context page - under construction
        </Typography>
      </Box>
    </div>
  );
}
