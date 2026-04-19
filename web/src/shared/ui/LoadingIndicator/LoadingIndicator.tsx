import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";
import Typography from "@mui/material/Typography";

/** Props for the `LoadingIndicator` component. */
interface LoadingIndicatorProps {
  /** Message displayed next to the spinner. Defaults to "Loading…". */
  message?: string;
}

/**
 * Inline spinner with an optional descriptive message.
 * @param props - Component props
 * @param props.message - Text shown next to the spinner
 * @returns The rendered loading indicator
 */
export function LoadingIndicator({
  message = "Loading…",
}: LoadingIndicatorProps) {
  return (
    <Box sx={{ mt: 3, display: "flex", alignItems: "center", gap: 2 }}>
      <CircularProgress size={20} />
      <Typography variant="body2" sx={{ color: "primary.main" }}>
        {message}
      </Typography>
    </Box>
  );
}
