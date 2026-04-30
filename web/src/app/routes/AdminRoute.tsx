import Typography from "@mui/material/Typography";
import { Navigate, Outlet } from "react-router";

import { useCurrentUser } from "@app/providers/useCurrentUser";

/**
 * Route guard that restricts access to PLATFORM_ADMIN users.
 * - Renders nothing while the user fetch is in progress.
 * - Shows an error message if the fetch failed (distinguishes network failure from not being logged in).
 * - Redirects unauthenticated users to /login.
 * - Redirects authenticated non-admins to /.
 * @returns The child routes via Outlet, a redirect, an error message, or null while loading
 */
export default function AdminRoute() {
  const { user, isLoading, error } = useCurrentUser();

  if (isLoading) return null;

  if (error) {
    return (
      <Typography color="error" sx={{ p: 4 }}>
        Could not verify your session. Please reload the page.
      </Typography>
    );
  }

  if (user === null) {
    return <Navigate to="/Login" replace />;
  }

  if (user.role !== "PLATFORM_ADMIN") {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
}
