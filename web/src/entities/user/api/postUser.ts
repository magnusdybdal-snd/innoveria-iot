import type {
  CreateUserRequest,
  UserApiResponse,
} from "@entities/user/model/userSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Creates a new user.
 * @param userData - The user data to create
 * @returns The created UserApiResponse
 */
export const postUser = async (
  userData: CreateUserRequest,
): Promise<UserApiResponse> => {
  return await apiRequest<UserApiResponse>(
    serviceClient,
    API_ROUTES.users,
    "POST",
    {
      company_id: userData.companyId,
      name: userData.name,
      email: userData.email,
      password: userData.password,
      role: userData.role,
    },
  );
};
