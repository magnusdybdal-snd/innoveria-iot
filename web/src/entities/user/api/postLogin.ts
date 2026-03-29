import type {
  LoginRequest,
  TokenApiResponse,
} from "@entities/user/model/userSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawTokenApiResponse = {
  access_token: string;
  expires_in: number;
  token_type: string;
};

/**
 * Request a login to the API.
 * @param loginData - An object containing the email and password needed to log in.
 * @returns RefreshRequest object, or nothing if the request fails
 */
export const postLogin = async (
  loginData: LoginRequest,
): Promise<TokenApiResponse> => {
  try {
    const data = await apiRequest<RawTokenApiResponse>(
      serviceClient,
      API_ROUTES.login,
      "POST",
      {
        email: loginData.email,
        password: loginData.password,
      },
    );

    return {
      accessToken: data.access_token,
      expiresIn: data.expires_in,
      tokenType: data.token_type,
    };
  } catch (error) {
    if (error instanceof Error) {
      console.error("Login failed:", error.message);
    } else {
      console.error("Login failed:", error);
    }
    throw error;
  }
};
