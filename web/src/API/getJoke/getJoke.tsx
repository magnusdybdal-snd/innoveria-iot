import React, { useEffect, useState } from "react";

import { type Joke } from "@/API/apiClient";
import { fetchJokeWithErrorHandling } from "@/API/fetchApi/fetchJoke";

export const JokeViewer: React.FC = () => {
  const [joke, setJoke] = useState<Joke | null>(null);

  useEffect(() => {
    const loadJoke = async () => {
      const result = await fetchJokeWithErrorHandling();
      setJoke(result);
    };

    loadJoke();
  }, []);

  if (!joke) return <p>Loading...</p>;

  return <p>{joke.value}</p>;
};
