import type { MeasurementTypeApiResponse } from "@entities/measurementType";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import Typography from "@mui/material/Typography";

type InfoMainProps = {
  payloadKey: string;
  measurementTypes: MeasurementTypeApiResponse[];
  value?: {
    measurementType: string;
    unit: string;
  };
  onChange: (value: { measurementType: string; unit: string }) => void;
};

{
  /*Format for single measurement type info*/
}
/**
 * Renders the primary measurement type row cells: default unit, description, display name, slug, and an action menu.
 * @param root0 - Component props
 * @param root0.payloadKey - Unit tied to the measurement type
 * @param root0.measurementTypes - Description explaining the measurements type's function
 * @param root0.value - Description explaining the measurements type's function
 * @param root0.onChange - Description explaining the measurements type's function
 * @returns The rendered measurement type row cells
 */
export function FixedSensorSchema({
  payloadKey,
  measurementTypes,
  value,
  onChange,
}: InfoMainProps) {
  const handleChange = (slug: string) => {
    const selected = measurementTypes.find((m) => m.slug === slug);

    if (!selected) return;

    onChange({
      measurementType: selected.slug,
      unit: selected.defaultUnit,
    });
  };

  return (
    <>
      <Typography>{payloadKey}</Typography>
      <Select
        value={value?.measurementType ?? ""}
        onChange={(e) => handleChange(e.target.value)}
        displayEmpty
      >
        {measurementTypes.map((option) => (
          <MenuItem key={option.slug} value={option.slug}>
            {option.displayName}
          </MenuItem>
        ))}
      </Select>
      <Typography>{value?.unit ?? "-"}</Typography>
    </>
  );
}
