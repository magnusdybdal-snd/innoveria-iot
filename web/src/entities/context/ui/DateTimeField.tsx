import Box from "@mui/material/Box";
import Checkbox from "@mui/material/Checkbox";
import FormControlLabel from "@mui/material/FormControlLabel";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

const textFieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&:hover .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
      borderColor: "primary.main",
    },
  },
  "& input::-webkit-calendar-picker-indicator": { filter: "invert(1)" },
};

interface DateTimeFieldProps {
  label: string;
  value: string;
  onChange: (iso: string) => void;
  useNow?: boolean;
  onUseNowChange?: (useNow: boolean) => void;
}

/**
 * Renders a labeled datetime input with an optional "Use current time" checkbox.
 * @param props - Component props
 * @param props.label - Text label rendered above the input
 * @param props.value - Current datetime value as an ISO 8601 string
 * @param props.onChange - Called with the updated ISO string when the input changes
 * @param props.useNow - When true, the input is disabled and the current time is used
 * @param props.onUseNowChange - Called with the updated useNow flag when the checkbox changes
 * @returns Labeled datetime input with an optional now-checkbox
 */
export function DateTimeField({
  label,
  value,
  onChange,
  useNow,
  onUseNowChange,
}: DateTimeFieldProps) {
  const localValue = value ? value.slice(0, 16) : "";

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const iso = e.target.value ? new Date(e.target.value).toISOString() : "";
    onChange(iso);
  };

  return (
    <Box sx={{ flex: 1 }}>
      <Typography variant="body2" sx={{ mb: 0.5 }}>
        {label}
      </Typography>
      <TextField
        type="datetime-local"
        fullWidth
        disabled={useNow}
        value={useNow ? "" : localValue}
        onChange={handleChange}
        sx={textFieldSx}
      />
      {onUseNowChange !== undefined && (
        <FormControlLabel
          control={
            <Checkbox
              checked={useNow ?? false}
              onChange={(e) => onUseNowChange(e.target.checked)}
              size="small"
            />
          }
          label={<Typography variant="caption">Use current time</Typography>}
          sx={{ mt: 0.5, ml: 0 }}
        />
      )}
    </Box>
  );
}
