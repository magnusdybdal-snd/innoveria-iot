import { useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import Button from "@mui/material/Button";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";
import { RenameDialog } from "@shared/ui/RenameDialog";

type InfoMainProps = {
  name: string;
  status: number;
  lastReading: string;
  onClick: () => void;
  onDelete: () => void;
};

{
  /*Format for main single sensor info*/
}
/**
 * Renders the primary sensor row cells: status indicator, name, last reading, and an action menu.
 * @param root0 - Component props
 * @param root0.name - Display name of the sensor
 * @param root0.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param root0.lastReading - Timestamp or relative time of the most recent sensor reading
 * @param root0.onClick - Called when the user clicks "Extra sensor info" to open the detail dialog
 * @param root0.onDelete - Called when the user clicks "Delete" to remove the sensor
 * @returns The rendered sensor row cells
 */
export function SensorMainInfo({
  name,
  status,
  lastReading,
  onClick,
  onDelete,
}: InfoMainProps) {
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
      <Typography>{lastReading}</Typography>
      <Button
        variant="outlined"
        sx={{
          backgroundColor: "primary.main",
          color: "primary.dark",
          "&:hover": { backgroundColor: "primary.main" },
          borderRadius: 2,
          textTransform: "none",
          fontSize: 15,
        }}
        onClick={onClick}
      >
        Extra sensor info
      </Button>
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
