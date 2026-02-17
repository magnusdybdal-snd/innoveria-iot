import axios, { type AxiosInstance, type AxiosResponse } from "axios";

// Instance of axios with some default configuration
const chuckApiClient = axios.create({
  baseURL: "https://api.chucknorris.io/jokes",
  headers: {
    "Content-Type": "application/json",
  },
});

const apiClient = axios.create({
  baseURL: "http://localhost:8090/api",
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

export interface Data {
  result: [
    {
      createdAt: string;
      description: string;
      id: string;
      name: string;
      updatedAt: string;
    },
  ];
  totalCount: number;
}

// Fetching data with our API client
export const fetchJoke = async (): Promise<Data> => {
  return await apiRequest<Data>(chuckApiClient, "/random", "GET");
};

export const fetchApi = async (): Promise<Data> => {
  return await apiRequest<Data>(apiClient, "/applications", "GET");
};

// Fetch with safe error handling
export const fetchJokeWithErrorHandling = async (): Promise<Joke | null> => {
  try {
    return await apiRequest<Joke>(chuckApiClient, "/random", "GET");
  } catch (error) {
    console.error("Failed to fetch joke:", error);
    return null;
  }
};

export const fetchApiWithErrorHandling = async (): Promise<Data | null> => {
  try {
    return await apiRequest<Data>(apiClient, "/applications", "GET");
  } catch (error) {
    console.error("Failed to fetch data:", error);
    return null;
  }
};
