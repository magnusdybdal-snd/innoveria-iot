/**
 *
 */
export type Sensor = {
  id: string;
  name: string;
  status: number;
  euid: string;
  machine: string;
};

export const mockSensors: Sensor[] = [
  {
    id: "1",
    name: "Sensor 1",
    status: 0,
    euid: "A1B2C3D4E5F6",
    machine: "Machine 1",
  },
  {
    id: "2",
    name: "Sensor 2",
    status: 2,
    euid: "B2C3D4E5F6A1",
    machine: "Machine 2",
  },
  {
    id: "3",
    name: "Sensor 3",
    status: 1,
    euid: "B2D3D5E5A6A1",
    machine: "Machine 3",
  },
];
