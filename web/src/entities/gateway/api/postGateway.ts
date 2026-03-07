import type { CreateGatewayRequest } from "@entities/gateway/model/gatewaySchema";

/**
 * Posts a new gateway to the API. This function sends a POST request to the /api/gateways endpoint with the provided gateway data.
 * @param gatewayData - An object containing the companyId, deviceEui, and name of the gateway to be created.
 */
export const postGateway = async (
  gatewayData: CreateGatewayRequest,
): Promise<void> => {
  const response = await fetch("/api/gateways", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    // The API expects snake_case keys, so we need to convert the camelCase keys from the CreateGatewayRequest to snake_case.
    body: JSON.stringify({
      company_id: gatewayData.companyId,
      gateway_eui: gatewayData.deviceEui,
      name: gatewayData.name,
    }),
  });

  if (!response.ok) {
    throw new Error(`Error posting gateway: ${response.statusText}`);
  }
};
