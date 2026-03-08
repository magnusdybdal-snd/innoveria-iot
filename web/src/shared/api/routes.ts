// API routes used across the aplication.
export const API_ROUTES = {
  gateways: "/api/v1/device/gateways",
  sensors: "/api/v1/device/sensors",
} as const; // makes the object readonly and preserves literal types.
