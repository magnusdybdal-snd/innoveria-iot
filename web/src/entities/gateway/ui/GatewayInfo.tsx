import { useState } from "react";

import { ActionMenu } from "@/shared/ui/actionMenu";
import CircleIcon from "@mui/icons-material/Circle";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";

import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";
import { RenameDialog } from "@shared/ui/RenameDialog";

type InfoProps = {
  name: string;
  status: number;
  device_eui: string;
  lastSeenAt: string;
  onDelete: () => void;
};

/**
 * Displays a single gateway row's data: online status, name, EUI, and last-seen time.
 * Includes an ActionMenu for renaming (local state only) and deleting the gateway.
 * Opens a Dialog to collect the new name on rename.
 * @param props - Component props
 * @param props.name - Display name of the gateway
 * @param props.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param props.device_eui
 * @param props.lastSeenAt
 * @param props.onDelete
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function GatewayInfo({
  name,
  status,
  device_eui,
  lastSeenAt,
  onDelete,
}: InfoProps) {
  const theme = useTheme();
  const [currentName, setCurrentName] = useState(name);
  const [editOpen, setEditOpen] = useState(false);
  const [editValue, setEditValue] = useState(name);
  const [deleteOpen, setDeleteOpen] = useState(false);

  const handleEditOpen = () => {
    setEditValue(currentName);
    setEditOpen(true);
  };

  const handleEditSave = () => {
    setCurrentName(editValue);
    setEditOpen(false);
  };

  const handleDeleteConfirm = () => {
    onDelete();
    setDeleteOpen(false);
  };

  const menuItems = [
    { label: "Rename", onClick: handleEditOpen },
    { label: "Delete", onClick: () => setDeleteOpen(true) },
  ];

  const statusColor = (status: number) => {
    const s = theme.palette.status;
    switch (status) {
      case 0:
        return s.online;
      case 1:
        return s.warning;
      case 2:
        return s.offline;
      default:
        return s.unknown;
    }
  };
  return (
    <>
      <CircleIcon
        sx={{
          color: statusColor(Number(status)),
          fontSize: 14,
          alignSelf: "center",
          filter: "drop-shadow(0 0 1px grey)",
        }}
      />
      <Typography>{currentName}</Typography>
      <Typography>{device_eui}</Typography>
      <Typography>{lastSeenAt}</Typography>
      <ActionMenu items={menuItems} />
      <RenameDialog
        open={editOpen}
        value={editValue}
        onChange={setEditValue}
        onClose={() => setEditOpen(false)}
        onSave={handleEditSave}
        label="Gateway name"
        title="Rename gateway"
      />
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
    </>
  );
}
