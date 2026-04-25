import { Navigate, Outlet } from "react-router";

import { useCurrentUser } from "@app/providers/useCurrentUser";

/**
 * Route guard that restricts access to authenticated users.
 * Redirects unauthenticated users to the login page. Renders nothing while the user is loading.
 * @returns The child routes via Outlet, a redirect to "/Login", or null while loading
 */
export default function ProtectedRoute() {
  const { user, isLoading } = useCurrentUser();

  if (isLoading) return null;

  if (!user) {
    return <Navigate to="/Login" replace />;
  }

  return <Outlet />;
}
