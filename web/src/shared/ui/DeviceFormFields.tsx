import Box from "@mui/material/Box";
import Checkbox from "@mui/material/Checkbox";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { DropDownSelect } from "@shared/ui/DropDownSelect";

interface DeviceFormFieldsProps {
  options: string[];
  values: Record<string, string>;
  profileOptions: { id: string; name: string }[];
  factoryOptions: { id: string; name: string }[];
  factoryAreaOptions: { id: string; name: string }[];
  voltageOptions: { id: string; name: string }[];
  lengthErrors: Record<string, boolean>;
  lengthErrorMessages: Record<string, string>;
  inputHints: Record<string, string>;
  onChange: (option: string, value: string) => void;
}

const dropdownOptions: Record<string, string> = {
  "Sensor profile": "profileOptions",
  Factory: "factoryOptions",
  "Factory area": "factoryAreaOptions",
  Voltage: "voltageOptions",
};

const checkBoxes: string[] = ["Electricity sensor"];

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
};

/**
 * Renders a vertical list of form fields for device registration.
 * @param props - Component props
 * @param props.options - Field names to render
 * @param props.values - Current field values
 * @param props.profileOptions - Sensor profile choices for the dropdown
 * @param props.factoryOptions - Factory choices for the dropdown
 * @param props.factoryAreaOptions - Factory area choices for the dropdown
 * @param props.voltageOptions
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
  factoryOptions,
  factoryAreaOptions,
  voltageOptions,
  lengthErrors,
  lengthErrorMessages,
  inputHints,
  onChange,
}: DeviceFormFieldsProps) {
  const dropdownData: Record<string, { id: string; name: string }[]> = {
    profileOptions,
    factoryOptions,
    factoryAreaOptions,
    voltageOptions,
  };

  const maxLengths: Record<string, number> = {
    Name: 100,
    DeviceEUI: 16,
    "Application key": 32,
    Machine: 100,
  };

  return (
    <Box display="flex" flexDirection="column" gap={2}>
      {options.map((option) => (
        <Box key={option}>
          {(option != "Voltage" || values["Electricity sensor"] === "true") && (
            <Typography variant="body2" color="primary.main" mb={0.5}>
              {option}
            </Typography>
          )}
          {option in dropdownOptions ? (
            (option != "Voltage" ||
              values["Electricity sensor"] === "true") && (
              <DropDownSelect
                options={dropdownData[dropdownOptions[option]]}
                value={typeof values[option] === "string" ? values[option] : ""}
                onChange={(value) => onChange(option, value)}
              />
            )
          ) : checkBoxes.includes(option) ? (
            <Checkbox
              checked={values[option] === "true"}
              onChange={(_, checked) => onChange(option, String(checked))}
              slotProps={{
                input: { "aria-label": "controlled" },
              }}
            />
          ) : (
            <TextField
              sx={fieldSx}
              fullWidth
              placeholder={inputHints[option]}
              helperText={
                lengthErrors[option] ? lengthErrorMessages[option] : ""
              }
              error={!!lengthErrors[option]}
              value={typeof values[option] === "string" ? values[option] : ""}
              slotProps={{
                htmlInput: {
                  maxLength: maxLengths[option],
                },
              }}
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
