import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawERPAgentTokenResponse = {
  erp_agent_token: string;
};

/**
 * Issues a new ERP agent token for a company.
 * @param companyId - The company ID to issue a token for.
 * @returns The ERP agent token string.
 */
export const postCompanyERPAgentToken = async (
  companyId: string,
): Promise<string> => {
  const data = await apiRequest<RawERPAgentTokenResponse>(
    serviceClient,
    API_ROUTES.issueCompanyERPAgentToken(companyId),
    "POST",
  );

  return data.erp_agent_token;
};
