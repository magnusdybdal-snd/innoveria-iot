import { useState } from "react";

import { CategoryHeader } from "@/components/CategoryHeader";
import {
  type SensorSortKey,
  type SortDirection,
} from "@/components/gatewayRow";
import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";

const companyMainDetails: string[] = ["Name"];
const sortableColumns: SensorSortKey[] = ["Name"];

/**
 * Full-page view listing all LoRaWAN sensors with sortable columns, summary statistics, and add/detail dialogs.
 * @returns The rendered Sensors page
 */
export default function Companies() {
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

  return (
    <Menu>
      <div className="flex h-screen">
        <PageContent>
          <SubPageHeader title="Companies" />
          <CategoryHeader
            categories={companyMainDetails}
            columns={companyMainDetails.length + 2}
            sortableColumns={sortableColumns}
            sortConfig={sortConfig}
            onSort={handleSort}
          ></CategoryHeader>
        </PageContent>
      </div>
    </Menu>
  );
}
