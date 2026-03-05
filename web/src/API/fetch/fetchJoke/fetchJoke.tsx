// Fetch with safe error handling
import { apiRequest, chuckApiClient, type Joke } from "@/API/apiClient";

/**
 * Fetches a random Chuck Norris joke from the Chuck Norris API with safe error handling.
 * @returns A Joke object on success, or null if the request fails
 */
export const fetchJokeWithErrorHandling = async (): Promise<Joke | null> => {
  try {
    return await apiRequest<Joke>(chuckApiClient, "/random", "GET");
  } catch (error) {
    console.error("Failed to fetch joke:", error);
    return null;
  }
};

/*/ Fetching data with our API client
export const fetchJoke = async (): Promise<Data> => {
    return await apiRequest<Data>(chuckApiClient, "/random", "GET");
};*/
