import Alert from "@mui/material/Alert";
import MuiSnackbar from "@mui/material/Snackbar";

type SnackbarSeverity = "error" | "success" | "info" | "warning";

interface AppSnackbarProps {
  open: boolean;
  message: string;
  onClose: () => void;
  severity?: SnackbarSeverity;
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
  severity = "info",
  autoHideDuration = 5000,
}: AppSnackbarProps) {
  return (
    <MuiSnackbar
      open={open}
      autoHideDuration={autoHideDuration}
      onClose={onClose}
      anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
    >
      <Alert onClose={onClose} severity={severity} variant="filled">
        {message}
      </Alert>
    </MuiSnackbar>
  );
}
