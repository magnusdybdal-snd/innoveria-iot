import type { GatewayApiResponse } from "@entities/gateway";

export const mockGateways: GatewayApiResponse[] = [
  {
    id: "1",
    companyId: "",
    name: "Gateway 1",
    status: 0,
    gatewayEui: "A1B2C3D4E5F6",
    lastSeenAt: "Thu, 05 Mar 2026 10:00:00 GMT",
  },
  {
    id: "2",
    companyId: "",
    name: "Gateway 2",
    status: 2,
    gatewayEui: "B2C3D4E5F6A1",
    lastSeenAt: "Thu, 05 Mar 2026 09:00:00 GMT",
  },
  {
    id: "3",
    companyId: "",
    name: "Gateway 3",
    status: 1,
    gatewayEui: "B2D3D5E5A6A1",
    lastSeenAt: "Mon, 02 Mar 2026 10:00:00 GMT",
  },
];
