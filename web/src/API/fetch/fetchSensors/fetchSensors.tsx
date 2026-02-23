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
      "/api/v1/device/sensors",
      "GET",
    );

    return data.sensors.map((gw) => ({
      id: gw.id,
      name: gw.name,
      status: gw.status,
      euid: gw.device_eui,
      machine: gw.machine,
    }));
  } catch (error) {
    console.error("Failed to fetch sensors:", error);
    return [];
  }
};
