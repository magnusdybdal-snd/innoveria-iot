import { ActionMenu } from "@/shared/ui/actionMenu";
import Typography from "@mui/material/Typography";

type FactoryAreaInfoProps = {
  name: string;
  description: string | null;
  onDelete: () => void;
};

/**
 * Displays a single factory area row: name, description, and a delete action.
 * @param props - Component props
 * @param props.name - Display name of the factory area
 * @param props.description - Description of the factory area, or null if none has been set
 * @param props.onDelete - Callback to delete the factory area
 * @returns A set of grid-aligned cells with a delete action menu
 */
export function FactoryAreaInfo({
  name,
  description,
  onDelete,
}: FactoryAreaInfoProps) {
  const menuItems = [{ label: "Delete", onClick: onDelete }];

  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{description ?? "—"}</Typography>
      <ActionMenu items={menuItems} />
    </>
  );
}
