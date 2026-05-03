import { useState } from "react";

import { ActionMenu } from "@/shared/ui/actionMenu";
import CircleIcon from "@mui/icons-material/Circle";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import { useTheme } from "@mui/material/styles";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoProps = {
  name: string;
  status: number;
  device_eui: string;
  lastSeenAt: string;
  description: string | null;
  onDelete: () => void;
  onEdit: () => void;
};

/**
 * Displays a single gateway row's data: online status, name, EUI, and last-seen time.
 * Includes an ActionMenu with Edit and Delete actions.
 * If a description is set, an info icon is shown that reveals it on hover.
 * @param props - Component props
 * @param props.name - Display name of the gateway
 * @param props.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param props.device_eui - LoRaWAN EUI of the gateway
 * @param props.lastSeenAt - Formatted timestamp of the last heartbeat
 * @param props.description - Optional description shown in a tooltip on hover
 * @param props.onDelete - Called when the user confirms deletion
 * @param props.onEdit - Called when the user clicks Edit in the action menu
 * @returns A set of grid-aligned cells with an action menu and delete confirmation dialog
 */
export function GatewayInfo({
  name,
  status,
  device_eui,
  lastSeenAt,
  description,
  onDelete,
  onEdit,
}: InfoProps) {
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
      <Typography>{device_eui}</Typography>
      <Typography>{lastSeenAt}</Typography>
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
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
    </>
  );
}
