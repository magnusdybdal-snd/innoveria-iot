import { apiRequest, serviceClient } from "@/shared/api";
import { API_ROUTES } from "@/shared/api/routes";

/**
 * Deletes a rule via the API gateway.
 * @param ruleId - the ID of the rule to be delete
 */
export const deleteRule = async (ruleId: string): Promise<void> => {
  await apiRequest(
    serviceClient,
    `${API_ROUTES.contextRules}/${encodeURIComponent(ruleId)}`,
    "DELETE",
  );
};
