/**
 * This file contains the types for the Company entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching company information.
 */
export interface UserApiResponse {
  companyId: string;
  email: string;
  name: string;
  role: string;
  userId: string;
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
