import axios, { type AxiosInstance, type AxiosResponse } from "axios";

// Client for all microservice requests — routed through the api-gateway
export const serviceClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL, // Imported api url from vite config
  headers: {
    "Content-Type": "application/json",
  },
});

/**
 * Generic API request helper
 * @param client - The Axios instance to use
 * @param url - The endpoint URL
 * @param method - The HTTP method
 * @param data - The request payload
 * @returns A promise resolving to the response data
 */
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
