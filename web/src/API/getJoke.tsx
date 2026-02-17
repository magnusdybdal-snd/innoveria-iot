import React, { useEffect, useState } from "react";
import { type Joke, fetchJokeWithErrorHandling } from "./apiClient.tsx";

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
