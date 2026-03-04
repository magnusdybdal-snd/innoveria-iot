import {
  apiRequest,
  serviceClient,
  type SensorListApiResponse,
} from "@/API/apiClient";
import type { Sensor } from "@/mocks/sensors";

export const fetchSensors = async (): Promise<Sensor[]> => {
  try {
    const data = await apiRequest<SensorListApiResponse>(
      serviceClient,
      "/v1/device/sensors",
      "GET",
    );

    return data.sensors.map((sensor) => ({
      id: sensor.id,
      name: sensor.name,
      status: sensor.status,
      euid: sensor.device_eui,
      machine: sensor.machine,
      lastReading: sensor.lastReading,
      appKey: sensor.appKey,
      senProf: sensor.senProf,
    }));
  } catch (error) {
    console.error("Failed to fetch sensors:", error);
    return [];
  }
};
