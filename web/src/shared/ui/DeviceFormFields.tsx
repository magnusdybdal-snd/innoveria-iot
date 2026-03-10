import Box from "@mui/material/Box";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

interface DeviceFormFieldsProps {
  options: string[];
  values: Record<string, string>;
  profileOptions: { id: string; name: string }[];
  lengthErrors: Record<string, boolean>;
  lengthErrorMessages: Record<string, string>;
  inputHints: Record<string, string>;
  onChange: (option: string, value: string) => void;
}

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
};

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
 * Renders a vertical list of form fields for device registration.
 * @param props - Component props
 * @param props.options - Field names to render
 * @param props.values - Current field values
 * @param props.profileOptions - Sensor profile choices for the dropdown
 * @param props.lengthErrors - Map of field name to whether it has a length error
 * @param props.lengthErrorMessages - Map of field name to its error message
 * @param props.inputHints - Map of field name to its placeholder hint
 * @param props.onChange - Called with the field name and new value on change
 * @returns A vertical stack of labeled form fields
 */
export function DeviceFormFields({
  options,
  values,
  profileOptions,
  lengthErrors,
  lengthErrorMessages,
  inputHints,
  onChange,
}: DeviceFormFieldsProps) {
  return (
    <Box display="flex" flexDirection="column" gap={2}>
      {options.map((option) => (
        <Box key={option}>
          <Typography variant="body2" color="primary.main" mb={0.5}>
            {option}
          </Typography>

          {option === "Sensor profile" ? (
            <>
              <InputLabel id={`label-${option}`} sx={{ display: "none" }}>
                {option}
              </InputLabel>
              <Select
                labelId={`label-${option}`}
                sx={selectSx}
                fullWidth
                value={values[option] ?? ""}
                displayEmpty
                onChange={(e) => onChange(option, e.target.value)}
              >
                {profileOptions.map((prof) => (
                  <MenuItem key={prof.id} value={prof.id}>
                    {prof.name}
                  </MenuItem>
                ))}
              </Select>
            </>
          ) : (
            <TextField
              sx={fieldSx}
              fullWidth
              placeholder={inputHints[option]}
              helperText={
                lengthErrors[option] ? lengthErrorMessages[option] : ""
              }
              error={!!lengthErrors[option]}
              value={values[option] ?? ""}
              onChange={(e) => {
                let value = e.target.value;
                if (option === "DeviceEUI" || option === "Application key") {
                  value = value.replace(/[^a-fA-F0-9]/g, "");
                }
                onChange(option, value);
              }}
            />
          )}
        </Box>
      ))}
    </Box>
  );
}
