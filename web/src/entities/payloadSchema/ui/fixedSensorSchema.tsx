import type { MeasurementTypeApiResponse } from "@entities/measurementType";
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
  return (
    <>
      <Typography>{payloadKey}</Typography>
      <Typography>{measurementTypes[0].displayName}</Typography>
      <Typography>{measurementTypes[0].defaultUnit}</Typography>
    </>
  );
}
