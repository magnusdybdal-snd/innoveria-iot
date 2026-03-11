import { useEffect, useState } from "react";

import {
  getFactories,
  sortFactories,
  type FactoryApiResponse,
  type FactorySortKey,
  type SortDirection,
} from "@entities/factory";
import { FactoryInfo } from "@entities/factory/ui/FactoryInfo";
import { formatTimestamp } from "@shared/lib";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

import { NoDeviceFoundCard } from "@/shared/ui/NoDeviceFoundCard";

//TODO: Add post request functionality through button and schema

const factoryDetails: string[] = [
  "ID",
  "Name",
  "Address",
  "Created",
  "Last update",
];

const sortableColumns: FactorySortKey[] = [
  "Name",
  "Address",
  "Created at",
  "Updated at",
];

/**
 * Full-page view listing all factories registered on the site.
 * @returns The rendered Factories page
 */
export default function Factories() {
  const [factories, setFactories] = useState<FactoryApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [sortConfig, setSortConfig] = useState<{
    key: FactorySortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });

  const fetchFactories = () => {
    getFactories().then((data) => {
      setFactories(data);
      setIsLoading(false);
    });
  };

  useEffect(() => {
    fetchFactories();
  }, []);

  function handleSort(column: string) {
    const col = column as FactorySortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  const sorted = sortFactories(factories, sortConfig.key, sortConfig.direction);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Factories" />
        <PageDivider />
        <CategoryHeader
          categories={factoryDetails}
          columns={factoryDetails.length + 1}
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {isLoading && <p>Loading...</p>}
          {/*TODO: make a better looking loading indicator */}
          {sorted.map((factory) => (
            <DeviceRow key={factory.id}>
              <FactoryInfo
                id={factory.id}
                name={factory.name}
                address={factory.address}
                createdAt={formatTimestamp(factory.createdAt.toString())}
                updatedAt={formatTimestamp(factory.updatedAt.toString())}
                addFactory={() => {}} // TODO: implement add user flow
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NoDeviceFoundCard />}
      </PageContent>
    </div>
  );
}
