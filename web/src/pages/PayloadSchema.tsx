import { useEffect, useState } from "react";

import {
  getMeasurementTypesAll,
  type MeasurementTypeApiResponse,
  type MeasurementTypeSortKey,
  type SortDirection,
} from "@entities/measurementType";
import {
  FixedSensorSchema,
  getPayloadTags,
  putPayloadSchema,
} from "@entities/payloadSchema";
import { getSensorProfiles, useSensors } from "@entities/sensor";
import { getDeviceEUI } from "@entities/sensor/api/getDeviceEUI.ts";
import { getSensorProfileConfig } from "@entities/sensor/api/getSensorProfileConfig.ts";
import type {
  SensorProfileApiResponse,
  SensorProfileConfigApiResponse,
} from "@entities/sensor/model/sensorSchema.ts";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import Typography from "@mui/material/Typography";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { DropDownSelect } from "@shared/ui/DropDownSelect";
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
  const [sensorProfileConfig, setSensorProfileConfig] =
    useState<SensorProfileConfigApiResponse | null>(null);
  const { sensors } = useSensors();
  //const [deviceEui, setDeviceEui] = useState<string>("");
  const [payloadKeys, setPayloadKeys] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [sortConfig, setSortConfig] = useState<{
    key: MeasurementTypeSortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });
  const [profile, setProfile] = useState<string>("");
  const [sensor, setSensor] = useState<string>("");
  const [perInstallation, setPerInstallation] = useState<boolean>(false);
  const [schemaRows, setSchemaRows] = useState<
    Record<
      string,
      {
        measurementType: string;
        unit: string;
      }
    >
  >({});

  const profileOptions = sensorProfiles.map((p) => ({
    id: p.id,
    name: p.id, // or p.name if you have it
  }));

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

  const handleSaveFixed = () => {
    const labels = Object.entries(schemaRows).map(([payloadKey, value]) => ({
      payloadKey,
      measurementType: value.measurementType,
      unit: value.unit,
    }));

    putPayloadSchema({ labels }, profile)
      .then(() => {
        show("Payload schema saved successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to save payload schema", SNACKBAR_SEVERITY.ERROR);
      });
  };

  useEffect(() => {
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
  }, []);

  useEffect(() => {
    getSensorProfiles().then(setSensorProfiles);
  }, []);

  useEffect(() => {
    if (!profile) return;
    getSensorProfileConfig(profile).then((eui) => {
      setSensorProfileConfig(eui);

      setPerInstallation(sensorProfileConfig?.configurableSchema ?? false);
    });
  }, [profile]);

  useEffect(() => {
    if (!profile) return;

    getDeviceEUI(profile)
      .then((eui) => {
        const finalEui = eui ?? "b000000000000001"; //TODO: Get actual eui
        //setDeviceEui(finalEui);

        return getPayloadTags(finalEui);
      })
      .then(setPayloadKeys);
  }, [profile]);

  function handleSort(column: string) {
    const col = column as MeasurementTypeSortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  const filteredSensors = sensors.filter((sensor) => {
    return sensor.sensorProfileId === profile;
  });

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Payload schema" />
        <PageDivider />
        {sensorProfileConfig?.chirpstackProfileId}
        <br />
        {sensorProfileConfig?.configurableSchema ? "true" : "false"}
        <br />
        <br />
        <Typography>Select a sensor profile</Typography>
        <Box sx={{ width: 200 }}>
          <DropDownSelect
            options={profileOptions}
            value={profile}
            onChange={() => setProfile("f0000000-0000-0000-0000-000000000001")} //TODO: Get actual profile.
          />
        </Box>
        <br />
        {profile.length ? (
          <>
            <Box display="flex" alignItems="center" gap={1}>
              <Typography>
                This profile requires per-installation configuration:
              </Typography>
              <Checkbox
                checked={perInstallation}
                onChange={(_, checked) => {
                  setPerInstallation(checked);
                }}
                slotProps={{
                  input: { "aria-label": "controlled" },
                }}
              />
            </Box>

            {!perInstallation ? (
              <>
                <CategoryHeader
                  categories={payloadDetails}
                  columns={payloadDetails.length}
                  sortConfig={sortConfig}
                  onSort={handleSort}
                  half={true}
                >
                  {isLoading && <p>Loading...</p>}
                  {/*TODO: make a better looking loading indicator */}
                  {payloadKeys.map((payloadKey) => (
                    <DeviceRow key={payloadKey}>
                      <FixedSensorSchema
                        payloadKey={payloadKey}
                        measurementTypes={measurementTypes}
                        value={schemaRows[payloadKey]}
                        onChange={(value) => {
                          setSchemaRows((prev) => ({
                            ...prev,
                            [payloadKey]: value,
                          }));
                        }}
                      />
                    </DeviceRow>
                  ))}
                </CategoryHeader>
                <Button
                  variant="outlined"
                  sx={{
                    backgroundColor: "primary.main",
                    color: "primary.dark",
                    "&:hover": { backgroundColor: "primary.main" },
                    borderRadius: 2,
                    textTransform: "none",
                    fontSize: 15,
                  }}
                  onClick={handleSaveFixed}
                >
                  Save
                </Button>
                {!isLoading && payloadKeys.length === 0 && (
                  <NotFoundCard page="measure types" isEmpty={true} />
                )}
              </>
            ) : (
              <Box sx={{ width: 200 }}>
                <DropDownSelect
                  options={filteredSensors}
                  value={sensor}
                  onChange={(e) => setSensor(e)}
                />
              </Box>
            )}
          </>
        ) : null}

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
