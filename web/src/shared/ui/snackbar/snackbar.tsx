import Alert from "@mui/material/Alert";
import Snackbar from "@mui/material/Snackbar";

export const SNACKBAR_SEVERITY = {
  ERROR: "error",
  SUCCESS: "success",
  INFO: "info",
  WARNING: "warning",
  UNDEFINED: undefined,
} as const;

export type SnackbarSeverity =
  (typeof SNACKBAR_SEVERITY)[keyof typeof SNACKBAR_SEVERITY];

interface AppSnackbarProps {
  open: boolean;
  message: string;
  onClose: () => void;
  severity?: (typeof SNACKBAR_SEVERITY)[keyof typeof SNACKBAR_SEVERITY];
  autoHideDuration?: number;
}

/**
 * Reusable snackbar notification with configurable severity and message.
 * @param props - Component props
 * @param props.open - Whether the snackbar is visible
 * @param props.message - The text to display
 * @param props.onClose - Called when the snackbar should close
 * @param props.severity - Controls the color: "error" | "success" | "info" | "warning". Defaults to "info".
 * @param props.autoHideDuration - Ms before auto-dismiss. Defaults to 5000.
 * @returns The rendered snackbar with an alert inside
 */
export function AppSnackbar({
  open,
  message,
  onClose,
  severity = SNACKBAR_SEVERITY.UNDEFINED,
  autoHideDuration = 5000, // Default to 5 seconds
}: AppSnackbarProps) {
  return (
    <Snackbar
      open={open}
      autoHideDuration={autoHideDuration}
      onClose={onClose}
      anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
    >
      <Alert onClose={onClose} severity={severity} variant="filled">
        {message}
      </Alert>
    </Snackbar>
  );
}
