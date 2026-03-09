// In dev, requests go directly to the api-gateway since the Vite proxy is not in use.
// In production, Caddy proxies /api/* so relative paths work.
const BASE_URL = import.meta.env.DEV ? "http://localhost:8081" : "";

export const API_ROUTES = {
  gateways: `${BASE_URL}/api/v1/device/gateways`,
  sensors: `${BASE_URL}/api/v1/device/sensors`,
  companies: `${BASE_URL}/api/v1/admin/companies`,
};
