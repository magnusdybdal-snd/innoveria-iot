import { useEffect, useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import Typography from "@mui/material/Typography";

import { ELECTRICITY_SENSOR } from "@shared/const";
import { DeviceFormFields } from "@shared/ui/DeviceFormFields";

export interface EditSensorProps {
  open: boolean;
  onClose: () => void;
  editOptions: string[];
  profileOptions?: { id: string; name: string }[];
  factoryOptions?: { id: string; name: string }[];
  factoryAreaOptions?: { id: string; name: string }[];
  voltageOptions?: { id: string; name: string }[];
  onFactoryChange?: (factoryId: string) => void;
  sensor: {
    id: string;
    name: string;
    deviceEui: string;
    electricitySensor: boolean;
    factory: string;
    factoryArea: string;
    productionResource: string | null;
    appKey: string;
    deviceProfile: string;
    voltage: number | null;
    description?: string;
  };
  onEdit: (
    deviceId: string,
    payload: {
      name?: string;
      electricitySensor?: boolean;
      factory?: string;
      factoryArea?: string;
      productionResource?: string | null;
      appKey?: string;
      deviceProfile?: string;
      voltage?: number | null;
    },
  ) => void;
  submitError?: string | null;
}

const inputHints: Record<string, string> = {
  Name: "Enter device name",
  DeviceEUI: "16 characters (hex)",
  "Production resource": "Enter production resource",
  "Application key": "32 characters (hex)",
};

const inputLengthError: Record<string, string> = {
  DeviceEUI: "DeviceEUI must be 16 characters",
  "Application key": "Application key must be 32 characters",
  "Production resource": "Production resource must be a positive number",
};

/**
 * Modal dialog for editing an existing sensor, with input validation for DeviceEUI and Application key lengths.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.addOptions - Field names to render as inputs inside the dialog
 * @param props.profileOptions - Available sensor profiles for the dropdown
 * @param props.onAdd - Called with the validated sensor data when the user confirms
 * @param props.submitError - Error message to display if the submission fails
 * @returns The rendered add-sensor dialog
 */
export function EditSensor(props: EditSensorProps) {
  const {
    onClose,
    open,
    sensor,
    onEdit,
    editOptions,
    profileOptions = [],
    factoryOptions = [],
    factoryAreaOptions = [],
    onFactoryChange,
    voltageOptions = [],
    submitError,
  } = props;
  const [values, setValues] = useState<Record<string, string>>({});
  const [lengthErrors, setLengthErrors] = useState<Record<string, boolean>>({});

  // Initialize form when sensor changes
  useEffect(() => {
    if (!sensor) return;

    // eslint-disable-next-line react-hooks/set-state-in-effect
    setValues({
      Name: sensor.name,
      DeviceEUI: sensor.deviceEui,
      "Electricity sensor": String(sensor.electricitySensor),
      Factory: sensor.factory,
      "Factory area": sensor.factoryArea,
      "Production resource": sensor.productionResource?.toString() ?? "",
      "Application key": sensor.appKey,
      "Sensor profile": sensor.deviceProfile,
      Voltage: sensor.voltage?.toString() ?? "",
    });
  }, [sensor]);

  const handleClose = () => {
    setLengthErrors({});
    onClose();
  };

  const handleSave = () => {
    const electricityEnabled = values[ELECTRICITY_SENSOR] === "true";

    const productionResourceRaw = (values["Production resource"] ?? "").trim();
    const productionResourceParsed = parseInt(productionResourceRaw, 10);

    const productionResourceInvalid =
      productionResourceRaw !== "" &&
      (isNaN(productionResourceParsed) || productionResourceParsed <= 0);

    const newLengthErrors = {
      DeviceEUI: (values["DeviceEUI"] ?? "").length !== 16,
      "Application key": (values["Application key"] ?? "").length !== 32,
      ProductionResource: productionResourceInvalid,
    };

    if (
      (editOptions.includes("DeviceEUI") && newLengthErrors.DeviceEUI) ||
      (editOptions.includes("Application key") &&
        newLengthErrors["Application key"]) ||
      (editOptions.includes("Production resource") &&
        newLengthErrors.ProductionResource)
    ) {
      setLengthErrors(newLengthErrors);
      return;
    }

    // Build payload only with changed values
    type EditSensorPayload = Parameters<EditSensorProps["onEdit"]>[1];
    const payload: EditSensorPayload = {};

    if (values["Name"] !== sensor.name && values["Name"].length > 0)
      payload.name = values["Name"];

    if ((values["Electricity sensor"] === "true") !== sensor.electricitySensor)
      payload.electricitySensor = values["Electricity sensor"] === "true";

    if (values["Factory"] !== sensor.factory && values["Factory"].length > 0)
      payload.factory = values["Factory"];

    if (
      values["Factory area"] !== sensor.factoryArea &&
      values["Factory area"].length > 0
    )
      payload.factoryArea = values["Factory area"];

    if (
      productionResourceRaw !== "" &&
      String(productionResourceParsed) !== sensor.productionResource
    ) {
      payload.productionResource = String(productionResourceParsed);
    }

    if (
      values["Application key"] !== sensor.appKey &&
      values["Application key"].length > 0
    )
      payload.appKey = values["Application key"];

    if (
      values["Sensor profile"] !== sensor.deviceProfile &&
      values["Sensor profile"].length > 0
    )
      payload.deviceProfile = values["Sensor profile"];

    const voltageValue = electricityEnabled ? Number(values["Voltage"]) : null;

    if (voltageValue !== sensor.voltage) {
      payload.voltage = voltageValue;
    }

    if (Object.keys(payload).length === 0) {
      onClose();
      return;
    }

    onEdit(sensor.id, payload);
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      maxWidth="sm"
      fullWidth
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.contrastText",
        },
      }}
    >
      <DialogTitle sx={{ color: "primary.main" }}>Edit sensor</DialogTitle>

      <DialogContent>
        <DeviceFormFields
          options={editOptions}
          values={values}
          profileOptions={profileOptions}
          factoryOptions={factoryOptions}
          factoryAreaOptions={factoryAreaOptions}
          voltageOptions={voltageOptions}
          lengthErrors={lengthErrors}
          lengthErrorMessages={inputLengthError}
          inputHints={inputHints}
          onChange={(option, value) => {
            if (option === "Factory") {
              setValues((prev) => ({
                ...prev,
                Factory: value,
                "Factory area": "",
              }));
              onFactoryChange?.(value);
            } else {
              setValues((prev) => ({ ...prev, [option]: value }));
            }
          }}
          productionResourceOptions={[]}
        />
        {submitError && (
          <Typography color="error" mt={1}>
            {submitError}
          </Typography>
        )}
      </DialogContent>
      <DialogActions>
        <Button
          sx={{
            backgroundColor: "primary.main",
            color: "primary.contrastText",
          }}
          onClick={handleClose}
        >
          Close
        </Button>
        <Button
          sx={{
            backgroundColor: "primary.main",
            color: "primary.contrastText",
          }}
          onClick={handleSave}
          autoFocus
        >
          Save
        </Button>
      </DialogActions>
    </Dialog>
  );
}
