import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Revokes the server-side refresh token cookie and clears it from the browser.
 * @returns A promise that resolves when the logout request completes
 */
export const postLogout = (): Promise<void> =>
  apiRequest<void>(serviceClient, API_ROUTES.logout, "POST");
