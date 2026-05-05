import { useMemo } from "react";

import Autocomplete from "@mui/material/Autocomplete";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import FormControlLabel from "@mui/material/FormControlLabel";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import {
  BUCKET_UNIT_OPTIONS,
  type AggregationRule,
  type BucketUnit,
} from "@entities/context";
import { DropDownSelect } from "@shared/ui/DropDownSelect";
import type { GraphWidgetConfig } from "@widgets/graphWidget/model/types";

interface GraphWidgetSettingsProps {
  open: boolean;
  draft: GraphWidgetConfig;
  onChange: (draft: GraphWidgetConfig) => void;
  onCancel: () => void;
  onConfirm: (config: GraphWidgetConfig) => void;
  deviceOptions: { id: string; name: string }[];
  rules: AggregationRule[];
  measurementTypeOptions: { id: string; name: string }[];
  measurementTypesError: string | null;
  optionsError: string | null;
}

/**
 * Settings dialog for configuring a GraphWidget's data parameters.
 * Not dismissable by backdrop click or Escape — requires explicit Cancel or Confirm.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.draft - The in-flight config being edited
 * @param props.onChange - Called with the updated draft on any field change
 * @param props.onCancel - Called when the user cancels; draft is discarded by the parent
 * @param props.onConfirm - Called with the committed config when the user confirms
 * @param props.deviceOptions - Available device options for the device selector
 * @param props.rules - Available aggregation rules for the rule selector
 * @param props.measurementTypeOptions - Measurement type options (label includes slug)
 * @param props.measurementTypesError - Non-null when measurement types failed to load
 * @param props.optionsError - Error message if device/rule options failed to load
 * @returns The rendered settings dialog
 */
export function GraphWidgetSettings({
  open,
  draft,
  onChange,
  onCancel,
  onConfirm,
  deviceOptions,
  rules,
  measurementTypeOptions,
  measurementTypesError,
  optionsError,
}: GraphWidgetSettingsProps) {
  const selectedDevice =
    deviceOptions.find((o) => o.id === draft.deviceEui) ?? null;

  const measurementTypeNameBySlug = useMemo(() => {
    return new Map(measurementTypeOptions.map((o) => [o.id, o.name]));
  }, [measurementTypeOptions]);

  const ruleOptions = useMemo(() => {
    // If measurement types failed to load, fall back to rule.name only.
    if (measurementTypesError) {
      return rules.map((r) => ({
        id: r.id,
        name: `${r.name} (${r.aggregationMethod})`,
      }));
    }

    return rules.map((r) => {
      const mtLabel =
        measurementTypeNameBySlug.get(r.measurementType) ?? r.measurementType;
      return {
        id: r.id,
        // Required: include measurement type label/slug + aggregation method in the label.
        name: `${mtLabel} - ${r.aggregationMethod}`,
      };
    });
  }, [measurementTypeNameBySlug, measurementTypesError, rules]);

  return (
    <Dialog
      open={open}
      onClose={(_e, reason) => {
        if (reason === "backdropClick" || reason === "escapeKeyDown") return;
        onCancel();
      }}
      disableEscapeKeyDown
      maxWidth="sm"
      fullWidth
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.main",
        },
      }}
    >
      <DialogTitle>Widget Settings</DialogTitle>
      <DialogContent>
        <Box
          sx={{
            display: "grid",
            gridTemplateColumns: "1fr 1fr",
            gap: 2,
            pt: 1,
          }}
        >
          {/* Title — full width */}
          <TextField
            label="Widget title"
            size="small"
            value={draft.title}
            onChange={(e) => onChange({ ...draft, title: e.target.value })}
            sx={{ gridColumn: "1 / -1" }}
          />

          {/* Device */}
          <Autocomplete
            size="small"
            options={deviceOptions}
            getOptionLabel={(o) => o.name}
            value={selectedDevice}
            onChange={(_, v) => onChange({ ...draft, deviceEui: v?.id ?? "" })}
            renderInput={(params) => (
              <TextField {...params} label="Device" size="small" />
            )}
          />

          {/* Aggregation rule */}
          <DropDownSelect
            label="Aggregation rule"
            size="small"
            options={ruleOptions}
            value={draft.ruleId}
            onChange={(value) => onChange({ ...draft, ruleId: value })}
          />

          {/* From */}
          <TextField
            label="From"
            type="datetime-local"
            size="small"
            value={draft.from}
            onChange={(e) => onChange({ ...draft, from: e.target.value })}
            slotProps={{ inputLabel: { shrink: true } }}
          />

          {/* To */}
          <TextField
            label="To"
            type="datetime-local"
            size="small"
            value={draft.to}
            onChange={(e) => {
              const isFuture = new Date(e.target.value) > new Date();
              onChange({
                ...draft,
                to: e.target.value,
                useCurrentTime: isFuture ? true : draft.useCurrentTime,
              });
            }}
            slotProps={{ inputLabel: { shrink: true } }}
            disabled={draft.useCurrentTime}
          />

          {/* Use current time checkbox — under To field */}
          <FormControlLabel
            sx={{ gridColumn: "2" }}
            control={
              <Checkbox
                size="small"
                checked={draft.useCurrentTime}
                onChange={(e) =>
                  onChange({ ...draft, useCurrentTime: e.target.checked })
                }
              />
            }
            label="Use current time as end"
          />

          {/* Bucket interval value */}
          <TextField
            label="Interval"
            type="number"
            size="small"
            value={draft.bucketValue}
            onChange={(e) =>
              onChange({ ...draft, bucketValue: e.target.value })
            }
            slotProps={{ htmlInput: { min: 1 } }}
          />

          {/* Bucket interval unit */}
          <Select
            size="small"
            value={draft.bucketUnit}
            onChange={(e) =>
              onChange({ ...draft, bucketUnit: e.target.value as BucketUnit })
            }
          >
            {BUCKET_UNIT_OPTIONS.map((o) => (
              <MenuItem key={o.id} value={o.id}>
                {o.name}
              </MenuItem>
            ))}
          </Select>
        </Box>

        {optionsError && (
          <Typography
            variant="caption"
            color="error"
            sx={{ mt: 1, display: "block" }}
          >
            {optionsError}
          </Typography>
        )}

        {measurementTypesError && (
          <Typography
            variant="caption"
            color="error"
            sx={{ mt: 1, display: "block" }}
          >
            Measurement types failed to load. Showing all rules.
          </Typography>
        )}
      </DialogContent>

      <DialogActions>
        <Button size="small" onClick={onCancel}>
          Cancel
        </Button>
        <Button
          size="small"
          variant="contained"
          sx={{ "&:hover": { backgroundColor: "primary.main" } }}
          onClick={() => onConfirm(draft)}
        >
          Confirm
        </Button>
      </DialogActions>
    </Dialog>
  );
}
