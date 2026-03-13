import Typography from "@mui/material/Typography";

import { ActionMenu } from "@/shared/ui/actionMenu";

type InfoProps = {
  name: string;
  address: string;
  createdAt: string;
  updatedAt: string;
  onDelete: () => void;
};

/**
 * Displays a single factory row's data: id, name, address, date created, and date updated.
 * Includes a button for adding the first factory admin user.
 * @param props - Component props
 * @param props.name - Display name of the factory
 * @param props.address - Main address of factory office
 * @param props.createdAt - Formatted date the factory was added
 * @param props.updatedAt - Formatted date of most recent change
 * @param props.onDelete - Callback to delete a factory
 * @returns A set of grid-aligned cells with an add user button
 */
export function FactoryInfo({
  name,
  address,
  createdAt,
  updatedAt,
  onDelete,
}: InfoProps) {
  const menuItems = [{ label: "Delete", onClick: onDelete }];

  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{createdAt}</Typography>
      <Typography>{updatedAt}</Typography>
      <ActionMenu items={menuItems} />
    </>
  );
}
