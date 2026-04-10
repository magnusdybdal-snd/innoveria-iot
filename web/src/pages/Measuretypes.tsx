import { useEffect, useState } from "react";

import {
  deprecateMeasureType,
  getMeasureTypesAll,
  MeasureTypeInfo,
  postMeasureType,
  sortMeasureTypes,
  type MeasureTypeApiResponse,
  type MeasureTypeSortKey,
  type SortDirection,
} from "@entities/measureType";
import { AddEntityDialog } from "@shared/ui/AddEntityDialog";
import { CustomButton } from "@shared/ui/Button";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
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

const measureTypeDetails: string[] = [
  "Slug",
  "Display name",
  "Description",
  "Default unit",
];

const sortableColumns: MeasureTypeSortKey[] = [
  "Slug",
  "Display name",
  "Description",
];

/**
 * Full-page view listing all measure types registered on the site.
 * @returns The rendered MeasureTypes page
 */
export default function MeasureTypes() {
  const [measureTypes, setMeasureTypes] = useState<MeasureTypeApiResponse[]>(
    [],
  );
  const [isLoading, setIsLoading] = useState(true);
  const [sortConfig, setSortConfig] = useState<{
    key: MeasureTypeSortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });

  // Adding a new measure type
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
    <CustomButton onClick={handleClickOpenAdd}>Add measure type</CustomButton>
  );
  const handleAddMeasureType = (measureTypeData: {
    defaultUnit: string;
    description: string;
    displayName: string;
    slug: string;
  }) => {
    setAddError(null);
    return postMeasureType({
      defaultUnit: measureTypeData.defaultUnit,
      description: measureTypeData.description,
      displayName: measureTypeData.displayName,
      slug: measureTypeData.slug,
    })
      .then(() => {
        fetchMeasureTypes();
        setOpenAdd(false);
        show("Measure type added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        setAddError(
          "Failed to add measure type.", // TODO: throw non-hardcoded error messages - based on actual error
        );
        show("Failed to add measure type", SNACKBAR_SEVERITY.ERROR);
      });
  };

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

  const fetchMeasureTypes = () => {
    getMeasureTypesAll().then((data) => {
      setMeasureTypes(data);
      setIsLoading(false);
    });
  };

  //TODO: use deletion confirmation dialog when it has been implemented.
  const handleDeprecateMeasureType = (id: string) => {
    deprecateMeasureType(id)
      .then(() => {
        fetchMeasureTypes();
        show("Measure type deprecated successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to deprecate measure type", SNACKBAR_SEVERITY.ERROR);
      });
  };

  useEffect(() => {
    fetchMeasureTypes();
  }, []);

  function handleSort(column: string) {
    const col = column as MeasureTypeSortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  const sorted = sortMeasureTypes(
    measureTypes,
    sortConfig.key,
    sortConfig.direction,
  );

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Measurement types" action={addButton} />
        <PageDivider />
        <CategoryHeader
          categories={measureTypeDetails}
          columns={measureTypeDetails.length + 1}
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {isLoading && <p>Loading...</p>}
          {/*TODO: make a better looking loading indicator */}
          {sorted.map((measureType) => (
            <DeviceRow key={measureType.slug} greyed={measureType.deprecated}>
              <MeasureTypeInfo
                defaultUnit={measureType.defaultUnit}
                description={measureType.description}
                displayName={measureType.displayName}
                slug={measureType.slug}
                deprecated={measureType.deprecated}
                onDeprecate={() => handleDeprecateMeasureType(measureType.slug)}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && (
          <NotFoundCard page="measure types" isEmpty={true} />
        )}
        <AddEntityDialog
          open={openAdd}
          title="Add measure type"
          fields={["Slug", "Default unit", "Description", "Display name"]}
          onClose={handleCloseAdd}
          onSubmit={(values) =>
            handleAddMeasureType({
              slug: values["Slug"],
              defaultUnit: values["Default unit"],
              description: values["Description"],
              displayName: values["Display name"],
            })
          }
          submitError={addError}
        />
        <AppSnackbar
          open={snackbar?.open ?? false}
          message={snackbar?.message ?? ""}
          severity={snackbar?.severity}
          onClose={hide}
        />
      </PageContent>
    </div>
  );
}
