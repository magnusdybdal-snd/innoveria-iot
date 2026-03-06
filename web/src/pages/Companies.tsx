import { useEffect, useState } from "react";

import Button from "@mui/material/Button";

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
import { SensorAllInfoPopUp } from "@/components/sensorAllInfo";
import { SensorMainInfo } from "@/components/sensorMainInfo";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";
import { mockSensors, type Sensor } from "@/mocks/sensors.ts";

const companyMainDetails: string[] = ["Name"];
const addCompanyDetails: string[] = ["Name"];
const sortableColumns: SensorSortKey[] = ["Name"];
type NewSensor = Omit<Sensor, "id" | "status" | "lastReading">;

/**
 * Full-page view listing all LoRaWAN sensors with sortable columns, summary statistics, and add/detail dialogs.
 * @returns The rendered Sensors page
 */
export default function Companies() {
  const [sensors, setSensors] = useState<Sensor[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [setMockSensors] = useState<Sensor[]>(mockSensors);

  useEffect(() => {
    fetchSensors().then((data) => {
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

  return (
    <Menu>
      <div className="flex h-screen">
        <PageContent>
          <SubPageHeader title="Companies" action={addButton} />
          <CategoryHeader
            categories={companyMainDetails}
            columns={companyMainDetails.length + 2}
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
          addOptions={addCompanyDetails}
          onAdd={handleAddSensor}
        />
      </div>
    </Menu>
  );
}
