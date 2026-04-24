import { useState } from "react";

import type { MeasurementTypeApiResponse } from "@entities/measurementType";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import Typography from "@mui/material/Typography";

type InfoMainProps = {
  payloadKey: string;
  measurementTypes: MeasurementTypeApiResponse[];
};

{
  /*Format for single measurement type info*/
}
/**
 * Renders the primary measurement type row cells: default unit, description, display name, slug, and an action menu.
 * @param root0 - Component props
 * @param root0.payloadKey - Unit tied to the measurement type
 * @param root0.measurementTypes - Description explaining the measurements type's function
 * @returns The rendered measurement type row cells
 */
export function FixedSensorSchema({
  payloadKey,
  measurementTypes,
}: InfoMainProps) {
  const [type, setType] = useState<string>("");
  const selectedType = measurementTypes.find((t) => t.slug === type);

  return (
    <>
      <Typography>{payloadKey}</Typography>
      <Select
        value={type}
        onChange={(e) => setType(e.target.value)}
        displayEmpty
      >
        {measurementTypes.map((option) => (
          <MenuItem key={option.slug} value={option.slug}>
            {option.displayName}
          </MenuItem>
        ))}
      </Select>
      <Typography>{selectedType?.defaultUnit}</Typography>
    </>
  );
}
