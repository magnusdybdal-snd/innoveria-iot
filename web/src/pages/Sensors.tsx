import { useEffect, useState } from "react";

import Button from "@mui/material/Button";

import {
  getSensors,
  SensorAllInfoPopUp,
  SensorMainInfo,
  SensorsGenInfo,
  sortSensors,
  type SensorSortKey,
  type SortDirection,
} from "@/entities/sensor";
import { AddDevice } from "@/features/addSensor";
import { mockSensors, type Sensor } from "@/mocks/sensors.ts";
import { CategoryHeader } from "@/shared/ui/CategoryHeader";
import { DeviceRow } from "@/shared/ui/DeviceRow";
import { NoDeviceFoundCard } from "@/shared/ui/NoDeviceFoundCard";
import { PageContent } from "@/shared/ui/PageContent";
import { PageDivider } from "@/shared/ui/PageDivider";
import { SubPageHeader } from "@/shared/ui/SubPageHeader";
import { Menu } from "@/widgets/menu";

const sensorMainDetails: string[] = ["Status", "Name", "Last reading"];
const addSensorDetails: string[] = [
  "Name",
  "DeviceEUI",
  "Machine",
  "Application key",
  "Sensor profile",
];
const sortableColumns: SensorSortKey[] = ["Status", "Name", "Last reading"];
type NewSensor = Omit<Sensor, "id" | "status" | "lastReading">;

export default function Sensors() {
  const [sensors, setSensors] = useState<Sensor[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [sensorsMocked, setMockSensors] = useState<Sensor[]>(mockSensors);

  useEffect(() => {
    getSensors().then((data) => {
      setSensors(data);
      setIsLoading(false);
    });
  }, []);

  const [openAdd, setOpenAdd] = useState(false);
  const [selectedSensor, setSelectedSensor] = useState<Sensor | null>(null);

  // Handler for opening and closing add sensor pop-up
  const handleClickOpenAdd = () => {
    setOpenAdd(true);
  };

  const handleCloseAdd = () => {
    setOpenAdd(false);
  };

  // Handler for opening and closing all info pop-up
  const handleRowClick = (sensor: Sensor) => {
    setSelectedSensor(sensor);
  };

  const handleCloseInfo = () => {
    setSelectedSensor(null);
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
      onClick={handleClickOpenAdd}
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
          categories={sensorMainDetails}
          columns={sensorMainDetails.length + 2}
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {sorted.map((sensor) => (
            <DeviceRow key={sensor.id}>
              <SensorMainInfo
                name={sensor.name}
                status={sensor.status}
                lastReading={sensor.lastReading}
                onClick={() => handleRowClick(sensor)}
              />
            </DeviceRow>
          ))}
          {sortedMock.map((sensor) => (
            <DeviceRow key={sensor.id}>
              <SensorMainInfo
                name={sensor.name}
                status={sensor.status}
                lastReading={sensor.lastReading}
                onClick={() => handleRowClick(sensor)}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {selectedSensor && (
          <SensorAllInfoPopUp
            open={true}
            onClose={handleCloseInfo}
            sensor={selectedSensor}
          />
        )}
        {!isLoading && sorted.length === 0 && <NoDeviceFoundCard />}
      </PageContent>

      <AddDevice
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addSensorDetails}
        onAdd={handleAddSensor}
      />
    </div>
  );
}
