export interface FactoryAreaListApiResponse {
  totalCount: number;
  factoryAreas: FactoryAreaApiResponse[];
}

export interface FactoryAreaApiResponse {
  id: string;
  factoryId: string;
  name: string;
  description: string | null;
  createdAt: Date;
  updatedAt: Date;
}
