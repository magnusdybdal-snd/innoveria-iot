import React, { useEffect, useState } from "react";
import { type Joke } from "./apiClient.tsx";
import { fetchJokeWithErrorHandling } from "./fetchApi.tsx";

const JokeViewer: React.FC = () => {
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

export default JokeViewer;
