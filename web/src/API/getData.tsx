import React, { useEffect, useState } from "react";
import { type Data, fetchApiWithErrorHandling } from "./apiClient.tsx";

const ApiViewer: React.FC = () => {
  const [data, setApi] = useState<Data | null>(null);

  useEffect(() => {
    const loadApi = async () => {
      const result = await fetchApiWithErrorHandling();
      setApi(result);
    };

    loadApi();
  }, []);

  if (!data) return <p>Loading...</p>;

  return <p>{data.totalCount}</p>;
};

export default ApiViewer;
