import { useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import TextField from "@mui/material/TextField";

import { CategoryHeader } from "@/components/CategoryHeader";

export interface AddDeviceProps {
  open: boolean;
  onClose: () => void;
  addOptions: string[];
  onAdd: (sensor: {
    name: string;
    euid: string;
    machine: string;
    appKey: string;
    senProf: string;
  }) => void;
}

export function AddDevice(props: AddDeviceProps) {
  const { onClose, open, addOptions } = props;
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
      newLengthErrors.DeviceEUI ||
      newLengthErrors["Application key"]
    ) {
      setFillError(false);
      setLengthErrors(newLengthErrors);
      return;
    }

    props.onAdd({
      name: values["Name"],
      euid: values["DeviceEUI"],
      machine: values["Machine"],
      appKey: values["Application key"],
      senProf: values["Sensor profile"],
    });

    setFillError(false);
    setLengthErrors({});
    onClose();
  };

  const deviceProfiles = [
    "Milesight EM300-CL",
    "Milesight EM300-DI",
    "Milesight EM300-MCS",
    "Milesight EM300-MLD",
    "Milesight EM300-SLD-ZLD",
    "Milesight EM300-TH",
    "Milesight EM310-TILT",
    "Milesight EM320-TH",
  ];

  const inputLength: Record<string, string> = {
    DeviceEUI: "16 characters",
    "Application key": "32 characters",
  };

  const inputLengthError: Record<string, string> = {
    DeviceEUI: "DeviceEUI must be 16 characters",
    "Application key": "Application key must be 32 characters",
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      maxWidth="lg"
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
        <CategoryHeader categories={addOptions} columns={addOptions.length}>
          {addOptions.map((option) => {
            // Textfield for all other than SenProf
            if (option !== "Sensor profile") {
              return (
                <TextField
                  sx={{
                    "& .MuiOutlinedInput-root": {
                      color: "primary.main", // input text color
                      "& fieldset": {
                        borderColor: "primary.main", // default border
                      },
                      "&:hover fieldset": {
                        borderColor: "primary.main", // hover border
                      },
                      "&.Mui-focused fieldset": {
                        borderColor: "primary.main", // focused border
                      },
                    },
                  }}
                  fullWidth
                  key={option}
                  placeholder={inputLength[option]}
                  helperText={
                    lengthErrors[option] ? inputLengthError[option] : ""
                  }
                  error={!!lengthErrors[option]}
                  value={values[option]}
                  onChange={(e) => {
                    let value = e.target.value;

                    // Only restrict DevEUI and AppKey
                    if (
                      option === "DeviceEUI" ||
                      option === "Application key"
                    ) {
                      value = value.replace(/[^a-zA-Z0-9]/g, "").toUpperCase();
                    }

                    setValues((prev) => ({
                      ...prev,
                      [option]: value,
                    }));
                  }}
                />
              );
            }
            // Dropdown selection for sensor profile
            return (
              <Select
                sx={{
                  color: "primary.main",
                  "& .MuiOutlinedInput-notchedOutline": {
                    borderColor: "primary.main",
                  },
                  "&:hover .MuiOutlinedInput-notchedOutline": {
                    borderColor: "primary.main",
                  },
                  "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
                    borderColor: "primary.main",
                  },
                  "& .MuiSelect-icon": {
                    color: "primary.main",
                  },
                }}
                fullWidth
                key={option}
                value={values[option]}
                displayEmpty
                onChange={(e) =>
                  setValues((prev) => ({
                    ...prev,
                    [option]: e.target.value,
                  }))
                }
              >
                {deviceProfiles.map((prof) => (
                  <MenuItem key={prof} value={prof}>
                    {prof}
                  </MenuItem>
                ))}
              </Select>
            );
          })}
        </CategoryHeader>
        {fillError && (
          <div style={{ color: "red", marginTop: 8 }}>
            All fields must be filled
          </div>
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
