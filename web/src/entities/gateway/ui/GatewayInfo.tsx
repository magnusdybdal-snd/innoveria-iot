import { useState } from "react";

import { ActionMenu } from "@/shared/ui/actionMenu";
import CircleIcon from "@mui/icons-material/Circle";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";

import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

import type { GatewayApiResponse } from "../model/gatewaySchema";

type InfoProps = {
  gateway: GatewayApiResponse;
  lastSeenAt: string;
  onDelete: () => void;
  onEdit: (sensor: {
    id: string;
    name: string;
    deviceEui: string;
    factory: string;
    factoryArea: string;
  }) => void;
};

/**
 * Displays a single gateway row's data: online status, name, EUI, and last-seen time.
 * Includes an ActionMenu for renaming (local state only) and deleting the gateway.
 * Opens a Dialog to collect the new name on rename.
 * @param props - Component props
 * @param props.gateway - Gateway to be edited
 * @param props.lastSeenAt - Timestamp or relative time of the most recent gateway reading
 * @param props.onDelete - Called when the user clicks "Delete" to remove the sensor
 * @param props.onEdit - Called when the user clicks "Edit" to edit the sensor
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function GatewayInfo({
  gateway,
  lastSeenAt,
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
    {
      label: "Edit",
      onClick: () =>
        onEdit({
          id: gateway.id,
          name: gateway.name,
          deviceEui: gateway.gatewayEui,
          factory: gateway.factory,
          factoryArea: gateway.factoryArea,
        }),
    },
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
          color: statusColor(Number(gateway.status)),
          fontSize: 14,
          alignSelf: "center",
          filter: "drop-shadow(0 0 1px grey)",
        }}
      />
      <Typography>{gateway.name}</Typography>
      <Typography>{gateway.gatewayEui}</Typography>
      <Typography>{lastSeenAt}</Typography>
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
    </>
  );
}
