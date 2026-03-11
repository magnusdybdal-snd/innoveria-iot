export interface FactoryListApiResponse {
  total_count: number;
  factories: FactoryApiResponse[];
}
export interface FactoryApiResponse {
  id: string;
  companyId: string;
  name: string;
  address: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface CreateFactoryRequest {
  companyId: string;
  name: string;
  address: string;
}
