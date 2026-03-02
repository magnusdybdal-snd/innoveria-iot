/**
 *
 */
export type Sensor = {
  id: string;
  name: string;
  status: number;
  euid: string;
  machine: string;
  lastReading: string;
  appKey: string;
  devProf: string;
};

export const mockSensors: Sensor[] = [
  {
    id: "1",
    name: "Sensor 1",
    status: 0,
    euid: "A1B2C3D4E5F6G7H8",
    machine: "Machine 1",
    lastReading: "2 min",
    appKey: "a9k3m7x2c5b8d1q4f6h2j9l3p7r5t8v",
    devProf: "TemperatureSensor",
  },
  {
    id: "2",
    name: "Sensor 2",
    status: 2,
    euid: "B8F2X7C4Q1W9E6R3",
    machine: "Machine 2",
    lastReading: "2 min",
    appKey: "3f8k1a7m2q5c9x4d6b8h2j1l7p3r5t9",
    devProf: "HumiditySensor",
  },
  {
    id: "3",
    name: "Sensor 3",
    status: 1,
    euid: "Z4P9L2R7T6H3K8M1",
    machine: "Machine 3",
    lastReading: "50 min",
    appKey: "d7m3a1x9c5k8q2b4f6h1j8l3p7r2t5v",
    devProf: "IndustrialVibrationSensor",
  },
];
