import { useEffect, useState } from "react";

import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";

import {
  deleteFactory,
  sortFactories,
  useFactories,
  type FactoryApiResponse,
  type FactorySortKey,
  type SortDirection,
} from "@entities/factory";
import { postFactory } from "@entities/factory/api/postFactory";
import { FactoryInfo } from "@entities/factory/ui/FactoryInfo";
import {
  deleteFactoryArea,
  postFactoryArea,
  useFactoryAreas,
  type FactoryAreaApiResponse,
} from "@entities/factoryArea";
import { FactoryAreaInfo } from "@entities/factoryArea/ui/FactoryAreaInfo";
import { formatTimestamp } from "@shared/lib";
import { AddEntityDialog } from "@shared/ui/AddEntityDialog";
import { CustomButton } from "@shared/ui/Button";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { NotFoundCard } from "@shared/ui/NotFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import {
  AppSnackbar,
  SNACKBAR_SEVERITY,
  useSnackbar,
} from "@shared/ui/snackbar";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

const factoryColumns: string[] = [
  "Name",
  "Address",
  "Created at",
  "Updated at",
];
const sortableFactoryColumns: FactorySortKey[] = [
  "Name",
  "Address",
  "Created at",
  "Updated at",
];

const areaColumns: string[] = ["Name", "Description"];

/**
 * Full-page view for factories and their areas.
 *
 * Default view lists all factories. Clicking a factory drills down into
 * that factory's areas where users can add or delete areas.
 * @returns The rendered Factories page
 */
