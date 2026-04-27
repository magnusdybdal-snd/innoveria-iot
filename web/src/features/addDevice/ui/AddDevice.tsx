import { useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import Typography from "@mui/material/Typography";

import { ELECTRICITY_SENSOR, VOLTAGE } from "@shared/const";
import { DeviceFormFields } from "@shared/ui/DeviceFormFields";

export interface AddDeviceProps {
  open: boolean;
  onClose: () => void;
  addOptions: string[];
  profileOptions?: { id: string; name: string }[];
  factoryOptions?: { id: string; name: string }[];
  factoryAreaOptions?: { id: string; name: string }[];
  voltageOptions?: { id: string; name: string }[];
  onFactoryChange?: (factoryId: string) => void;
  onAdd: (sensor: {
    name: string;
    deviceEui: string;
    electricitySensor: boolean;
    factory: string;
    factoryArea: string;
    productionResource: number | null;
    appKey: string;
    senProf: string;
    voltage: number | null;
  }) => Promise<void>;
  submitError?: string | null;
}

const inputHints: Record<string, string> = {
  Name: "Enter device name",
  DeviceEUI: "16 characters (hex)",
  ProductionResource: "Enter production resource",
  "Application key": "32 characters (hex)",
};

const inputLengthError: Record<string, string> = {
  DeviceEUI: "DeviceEUI must be 16 characters",
  "Application key": "Application key must be 32 characters",
  ProductionResource: "Production resource must be a positive number",
};

/**
 * Modal dialog for registering a new sensor device, with input validation for DeviceEUI and Application key lengths.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.addOptions - Field names to render as inputs inside the dialog
 * @param props.profileOptions - Available sensor profiles for the dropdown
 * @param props.onAdd - Called with the validated sensor data when the user confirms
 * @param props.submitError - Error message to display if the submission fails
 * @returns The rendered add-device dialog
 */
export function AddDevice(props: AddDeviceProps) {
  const {
    onClose,
    open,
    addOptions,
    profileOptions = [],
    factoryOptions = [],
    factoryAreaOptions = [],
    onFactoryChange,
    voltageOptions = [],
    submitError,
  } = props;
  const [values, setValues] = useState<Record<string, string>>({});
  const [fillError, setFillError] = useState(false);
  const [lengthErrors, setLengthErrors] = useState<Record<string, boolean>>({});

  const handleClose = () => {
    setValues({});
    setFillError(false);
    setLengthErrors({});
    onClose();
  };

  const handleSafeClose = () => {
    const electricityEnabled = values[ELECTRICITY_SENSOR] === "true";

    const allFilled = addOptions
      .filter(
        (option) =>
          option !== ELECTRICITY_SENSOR &&
          option !== "Production resource" &&
          (option !== VOLTAGE || electricityEnabled),
      )
      .every((option) => (values[option] ?? "").trim() !== "");

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

    if (!allFilled) {
      setFillError(true);
      return;
    } else if (
      (addOptions.includes("DeviceEUI") && newLengthErrors.DeviceEUI) ||
      (addOptions.includes("Application key") &&
        newLengthErrors["Application key"]) ||
      (addOptions.includes("Production resource") &&
        newLengthErrors.ProductionResource)
    ) {
      setFillError(false);
      setLengthErrors(newLengthErrors);
      return;
    }

    props
      .onAdd({
        name: values["Name"],
        deviceEui: values["DeviceEUI"],
        electricitySensor: values["Electricity sensor"] === "true",
        factory: values["Factory"],
        factoryArea: values["Factory area"],
        productionResource:
          productionResourceRaw !== "" ? productionResourceParsed : null,
        appKey: values["Application key"],
        senProf: values["Sensor profile"],
        voltage: electricityEnabled ? Number(values["Voltage"]) : null,
      })
      .then(() => {
        setValues({});
        setFillError(false); // only clear on success
        setLengthErrors({});
      })
      .catch(() => {});
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
      <DialogTitle id="alert-dialog-title" sx={{ color: "primary.main" }}>
        {"Insert device info"}
      </DialogTitle>
      <DialogContent>
        <DeviceFormFields
          options={addOptions}
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
        />
        {fillError && (
          <Typography color="error" mt={1}>
            All fields must be filled
          </Typography>
        )}
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
          onClick={handleSafeClose}
          autoFocus
        >
          Add
        </Button>
      </DialogActions>
    </Dialog>
  );
}
