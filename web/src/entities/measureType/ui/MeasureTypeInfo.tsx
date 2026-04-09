import { useState } from "react";

import Typography from "@mui/material/Typography";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoMainProps = {
  defaultUnit: string;
  description: string;
  displayName: string;
  slug: string;
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
 * @param root0.onDeprecate - Called when the user clicks "Deprecate" to deprecates the measurement type
 * @returns The rendered measurement type row cells
 */
export function MeasureTypeInfo({
  defaultUnit,
  description,
  displayName,
  slug,
  onDeprecate,
}: InfoMainProps) {
  const [deprecateOpen, setDeprecateOpen] = useState(false);

  const handleDeprecateConfirm = () => {
    onDeprecate();
    setDeprecateOpen(false);
  };

  const menuItems = [
    { label: "Deprecate", onClick: () => setDeprecateOpen(true) },
  ];

  return (
    <>
      <Typography>{slug}</Typography>
      <Typography>{displayName}</Typography>
      <Typography>{description}</Typography>
      <Typography>{defaultUnit}</Typography>
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deprecateOpen}
        onClose={() => setDeprecateOpen(false)}
        onConfirm={handleDeprecateConfirm}
      />
    </>
  );
}
