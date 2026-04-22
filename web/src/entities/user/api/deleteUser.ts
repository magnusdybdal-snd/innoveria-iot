import { serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Deletes a user by ID.
 * @param userId - The ID of the user to delete
 * @returns True if deletion was successful
 */
export const deleteUser = async (userId: string): Promise<boolean> => {
  const response = await serviceClient.delete(
    `${API_ROUTES.users}/${encodeURIComponent(userId)}`,
  );
  return response.status === 204;
};
