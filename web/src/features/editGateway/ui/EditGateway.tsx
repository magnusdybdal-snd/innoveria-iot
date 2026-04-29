import { useEffect, useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import Typography from "@mui/material/Typography";

import { DeviceFormFields } from "@shared/ui/DeviceFormFields";

export interface EditGatewayProps {
  open: boolean;
  onClose: () => void;
  editOptions: string[];
  factoryOptions?: { id: string; name: string }[];
  factoryAreaOptions?: { id: string; name: string }[];
  onFactoryChange?: (factoryId: string) => void;
  gateway: {
    id: string;
    name: string;
    deviceEui: string;
    factory: string;
    factoryArea: string;
  };
  onEdit: (
    deviceId: string,
    payload: {
      name?: string;
      factory?: string;
      factoryArea?: string;
    },
  ) => void;
  submitError?: string | null;
}

const inputHints: Record<string, string> = {
  Name: "Enter gateway name",
  DeviceEUI: "16 characters (hex)",
};

const inputLengthError: Record<string, string> = {
  DeviceEUI: "DeviceEUI must be 16 characters",
};

/**
 * Modal dialog for editing an existing gateway, with input validation for DeviceEUI and Application key lengths.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.addOptions - Field names to render as inputs inside the dialog
 * @param props.onAdd - Called with the validated sensor data when the user confirms
 * @param props.submitError - Error message to display if the submission fails
 * @returns The rendered edit-gateway dialog
 */
export function EditGateway(props: EditGatewayProps) {
  const {
    onClose,
    open,
    gateway,
    onEdit,
    editOptions,
    factoryOptions = [],
    factoryAreaOptions = [],
    onFactoryChange,
    submitError,
  } = props;
  const [values, setValues] = useState<Record<string, string>>({});
  const [lengthErrors, setLengthErrors] = useState<Record<string, boolean>>({});

  // Initialize form when gateway changes
  useEffect(() => {
    if (!gateway) return;

    // eslint-disable-next-line react-hooks/set-state-in-effect
    setValues({
      Name: gateway.name,
      DeviceEUI: gateway.deviceEui,
      Factory: gateway.factory,
      "Factory area": gateway.factoryArea,
    });
  }, [gateway]);

  const handleClose = () => {
    setLengthErrors({});
    onClose();
  };

  const handleSave = () => {
    const newLengthErrors = {
      DeviceEUI: (values["DeviceEUI"] ?? "").length !== 16,
      "Application key": (values["Application key"] ?? "").length !== 32,
    };

    if (
      (editOptions.includes("DeviceEUI") && newLengthErrors.DeviceEUI) ||
      (editOptions.includes("Application key") &&
        newLengthErrors["Application key"])
    ) {
      setLengthErrors(newLengthErrors);
      return;
    }

    // Build payload only with changed values
    type EditGatewayPayload = Parameters<EditGatewayProps["onEdit"]>[1];
    const payload: EditGatewayPayload = {};

    if (values["Name"] !== gateway.name && values["Name"].length > 0)
      payload.name = values["Name"];

    if (values["Factory"] !== gateway.factory && values["Factory"].length > 0)
      payload.factory = values["Factory"];

    if (
      values["Factory area"] !== gateway.factoryArea &&
      values["Factory area"].length > 0
    )
      payload.factoryArea = values["Factory area"];

    if (Object.keys(payload).length === 0) {
      onClose();
      return;
    }

    onEdit(gateway.id, payload);
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
      <DialogTitle sx={{ color: "primary.main" }}>Edit gateway</DialogTitle>

      <DialogContent>
        <DeviceFormFields
          options={editOptions}
          values={values}
          factoryOptions={factoryOptions}
          factoryAreaOptions={factoryAreaOptions}
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
          profileOptions={[]}
          voltageOptions={[]}
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
