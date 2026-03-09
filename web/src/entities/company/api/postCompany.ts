import type { CreateCompanyRequest } from "@entities/company/model/companySchema";
import { API_ROUTES } from "@shared/api/routes";

/**
 * Posts a new company to the API. This function sends a POST request to the /api/v1/device/companies endpoint with the provided company data.
 * @param companyData - An object containing the companyId, deviceEui, and name of the company to be created.
 */
export const postCompany = async (
  companyData: CreateCompanyRequest,
): Promise<void> => {
  const response = await fetch(API_ROUTES.companies, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    // The API expects snake_case keys, so we need to convert the camelCase keys from the CreateCompanyRequest to snake_case.
    body: JSON.stringify({
      address: companyData.address,
      name: companyData.name,
    }),
  });

  if (!response.ok) {
    throw new Error(`Error posting company: ${response.statusText}`);
  }
};
