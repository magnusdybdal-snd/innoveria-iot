import React, { useEffect, useMemo, useState } from "react";

import Box from "@mui/material/Box";
import Tab from "@mui/material/Tab";
import Tabs from "@mui/material/Tabs";
import axios from "axios";

import { useFactories } from "@entities/factory";
import { useFactoryAreas } from "@entities/factoryArea";
import { useProductionResources } from "@entities/productionResource";
import {
  deleteSensor,
  patchSensor,
  postSensor,
  SensorAllInfoPopUp,
  SensorMainInfo,
  SensorsGenInfo,
  sortSensors,
  useSensorProfiles,
  useSensors,
  type SensorApiResponse,
  type SensorSortKey,
  type SortDirection,
  type UpdateSensorRequest,
} from "@entities/sensor";
import { AddDevice } from "@features/addDevice";
import { EditSensor } from "@features/editSensor";
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

const sensorMainDetails: string[] = ["Status", "Name", "Last reading"];
const addSensorDetails: string[] = [
  "Name",
  "DeviceEUI",
  "Description",
  "Factory",
  "Factory area",
  "Production resource",
  "Application key",
  "Sensor profile",
  "Electricity sensor",
  "Voltage",
];
const sortableColumns: SensorSortKey[] = [
  "Status",
  "Factory",
  "Name",
  "Last reading",
];
const voltageOptions = [
  { id: "230", name: "230V" },
  { id: "400", name: "400V" },
];

/**
 * Full-page view listing all LoRaWAN sensors with sortable columns, summary statistics, and add/detail dialogs.
 * @returns The rendered Sensors page
 */
