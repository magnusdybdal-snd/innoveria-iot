import { BUCKET_UNIT_OPTIONS } from "@entities/context/lib";
import type { BucketUnit } from "@entities/context/model/contextSchema";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { DropDownSelect } from "@shared/ui/DropDownSelect";

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

interface BucketIntervalFieldProps {
  value: string;
  onValueChange: (value: string) => void;
  unit: BucketUnit;
  onUnitChange: (unit: BucketUnit) => void;
}

/**
 * Renders a numeric input paired with a time unit selector for configuring a bucket interval.
 * @param props - Component props
 * @param props.value - Numeric interval value as a string
 * @param props.onValueChange - Called with the updated numeric value string when the input changes
 * @param props.unit - Currently selected time unit for the interval
 * @param props.onUnitChange - Called with the updated BucketUnit when the unit selection changes
 * @returns Time interval input with a paired unit dropdown
 */
export function BucketIntervalField({
  value,
  onValueChange,
  unit,
  onUnitChange,
}: BucketIntervalFieldProps) {
  return (
    <Box>
      <Typography variant="body2" sx={{ mb: 0.5 }}>
        Time interval
      </Typography>
      <Box sx={{ display: "flex", gap: 1 }}>
        <TextField
          type="number"
          value={value}
          onChange={(e) => onValueChange(e.target.value)}
          slotProps={{ htmlInput: { min: 1 } }}
          sx={{ ...textFieldSx, width: 100 }}
        />
        <Box sx={{ flex: 1 }}>
          <DropDownSelect
            options={BUCKET_UNIT_OPTIONS}
            value={unit}
            onChange={(v) => onUnitChange(v as BucketUnit)}
          />
        </Box>
      </Box>
    </Box>
  );
}
