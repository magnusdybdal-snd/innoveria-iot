import type {
  CompanyApiResponse,
  CompanyListApiResponse,
} from "@entities/company/model/companySchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Fetches all companies from the collection-service via the API company.
 * @returns Array of CompanyApiResponse objects, or an empty array if the request fails
 */
export const getCompanies = async (): Promise<CompanyApiResponse[]> => {
  try {
    const data = await apiRequest<CompanyListApiResponse>(
      serviceClient,
      API_ROUTES.companies,
      "GET",
    );

    return data.companies ?? [];
  } catch (error) {
    console.error("Failed to fetch companies:", error);
    return [];
  }
};
