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
 * Renders a single payload key row with a measurement type dropdown and unit display.
 * @param root0 - Component props
 * @param root0.payloadKey - The raw payload key name from the sensor
 * @param root0.measurementTypes - Available measurement types to select from
 * @param root0.value - Currently selected measurement type and unit
 * @param root0.onChange - Called with the new measurement type and unit when selection changes
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
