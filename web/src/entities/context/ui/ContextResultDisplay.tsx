import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";

import type { ContextDataResponse } from "@entities/context/model/contextSchema";

import { BucketBarChart } from "./BucketBarChart";
import { BucketLineChart } from "./BucketLineChart";
import { LabeledSelect } from "./LabeledSelect";

const textFieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&:hover .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
      borderColor: "primary.main",
    },
  },
};

const fieldOptions: { id: keyof ContextDataResponse; name: string }[] = [
  { id: "deviceEui", name: "Device EUI" },
  { id: "companyId", name: "Company ID" },
  { id: "contextType", name: "Context type" },
  { id: "unit", name: "Unit" },
  { id: "periodStart", name: "Period start" },
  { id: "periodEnd", name: "Period end" },
  { id: "totalValue", name: "Total value" },
  { id: "buckets", name: "Buckets (JSON)" },
  { id: "calculatedAt", name: "Calculated at" },
];

interface ContextResultDisplayProps {
  result: ContextDataResponse[];
  selectedField: keyof ContextDataResponse;
  onFieldChange: (field: keyof ContextDataResponse) => void;
}

/**
 * Displays context query results as formatted text for a user-selected response field.
 * @param props - Component props
 * @param props.result - Array of context data responses returned from the API
 * @param props.selectedField - Key of ContextDataResponse currently chosen for display
 * @param props.onFieldChange - Called with the newly selected field key when the user changes the field
 * @returns Field selector and read-only text output of the selected field values
 */
export function ContextResultDisplay({
  result,
  selectedField,
  onFieldChange,
}: ContextResultDisplayProps) {
  const outputText = result
    .map((item) => {
      const val = item[selectedField];
      return typeof val === "object"
        ? JSON.stringify(val, null, 2)
        : String(val);
    })
    .join("\n---\n");

  const buckets = result[0]?.buckets ?? [];
  const unit = result[0]?.unit;

  return (
    <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
      {buckets.length > 0 && (
        <>
          <BucketBarChart buckets={buckets} unit={unit} />
          <BucketLineChart buckets={buckets} unit={unit} />
        </>
      )}
      <LabeledSelect
        label="Field to display"
        options={fieldOptions.map((f) => ({ id: f.id, name: f.name }))}
        value={selectedField}
        onChange={(v) => onFieldChange(v as keyof ContextDataResponse)}
      />
      <TextField
        multiline
        fullWidth
        minRows={6}
        value={outputText}
        slotProps={{ input: { readOnly: true } }}
        sx={{
          ...textFieldSx,
          "& .MuiInputBase-input": {
            fontFamily: "monospace",
            fontSize: "0.75rem",
          },
        }}
      />
    </Box>
  );
}
