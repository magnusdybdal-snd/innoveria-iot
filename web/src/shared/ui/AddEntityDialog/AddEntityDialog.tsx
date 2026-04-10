import { useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
};

export interface AddEntityDialogProps {
  open: boolean;
  title: string;
  fields: string[];
  optionalFields?: string[];
  onClose: () => void;
  onSubmit: (values: Record<string, string>) => Promise<void>;
  submitError?: string | null;
}

/**
 * Generic dialog for adding a new entity via a set of text fields.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.title - Title displayed at the top of the dialog
 * @param props.fields - Field labels to render as text inputs
 * @param props.optionalFields
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.onSubmit - Called with a map of field label to value when the user confirms
 * @param props.submitError - Error message to display if the submission fails
 * @returns The rendered add-entity dialog
 */
export function AddEntityDialog({
  open,
  title,
  fields,
  optionalFields,
  onClose,
  onSubmit,
  submitError,
}: AddEntityDialogProps) {
  const [values, setValues] = useState<Record<string, string>>({});
  const [fillError, setFillError] = useState(false);

  const handleClose = () => {
    onClose();
  };

  const handleSubmit = () => {
    const allFilled = fields.every((field) => {
      const isOptional = optionalFields?.includes(field);
      if (isOptional) return true;

      return (values[field] ?? "").trim() !== "";
    });

    if (!allFilled) {
      setFillError(true);
      return;
    }

    setFillError(false);
    onSubmit(values)
      .then(() => {
        setValues({});
      })
      .catch(() => {});
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      maxWidth="sm"
      fullWidth
      aria-labelledby="add-entity-dialog-title"
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.contrastText",
        },
      }}
    >
      <DialogTitle id="add-entity-dialog-title" sx={{ color: "primary.main" }}>
        {title}
      </DialogTitle>
      <DialogContent>
        <Box display="flex" flexDirection="column" gap={2}>
          {fields.map((field) => (
            <Box key={field}>
              <Typography variant="body2" color="primary.main" mb={0.5}>
                {field}
              </Typography>
              <TextField
                sx={fieldSx}
                fullWidth
                value={values[field] ?? ""}
                onChange={(e) => {
                  let value = e.target.value;

                  if (field === "Slug") {
                    value = value
                      .replace(/ /g, "_") // spaces → underscores
                      .replace(/[^a-zA-Z_]/g, "") // only letters
                      .toLowerCase();
                  }

                  setValues((prev) => ({ ...prev, [field]: value }));
                }}
              />
            </Box>
          ))}
        </Box>
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
          onClick={handleSubmit}
          autoFocus
        >
          Add
        </Button>
      </DialogActions>
    </Dialog>
  );
}
