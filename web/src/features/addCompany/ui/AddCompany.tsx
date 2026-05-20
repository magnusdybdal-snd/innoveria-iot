import { useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";

import { CategoryHeader } from "@shared/ui/CategoryHeader";

export interface AddCompanyProps {
  open: boolean;
  onClose: () => void;
  addOptions: string[];
  onAdd: (sensor: { name: string; address: string }) => void;
  submitError?: string | null;
}

/**
 * Modal dialog for registering a new company.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.addOptions - Field names to render as inputs inside the dialog
 * @param props.onAdd - Called with the validated sensor data when the user confirms
 * @returns The rendered add-company dialog
 */
export function AddCompany(props: AddCompanyProps) {
  const { onClose, open, addOptions, submitError } = props;
  const [values, setValues] = useState<Record<string, string>>({});
  const [fillError, setFillError] = useState(false);

  const handleClose = () => {
    onClose();
  };

  const handleSafeClose = () => {
    const allFilled = addOptions.every(
      (option) => (values[option] ?? "").trim() !== "",
    );

    if (!allFilled) {
      setFillError(true);
      return;
    }

    props.onAdd({
      name: values["Name"],
      address: values["Address"],
    });

    setFillError(false);
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
        {"Insert company info"}
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
              }}
              fullWidth
              key={option}
              value={values[option]}
              onChange={(e) => {
                const value = e.target.value;

                setValues((prev) => ({
                  ...prev,
                  [option]: value,
                }));
              }}
            />
          ))}
        </CategoryHeader>
        {fillError && (
          <div style={{ color: "red", marginTop: 8 }}>
            All fields must be filled
          </div>
        )}
        {submitError && (
          <div style={{ color: "red", marginTop: 8 }}>{submitError}</div>
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
