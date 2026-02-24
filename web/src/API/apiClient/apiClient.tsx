import axios, { type AxiosInstance, type AxiosResponse } from "axios";

// Instance of axios with some default configuration
export const chuckApiClient = axios.create({
  baseURL: "https://api.chucknorris.io/jokes",
  headers: {
    "Content-Type": "application/json",
  },
});

// Client for all microservice requests — routed through the api-gateway
export const serviceClient = axios.create({
  baseURL: "http://localhost:8081",
  headers: {
    "Content-Type": "application/json",
  },
});

// Generic API helper
export const apiRequest = async <T,>(
  client: AxiosInstance,
  url: string,
  method: "GET" | "POST" | "PUT" | "DELETE",
  data?: unknown,
): Promise<T> => {
  const response: AxiosResponse<T> = await client({
    method,
    url,
    data,
  });

  return response.data;
};

// Data shape
export interface Joke {
  categories: [];
  created_at: string;
  icon_url: string;
  id: string;
  updated_at: string;
  url: string;
  value: string;
}

export interface GatewayApiResponse {
  id: string;
  device_eui: string;
  name: string;
  status: number;
  lastSeenAt: string; // RFC1123 string from go - check format
}

export interface GatewayListApiResponse {
  totalCount: number;
  gateways: GatewayApiResponse[];
}

export interface SensorApiResponse {
  id: string;
  device_eui: string;
  name: string;
  status: number;
  machine: string;
}

export interface SensorListApiResponse {
  totalCount: number;
  sensors: SensorApiResponse[];
}
