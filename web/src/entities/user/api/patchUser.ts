import { serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

export interface UpdateUserRequest {
  name?: string;
  email?: string;
  role?: string;
  password?: string;
}

/**
 * Updates a user's fields by ID. Only provided fields are updated.
 * @param userId - The ID of the user to update
 * @param payload - The fields to update
 * @returns void
 */
export const patchUser = async (
  userId: string,
  payload: UpdateUserRequest,
): Promise<void> => {
  await serviceClient.patch(
    `${API_ROUTES.users}/${encodeURIComponent(userId)}`,
    payload,
  );
};
