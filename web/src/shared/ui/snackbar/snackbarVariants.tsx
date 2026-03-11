import { AppSnackbar, SNACKBAR_SEVERITY } from "./snackbar";

interface SnackbarProps {
  open: boolean;
  message: string;
  onClose: () => void;
}
/**
 *  Pre-configured snackbar for success messages. Uses AppSnackbar with severity="success".
 * @param props - Component props
 * @param props.open - Whether the snackbar is visible
 * @param props.message - The text to display
 * @param props.onClose - Called when the snackbar should close
 * @returns The rendered success snackbar
 */
export function SuccessSnackbar(props: SnackbarProps) {
  return <AppSnackbar {...props} severity={SNACKBAR_SEVERITY.SUCCESS} />;
}

/**
 *  Pre-configured snackbar for error messages. Uses AppSnackbar with severity="error".
 * @param props - Component props
 * @param props.open - Whether the snackbar is visible
 * @param props.message - The text to display
 * @param props.onClose - Called when the snackbar should close
 * @returns The rendered error snackbar
 */
export function ErrorSnackbar(props: SnackbarProps) {
  return <AppSnackbar {...props} severity={SNACKBAR_SEVERITY.ERROR} />;
}

/**
 *  Pre-configured snackbar for info messages. Uses AppSnackbar with severity="info".
 * @param props - Component props
 * @param props.open - Whether the snackbar is visible
 * @param props.message - The text to display
 * @param props.onClose - Called when the snackbar should close
 * @returns The rendered info snackbar
 */
export function InfoSnackbar(props: SnackbarProps) {
  return <AppSnackbar {...props} severity={SNACKBAR_SEVERITY.INFO} />;
}

/**
 *  Pre-configured snackbar for warning messages. Uses AppSnackbar with severity="warning".
 * @param props - Component props
 * @param props.open - Whether the snackbar is visible
 * @param props.message - The text to display
 * @param props.onClose - Called when the snackbar should close
 * @returns The rendered warning snackbar
 */
export function WarningSnackbar(props: SnackbarProps) {
  return <AppSnackbar {...props} severity={SNACKBAR_SEVERITY.WARNING} />;
}
