import { useEffect, useState } from "react";

import {
  getMeasurementTypesAll,
  type MeasurementTypeApiResponse,
} from "@entities/measurementType";
import {
  FixedSensorSchema,
  getPayloadSchema,
  getPayloadTags,
  putPayloadSchema,
} from "@entities/payloadSchema";
import {
  getSensorMetrics,
  getSensorProfileConfig,
  getSensorProfiles,
  putSensorMetrics,
  putSensorProfileConfig,
  useSensors,
} from "@entities/sensor";
import { getDeviceEUI } from "@entities/sensor/api/getDeviceEUI.ts";
import type { SensorProfileApiResponse } from "@entities/sensor/model/sensorSchema.ts";
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
 * Full-page view listing sensor schema from a given profile registered on the site.
 * @returns The rendered SensorSchema page
 */
export default function PayloadSchema() {
  const [sensorProfiles, setSensorProfiles] = useState<
    SensorProfileApiResponse[]
  >([]);
  const [measurementTypes, setMeasurementTypes] = useState<
    MeasurementTypeApiResponse[]
  >([]);
  const { sensors } = useSensors();
  const [payloadKeys, setPayloadKeys] = useState<string[]>([]);
  const [isLoadingKeys, setIsLoadingKeys] = useState(false);
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
    name: p.name,
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

  const handleSaveConfigurable = () => {
    const metrics = Object.entries(schemaRows).map(([payloadKey, value]) => ({
      payloadKey,
      measurementType: value.measurementType,
      unit: value.unit,
    }));

    putSensorMetrics({ metrics }, sensor)
      .then(() => {
        show("Sensor metrics saved successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to save sensor metrics", SNACKBAR_SEVERITY.ERROR);
      });
  };

  useEffect(() => {
    getMeasurementTypesAll()
      .then((data) => {
        setMeasurementTypes(data);
      })
      .catch(() => {
        show("Failed to fetch measurement types", SNACKBAR_SEVERITY.ERROR);
      });
  }, []);

  useEffect(() => {
    getSensorProfiles().then(setSensorProfiles);
  }, []);

  useEffect(() => {
    if (!profile) return;

    getSensorProfileConfig(profile)
      .then((config) => {
        setPerInstallation(config.configurableSchema);
      })
      .catch(() => {
        show("Failed to fetch profile configuration", SNACKBAR_SEVERITY.ERROR);
      });
  }, [profile]);

  useEffect(() => {
    if (!profile || perInstallation) return;

    getDeviceEUI(profile)
      .then((eui) => {
        setIsLoadingKeys(true);
        if (!eui) {
          setPayloadKeys([]);
          setSchemaRows({});
          return Promise.all([[], []]);
        }
        return Promise.all([getPayloadTags(eui), getPayloadSchema(profile)]);
      })
      .then((result) => {
        if (!result) return;
        const [keys, existing] = result;
        setPayloadKeys(keys);
        setSchemaRows(
          Object.fromEntries(
            existing.map((row) => [
              row.payloadKey,
              { measurementType: row.measurementType, unit: row.unit },
            ]),
          ),
        );
        setIsLoadingKeys(false);
      });
  }, [profile, perInstallation]);

  useEffect(() => {
    if (!sensor || !perInstallation) return;

    getSensorMetrics(sensor).then((metrics) => {
      setIsLoadingKeys(true);
      if (metrics.length > 0) {
        setPayloadKeys(metrics.map((m) => m.payloadKey));
        setSchemaRows(
          Object.fromEntries(
            metrics.map((m) => [
              m.payloadKey,
              { measurementType: m.measurementType, unit: m.unit },
            ]),
          ),
        );
      } else {
        getPayloadTags(sensor).then((keys) => {
          setPayloadKeys(keys);
          setSchemaRows({});
        });
      }
      setIsLoadingKeys(false);
    });
  }, [sensor, perInstallation]);

  const filteredSensors = sensors.filter((sensor) => {
    return sensor.sensorProfileId === profile;
  });

  const filteredSensorOptions = filteredSensors.map((s) => ({
    id: s.deviceEui,
    name: s.name,
  }));

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Payload schema" />
        <PageDivider />
        <Typography>Select a sensor profile</Typography>
        <Box sx={{ width: 200 }}>
          <DropDownSelect
            options={profileOptions}
            value={profile}
            onChange={(value) => setProfile(value)}
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
                  putSensorProfileConfig(
                    { configurableSchema: checked },
                    profile,
                  );
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
                  half={true}
                >
                  {isLoadingKeys && <p>Loading...</p>}
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
                {!isLoadingKeys && payloadKeys.length === 0 && (
                  <NotFoundCard page="measure types" isEmpty={true} />
                )}
              </>
            ) : (
              <>
                <Box sx={{ width: 200 }}>
                  <DropDownSelect
                    options={filteredSensorOptions}
                    value={sensor}
                    onChange={(e) => {
                      setSensor(e);
                    }}
                  />
                </Box>
                {sensor.length ? (
                  <>
                    <CategoryHeader
                      categories={payloadDetails}
                      columns={payloadDetails.length}
                      half={true}
                    >
                      {isLoadingKeys && <p>Loading...</p>}
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
                      onClick={handleSaveConfigurable}
                    >
                      Save
                    </Button>
                    {!isLoadingKeys && payloadKeys.length === 0 && (
                      <NotFoundCard page="measure types" isEmpty={true} />
                    )}
                  </>
                ) : null}
              </>
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
