import axios from "axios";

import { apiRequest } from "@/shared/api";

import type { Joke } from "../model/jokeSchema";

// Instance of axios with some default configuration
// Ideally the client should be in a separate file, but for simplicity it's included here
const chuckApiClient = axios.create({
  baseURL: "https://api.chucknorris.io/jokes",
  headers: {
    "Content-Type": "application/json",
  },
});

export const getJoke = async (): Promise<Joke | null> => {
  try {
    return await apiRequest<Joke>(chuckApiClient, "/random", "GET");
  } catch (error) {
    console.error("Failed to fetch joke:", error);
    return null;
  }
};
