import { useState } from "react";

import Button from "@mui/material/Button";

import { AddDevice } from "@/components/addDevicePopup";
import { CategoryHeader } from "@/components/CategoryHeader";
import {
  DeviceRow,
  sortSensors,
  type SensorSortKey,
  type SortDirection,
} from "@/components/gatewayRow";
import { PageContent } from "@/components/pageContent";
import { PageDivider } from "@/components/pageDivider";
import { SensorInfo } from "@/components/sensorInfo";
import { SensorsGenInfo } from "@/components/sensorsGenInfo";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";
import { mockSensors } from "@/mocks/sensors.ts";

const sensorDetails: string[] = ["Status", "Name", "DebEUI", "Machine"];
const addSensorDetails: string[] = ["Name", "DebEUI", "Machine"];
const sortableColumns: SensorSortKey[] = ["Status", "Name", "Machine"];

export default function Sensors() {
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
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

  const sorted = sortSensors(mockSensors, sortConfig.key, sortConfig.direction);

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
          columns={sensorDetails.length + 1}
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
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
      </PageContent>

      <AddDevice
        open={open}
        onClose={handleClose}
        addOptions={addSensorDetails}
      />
    </div>
  );
}
