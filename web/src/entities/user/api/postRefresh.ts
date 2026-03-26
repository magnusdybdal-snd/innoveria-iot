import type { TokenApiResponse } from "@entities/user";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawTokenApiResponse = {
  access_token: string;
  expires_in: number;
  token_type: string;
};

/**
 * Request a refresh to the API.
 * @returns RefreshRequest object, or nothing if the request fails
 */
export const postRefresh = async (): Promise<TokenApiResponse> => {
  try {
    const data = await apiRequest<RawTokenApiResponse>(
      serviceClient,
      API_ROUTES.refresh,
      "POST",
      {},
    );

    return {
      accessToken: data.access_token,
      expiresIn: data.expires_in,
      tokenType: data.token_type,
    };
  } catch (error) {
    if (error instanceof Error) {
      console.error("Refresh failed:", error.message);
    } else {
      console.error("Refresh failed:", error);
    }
    throw error;
  }
};
