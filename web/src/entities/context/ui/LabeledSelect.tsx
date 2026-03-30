import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { DropDownSelect } from "@shared/ui/DropDownSelect";

interface LabeledSelectProps {
  label: string;
  options: { id: string; name: string }[];
  value: string;
  onChange: (value: string) => void;
  flex?: number;
  error?: string;
}

/**
 * Renders a labeled dropdown select with an optional flex layout weight.
 * @param props - Component props
 * @param props.label - Text label rendered above the select
 * @param props.options - List of selectable options with id and display name
 * @param props.value - Currently selected option id
 * @param props.onChange - Called with the newly selected option id
 * @param props.flex - CSS flex grow value applied to the wrapping Box
 * @param props.error
 * @returns Labeled dropdown select element
 */
export function LabeledSelect({
  label,
  options,
  value,
  onChange,
  flex,
  error,
}: LabeledSelectProps) {
  return (
    <Box sx={{ flex }}>
      <Typography variant="body2" sx={{ mb: 0.5 }}>
        {label}
      </Typography>
      <DropDownSelect options={options} value={value} onChange={onChange} />
      {error && (
        <Typography
          variant="caption"
          color="error"
          sx={{ mt: 0.5, display: "block" }}
        >
          {error}
        </Typography>
      )}
    </Box>
  );
}
