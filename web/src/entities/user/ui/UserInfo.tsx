import { useState } from "react";

import Typography from "@mui/material/Typography";

import type { UserRole } from "@entities/user";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

const ROLE_LABELS: Record<UserRole, string> = {
  FACTORY_WORKER: "User",
  FACTORY_SUPERUSER: "Superuser",
  PLATFORM_ADMIN: "Admin",
};

type InfoProps = {
  id: string;
  name: string;
  email: string;
  role: string;
  createdAt: string;
  onDelete: (id: string) => void;
  onEdit: (user: { id: string; name: string; email: string }) => void;
};

/**
 * Displays a single user row: name, email, role, created date, and an action menu.
 * @param props - Component props
 * @param props.id - User ID used for delete and edit
 * @param props.name - Display name of the user
 * @param props.email - Email address of the user
 * @param props.role - Role assigned to the user
 * @param props.createdAt - Formatted creation timestamp
 * @param props.onDelete - Called with the user ID when delete is selected
 * @param props.onEdit - Called with the user data when edit is selected
 * @returns A set of grid-aligned cells with an action menu
 */
export function UserInfo({
  id,
  name,
  email,
  role,
  createdAt,
  onDelete,
  onEdit,
}: InfoProps) {
  const [deleteOpen, setDeleteOpen] = useState(false);

  const menuItems = [
    { label: "Edit", onClick: () => onEdit({ id, name, email }) },
    { label: "Delete", onClick: () => setDeleteOpen(true) },
  ];

  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{email}</Typography>
      <Typography>{ROLE_LABELS[role] ?? role}</Typography>
      <Typography>{createdAt}</Typography>
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={() => {
          onDelete(id);
          setDeleteOpen(false);
        }}
      />
    </>
  );
}
