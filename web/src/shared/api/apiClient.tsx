import axios, { type AxiosInstance, type AxiosResponse } from "axios";

import { postRefresh, type TokenApiResponse } from "@entities/user";

// Client for all microservice requests — routed through the api-gateway
export const serviceClient = axios.create({
  baseURL: "/api",
  headers: {
    "Content-Type": "application/json",
  },
});

let isRefreshing = false;
let refreshPromise: Promise<TokenApiResponse> | null = null;

serviceClient.interceptors.request.use((config) => {
  const token = localStorage.getItem("access_token");

  const isLoginRequest = config.url?.includes("/login");
  const isRefreshRequest = config.url?.includes("/refresh");

  if (token && config.headers && !isLoginRequest && !isRefreshRequest) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

serviceClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (!originalRequest) return Promise.reject(error);

    const isRefreshRequest = originalRequest.url?.includes("/refresh");
    const isLoginRequest = originalRequest.url?.includes("/login");

    // No token at all → redirect to login immediately
    if (
      error.response?.status === 401 &&
      !isRefreshRequest &&
      !isLoginRequest &&
      !localStorage.getItem("access_token")
    ) {
      window.location.href = "/Login";
      return Promise.reject(error);
    }

    // If refresh fails → logout
    if (error.response?.status === 401 && isRefreshRequest) {
      isRefreshing = false;
      refreshPromise = null;
      localStorage.removeItem("access_token");
      window.location.href = "/Login";
      return Promise.reject(error);
    }

    // Only refresh for normal API calls (not login/refresh) and only when a token exists
    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      !isRefreshRequest &&
      !isLoginRequest &&
      localStorage.getItem("access_token")
    ) {
      originalRequest.headers = originalRequest.headers || {};
      originalRequest._retry = true;

      try {
        // Prevent multiple refresh calls
        if (!isRefreshing) {
          isRefreshing = true;
          refreshPromise = postRefresh();
        }

        if (!refreshPromise) {
          throw new Error("Refresh promise was not initialized");
        }

        const data = await refreshPromise;

        isRefreshing = false;
        refreshPromise = null;

        localStorage.setItem("access_token", data.accessToken);

        // Retry original request with new token
        originalRequest.headers.Authorization = `Bearer ${data.accessToken}`;

        return serviceClient(originalRequest);
      } catch (refreshError) {
        isRefreshing = false;
        refreshPromise = null;

        // Refresh fails → logout
        localStorage.removeItem("access_token");
        window.location.href = "/Login";

        return Promise.reject(refreshError);
      }
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
  method: "GET" | "POST" | "PUT" | "PATCH" | "DELETE",
  data?: unknown,
): Promise<T> => {
  const response: AxiosResponse<T> = await client({
    method,
    url,
    data,
  });

  return response.data;
};
