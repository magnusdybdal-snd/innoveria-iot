import { postRefresh } from "@entities/user";
import axios, { type AxiosInstance, type AxiosResponse } from "axios";

// Client for all microservice requests — routed through the api-gateway
export const serviceClient = axios.create({
  baseURL: "/api",
  headers: {
    "Content-Type": "application/json",
  },
});

serviceClient.interceptors.request.use((config) => {
  const token = localStorage.getItem("access_token");

  const isLoginRequest = config.url?.includes("/login");

  if (token && config.headers && !isLoginRequest) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

serviceClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      const data = await postRefresh();

      localStorage.setItem("access_token", data.accessToken);
    }

    return Promise.reject(error);
  },
);

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
