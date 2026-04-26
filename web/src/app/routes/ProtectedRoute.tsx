import Button from "@mui/material/Button";
import CircularProgress from "@mui/material/CircularProgress";
import Typography from "@mui/material/Typography";
import { Navigate, Outlet } from "react-router";

import { useCurrentUser } from "@app/providers/useCurrentUser";

/**
 * Route guard that restricts access to authenticated users.
 * - Shows a spinner while the user fetch is in progress.
 * - Shows an error message if the fetch failed (distinguishes backend failure from not being logged in).
 * - Redirects unauthenticated users to /Login.
 * @returns The child routes via Outlet, a redirect to "/Login", an error message, or a spinner while loading
 */
export default function ProtectedRoute() {
  const { user, isLoading, error } = useCurrentUser();

  if (isLoading) return <CircularProgress />;

  if (error) {
    return (
      <div style={{ padding: 32 }}>
        <Typography color="error">Could not verify your session.</Typography>
        <Button onClick={() => window.location.reload()} sx={{ mt: 2 }}>
          Retry
        </Button>
      </div>
    );
  }

  if (!user) {
    return <Navigate to="/Login" replace />;
  }

  return <Outlet />;
}
