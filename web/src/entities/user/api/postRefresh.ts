import type { TokenApiResponse } from "@entities/user/model/userSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Request a refresh to the API.
 * @param refreshData - An object containing the email and password needed to log in.
 */
export const postRefresh = async (
  refreshData: TokenApiResponse,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.login, "POST", {
    access_token: refreshData.accessToken,
    expires_in: refreshData.expiresIn,
    token_type: refreshData.tokenType,
  });
};
