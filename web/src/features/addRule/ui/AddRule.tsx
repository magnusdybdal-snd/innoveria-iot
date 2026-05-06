import { useEffect, useMemo, useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import FormControlLabel from "@mui/material/FormControlLabel";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import type { CreateRuleRequest } from "@entities/context/model/contextSchema";
import { getMeasurementTypesAll } from "@entities/measurementType";
import { DropDownSelect } from "@shared/ui/DropDownSelect";

const AGGREGATION_METHOD_OPTIONS = [
  { id: "AVG", name: "AVG" },
  { id: "SUM", name: "SUM" },
  { id: "MAX", name: "MAX" },
  { id: "MIN", name: "MIN" },
];

interface AddRuleProps {
  open: boolean;
  companyId: string;
  onClose: () => void;
  onAdd: (rule: CreateRuleRequest) => Promise<void>;
  submitError?: string | null;
}

interface FormState {
  name: string;
  contextType: string;
  /** Measurement type slug (must exist in configured vocabulary). */
  measurementType: string;
  aggregationMethod: string;
  timeBucketMinutes: string;
  isActive: boolean;
}

const INITIAL_FORM: FormState = {
  name: "",
  contextType: "",
  measurementType: "",
  aggregationMethod: "AVG",
  timeBucketMinutes: "",
  isActive: true,
};

/**
 * Dialog form for creating a new aggregation rule.
 * @param props - Component props.
 * @param props.open - Whether the dialog is open.
 * @param props.companyId - UUID of the company the rule belongs to.
 * @param props.onClose - Callback invoked when the dialog is closed.
 * @param props.onAdd - Async callback invoked with the new rule data on submit.
 * @param props.submitError - Optional error message to display below the form fields.
 * @returns The rendered dialog.
 */
export function AddRule({
  open,
  companyId,
  onClose,
  onAdd,
  submitError,
}: AddRuleProps) {
  const [form, setForm] = useState<FormState>(INITIAL_FORM);
  const [fieldErrors, setFieldErrors] = useState<
    Partial<Record<keyof FormState, string>>
  >({});

  const [measurementTypes, setMeasurementTypes] = useState<
    { slug: string; displayName: string }[]
  >([]);
  const [measurementTypesError, setMeasurementTypesError] = useState<
    string | null
  >(null);
  const [prevOpen, setPrevOpen] = useState(open);

  // Avoid setState-in-effect lint: reset during render when the dialog opens.
  if (prevOpen !== open) {
    setPrevOpen(open);
    if (open) {
      setMeasurementTypes([]);
      setMeasurementTypesError(null);
    }
  }

  useEffect(() => {
    if (!open) return;
    getMeasurementTypesAll()
      .then((types) => {
        setMeasurementTypes(types);
      })
      .catch((err: unknown) => {
        setMeasurementTypes([]);
        setMeasurementTypesError(
          err instanceof Error
            ? err.message
            : "Failed to load measurement types.",
        );
      });
  }, [open]);

  const measurementTypeOptions = useMemo(() => {
    return measurementTypes
      .slice()
      .sort((a, b) => a.displayName.localeCompare(b.displayName))
      .map((mt) => ({
        id: mt.slug,
        name: `${mt.displayName} (${mt.slug})`,
      }));
  }, [measurementTypes]);

  const set = <K extends keyof FormState>(key: K, value: FormState[K]) => {
    setForm((prev) => ({ ...prev, [key]: value }));
    setFieldErrors((prev) => ({ ...prev, [key]: undefined }));
  };

  const validate = (): boolean => {
    const errors: Partial<Record<keyof FormState, string>> = {};
    if (!form.name.trim()) errors.name = "Name is required";
    if (!form.contextType.trim())
      errors.contextType = "Context type is required";
    if (!form.measurementType.trim())
      errors.measurementType = "Measurement type is required";
    const minutes = parseInt(form.timeBucketMinutes, 10);
    if (!form.timeBucketMinutes || isNaN(minutes) || minutes <= 0) {
      errors.timeBucketMinutes = "Must be a positive number";
    }
    setFieldErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleSubmit = () => {
    if (!validate()) return;
    onAdd({
      companyId,
      name: form.name,
      contextType: form.contextType,
      measurementType: form.measurementType,
      aggregationMethod: form.aggregationMethod,
      timeBucketMinutes: parseInt(form.timeBucketMinutes, 10),
      isActive: form.isActive,
    })
      .then(() => {
        setForm(INITIAL_FORM);
        setFieldErrors({});
      })
      .catch(() => {
        // error state is surfaced to the user via the submitError prop
      });
  };

  const handleClose = () => {
    setForm(INITIAL_FORM);
    setFieldErrors({});
    onClose();
  };

  const textFieldSx = {
    mb: 2,
    "& .MuiOutlinedInput-root": {
      color: "primary.main",
      "& .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
      "&:hover .MuiOutlinedInput-notchedOutline": {
        borderColor: "primary.main",
      },
      "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
        borderColor: "primary.main",
      },
    },
    "& .MuiInputLabel-root": { color: "primary.main" },
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      maxWidth="sm"
      fullWidth
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.contrastText",
        },
      }}
    >
      <DialogTitle sx={{ color: "primary.main" }}>Add context rule</DialogTitle>
      <DialogContent>
        <TextField
          label="Name"
          fullWidth
          sx={textFieldSx}
          value={form.name}
          onChange={(e) => set("name", e.target.value)}
          error={!!fieldErrors.name}
          helperText={fieldErrors.name}
        />
        <TextField
          label="Context type"
          fullWidth
          sx={textFieldSx}
          value={form.contextType}
          onChange={(e) => set("contextType", e.target.value)}
          error={!!fieldErrors.contextType}
          helperText={fieldErrors.contextType}
        />
        <Box sx={{ mb: 2 }}>
          <DropDownSelect
            label="Measurement type"
            size="small"
            options={measurementTypeOptions}
            value={form.measurementType}
            onChange={(value) => set("measurementType", value)}
          />
        </Box>
        {fieldErrors.measurementType && (
          <Typography variant="caption" color="error" sx={{ mt: 0.5, mb: 2 }}>
            {fieldErrors.measurementType}
          </Typography>
        )}
        {measurementTypesError && (
          <Typography variant="caption" color="error" sx={{ mt: 0.5, mb: 2 }}>
            {measurementTypesError}
          </Typography>
        )}
        <Typography variant="body2" sx={{ mb: 0.5, color: "primary.main" }}>
          Aggregation method
        </Typography>
        <DropDownSelect
          size="small"
          options={AGGREGATION_METHOD_OPTIONS}
          value={form.aggregationMethod}
          onChange={(value) => set("aggregationMethod", value)}
        />
        <TextField
          label="Time bucket (minutes)"
          fullWidth
          type="number"
          sx={{ ...textFieldSx, mt: 2 }}
          value={form.timeBucketMinutes}
          onChange={(e) => set("timeBucketMinutes", e.target.value)}
          error={!!fieldErrors.timeBucketMinutes}
          helperText={fieldErrors.timeBucketMinutes}
          slotProps={{ htmlInput: { min: 1 } }}
        />
        <FormControlLabel
          control={
            <Checkbox
              checked={form.isActive}
              onChange={(e) => set("isActive", e.target.checked)}
              sx={{ color: "primary.main" }}
            />
          }
          label="Active"
          sx={{ color: "primary.main" }}
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
          onClick={handleSubmit}
          autoFocus
        >
          Add
        </Button>
      </DialogActions>
    </Dialog>
  );
}
