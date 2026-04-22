/**
 * This file contains the types for the User entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching user information.
 */
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
