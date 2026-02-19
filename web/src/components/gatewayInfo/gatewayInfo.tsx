import { useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import { ActionMenu } from "@/components/actionMenu";

type InfoProps = {
  name: string;
  online?: boolean;
  euid: string;
  lastSeen: string;
};

/**
 * Displays a single gateway row's data: online status, name, EUI, and last-seen time.
 * Includes an ActionMenu for renaming (local state only) and deleting the gateway.
 * Opens a Dialog to collect the new name on rename.
 *
 * @param props - Component props
 * @param props.name - Display name of the gateway
 * @param props.online - Whether the gateway is currently online
 * @param props.euid - EUI (Extended Unique Identifier) of the gateway
 * @param props.lastSeen - Human-readable time since last contact (e.g. "2 min")
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function GatewayInfo({ name, online, euid, lastSeen }: InfoProps) {
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

  return (
    <>
      <CircleIcon
        color={online ? "success" : "error"}
        sx={{ fontSize: 14, alignSelf: "center" }}
      />
      <Typography>{currentName}</Typography>
      <Typography>{euid}</Typography>
      <Typography>{lastSeen}</Typography>
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
