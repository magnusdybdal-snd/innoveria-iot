// API base URL is handled via environment variables (VITE_API_URL) in Vite and Docker.

// In production, Caddy proxies /api/* so relative paths work.
export const API_ROUTES = {
  gateways: `/api/v1/device/gateways`,
  sensors: `/api/v1/device/sensors`,
};
