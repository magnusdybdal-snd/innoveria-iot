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
 * @returns The rendered autocomplete select element
 */
export function DropDownSelect({
  options,
  value,
  onChange,
}: DropDownSelectProps) {
  const selected = options.find((o) => o.id === value);

  return (
    <Autocomplete
      sx={autocompleteSx}
      fullWidth
      disableClearable
      options={options}
      getOptionKey={(option) => option.id}
      getOptionLabel={(option) => option.name}
      isOptionEqualToValue={(option, val) => option.id === val.id}
      value={selected ?? null}
      onChange={(_, newValue) => {
        if (newValue) onChange(newValue.id);
      }}
      renderInput={(params) => <TextField {...params} />}
    />
  );
}
