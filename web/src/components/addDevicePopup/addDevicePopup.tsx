import { useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
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
    devProf: string;
  }) => void;
}

export function AddDevice(props: AddDeviceProps) {
  const { onClose, open, addOptions } = props;
  const [values, setValues] = useState<Record<string, string>>({});
  const [showError, setShowError] = useState(false);

  const handleClose = () => {
    onClose();
  };

  const handleSafeClose = () => {
    if (!allFilled) {
      setShowError(true);
      return;
    }

    props.onAdd({
      name: values["Name"] ?? "",
      euid: values["DeviceEUI"] ?? "",
      machine: values["Machine"] ?? "",
      appKey: values["AppKey"] ?? "",
      devProf: values["DevProf"] ?? "",
    });

    setShowError(false);
    onClose();
  };

  const allFilled = addOptions.every(
    (option) => (values[option] ?? "").trim() !== "",
  );

  const inputLength: Record<string, string> = {
    DeviceEUI: "16 characters",
    "Application key": "32 characters",
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
          {addOptions.map((option) => (
            <TextField
              fullWidth
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
              key={option}
              placeholder={inputLength[option] ?? ""}
              value={values[option] ?? ""}
              onChange={(e) =>
                setValues((prev) => ({
                  ...prev,
                  [option]: e.target.value,
                }))
              }
            />
          ))}
        </CategoryHeader>
        {showError && (
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
