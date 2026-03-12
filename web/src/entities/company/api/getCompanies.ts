import type { CompanyApiResponse } from "@entities/company";
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
      companyId: s.company_id,
      name: s.name,
      address: s.address,
      createdAt: s.created_at,
      updatedAt: s.updated_at,
    }));
  } catch (error) {
    console.error("Failed to fetch companies:", error);
    return [];
  }
};
