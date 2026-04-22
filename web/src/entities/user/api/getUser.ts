import type { UserApiResponse } from "@entities/user/model/userSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawUserApiResponse = {
  company_id: string;
  email: string;
  name: string;
  role: string;
  user_id: string;
};

/**
 * Fetches user from the authentication-service via the API user.
 * @returns Array of UserApiResponse objects, or nothing if the request fails
 */
export const getUser = async (): Promise<UserApiResponse> => {
  try {
    const data = await apiRequest<RawUserApiResponse>(
      serviceClient,
      API_ROUTES.user,
      "GET",
    );

    return {
      companyId: data.company_id,
      email: data.email,
      name: data.name,
      role: data.role,
      id: data.user_id,
      lastLoggedIn: null,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
  } catch (error) {
    if (error instanceof Error) {
      console.error("Failed to fetch user:", error.message);
    } else {
      console.error("Failed to fetch user:", error);
    }

    throw error;
  }
};
