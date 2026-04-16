import { useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import Typography from "@mui/material/Typography";

import { DeviceFormFields } from "@shared/ui/DeviceFormFields";

export interface AddDeviceProps {
  open: boolean;
  onClose: () => void;
  addOptions: string[];
  profileOptions?: { id: string; name: string }[];
  factoryOptions?: { id: string; name: string }[];
  factoryAreaOptions?: { id: string; name: string }[];
  onAdd: (sensor: {
    name: string;
    deviceEui: string;
    factory: string;
    factoryArea: string;
    productionResource: number;
    appKey: string;
    senProf: string;
  }) => Promise<void>;
  submitError?: string | null;
}

const inputHints: Record<string, string> = {
  Name: "Enter device name",
  DeviceEUI: "16 characters (hex)",
  Machine: "Enter machine name",
  "Application key": "32 characters (hex)",
};

const inputLengthError: Record<string, string> = {
  DeviceEUI: "DeviceEUI must be 16 characters",
  "Application key": "Application key must be 32 characters",
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
    submitError,
  } = props;
  const [values, setValues] = useState<Record<string, string>>({});
  const [fillError, setFillError] = useState(false);
  const [lengthErrors, setLengthErrors] = useState<Record<string, boolean>>({});

  const handleClose = () => {
    onClose();
  };

  const handleSafeClose = () => {
    const allFilled = addOptions.every(
      (option) => (values[option] ?? "").trim() !== "",
    );

    const newLengthErrors = {
      DeviceEUI: (values["DeviceEUI"] ?? "").length !== 16,
      "Application key": (values["Application key"] ?? "").length !== 32,
    };

    if (!allFilled) {
      setFillError(true);
      return;
    } else if (
      (addOptions.includes("DeviceEUI") && newLengthErrors.DeviceEUI) ||
      (addOptions.includes("Application key") &&
        newLengthErrors["Application key"])
    ) {
      setFillError(false);
      setLengthErrors(newLengthErrors);
      return;
    }

    props
      .onAdd({
        name: values["Name"],
        deviceEui: values["DeviceEUI"],
        factory: values["Factory"],
        factoryArea: values["Factory area"],
        productionResource: parseInt(values["Machine"], 10),
        appKey: values["Application key"],
        senProf: values["Sensor profile"],
      })
      .then(() => {
        setValues({});
      });

    setFillError(false);
    setLengthErrors({});
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
          lengthErrors={lengthErrors}
          lengthErrorMessages={inputLengthError}
          inputHints={inputHints}
          onChange={(option, value) =>
            setValues((prev) => ({ ...prev, [option]: value }))
          }
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
