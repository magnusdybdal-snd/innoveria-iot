import type { UserApiResponse } from "@entities/user/model/userSchema.ts";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawUser = {
  id: string;
  company_id: string;
  name: string;
  email: string;
  role: string;
  last_logged_in: string | null;
  created_at: string;
  updated_at: string;
};

type RawUserListApiResponse = {
  total_count: number;
  users: RawUser[];
};

/**
 * Fetches all users, optionally filtered by company.
 * @param companyId - If provided, only users belonging to this company are returned
 * @returns Array of UserApiResponse objects
 */
export const getUsers = async (
  companyId?: string,
): Promise<UserApiResponse[]> => {
  const url = companyId
    ? `${API_ROUTES.users}?company_id=${encodeURIComponent(companyId)}`
    : API_ROUTES.users;

  const data = await apiRequest<RawUserListApiResponse>(
    serviceClient,
    url,
    "GET",
  );

  return (data.users ?? []).map((u) => ({
    id: u.id,
    companyId: u.company_id,
    name: u.name,
    email: u.email,
    role: u.role,
    lastLoggedIn: u.last_logged_in,
    createdAt: u.created_at,
    updatedAt: u.updated_at,
  }));
};
