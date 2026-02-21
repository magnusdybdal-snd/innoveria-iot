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
  onAdd: (sensor: { name: string; euid: string; machine: string }) => void;
}

export function AddDevice(props: AddDeviceProps) {
  const { onClose, open, addOptions } = props;
  const [values, setValues] = useState<Record<string, string>>({});

  const handleClose = () => {
    onClose();
  };

  const handleSafeClose = () => {
    props.onAdd({
      name: values["Name"] ?? "",
      euid: values["DebEUI"] ?? "",
      machine: values["Machine"] ?? "",
    });

    onClose();
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
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
                "& .MuiInputLabel-root": {
                  color: "primary.main", // default label
                },
                "& .MuiInputLabel-root.Mui-focused": {
                  color: "primary.main", // focused label
                },
              }}
              key={option}
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
      </DialogContent>
      <DialogActions>
        <Button sx={{ color: "primary.dark" }} onClick={handleClose}>
          Close
        </Button>
        <Button
          sx={{ color: "primary.dark" }}
          onClick={handleSafeClose}
          autoFocus
        >
          Add
        </Button>
      </DialogActions>
    </Dialog>
  );
}
