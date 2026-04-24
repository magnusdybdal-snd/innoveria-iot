import { useEffect, useState } from "react";

import {
  deprecateMeasurementType,
  getMeasurementTypesAll,
  MeasurementTypeInfo,
  postMeasurementType,
  sortMeasurementTypes,
  type MeasurementTypeApiResponse,
  type MeasurementTypeSortKey,
  type SortDirection,
} from "@entities/measurementType";
import { getPayloadTags } from "@entities/payloadSchema";
import { getSensorProfiles } from "@entities/sensor";
import { getDeviceEUI } from "@entities/sensor/api/getDeviceEUI.ts";
import type { SensorProfileApiResponse } from "@entities/sensor/model/sensorSchema.ts";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
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

const payloadDetails: string[] = ["Payload key", "Measurement type", "Unit"];

/**
 * Full-page view listing all measure types registered on the site.
 * @returns The rendered MeasurementTypes page
 */
export default function PayloadSchema() {
  const [sensorProfiles, setSensorProfiles] = useState<
    SensorProfileApiResponse[]
  >([]);
  const [measurementTypes, setMeasurementTypes] = useState<
    MeasurementTypeApiResponse[]
  >([]);
  const [deviceEui, setDeviceEui] = useState<string | null>(null);
  const [payloadKeys, setPayloadKeys] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [sortConfig, setSortConfig] = useState<{
    key: MeasurementTypeSortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });
  const [profile, setProfile] = useState<string>("");
  const [type, setType] = useState<string>("");
  const [payloadKey, setPayloadKey] = useState<string>("");

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
  const handleAddMeasurementType = (measurementTypeData: {
    defaultUnit: string;
    description: string;
    displayName: string;
    slug: string;
  }) => {
    setAddError(null);
    return postMeasurementType({
      defaultUnit: measurementTypeData.defaultUnit,
      description: measurementTypeData.description,
      displayName: measurementTypeData.displayName,
      slug: measurementTypeData.slug,
    })
      .then(() => {
        fetchMeasurementTypes();
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

  const fetchMeasurementTypes = () => {
    setIsLoading(true);

    getMeasurementTypesAll()
      .then((data) => {
        setMeasurementTypes(data);
      })
      .catch(() => {
        show("Failed to fetch measurement types", SNACKBAR_SEVERITY.ERROR);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  //TODO: use deletion confirmation dialog when it has been implemented.
  const handleDeprecateMeasurementType = (id: string) => {
    deprecateMeasurementType(id)
      .then(() => {
        fetchMeasurementTypes();
        show("Measure type deprecated successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to deprecate measure type", SNACKBAR_SEVERITY.ERROR);
      });
  };

  useEffect(() => {
    fetchMeasurementTypes();
  }, []);

  useEffect(() => {
    getSensorProfiles().then(setSensorProfiles);
  }, []);

  useEffect(() => {
    if (!profile) return;

    getDeviceEUI(profile).then((eui) => {
      setDeviceEui(eui);
    });
  }, [profile]);

  useEffect(() => {
    if (!deviceEui) return;

    getPayloadTags(deviceEui).then((keys) => {
      setPayloadKeys(keys);
    });
  }, [deviceEui]);

  function handleSort(column: string) {
    const col = column as MeasurementTypeSortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  const sorted = sortMeasurementTypes(
    measurementTypes,
    sortConfig.key,
    sortConfig.direction,
  );

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Measurement types" action={addButton} />
        <PageDivider />
        {profile}
        <br />
        {deviceEui ? deviceEui : "null"}
        <br />
        {payloadKeys[0] ? "full" : "null"}
        <br />
        <Select
          value={profile}
          onChange={(e) => {
            setProfile(e.target.value);
          }}
          displayEmpty
        >
          {sensorProfiles.map((option) => (
            <MenuItem key={option.id} value={option.id}>
              {option.name}
            </MenuItem>
          ))}
        </Select>
        <Select
          value={type}
          onChange={(e) => setType(e.target.value)}
          displayEmpty
        >
          {measurementTypes.map((option) => (
            <MenuItem key={option.slug} value={option.slug}>
              {option.displayName}
            </MenuItem>
          ))}
        </Select>
        <Select
          value={payloadKey}
          onChange={(e) => {
            const newValue = e.target.value;
            if (payloadKeys.includes(newValue)) {
              setPayloadKey(newValue);
            }
          }}
          displayEmpty
        >
          {payloadKeys.map((option) => (
            <MenuItem key={option} value={option}>
              {option}
            </MenuItem>
          ))}
        </Select>
        <CategoryHeader
          categories={payloadDetails}
          columns={payloadDetails.length}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {isLoading && <p>Loading...</p>}
          {/*TODO: make a better looking loading indicator */}
          {sorted.map((measurementType) => (
            <DeviceRow
              key={measurementType.slug}
              greyed={measurementType.deprecated}
            >
              <MeasurementTypeInfo
                defaultUnit={measurementType.defaultUnit}
                description={measurementType.description}
                displayName={measurementType.displayName}
                slug={measurementType.slug}
                deprecated={measurementType.deprecated}
                onDeprecate={() =>
                  handleDeprecateMeasurementType(measurementType.slug)
                }
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
          optionalFields={["Default unit", "Description"]}
          onClose={handleCloseAdd}
          onSubmit={(values) =>
            handleAddMeasurementType({
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
