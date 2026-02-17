// Fetch with safe error handling
import {
  apiClient,
  apiRequest,
  chuckApiClient,
  type Joke,
  type Data,
} from "./apiClient.tsx";

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

/*/ Fetching data with our API client
export const fetchJoke = async (): Promise<Data> => {
    return await apiRequest<Data>(chuckApiClient, "/random", "GET");
};

export const fetchApi = async (): Promise<Data> => {
    return await apiRequest<Data>(apiClient, "/applications", "GET");
};*/
