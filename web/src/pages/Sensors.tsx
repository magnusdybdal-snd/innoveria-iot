import React, { useEffect, useState } from "react";

import {
  getSensorProfiles,
  postSensor,
  SensorAllInfoPopUp,
  SensorMainInfo,
  SensorsGenInfo,
  sortSensors,
  useSensors,
  type SensorApiResponse,
  type SensorProfileApiResponse,
  type SensorSortKey,
  type SortDirection,
} from "@entities/sensor";
import { deleteSensor } from "@entities/sensor/api/deleteSensor";
import { AddDevice } from "@features/addDevice";
import Box from "@mui/material/Box";
import Tab from "@mui/material/Tab";
import Tabs from "@mui/material/Tabs";
import { formatTimestamp } from "@shared/lib";
import { CustomButton } from "@shared/ui/Button";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { NotFoundCard } from "@shared/ui/NotFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import {
  AppSnackbar,
  SNACKBAR_SEVERITY,
  useSnackbar,
} from "@shared/ui/snackbar";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

import { getFactories, type FactoryApiResponse } from "@/entities/factory";

const sensorMainDetails: string[] = ["Status", "Name", "Last reading"];
const addSensorDetails: string[] = [
  "Name",
  "DeviceEUI",
  "Factory",
  "Machine",
  "Application key",
  "Sensor profile",
];
const sortableColumns: SensorSortKey[] = [
  "Status",
  "Factory",
  "Name",
  "Last reading",
];

/**
 * Full-page view listing all LoRaWAN sensors with sortable columns, summary statistics, and add/detail dialogs.
 * @returns The rendered Sensors page
 */
export default function Sensors() {
  const { sensors, isLoading, refetch } = useSensors();
  const [openAdd, setOpenAdd] = useState(false);
  const [addError, setAddError] = useState<string | null>(null);
  const [selectedSensor, setSelectedSensor] =
    useState<SensorApiResponse | null>(null);
  const [sensorProfiles, setSensorProfiles] = useState<
    SensorProfileApiResponse[]
  >([]);
  const [factory, setFactory] = useState<FactoryApiResponse[]>([]); // factory location sensor

  const { show, hide, snackbar } = useSnackbar();
  const [tabValue, setTabValue] = useState<number | string>(0);

  const handleTabChange = (
    _event: React.SyntheticEvent,
    newValue: number | string,
  ) => {
    setTabValue(newValue);
  };

  useEffect(() => {
    getSensorProfiles().then(setSensorProfiles);
  }, []);

  useEffect(() => {
    getFactories().then(setFactory);
  }, []);

  // Handler for deleting a sensor; refreshes list on success
  const handleDeleteSensor = (id: string) => {
    deleteSensor(id)
      .then(() => {
        refetch();
        show("Sensor deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to delete sensor.", SNACKBAR_SEVERITY.ERROR);
      });
  };

  // Handler for opening and closing add sensor pop-up
  const handleClickOpenAdd = () => {
    setOpenAdd(true);
  };

  const handleCloseAdd = () => {
    setOpenAdd(false);
    setAddError(null);
  };

  const handleAddSensor = (sensorData: {
    name: string;
    deviceEui: string;
    factory: string;
    machine: string;
    appKey: string;
    senProf: string;
  }): Promise<void> => {
    setAddError(null);
    return postSensor({
      companyId: "a0000000-0000-0000-0000-000000000001", // TODO: replace with real company ID from auth
      factoryId: sensorData.factory,
      deviceEui: sensorData.deviceEui,
      sensorProfileId: sensorData.senProf,
      appKey: sensorData.appKey,
      name: sensorData.name,
    })
      .then(() => {
        refetch();
        setOpenAdd(false);
        show("Sensor added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        setAddError("Something went wrong adding sensor"); // TODO: improve error handling with specific messages based on error type
        show("Failed to add sensor.", SNACKBAR_SEVERITY.ERROR);
        throw err;
      });
  };

  // Handler for opening and closing all info pop-up
  const handleRowClick = (sensor: SensorApiResponse) => {
    setSelectedSensor(sensor);
  };

  const handleCloseInfo = () => {
    setSelectedSensor(null);
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

  // Filter sensors to only show thos from chosen factory
  const filteredSensors = sorted.filter((sensor) => {
    if (tabValue === 0) return true;
    return sensor.factory === tabValue;
  });

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
      <PageContent>
        <SubPageHeader title="Sensor devices" action={addButton} />
        <div className="flex justify-between flex-wrap">
          {Array.from(sensorInfos.entries()).map(([key, value]) => (
            <SensorsGenInfo key={key} title={key} count={value} />
          ))}
        </div>
        <Box
          sx={{
            borderBottom: 1,
            borderColor: "divider",
            marginTop: 2,
            marginBottom: 5,
          }}
        >
          <Tabs
            value={tabValue}
            onChange={handleTabChange}
            variant="scrollable"
            scrollButtons="auto"
            aria-label="scrollable auto tabs example"
          >
            <Tab label="All" value={0} />
            {factory.map((factory) => (
              <Tab label={factory.name} value={factory.id} />
            ))}
          </Tabs>
        </Box>
        <CategoryHeader
          categories={sensorMainDetails}
          columns={sensorMainDetails.length + 2}
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {filteredSensors.map((sensor) => (
            <DeviceRow key={sensor.id}>
              <SensorMainInfo
                name={sensor.name}
                status={sensor.status}
                lastReading={formatTimestamp(sensor.lastReading)}
                onClick={() => handleRowClick(sensor)}
                onDelete={() => handleDeleteSensor(sensor.id)}
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
        {!isLoading && filteredSensors.length === 0 && (
          <NotFoundCard page="sensors" isEmpty={sorted.length === 0} />
        )}
      </PageContent>

      <AddDevice
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addSensorDetails}
        profileOptions={sensorProfiles}
        onAdd={handleAddSensor}
        submitError={addError}
        factoryOptions={factory}
      />
      <AppSnackbar
        open={snackbar?.open ?? false}
        message={snackbar?.message ?? ""}
        severity={snackbar?.severity}
        onClose={hide}
      />
    </div>
  );
}
