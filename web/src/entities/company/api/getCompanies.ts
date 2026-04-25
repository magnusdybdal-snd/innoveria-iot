import type { CompanyApiResponse } from "@entities/company";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawCompany = {
  id?: string;
  company_id?: string;
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
 * Fetches all companies from the auth-service API.
 * @returns Array of CompanyApiResponse objects
 */
export const getCompanies = async (): Promise<CompanyApiResponse[]> => {
  const data = await apiRequest<RawCompanyListApiResponse>(
    serviceClient,
    API_ROUTES.companiesGet,
    "GET",
  );

  return (data.companies ?? []).map((s) => {
    const companyId = s.id ?? s.company_id ?? "";

    return {
      companyId,
      name: s.name,
      address: s.address,
      createdAt: s.created_at,
      updatedAt: s.updated_at,
    };
  });
};
