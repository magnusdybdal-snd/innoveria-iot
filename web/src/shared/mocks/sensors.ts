import type { SensorApiResponse } from "@entities/sensor";

export const mockSensors: SensorApiResponse[] = [
  {
    id: "1",
    name: "Sensor 1",
    factory: "Factory 1",
    status: 0,
    deviceEui: "A1B2C3D4E5F6G7H8",
    machine: "Machine 1",
    lastReading: "2 min",
    deviceProfileId: "Milesight EM300-MLD",
  },
  {
    id: "2",
    name: "Sensor 2",
    factory: "Factory 2",
    status: 2,
    deviceEui: "B8F2X7C4Q1W9E6R3",
    machine: "Machine 2",
    lastReading: "2 min",
    deviceProfileId: "Milesight EM310-TILT",
  },
  {
    id: "3",
    name: "Sensor 3",
    factory: "Factory 3",
    status: 1,
    deviceEui: "Z4P9L2R7T6H3K8M1",
    machine: "Machine 3",
    lastReading: "50 min",
    deviceProfileId: "Milesight EM300-CL",
  },
];
