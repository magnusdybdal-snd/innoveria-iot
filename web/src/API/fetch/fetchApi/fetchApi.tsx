// Fetch with safe error handling
import { apiClient, apiRequest, type Data } from "@/API/apiClient";

export const fetchApiWithErrorHandling = async (): Promise<Data | null> => {
  try {
    return await apiRequest<Data>(apiClient, "/sensors", "GET");
  } catch (error) {
    console.error("Failed to fetch data:", error);
    return null;
  }
};

/*
export const fetchApi = async (): Promise<Data> => {
    return await apiRequest<Data>(apiClient, "/sensors", "GET");
};*/
