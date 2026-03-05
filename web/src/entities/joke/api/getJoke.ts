import { apiRequest } from "@shared/api";
import axios from "axios";

import type { Joke } from "../model/jokeSchema";

// Instance of axios with some default configuration
// Ideally the client should be in a separate file, but for simplicity it's included here
const chuckApiClient = axios.create({
  baseURL: "https://api.chucknorris.io/jokes",
  headers: {
    "Content-Type": "application/json",
  },
});
/**
 *  Fetches a random joke from the Chuck Norris API using the apiRequest wrapper and returns it as a Joke object.
 * @returns  A Joke object containing the joke data, or null if the request fails
 */
export const getJoke = async (): Promise<Joke | null> => {
  try {
    return await apiRequest<Joke>(chuckApiClient, "/random", "GET");
  } catch (error) {
    console.error("Failed to fetch joke:", error);
    return null;
  }
};
