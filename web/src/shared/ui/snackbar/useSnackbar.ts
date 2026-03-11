import { useState } from "react";

import { type SnackbarSeverity } from "./snackbar";

/**
 * Custom hook to manage snackbar state and provide a function to show snackbars with different severities.
 * @returns An object with a `show` function to display a snackbar and the current `snackbar` state.
 */
export function useSnackbar() {
  const [snackbar, setSnackbar] = useState<{
    message: string;
    severity: SnackbarSeverity;
  } | null>(null);
  const show = (message: string, severity: SnackbarSeverity) => {
    setSnackbar({ message, severity });
  };

  const hide = () => setSnackbar(null);

  return { show, hide, snackbar };
}
