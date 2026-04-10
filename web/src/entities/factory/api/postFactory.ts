import { apiRequest, serviceClient } from "@/shared/api";
import { API_ROUTES } from "@/shared/api/routes";

import type { CreateFactoryRequest } from "@entities/factory";

/**
 * Post new Factory (POST REQUEST)
 * @param factoryData - an object containing factory data
 */
export const postFactory = async (
  factoryData: CreateFactoryRequest,
): Promise<void> => {
  await apiRequest(serviceClient, API_ROUTES.factories, "POST", {
    company_id: factoryData.companyId,
    name: factoryData.name,
    address: factoryData.address,
  });
};
