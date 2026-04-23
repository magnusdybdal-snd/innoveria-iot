/**
 * This file contains the types for the User entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching user information.
 */

// Returned by GET /me — minimal profile used to identify and authorize the current session.
export interface CurrentUserApiResponse {
  id: string;
  companyId: string;
  name: string;
  email: string;
  role: string;
}

// Returned by GET /users — full record used in the admin user management table.
export interface UserApiResponse {
  id: string;
  companyId: string;
  name: string;
  email: string;
  role: string;
  lastLoggedIn: string | null;
  createdAt: string;
  updatedAt: string;
}

export interface CreateUserRequest {
  companyId: string;
  name: string;
  email: string;
  password: string;
  role: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface TokenApiResponse {
  accessToken: string;
  expiresIn: number;
  tokenType: string;
}
