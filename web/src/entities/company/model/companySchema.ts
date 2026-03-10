/**
 * This file contains the types for the Company entity, as well as the API response types.
 * It defines the structure of the data returned by the API when fetching company information.
 */
export interface CompanyApiResponse {
  company_id: string;
  name: string;
  address: string;
  created_at: string; // RFC1123 - directly parsable in JS.
  updated_at: string;
}

export interface CompanyListApiResponse {
  totalCount: number; // TODO: check id backens uses batching of max fetched in one fetch - if so update logic to fetch again. (pagination)
  companies: CompanyApiResponse[];
}

export interface CreateCompanyRequest {
  name: string;
  address: string;
}
