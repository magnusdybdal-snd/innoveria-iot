import React, { useEffect, useState } from "react";

import { getJoke } from "../api";
import type { Joke } from "../model/jokeSchema";

/**
 * Fetches and displays a random Chuck Norris joke on mount.
 * @returns The rendered joke component, or a loading paragraph while the joke is being fetched
 */
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
