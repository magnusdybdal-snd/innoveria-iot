import axios, { type AxiosResponse } from "axios";

// Instance of axios with some default configuration
const apiClient = axios.create({
  baseURL: "https://api.chucknorris.io/jokes",
  headers: {
    "Content-Type": "application/json",
  },
});

// Generic API helper
export const apiRequest = async <T,>(
  url: string,
  method: "GET" | "POST" | "PUT" | "DELETE",
  data?: unknown,
): Promise<T> => {
  const response: AxiosResponse<T> = await apiClient({
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

/*/ Shape of the response for fetching list
interface UserListResponse {
    jokes: Joke[];
    total: number;
}*/

// Fetching joke with our API client
export const fetchJoke = async (): Promise<Joke> => {
  return await apiRequest<Joke>("/random", "GET");
};

// Fetch joke with safe error handling
export const fetchJokeWithErrorHandling = async (): Promise<Joke | null> => {
  try {
    return await apiRequest<Joke>("/random", "GET");
  } catch (error) {
    console.error("Failed to fetch joke:", error);
    return null;
  }
};
