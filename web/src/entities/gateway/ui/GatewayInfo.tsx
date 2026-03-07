import { useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import { useTheme } from "@mui/material/styles";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import { ActionMenu } from "@/shared/ui/actionMenu";

type InfoProps = {
  name: string;
  status: number;
  device_eui: string;
  lastSeenAt: string;
};

/**
 * Displays a single gateway row's data: online status, name, EUI, and last-seen time.
 * Includes an ActionMenu for renaming (local state only) and deleting the gateway.
 * Opens a Dialog to collect the new name on rename.
 * @param props - Component props
 * @param props.name - Display name of the gateway
 * @param props.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param props.euid - EUI (Extended Unique Identifier) of the gateway
 * @param props.lastSeen - Human-readable time since last contact (e.g. "2 min ago")
 * @param props.device_eui
 * @param props.lastSeenAt
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function GatewayInfo({
  name,
  status,
  device_eui,
  lastSeenAt,
}: InfoProps) {
  const theme = useTheme();
  const [currentName, setCurrentName] = useState(name);
  const [editOpen, setEditOpen] = useState(false);
  const [editValue, setEditValue] = useState(name);

  const handleEditOpen = () => {
    setEditValue(currentName);
    setEditOpen(true);
  };

  const handleEditSave = () => {
    setCurrentName(editValue);
    setEditOpen(false);
  };

  const menuItems = [
    { label: "Rename", onClick: handleEditOpen },
    { label: "Delete", onClick: () => {} },
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
      <Dialog open={editOpen} onClose={() => setEditOpen(false)}>
        <DialogTitle>Rename gateway</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            value={editValue}
            onChange={(e) => setEditValue(e.target.value)}
            label="Gateway name"
            fullWidth
            sx={{ mt: 1 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setEditOpen(false)}>Cancel</Button>
          <Button variant="contained" onClick={handleEditSave}>
            Save
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
