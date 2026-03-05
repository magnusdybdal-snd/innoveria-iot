import { useEffect, useState } from "react";

import {
  GatewayInfo,
  getGateways,
  sortGateways,
  type GatewaySortKey,
  type SortDirection,
} from "@/entities/gateway";
import type { Gateway } from "@/mocks/gateways";
import { CategoryHeader } from "@/shared/ui/CategoryHeader";
import { DeviceRow } from "@/shared/ui/DeviceRow";
import { NoDeviceFoundCard } from "@/shared/ui/NoDeviceFoundCard";
import { PageContent } from "@/shared/ui/PageContent";
import { PageDivider } from "@/shared/ui/PageDivider";
import { SubPageHeader } from "@/shared/ui/SubPageHeader";
import { Menu } from "@/widgets/menu";

// Column labels rendered by CategoryHeader; order determines grid layout
const gatewayDetails: string[] = ["Status", "Name", "EUI", "Last seen"];

// EUI is display-only; excluded because the ChirpStack identifier is not a meaningful value to sort by.
const sortableColumns: GatewaySortKey[] = ["Status", "Name", "Last seen"];

/**
 * Full-page view listing all LoRaWAN gateways registered in ChirpStack.
 *
 * Fetches live gateway data from the device-service on mount and manages
 * column sort state. Delegates row rendering to GatewayRow/GatewayInfo.
 */
export default function Gateways() {
  const [gateways, setGateways] = useState<Gateway[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    getGateways().then((data) => {
      setGateways(data);
      setIsLoading(false);
    });
  }, []);

  const [sortConfig, setSortConfig] = useState<{
    key: GatewaySortKey | null;
    direction: SortDirection;
  }>({
    key: null,
    direction: "asc",
  });

  /**
   * Updates sort state when a column header is clicked.
   *
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
      <Menu />
      <PageContent>
        <SubPageHeader title="Gateways" />
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
                euid={gateway.euid}
                lastSeen={gateway.lastSeen}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NoDeviceFoundCard />}
      </PageContent>
    </div>
  );
}
