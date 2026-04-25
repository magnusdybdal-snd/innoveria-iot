export interface FactoryListApiResponse {
  totalCount: number;
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
  name: string;
  address: string;
}
