import { useEffect, useState } from "react";

import {
  deleteFactory,
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

import { postFactory } from "@/entities/factory/api/postFactory";
import { AddEntityDialog } from "@/shared/ui/AddEntityDialog";
import { CustomButton } from "@/shared/ui/Button";
import { NoDeviceFoundCard } from "@/shared/ui/NoDeviceFoundCard";

//TODO: Add post request functionality through button and schema

const factoryDetails: string[] = [
  "Name",
  "Address",
  "Created at",
  "Updated at",
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

  // Adding a new factory
  const [openAdd, setOpenAdd] = useState(false);
  const [addError, setAddError] = useState<string | null>(null);
  const handleClickOpenAdd = () => {
    setOpenAdd(true);
  };
  const handleCloseAdd = () => {
    setOpenAdd(false);
    setAddError(null);
  };
  const addButton = (
    <CustomButton onClick={handleClickOpenAdd}>Add factory</CustomButton>
  );
  const handleAddFactory = (factoryData: {
    factoryId: string;
    name: string;
    address: string;
  }) => {
    setAddError(null);
    return postFactory({
      companyId: "a0000000-0000-0000-0000-000000000001", // TODO: replace with real company ID from auth
      name: factoryData.name,
      address: factoryData.address,
    })
      .then(() => {
        fetchFactories();
        setOpenAdd(false);
        //show("Gateway added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        setAddError(
          "Failed to add gateway.", // TODO: throw non-hardcoded error messages - based on actual error
        );
        //show("Failed to add gateway", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const fetchFactories = () => {
    getFactories().then((data) => {
      setFactories(data);
      setIsLoading(false);
    });
  };

  //TODO: impement snackbar.
  //TODO: use deletion confirmation dialog when it has been implemented.
  const handleDeleteFactory = (id: string) => {
    deleteFactory(id)
      .then(() => {
        fetchFactories();
        //show("Factory deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        //show("Failed to delete factory", SNACKBAR_SEVERITY.ERROR);
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
        <SubPageHeader title="Factories" action={addButton} />
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
                name={factory.name}
                address={factory.address}
                createdAt={formatTimestamp(factory.createdAt.toString())}
                updatedAt={formatTimestamp(factory.updatedAt.toString())}
                onDelete={() => handleDeleteFactory(factory.id)}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NoDeviceFoundCard />}
        <AddEntityDialog
          open={openAdd}
          title="Add factory"
          fields={["Name", "Address"]}
          onClose={handleCloseAdd}
          onSubmit={(values) =>
            handleAddFactory({
              name: values["Name"],
              factoryId: values["Factory ID"],
              address: values["Address"],
            })
          }
          submitError={addError}
        />
      </PageContent>
    </div>
  );
}