export default function Sensors() {
  const { sensors, isLoading, refetch } = useSensors();
  const { factories } = useFactories();
  const [selectedFactoryId, setSelectedFactoryId] = useState<
    string | undefined
  >();
  const { factoryAreas, error: areasError } =
    useFactoryAreas(selectedFactoryId);
  const [editSelectedFactoryId, setEditSelectedFactoryId] = useState<
    string | undefined
  >();
  const { factoryAreas: editFactoryAreas } = useFactoryAreas(
    editSelectedFactoryId,
  );
  const { sensorProfiles } = useSensorProfiles();
  const { productionResources } = useProductionResources();
  const productionResourceOptions = useMemo(
    () => [
      { id: "", name: "No machine" },
      ...productionResources.map((r) => ({
        id: String(r.id),
        name: `${r.number} – ${r.description}`,
      })),
    ],
    [productionResources],
  );
  const [openAdd, setOpenAdd] = useState(false);
  const [addError, setAddError] = useState<string | null>(null);
  const [selectedSensor, setSelectedSensor] =
    useState<SensorApiResponse | null>(null);
  const [editingSensor, setEditingSensor] = useState<SensorApiResponse | null>(
    null,
  );
  const [editError, setEditError] = useState<string | null>(null);

  const { show, hide, snackbar } = useSnackbar();

  useEffect(() => {
    if (areasError)
      show("Failed to load factory areas", SNACKBAR_SEVERITY.ERROR);
  }, [areasError, show]);

  const [tabValue, setTabValue] = useState<number | string>(0);

  const handleTabChange = (
    _event: React.SyntheticEvent,
    newValue: number | string,
  ) => {
    setTabValue(newValue);
  };

  // Handler for deleting a sensor; refreshes list on success
  const handleDeleteSensor = (id: string) => {
    deleteSensor(id)
      .then(() => {
        refetch();
        show("Sensor deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to delete sensor:", err);
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
    setSelectedFactoryId(undefined);
  };

  const handleAddSensor = (sensorData: {
    name: string;
    deviceEui: string;
    description: string;
    electricitySensor: boolean;
    factory: string;
    factoryArea: string;
    productionResource: number | null;
    appKey: string;
    senProf: string;
    voltage: number | null;
  }): Promise<void> => {
    setAddError(null);
    return postSensor({
      electricitySensor: sensorData.electricitySensor,
      factoryId: sensorData.factory,
      factoryAreaId: sensorData.factoryArea,
      deviceEui: sensorData.deviceEui,
      sensorProfileId: sensorData.senProf,
      productionResource: sensorData.productionResource,
      appKey: sensorData.appKey,
      name: sensorData.name,
      description: sensorData.description || undefined,
      voltage: sensorData.voltage,
    })
      .then(() => {
        refetch();
        setOpenAdd(false);
        show("Sensor added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to add sensor:", err);
        const message =
          axios.isAxiosError(err) && err.response?.data?.message
            ? (err.response.data.message as string)
            : "Something went wrong adding sensor.";
        setAddError(message);
        show("Failed to add sensor.", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const handleEditSensor = (
    id: string,
    payload: UpdateSensorRequest,
  ): Promise<void> => {
    setEditError(null);
    return patchSensor(id, payload)
      .then(() => {
        refetch();
        setEditingSensor(null);
        show("Sensor updated successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to update sensor:", err);
        const message =
          axios.isAxiosError(err) && err.response?.data?.message
            ? (err.response.data.message as string)
            : "Failed to update sensor.";
        setEditError(message);
        show("Failed to update sensor.", SNACKBAR_SEVERITY.ERROR);
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

  // Filter sensors to only show those from chosen factory
  const filteredSensors = sorted.filter((sensor) => {
    if (tabValue === 0) return true;
    return sensor.factory === tabValue;
  });

  const sensorInfos = new Map<string, number>();
  sensorInfos.set("Total sensors", filteredSensors.length);
  sensorInfos.set(
    "Online sensors",
    filteredSensors.filter((sensor) => sensor.status === 0).length,
  );
  sensorInfos.set(
    "Offline sensors",
    filteredSensors.filter((sensor) => sensor.status === 2).length,
  );
  sensorInfos.set("Last seen 24hr", 0);
  sensorInfos.set("Error last 24hr", 0);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Sensor devices" action={addButton} />
        <Box
          sx={{
            borderBottom: 1,
            borderColor: "divider",
            marginTop: 2,
            marginBottom: 2,
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
            {factories.map((factory) => (
              <Tab label={factory.name} value={factory.id} />
            ))}
          </Tabs>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            flexWrap: "wrap",
            marginBottom: 5,
          }}
        >
          {Array.from(sensorInfos.entries()).map(([key, value]) => (
            <SensorsGenInfo key={key} title={key} count={value} />
          ))}
        </Box>
        <CategoryHeader
          categories={sensorMainDetails}
          gridTemplateColumns="auto 55ch auto auto max-content auto"
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
                description={sensor.description}
                onClick={() => handleRowClick(sensor)}
                onDelete={() => handleDeleteSensor(sensor.id)}
                onEdit={() => {
                  setEditSelectedFactoryId(sensor.factory);
                  setEditingSensor(sensor);
                }}
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
          <NotFoundCard
            page="Sensors"
            action={sorted.length === 0 ? "found" : "registered"}
          />
        )}
      </PageContent>

      <AddDevice
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addSensorDetails}
        profileOptions={sensorProfiles}
        voltageOptions={voltageOptions}
        onAdd={handleAddSensor}
        submitError={addError}
        factoryOptions={factories}
        factoryAreaOptions={factoryAreas}
        onFactoryChange={setSelectedFactoryId}
        productionResourceOptions={productionResourceOptions}
      />
      {editingSensor && (
        <EditSensor
          open={!!editingSensor}
          sensor={editingSensor}
          onClose={() => {
            setEditingSensor(null);
            setEditError(null);
            setEditSelectedFactoryId(undefined);
          }}
          onEdit={handleEditSensor}
          factoryOptions={factories}
          factoryAreaOptions={editFactoryAreas}
          sensorProfileOptions={sensorProfiles}
          onErrorClear={() => setEditError(null)}
          productionResourceOptions={productionResourceOptions}
          voltageOptions={voltageOptions}
          onFactoryChange={setEditSelectedFactoryId}
          submitError={editError}
        />
      )}
      <AppSnackbar
        open={snackbar?.open ?? false}
        message={snackbar?.message ?? ""}
        severity={snackbar?.severity}
        onClose={hide}
      />
    </div>
  );
}
