import Typography from "@mui/material/Typography";

import { ActionMenu } from "@shared/ui/actionMenu";

type InfoProps = {
  id: string;
  name: string;
  email: string;
  role: string;
  createdAt: string;
  onDelete: (id: string) => void;
};

/**
 * Displays a single user row: name, email, role, created date, and an action menu.
 * @param props - Component props
 * @param props.id - User ID used for delete
 * @param props.name - Display name of the user
 * @param props.email - Email address of the user
 * @param props.role - Role assigned to the user
 * @param props.createdAt - Formatted creation timestamp
 * @param props.onDelete - Called with the user ID when delete is selected
 * @returns A set of grid-aligned cells with an action menu
 */
export function UserInfo({
  id,
  name,
  email,
  role,
  createdAt,
  onDelete,
}: InfoProps) {
  const menuItems = [
    { label: "Edit", disabled: true },
    { label: "Delete", onClick: () => onDelete(id) },
  ];

  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{email}</Typography>
      <Typography>{role}</Typography>
      <Typography>{createdAt}</Typography>
      <ActionMenu items={menuItems} />
    </>
  );
}
