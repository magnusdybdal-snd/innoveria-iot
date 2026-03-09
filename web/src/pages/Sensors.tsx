import { useState } from "react";

import {
  SensorAllInfoPopUp,
  SensorMainInfo,
  SensorsGenInfo,
  sortSensors,
  useSensors,
  type SensorApiResponse,
  type SensorSortKey,
  type SortDirection,
} from "@entities/sensor";
import { AddDevice } from "@features/addSensor";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { NoDeviceFoundCard } from "@shared/ui/NoDeviceFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { Menu } from "@widgets/menu";

import { mockSensors } from "@/shared/mocks/sensors";
import { CustomButton } from "@/shared/ui/Button";

const sensorMainDetails: string[] = ["Status", "Name", "Last reading"];
const addSensorDetails: string[] = [
  "Name",
  "DeviceEUI",
  "Machine",
  "Application key",
  "Sensor profile",
];
const sortableColumns: SensorSortKey[] = ["Status", "Name", "Last reading"];
type NewSensor = Omit<SensorApiResponse, "id" | "status" | "lastReading">;

/**
 * Full-page view listing all LoRaWAN sensors with sortable columns, summary statistics, and add/detail dialogs.
 * @returns The rendered Sensors page
 */
export default function Sensors() {
  const { sensors, isLoading } = useSensors();
  const [sensorsMocked, setMockSensors] =
    useState<SensorApiResponse[]>(mockSensors);

  const [openAdd, setOpenAdd] = useState(false);
  const [selectedSensor, setSelectedSensor] =
    useState<SensorApiResponse | null>(null);

  // Handler for opening and closing add sensor pop-up
  const handleClickOpenAdd = () => {
    setOpenAdd(true);
  };

  const handleCloseAdd = () => {
    setOpenAdd(false);
  };

  // Handler for opening and closing all info pop-up
  const handleRowClick = (sensor: SensorApiResponse) => {
    setSelectedSensor(sensor);
  };

  const handleCloseInfo = () => {
    setSelectedSensor(null);
  };

  const handleAddSensor = (sensorData: NewSensor): Promise<void> => {
    setMockSensors((prev) => [
      ...prev,
      {
        id: crypto.randomUUID(),
        status: 0,
        lastReading: "0 min",
        ...sensorData,
      },
    ]);
    setOpenAdd(false);
    return Promise.resolve();
  };

  const addButton = (
    <CustomButton onClick={handleClickOpenAdd}>Add sensor</CustomButton>
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
    <Menu>
      <div className="flex h-screen">
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
    </Menu>
  );
}
