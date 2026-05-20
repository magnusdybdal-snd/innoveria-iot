import { ActionMenu } from "@/shared/ui/actionMenu";
import { OutlinedButton } from "@/shared/ui/Button";
import Typography from "@mui/material/Typography";

type InfoProps = {
  name: string;
  address: string;
  createdAt: string;
  updatedAt: string;
  onDelete: () => void;
  onViewAreas: () => void;
};

/**
 * Displays a single factory row's data: name, address, dates, a link to its areas, and a delete action.
 * @param props - Component props
 * @param props.name - Display name of the factory
 * @param props.address - Main address of factory office
 * @param props.createdAt - Formatted date the factory was added
 * @param props.updatedAt - Formatted date of most recent change
 * @param props.onDelete - Callback to delete the factory
 * @param props.onViewAreas - Callback to drill down into the factory's areas
 * @returns A set of grid-aligned cells with a factory areas button and delete action menu
 */
export function FactoryInfo({
  name,
  address,
  createdAt,
  updatedAt,
  onDelete,
  onViewAreas,
}: InfoProps) {
  const menuItems = [{ label: "Delete", onClick: onDelete }];

  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{createdAt}</Typography>
      <Typography>{updatedAt}</Typography>
      <OutlinedButton
        onClick={(e) => {
          e.stopPropagation();
          onViewAreas();
        }}
      >
        Factory areas
      </OutlinedButton>
      <ActionMenu items={menuItems} />
    </>
  );
}
