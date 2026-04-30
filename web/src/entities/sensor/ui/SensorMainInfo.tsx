import { useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import Button from "@mui/material/Button";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";

import type { SensorApiResponse } from "@entities/sensor";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoMainProps = {
  sensor: SensorApiResponse;
  status: number;
  lastReading: string;
  onClick: () => void;
  onDelete: () => void;
  onEdit: (sensor: {
    id: string;
    name: string;
    deviceEui: string;
    electricitySensor: boolean;
    factory: string;
    factoryArea: string;
    productionResource: string | null;
    appKey: string;
    deviceProfile: string;
    voltage: number | null;
  }) => void;
};

/**
 * Renders the primary sensor row cells: status indicator, name, last reading, and an action menu.
 * @param root0 - Component props
 * @param root0.sensor - Sensor to be displayed
 * @param root0.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param root0.lastReading - Timestamp or relative time of the most recent sensor reading
 * @param root0.onClick - Called when the user clicks "Extra sensor info" to open the detail dialog
 * @param root0.onDelete - Called when the user clicks "Delete" to remove the sensor
 * @param root0.onEdit - Called when the user clicks "Edit" to edit the sensor
 * @returns The rendered sensor row cells
 */
export function SensorMainInfo({
  sensor,
  status,
  lastReading,
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
    {
      label: "Edit",
      onClick: () =>
        onEdit({
          id: sensor.id,
          name: sensor.name,
          deviceEui: sensor.deviceEui,
          electricitySensor: sensor.electricitySensor,
          factory: sensor.factory,
          factoryArea: sensor.factoryArea,
          productionResource: sensor.productionResource
            ? sensor.productionResource
            : null,
          appKey: sensor.appKey,
          deviceProfile: sensor.sensorProfile,
          voltage: sensor.voltage,
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
          color: statusColor(Number(status)),
          fontSize: 14,
          alignSelf: "center",
          filter: "drop-shadow(0 0 1px grey)",
        }}
      />
      <Typography>{sensor.name}</Typography>
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
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
    </>
  );
}
