import Autocomplete from "@mui/material/Autocomplete";
import TextField from "@mui/material/TextField";

interface Option {
  id: string;
  name: string;
}

interface DropDownSelectProps {
  options: Option[];
  value: string;
  onChange: (value: string) => void;
  /** Optional TextField label (renders inside the input). */
  label?: string;
  /** Optional sizing to match surrounding MUI fields. */
  size?: "small" | "medium";
}

const autocompleteSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&:hover .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
      borderColor: "primary.main",
    },
  },
  "& .MuiSvgIcon-root": { color: "primary.main" },
};

/**
 * Generic styled dropdown select with keyboard/type-to-filter support.
 * Only emits ids that exist in the options array.
 * @param props - Component props
 * @param props.options - Items to display in the dropdown
 * @param props.value - Currently selected item id
 * @param props.onChange - Called with the selected item id on change
 * @param props.label - Optional input label
 * @param props.size - Optional MUI input size
 * @returns The rendered autocomplete select element
 */
export function DropDownSelect({
  options,
  value,
  onChange,
  label,
  size = "medium",
}: DropDownSelectProps) {
  const selected = options.find((o) => o.id === value) ?? null;

  return (
    <Autocomplete
      sx={autocompleteSx}
      fullWidth
      size={size}
      options={options}
      getOptionKey={(option) => option.id}
      getOptionLabel={(option) => option.name}
      isOptionEqualToValue={(option, val) => option.id === val.id}
      value={selected ?? null}
      onChange={(_, newValue) => {
        onChange(newValue?.id ?? "");
      }}
      renderInput={(params) => <TextField {...params} label={label} />}
    />
  );
}
