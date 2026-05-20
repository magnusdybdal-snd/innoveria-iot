import type {
  CurrentUserApiResponse,
  UserRole,
} from "@entities/user/model/userSchema.ts";
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
 * Fetches the currently authenticated user's profile from the auth service (/me).
 * Throws if the request fails — callers are responsible for error handling.
 * @returns The current user's minimal profile
 */
export const getUser = async (): Promise<CurrentUserApiResponse> => {
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
      role: data.role as UserRole,
      id: data.user_id,
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
