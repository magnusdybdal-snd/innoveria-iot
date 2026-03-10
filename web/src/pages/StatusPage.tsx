import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";
import { Link as RouterLink } from "react-router";

/**
 * StatusPage component to display status information.
 * @param status - An object containing the status code and message to be displayed on the page.
 * @param status.code - The status code to display
 * @param status.message - The status message to display
 * @returns A simple status page with the provided status code and message
 */
export default function StatusPage(status: { code: string; message: string }) {
  return (
    <Box
      display="flex"
      height="100vh"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
    >
      <Typography variant="h1" fontWeight="bold" mb={1}>
        {status.code}
      </Typography>
      <Typography variant="h4">{status.message}</Typography>
      <Link component={RouterLink} to="/" mt={3} variant="h6">
        Go to home page
      </Link>
    </Box>
  );
}
