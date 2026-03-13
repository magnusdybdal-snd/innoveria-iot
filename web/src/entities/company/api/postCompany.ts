import type { CreateCompanyRequest } from "@entities/company";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new company to the API.
 * @param companyData - An object containing the name and address of the company to be created.
 */
export const postCompany = async (
  companyData: CreateCompanyRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.companiesPost, "POST", {
    name: companyData.name,
    address: companyData.address,
  });
};
