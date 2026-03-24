// API base URL is handled via environment variables (VITE_API_URL) in Vite and Docker.

// In production, Caddy proxies /api/* so relative paths work.
export const API_ROUTES = {
  gateways: `/v1/device/gateways`,
  sensors: `/v1/device/sensors`,
  collection: `/v1/collection`,
  sensorProfile: `/v1/device/sensor-profiles`,
  sensorLatest: `/v1/collection/latest?device_eui=`,
  companies: `/v1/auth/companies`,
  factories: `/v1/auth/factories`,
  companiesGet: `/v1/auth/companies`,
  companiesPost: `/v1/onboarding/company`,
  user: `/v1/auth/me`,
  login: `/v1/auth/login`,
  refresh: `/v1/auth/refresh`,
};
