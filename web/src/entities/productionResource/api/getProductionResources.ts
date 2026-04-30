import type { ProductionResourceApiResponse } from "@entities/productionResource/model/productionResourceSchema";
import { apiRequest, serviceClient } from "@shared/api";
import { API_ROUTES } from "@shared/api/routes";

type RawProductionResource = {
  id: number;
  number: string;
  description: string;
  type: string;
};

/**
 * Fetches all ERP production resources (work centers) from the context-service.
 * @returns Array of ProductionResourceApiResponse objects
 */
export const getProductionResources = async (): Promise<
  ProductionResourceApiResponse[]
> => {
  const data = await apiRequest<RawProductionResource[]>(
    serviceClient,
    API_ROUTES.contextProductionResources,
    "GET",
  );

  return (data ?? []).map((r) => ({
    id: r.id,
    number: r.number,
    description: r.description,
    type: r.type,
  }));
};
