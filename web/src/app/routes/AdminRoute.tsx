import { Navigate, Outlet } from "react-router";

import { useCurrentUser } from "@app/providers/useCurrentUser";

/**
 * Route guard that restricts access to PLATFORM_ADMIN users.
 * Redirects non-admins to the home page. Renders nothing while the user is loading.
 * @returns The child routes via Outlet, a redirect to "/", or null while loading
 */
export default function AdminRoute() {
  const { user, isLoading } = useCurrentUser();

  if (isLoading) return null;

  if (user?.role !== "PLATFORM_ADMIN") {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
}
