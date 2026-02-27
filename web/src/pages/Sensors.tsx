import { useEffect, useState } from "react";

import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import { fetchSensors } from "@/API/fetch/fetchSensors";
import { AddDevice } from "@/components/addDevicePopup";
import { CategoryHeader } from "@/components/CategoryHeader";
import {
  DeviceRow,
  sortSensors,
  type SensorSortKey,
  type SortDirection,
} from "@/components/gatewayRow";
import { NoDeviceFoundCard } from "@/components/noDeviceFoundCard";
import { PageContent } from "@/components/pageContent";
import { PageDivider } from "@/components/pageDivider";
import { SensorInfo } from "@/components/sensorInfo";
import { SensorsGenInfo } from "@/components/sensorsGenInfo";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";
import { mockSensors, type Sensor } from "@/mocks/sensors.ts";

const sensorDetails: string[] = [
  "Status",
  "Name",
  "DeviceEUI",
  "Machine",
  "Last reading",
  "Application key",
  "Device profile",
];
const addSensorDetails: string[] = [
  "Name",
  "DeviceEUI",
  "Machine",
  "Application key",
  "Device profile",
];
const sortableColumns: SensorSortKey[] = [
  "Status",
  "Name",
  "Machine",
  "Last reading",
  "Application key",
  "Device profile",
];
type NewSensor = Omit<Sensor, "id" | "status" | "lastReading">;

export default function Sensors() {
  const [sensors, setSensors] = useState<Sensor[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [sensorsMocked, setMockSensors] = useState<Sensor[]>(mockSensors);
  const [show, setShow] = useState(false);

  useEffect(() => {
    fetchSensors().then((data) => {
      setSensors(data);
      setIsLoading(false);
    });
  }, []);

  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleAddSensor = (sensorData: NewSensor) => {
    setMockSensors((prev) => [
      ...prev,
      {
        id: crypto.randomUUID(),
        status: 0,
        lastReading: "0 min",
        ...sensorData,
      },
    ]);
  };

  const addButton = (
    <Button
      variant="outlined"
      sx={{
        backgroundColor: "primary.main",
        color: "primary.dark",
        "&:hover": { backgroundColor: "primary.main" },
        borderRadius: 2,
        textTransform: "none",
        fontSize: 20,
      }}
      onClick={handleClickOpen}
    >
      Add device +
    </Button>
  );

  const [sortConfig, setSortConfig] = useState<{
    key: SensorSortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });

  function handleSort(column: string) {
    const col = column as SensorSortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  const sorted = sortSensors(sensors, sortConfig.key, sortConfig.direction);
  const sortedMock = sortSensors(
    sensorsMocked,
    sortConfig.key,
    sortConfig.direction,
  );

  const sensorInfos = new Map<string, number>();
  sensorInfos.set("Total sensors", sorted.length);
  sensorInfos.set(
    "Online sensors",
    sorted.filter((sensor) => sensor.status === 0).length,
  );
  sensorInfos.set(
    "Offline sensors",
    sorted.filter((sensor) => sensor.status === 2).length,
  );
  sensorInfos.set("Last seen 24hr", 0);
  sensorInfos.set("Error last 24hr", 0);

  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader title="Sensor devices" action={addButton} />
        <div className="flex justify-between flex-wrap">
          {Array.from(sensorInfos.entries()).map(([key, value]) => (
            <SensorsGenInfo key={key} title={key} count={value} />
          ))}
        </div>
        <PageDivider />
        <CategoryHeader
          categories={sensorDetails}
          columns={sensorDetails.length + 2}
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {sorted.map((sensor) => (
            <DeviceRow key={sensor.id}>
              <SensorInfo
                name={sensor.name}
                status={sensor.status}
                euid={sensor.euid}
                machine={sensor.machine}
                lastReading={sensor.lastReading}
                appKey={sensor.appKey}
                devProf={sensor.devProf}
              />
            </DeviceRow>
          ))}
          {sortedMock.map((sensor) => (
            <DeviceRow key={sensor.id}>
              <SensorInfo
                name={sensor.name}
                status={sensor.status}
                euid={sensor.euid}
                machine={sensor.machine}
                lastReading={sensor.lastReading}
                appKey={sensor.appKey}
                devProf={sensor.devProf}
              />
              <button onClick={() => setShow((prev) => !prev)}>Expand</button>
              {show && <Typography>This is your component</Typography>}
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NoDeviceFoundCard />}
      </PageContent>

      <AddDevice
        open={open}
        onClose={handleClose}
        addOptions={addSensorDetails}
        onAdd={handleAddSensor}
      />
    </div>
  );
}
