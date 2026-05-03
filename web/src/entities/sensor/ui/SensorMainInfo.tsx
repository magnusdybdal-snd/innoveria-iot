import { useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import Button from "@mui/material/Button";
import { useTheme } from "@mui/material/styles";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoMainProps = {
  name: string;
  status: number;
  lastReading: string;
  description: string | null;
  onClick: () => void;
  onDelete: () => void;
  onEdit: () => void;
};

/**
 * Renders the primary sensor row cells: status indicator, name, last reading, and an action menu.
 * @param props - Component props
 * @param props.name - Display name of the sensor
 * @param props.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param props.lastReading - Timestamp of the most recent sensor reading
 * @param props.description - Optional description shown in a tooltip on hover
 * @param props.onClick - Called when the user clicks "Extra sensor info" to open the detail dialog
 * @param props.onDelete - Called when the user confirms deletion
 * @param props.onEdit - Called when the user clicks Edit in the action menu
 * @returns The rendered sensor row cells
 */
export function SensorMainInfo({
  name,
  status,
  lastReading,
  description,
  onClick,
  onDelete,
  onEdit,
}: InfoMainProps) {
  const theme = useTheme();
  const [deleteOpen, setDeleteOpen] = useState(false);

  const handleDeleteConfirm = () => {
    onDelete();
    setDeleteOpen(false);
  };

  const menuItems = [
    { label: "Edit", onClick: onEdit },
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
      <Typography>{name}</Typography>
      <Typography>{lastReading}</Typography>
      {description ? (
        <Tooltip title={description} enterDelay={800} arrow>
          <InfoOutlinedIcon
            sx={{
              fontSize: 18,
              alignSelf: "center",
              color: "primary.main",
              cursor: "default",
            }}
          />
        </Tooltip>
      ) : (
        <span />
      )}
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
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
    </>
  );
}
