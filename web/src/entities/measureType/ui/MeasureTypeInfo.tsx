import { useState } from "react";

import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoMainProps = {
  defaultUnit: string;
  description: string;
  displayName: string;
  slug: string;
  deprecated: boolean;
  onDeprecate: () => void;
};

{
  /*Format for single measurement type info*/
}
/**
 * Renders the primary measurement type row cells: default unit, description, display name, slug, and an action menu.
 * @param root0 - Component props
 * @param root0.defaultUnit - Unit tied to the measurement type
 * @param root0.description - Description explaining the measurements type's function
 * @param root0.displayName - The displayed name of the measurements type
 * @param root0.slug - The id of the measurement type
 * @param root0.deprecated
 * @param root0.onDeprecate - Called when the user clicks "Deprecate" to deprecates the measurement type
 * @returns The rendered measurement type row cells
 */
export function MeasureTypeInfo({
  defaultUnit,
  description,
  displayName,
  slug,
  deprecated,
  onDeprecate,
}: InfoMainProps) {
  const [deprecateOpen, setDeprecateOpen] = useState(false);

  const handleDeprecateConfirm = () => {
    onDeprecate();
    setDeprecateOpen(false);
  };
  return (
    <>
      <Typography>{slug}</Typography>
      <Typography>{displayName}</Typography>
      <Typography>{description}</Typography>
      <Typography>{defaultUnit}</Typography>

      {!deprecated && (
        <Button
          variant="outlined"
          sx={{
            backgroundColor: "primary.main",
            color: "primary.dark",
            "&:hover": { backgroundColor: "primary.main" },
            borderRadius: 2,
            textTransform: "none",
            fontSize: 15,
            width: "fit-content",
          }}
          onClick={() => setDeprecateOpen(true)}
        >
          Deprecate
        </Button>
      )}
      <DeleteConfirmation
        open={deprecateOpen}
        onClose={() => setDeprecateOpen(false)}
        onConfirm={handleDeprecateConfirm}
      />
    </>
  );
}
