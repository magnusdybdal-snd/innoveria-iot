import { useEffect, useState } from "react";

import {
  getMeasurementTypesAll,
  type MeasurementTypeApiResponse,
  type MeasurementTypeSortKey,
  type SortDirection,
} from "@entities/measurementType";
import { FixedSensorSchema, getPayloadTags } from "@entities/payloadSchema";
import { getSensorProfiles } from "@entities/sensor";
import { getDeviceEUI } from "@entities/sensor/api/getDeviceEUI.ts";
import { getSensorProfileConfig } from "@entities/sensor/api/getSensorProfileConfig.ts";
import type {
  SensorProfileApiResponse,
  SensorProfileConfigApiResponse,
} from "@entities/sensor/model/sensorSchema.ts";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
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
  const [sensorProfileConfig, setSensorProfileConfig] =
    useState<SensorProfileConfigApiResponse | null>(null);
  const [deviceEui, setDeviceEui] = useState<string>("");
  const [payloadKeys, setPayloadKeys] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [sortConfig, setSortConfig] = useState<{
    key: MeasurementTypeSortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });
  const [profile, setProfile] = useState<string>("");
  const [payloadKey, setPayloadKey] = useState<string>("");

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

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
    });
  }, [profile]);

  useEffect(() => {
    if (!profile) return;

    getDeviceEUI(profile).then((eui) => {
      setDeviceEui(eui ?? "b000000000000001");
    });

    getPayloadTags(deviceEui).then((keys) => {
      setPayloadKeys(keys);
    });
  }, [profile]);

  function handleSort(column: string) {
    const col = column as MeasurementTypeSortKey;
    setSortConfig((prev) =>
      prev.key === col
        ? { key: col, direction: prev.direction === "asc" ? "desc" : "asc" }
        : { key: col, direction: "asc" },
    );
  }

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Payload schema" />
        <PageDivider />
        {profile ? "full" : "empty"}
        <br />
        {sensorProfileConfig?.chirpstackProfileId}
        <br />
        {sensorProfileConfig?.configurableSchema ? "true" : "false"}
        <br />
        {profile}
        <br />
        {deviceEui.length > 0 ? deviceEui : "null"}
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
              {option.id}
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
          half={true}
        >
          {isLoading && <p>Loading...</p>}
          {/*TODO: make a better looking loading indicator */}
          {payloadKeys.map((payloadKey) => (
            <DeviceRow key={payloadKey}>
              <FixedSensorSchema
                payloadKey={payloadKey}
                measurementTypes={measurementTypes}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && payloadKeys.length === 0 && (
          <NotFoundCard page="measure types" isEmpty={true} />
        )}
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
