import React, { useEffect, useState } from "react";

import Box from "@mui/material/Box";
import Tab from "@mui/material/Tab";
import Tabs from "@mui/material/Tabs";
import axios from "axios";

import { useFactories } from "@entities/factory";
import { useFactoryAreas } from "@entities/factoryArea";
import {
  deleteGateway,
  GatewayInfo,
  patchGateway,
  postGateway,
  sortGateways,
  useGateways,
  type GatewayApiResponse,
  type GatewaySortKey,
  type SortDirection,
  type UpdateGatewayRequest,
} from "@entities/gateway";
import { AddDevice } from "@features/addDevice";
import { EditGateway } from "@features/editGateway";
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

const gatewayDetails: string[] = ["Status", "Name", "EUI", "Last seen"];
const sortableColumns: GatewaySortKey[] = ["Status", "Name", "Last seen"];
const addGatewayDetails: string[] = [
  "Name",
  "DeviceEUI",
  "Description",
  "Factory",
  "Factory area",
];

/**
 * Full-page view listing all LoRaWAN gateways registered in database.
 *
 * Fetches live gateway data from the device-service on mount and manages
 * column sort state. Delegates row rendering to GatewayRow/GatewayInfo.
 * @returns The rendered Gateways page
 */
export default function Gateways() {
  const { gateways, isLoading, refetch } = useGateways();
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
  const [addError, setAddError] = useState<string | null>(null);
  const [editingGateway, setEditGateway] = useState<GatewayApiResponse | null>(
    null,
  );
  const [editError, setEditError] = useState<string | null>(null);

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

  useEffect(() => {
    if (areasError)
      show("Failed to load factory areas", SNACKBAR_SEVERITY.ERROR);
  }, [areasError, show]);

  // Factory tabs
  const [tabValue, setTabValue] = useState<number | string>(0);

  const handleTabChange = (
    _event: React.SyntheticEvent,
    newValue: number | string,
  ) => {
    setTabValue(newValue);
  };

  const handleDeleteGateway = (id: string) => {
    deleteGateway(id)
      .then(() => {
        refetch();
        show("Gateway deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to delete gateway:", err);
        show("Failed to delete gateway.", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const [openAdd, setOpenAdd] = useState(false);
  const [sortConfig, setSortConfig] = useState<{
    key: GatewaySortKey | null;
    direction: SortDirection;
  }>({
    key: null,
    direction: "asc",
  });

  // Handler for opening and closing add gateway pop-up
  const handleClickOpenAdd = () => {
    setOpenAdd(true);
  };
  const handleCloseAdd = () => {
    setOpenAdd(false);
    setAddError(null);
    setSelectedFactoryId(undefined);
  };
  // Handler for submitting add gateway form; shows success or error snackbar based on result.
  const handleAddGateway = (gatewayData: {
    name: string;
    deviceEui: string;
    description: string;
    factory: string;
    factoryArea: string;
  }) => {
    setAddError(null);
    return postGateway({
      gatewayEui: gatewayData.deviceEui,
      name: gatewayData.name,
      description: gatewayData.description || undefined,
      factoryId: gatewayData.factory,
      factoryAreaId: gatewayData.factoryArea,
    })
      .then(() => {
        refetch();
        setOpenAdd(false);
        show("Gateway added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to add gateway:", err);
        setAddError(
          "Failed to add gateway. The EUI may already be registered.", // TODO: throw non-hardcoded error messages - based on actual error
        );
        show("Failed to add gateway", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const handleEditGateway = (
    id: string,
    payload: UpdateGatewayRequest,
  ): Promise<void> => {
    setEditError(null);
    return patchGateway(id, payload)
      .then(() => {
        refetch();
        setEditGateway(null);
        show("Gateway updated successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to edit gateway:", err);
        const message =
          axios.isAxiosError(err) && err.response?.data?.message
            ? (err.response.data.message as string)
            : "Failed to update gateway. Please try again.";
        setEditError(message);
        show("Failed to update gateway", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const addButton = (
    <CustomButton onClick={handleClickOpenAdd}>Add gateway</CustomButton>
  );

  /**
   * Updates sort state when a column header is clicked.
   * @param column - The column label passed up from CategoryHeader's onSort callback
   */
  function handleSort(column: string) {
    const col = column as GatewaySortKey;
    setSortConfig((prev) => {
      if (prev.key !== col) return { key: col, direction: "asc" };
      if (prev.direction === "asc") return { key: col, direction: "desc" };
      return { key: null, direction: "asc" }; // third click resets to initial sorting (unsorted?) //TODO check if default is unsorted or sorted when connected to API.
    });
  }

  // Derive sorted list on every render; sortGateways returns a new array and does not mutate gateways
  const sorted = sortGateways(gateways, sortConfig.key, sortConfig.direction);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Gateways" action={addButton} />
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
            {factories.map((factory) => (
              <Tab label={factory.name} value={factory.id} />
            ))}
          </Tabs>
        </Box>
        <CategoryHeader
          categories={gatewayDetails}
          gridTemplateColumns="auto 55ch auto auto auto auto"
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {isLoading && <p>Loading...</p>}{" "}
          {/*TODO: make a better looking loading indicator */}
          {sorted.map((gateway) => (
            <DeviceRow key={gateway.id}>
              <GatewayInfo
                name={gateway.name}
                status={gateway.status}
                device_eui={gateway.gatewayEui}
                lastSeenAt={formatTimestamp(gateway.lastSeenAt)}
                description={gateway.description}
                onDelete={() => handleDeleteGateway(gateway.id)}
                onEdit={() => {
                  setEditSelectedFactoryId(gateway.factoryId);
                  setEditGateway(gateway);
                }}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NotFoundCard page="gateways" />}
      </PageContent>
      <AddDevice
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addGatewayDetails}
        onAdd={handleAddGateway}
        submitError={addError}
        factoryOptions={factories}
        factoryAreaOptions={factoryAreas}
        onFactoryChange={setSelectedFactoryId}
      />
      {editingGateway && (
        <EditGateway
          open={!!editingGateway}
          gateway={editingGateway}
          onClose={() => {
            setEditGateway(null);
            setEditError(null);
            setEditSelectedFactoryId(undefined);
          }}
          onEdit={handleEditGateway}
          factoryOptions={factories}
          factoryAreaOptions={editFactoryAreas}
          onFactoryChange={setEditSelectedFactoryId}
          onErrorClear={() => setEditError(null)}
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
