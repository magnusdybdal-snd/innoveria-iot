import { useEffect, useState } from "react";

import {
  deleteGateway,
  GatewayInfo,
  getGateways,
  postGateway,
  sortGateways,
  type GatewayApiResponse,
  type GatewaySortKey,
  type SortDirection,
} from "@entities/gateway";
import { AddDevice } from "@features/addDevice";
import { formatTimestamp } from "@shared/lib";
import { CustomButton } from "@shared/ui/Button";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { NoDeviceFoundCard } from "@shared/ui/NoDeviceFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import {
  AppSnackbar,
  SNACKBAR_SEVERITY,
  useSnackbar,
} from "@shared/ui/snackbar";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

const gatewayDetails: string[] = ["Status", "Name", "EUI", "Last seen"];
const sortableColumns: GatewaySortKey[] = ["Status", "Name", "Last seen"];
const addGatewayDetails: string[] = ["Name", "DeviceEUI"];

/**
 * Full-page view listing all LoRaWAN gateways registered in database.
 *
 * Fetches live gateway data from the device-service on mount and manages
 * column sort state. Delegates row rendering to GatewayRow/GatewayInfo.
 * @returns The rendered Gateways page
 */
export default function Gateways() {
  const [gateways, setGateways] = useState<GatewayApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [addError, setAddError] = useState<string | null>(null);

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

  const fetchGateways = () => {
    getGateways().then((data) => {
      setGateways(data);
      setIsLoading(false);
    });
  };

  const handleDeleteGateway = (id: string) => {
    deleteGateway(id)
      .then(() => {
        fetchGateways();
        show("Gateway deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to delete gateway.", SNACKBAR_SEVERITY.ERROR);
      });
  };

  useEffect(() => {
    fetchGateways();
  }, []);
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
  };
  // Handler for submitting add gateway form; shows success or error snackbar based on result.
  const handleAddGateway = (gatewayData: {
    name: string;
    deviceEui: string;
  }) => {
    setAddError(null);
    return postGateway({
      companyId: "a0000000-0000-0000-0000-000000000001", // TODO: replace with real company ID from auth
      deviceEui: gatewayData.deviceEui,
      name: gatewayData.name,
    })
      .then(() => {
        fetchGateways();
        setOpenAdd(false);
        show("Gateway added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        setAddError(
          "Failed to add gateway. The EUI may already be registered.", // TODO: throw non-hardcoded error messages - based on actual error
        );
        show("Failed to add gateway", SNACKBAR_SEVERITY.ERROR);
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
        <PageDivider />
        <CategoryHeader
          categories={gatewayDetails}
          columns={gatewayDetails.length + 1}
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
                device_eui={gateway.deviceEui}
                lastSeenAt={formatTimestamp(gateway.lastSeenAt)}
                onDelete={() => handleDeleteGateway(gateway.id)}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NoDeviceFoundCard />}
      </PageContent>
      <AddDevice
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addGatewayDetails}
        onAdd={handleAddGateway}
        submitError={addError}
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
