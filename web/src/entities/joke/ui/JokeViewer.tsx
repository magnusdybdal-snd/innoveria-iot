import React, { useEffect, useState } from "react";

import { getJoke } from "../api";
import type { Joke } from "../model/jokeSchema";

export const JokeViewer: React.FC = () => {
  const [joke, setJoke] = useState<Joke | null>(null);

  useEffect(() => {
    const loadJoke = async () => {
      const result = await getJoke();
      setJoke(result);
    };

    loadJoke();
  }, []);

  if (!joke) return <p>Loading...</p>;

  return <p>{joke.value}</p>;
};
