import { useState } from "react";

import Typography from "@mui/material/Typography";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoMainProps = {
  defaultUnit: string;
  deprecated: boolean;
  description: string;
  displayName: string;
  slug: string;
  onDeprecate: () => void;
};

{
  /*Format for single measurement type info*/
}
/**
 * Renders the primary measurement type row cells: status indicator, name, last reading, and an action menu.
 * @param root0 - Component props
 * @param root0.defaultUnit - Display name of the measurement type
 * @param root0.deprecated - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param root0.description - Timestamp or relative time of the most recent measurement type reading
 * @param root0.displayName - Timestamp or relative time of the most recent measurement type reading
 * @param root0.slug - Timestamp or relative time of the most recent measurement type reading
 * @param root0.onDeprecate - Called when the user clicks "Deprecate" to deprecates the measurement type
 * @returns The rendered measurement type row cells
 */
export function MeasureTypeMainInfo({
  defaultUnit,
  deprecated,
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
      <Typography>{deprecated}</Typography>
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deprecateOpen}
        onClose={() => setDeprecateOpen(false)}
        onConfirm={handleDeprecateConfirm}
      />
    </>
  );
}
