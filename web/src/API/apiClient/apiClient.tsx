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
  baseURL: import.meta.env.VITE_API_URL, // Imported api url from vite config
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

// TODO: move ALL data shapes to their own layer (issue #106)
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
