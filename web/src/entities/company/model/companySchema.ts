/**
 * This file contains the types for the Company entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching company information.
 */
export interface CompanyApiResponse {
  companyId: string;
  name: string;
  address: string;
  createdAt: string; // RFC1123 - directly parsable in JS.
  updatedAt: string;
}

export interface CompanyListApiResponse {
  totalCount: number;
  companies: CompanyApiResponse[];
}

export interface CreateCompanyRequest {
  name: string;
  address: string;
}
