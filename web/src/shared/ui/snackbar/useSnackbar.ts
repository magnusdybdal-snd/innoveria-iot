import { useState } from "react";

import { type SnackbarSeverity } from "./snackbar";

/**
 * Custom hook to manage snackbar state and provide a function to show snackbars with different severities.
 * @returns An object with a `show` function to display a snackbar and the current `snackbar` state.
 */
export function useSnackbar() {
  // Null means no snackbar has been shown yet. Once show() is called,
  // the object persists with open:false during the exit animation so
  // the message and severity are preserved and don't flash to defaults.
  const [snackbar, setSnackbar] = useState<{
    open: boolean;
    message: string;
    severity: SnackbarSeverity;
  } | null>(null);

  // Opens the snackbar with the given message and severity.
  const show = (message: string, severity: SnackbarSeverity) => {
    setSnackbar({ open: true, message, severity });
  };

  // Closes the snackbar by setting open:false, keeping message/severity
  // intact so MUI's exit animation renders the correct color and text.
  const hide = () =>
    setSnackbar((prev) => (prev ? { ...prev, open: false } : null));

  return { show, hide, snackbar };
}
