import { AppSnackbar } from "@shared/ui/snackbar";

interface ErrorSnackbarProps {
  open: boolean;
  message: string;
  onClose: () => void;
}

/**
 * Success snackbar that displays a green alert with the provided message.
 * @param props - Component props
 * @param props.open - Whether the snackbar is visible
 * @param props.message - The success message to display
 * @param props.onClose - Called when the snackbar should close
 * @returns A pre-configured error-severity snackbar
 */
export function SuccessSnackbar({
  open,
  message,
  onClose,
}: ErrorSnackbarProps) {
  return (
    <AppSnackbar
      open={open}
      message={message}
      onClose={onClose}
      severity="success"
    />
  );
}
