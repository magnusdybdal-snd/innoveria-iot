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

const gatewayDetails: string[] = ["Status", "Name", "EUI", "Last seen"];
const sortableColumns: GatewaySortKey[] = ["Status", "Name", "Last seen"];

export default function Gateways() {
  const [sortConfig, setSortConfig] = useState<{
    key: GatewaySortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });

  function handleSort(column: string) {
    const col = column as GatewaySortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

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
