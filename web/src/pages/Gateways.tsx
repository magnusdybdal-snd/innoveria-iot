import { useState } from "react";

import { CategoryHeader } from "@/components/CategoryHeader";
import { GatewayInfo } from "@/components/gatewayInfo";
import {
  GatewayRow,
  sortGateways,
  type GatewaySortKey,
  type SortDirection,
} from "@/components/gatewayRow";
import { PageContent } from "@/components/pageContent";
import { PageDivider } from "@/components/pageDivider";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu";
import { mockGateways } from "@/mocks/gateways";

// Column labels rendered by CategoryHeader; order determines grid layout
const gatewayDetails: string[] = ["Status", "Name", "EUI", "Last seen"];

// EUI is display-only; excluded because the ChirpStack identifier is not a meaningful value to sort by.
const sortableColumns: GatewaySortKey[] = ["Status", "Name", "Last seen"];

/**
 * Full-page view listing all LoRaWAN gateways registered in ChirpStack.
 *
 * Manages column sort state and delegates rendering to GatewayRow/GatewayInfo.
 * Data is currently sourced from mock fixtures; // TODO: replace with a live API call when the collection-service gateway endpoint is available.
 */
export default function Gateways() {
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
    // Clicking the active column toggles direction; clicking a new column resets to ascending
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  // Derive sorted list on every render; sortGateways returns a new array and does not mutate mockGateways
  const sorted = sortGateways(
    mockGateways,
    sortConfig.key,
    sortConfig.direction,
  );

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
          {sorted.map((gateway) => (
            <GatewayRow key={gateway.id}>
              <GatewayInfo
                name={gateway.name}
                online={gateway.online}
                euid={gateway.euid}
                lastSeen={gateway.lastSeen}
              />
            </GatewayRow>
          ))}
        </CategoryHeader>
      </PageContent>
    </div>
  );
}
