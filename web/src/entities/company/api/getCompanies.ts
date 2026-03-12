import type { CompanyApiResponse } from "@entities/company/model/companySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawCompany = {
  company_id: string;
  name: string;
  address: string;
  created_at: string;
  updated_at: string;
};

type RawCompanyListApiResponse = {
  total_count: number;
  companies: RawCompany[];
};

/**
 * Fetches all companies from the collection-service via the API company.
 * @returns Array of CompanyApiResponse objects, or an empty array if the request fails
 */
export const getCompanies = async (): Promise<CompanyApiResponse[]> => {
  try {
    const data = await apiRequest<RawCompanyListApiResponse>(
      serviceClient,
      API_ROUTES.companiesGet,
      "GET",
    );

    return (data.companies ?? []).map((s) => ({
      company_id: s.company_id,
      name: s.name,
      address: s.address,
      created_at: s.created_at,
      updated_at: s.updated_at,
    }));
  } catch (error) {
    console.error("Failed to fetch companies:", error);
    return [];
  }
};