export default function Factories() {
  const [selectedFactory, setSelectedFactory] =
    useState<FactoryApiResponse | null>(null);

  const {
    factories,
    isLoading: factoriesLoading,
    error: factoriesError,
    refetch: refetchFactories,
  } = useFactories();
  const {
    factoryAreas,
    isLoading: areasLoading,
    error: areasError,
    refetch: refetchAreas,
  } = useFactoryAreas(selectedFactory?.id);

  const [sortConfig, setSortConfig] = useState<{
    key: FactorySortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });

  const [openAddFactory, setOpenAddFactory] = useState(false);
  const [addFactoryError, setAddFactoryError] = useState<string | null>(null);

  const [openAddArea, setOpenAddArea] = useState(false);
  const [addAreaError, setAddAreaError] = useState<string | null>(null);
  const [pendingDeleteArea, setPendingDeleteArea] =
    useState<FactoryAreaApiResponse | null>(null);
  const [pendingDeleteFactory, setPendingDeleteFactory] = useState<
    string | null
  >(null);

  const { show, hide, snackbar } = useSnackbar();

  useEffect(() => {
    if (areasError)
      show("Failed to load factory areas", SNACKBAR_SEVERITY.ERROR);
  }, [areasError, show]);

  const handleSelectFactory = (factory: FactoryApiResponse) => {
    setSelectedFactory(factory);
  };

  const handleBack = () => {
    setSelectedFactory(null);
  };

  function handleSort(column: string) {
    const col = column as FactorySortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  const handleAddFactory = (values: Record<string, string>) => {
    setAddFactoryError(null);
    return postFactory({ name: values["Name"], address: values["Address"] })
      .then(() => {
        refetchFactories();
        setOpenAddFactory(false);
        show("Factory added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to add factory:", err);
        setAddFactoryError("Failed to add factory.");
        show("Failed to add factory", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const handleConfirmDeleteFactory = () => {
    if (!pendingDeleteFactory) return;
    deleteFactory(pendingDeleteFactory)
      .then(() => {
        refetchFactories();
        setPendingDeleteFactory(null);
        show("Factory deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to delete factory:", err);
        show("Failed to delete factory", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const handleAddArea = (values: Record<string, string>) => {
    if (!selectedFactory) return Promise.resolve();
    setAddAreaError(null);
    return postFactoryArea({
      factoryId: selectedFactory.id,
      name: values["Name"],
      description: values["Description"] || undefined,
    })
      .then(() => {
        refetchAreas();
        setOpenAddArea(false);
        show("Factory area added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to add factory area:", err);
        setAddAreaError("Failed to add factory area.");
        show("Failed to add factory area", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const handleConfirmDeleteArea = () => {
    if (!pendingDeleteArea) return;
    deleteFactoryArea(pendingDeleteArea.id)
      .then(() => {
        refetchAreas();
        setPendingDeleteArea(null);
        show("Factory area deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        console.error("Failed to delete factory area:", err);
        show("Failed to delete factory area", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const sorted = sortFactories(factories, sortConfig.key, sortConfig.direction);

  // — Factories view —
  if (!selectedFactory) {
    return (
      <div className="flex h-screen">
        <PageContent>
          <SubPageHeader
            title="Factories"
            action={
              <CustomButton onClick={() => setOpenAddFactory(true)}>
                Add factory
              </CustomButton>
            }
          />
          <PageDivider />
          <CategoryHeader
            categories={factoryColumns}
            columns={factoryColumns.length + 2}
            sortableColumns={sortableFactoryColumns}
            sortConfig={sortConfig}
            onSort={handleSort}
          >
            {factoriesLoading && <p>Loading...</p>}
            {factoriesError && (
              <Box sx={{ mt: 2 }}>
                <Typography variant="body2" sx={{ color: "error.main", mb: 1 }}>
                  Failed to load factories. {factoriesError.message}
                </Typography>
                <Button
                  variant="outlined"
                  size="small"
                  onClick={refetchFactories}
                >
                  Retry
                </Button>
              </Box>
            )}
            {sorted.map((factory) => (
              <DeviceRow key={factory.id}>
                <FactoryInfo
                  name={factory.name}
                  address={factory.address}
                  createdAt={formatTimestamp(factory.createdAt.toString())}
                  updatedAt={formatTimestamp(factory.updatedAt.toString())}
                  onDelete={() => setPendingDeleteFactory(factory.id)}
                  onViewAreas={() => handleSelectFactory(factory)}
                />
              </DeviceRow>
            ))}
          </CategoryHeader>
          {!factoriesLoading && !factoriesError && sorted.length === 0 && (
            <NotFoundCard page="factories" isEmpty={true} />
          )}
        </PageContent>
        <AddEntityDialog
          open={openAddFactory}
          title="Add factory"
          fields={["Name", "Address"]}
          onClose={() => {
            setOpenAddFactory(false);
            setAddFactoryError(null);
          }}
          onSubmit={handleAddFactory}
          submitError={addFactoryError}
        />
        <DeleteConfirmation
          open={!!pendingDeleteFactory}
          onClose={() => setPendingDeleteFactory(null)}
          onConfirm={handleConfirmDeleteFactory}
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

  // — Factory areas view —
  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader
          title={`Factory areas — ${selectedFactory.name}`}
          action={
            <CustomButton onClick={() => setOpenAddArea(true)}>
              Add factory area
            </CustomButton>
          }
        />
        <div className="flex items-center gap-2 mt-1 mb-2">
          <IconButton
            onClick={handleBack}
            size="small"
            sx={{ color: "primary.main" }}
            aria-label="back to factories"
          >
            <ArrowBackIcon />
          </IconButton>
          <Typography variant="body2" sx={{ color: "primary.main" }}>
            Back to factories
          </Typography>
        </div>
        <PageDivider />
        <CategoryHeader
          categories={areaColumns}
          columns={areaColumns.length + 1}
          sortableColumns={[]}
          sortConfig={{ key: null, direction: "asc" }}
          onSort={() => {}}
        >
          {areasLoading && <p>Loading...</p>}
          {factoryAreas.map((area) => (
            <DeviceRow key={area.id}>
              <FactoryAreaInfo
                name={area.name}
                description={area.description}
                onDelete={() => setPendingDeleteArea(area)}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!areasLoading && !areasError && factoryAreas.length === 0 && (
          <NotFoundCard page="factory areas" isEmpty={false} />
        )}
      </PageContent>
      <AddEntityDialog
        open={openAddArea}
        title="Add factory area"
        fields={["Name", "Description"]}
        onClose={() => {
          setOpenAddArea(false);
          setAddAreaError(null);
        }}
        onSubmit={handleAddArea}
        submitError={addAreaError}
      />
      <DeleteConfirmation
        open={!!pendingDeleteArea}
        onClose={() => setPendingDeleteArea(null)}
        onConfirm={handleConfirmDeleteArea}
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
