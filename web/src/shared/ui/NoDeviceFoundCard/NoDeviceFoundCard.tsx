import ErrorIcon from "@mui/icons-material/Error";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import { yellow } from "@mui/material/colors";

/**
 * NoDeviceFoundCard
 * @returns card with message that no device has been found
 */
export function NoDeviceFoundCard() {
  return (
    <>
      <Card
        sx={{
          marginTop: 2,
          color: "primary.main",
          backgroundColor: "primary.light",
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            justifyContent: "center",
            padding: 2,
          }}
        >
          <ErrorIcon sx={{ margin: 1, color: yellow[500] }} />
          <p>No devices found</p>
        </Box>
      </Card>
    </>
  );
}
