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
  /*Format for single sensor schema info*/
}
/**
 * Renders the primary sensor schema row cells: default unit, description, display name, slug, and an action menu.
 * @param root0 - Component props
 * @param root0.payloadKey - Key to be identified
 * @param root0.measurementTypes - Types choose between to define the payload key
 * @param root0.value - Type and unit for each chosen measurement type
 * @param root0.onChange - Called with the chosen value when the input changes
 * @returns The rendered sensor schema type row cells
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
