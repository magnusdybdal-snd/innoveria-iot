// API base URL is handled via environment variables (VITE_API_URL) in Vite and Docker.

// In production, Caddy proxies /api/* so relative paths work.
export const API_ROUTES = {
  gateways: `/v1/device/gateways`,
  sensors: `/v1/device/sensors`,
  collection: `/v1/collection`,
  sensorProfile: `/v1/device/sensor-profiles`,
  sensorLatest: `/v1/latest?device_eui=`,
  companies: `/v1/auth/companies`,
};
