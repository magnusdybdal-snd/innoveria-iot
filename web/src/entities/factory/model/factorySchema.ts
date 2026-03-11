export interface FactoryListApiResponse {
  total_count: number;
  factories: FactoryApiResponse[];
}
export interface FactoryApiResponse {
  id: string;
  factoryId: string;
  name: string;
  address: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface CreateFactoryRequest {
  factoryId: string;
  name: string;
  address: string;
}
