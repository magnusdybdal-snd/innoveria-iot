import { useState } from "react";

import { type SensorSortKey, type SortDirection } from "@entities/sensor";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { Menu } from "@widgets/menu";

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
