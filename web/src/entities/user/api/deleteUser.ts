import { serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Deletes a user by ID. Throws if the server does not respond with 204.
 * @param userId - The ID of the user to delete
 */
export const deleteUser = async (userId: string): Promise<void> => {
  await serviceClient.delete(
    `${API_ROUTES.users}/${encodeURIComponent(userId)}`,
    { validateStatus: (status) => status === 204 },
  );
};
