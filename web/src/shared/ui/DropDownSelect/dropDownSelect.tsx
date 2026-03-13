import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";

interface DropDownSelectProps {
  options: { id: string; name: string }[];
  value: string;
  onChange: (value: string) => void;
}

const selectSx = {
  color: "primary.main",
  "& .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
  "&:hover .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
  "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
    borderColor: "primary.main",
  },
  "& .MuiSelect-icon": { color: "primary.main" },
};

/**
 * Generic styled dropdown select.
 * @param props - Component props
 * @param props.options - Items to display in the dropdown
 * @param props.value - Currently selected item id
 * @param props.onChange - Called with the selected item id on change
 * @returns The rendered select element
 */
export function DropDownSelect({
  options,
  value,
  onChange,
}: DropDownSelectProps) {
  return (
    <Select
      sx={selectSx}
      fullWidth
      value={value}
      displayEmpty
      onChange={(e) => onChange(e.target.value)}
    >
      {options.map((option) => (
        <MenuItem key={option.id} value={option.id}>
          {option.name}
        </MenuItem>
      ))}
    </Select>
  );
}
