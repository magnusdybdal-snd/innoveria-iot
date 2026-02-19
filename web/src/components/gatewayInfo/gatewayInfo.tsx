import { useState } from "react";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogActions from "@mui/material/DialogActions";
import TextField from "@mui/material/TextField";
import CircleIcon from "@mui/icons-material/Circle";
import { ActionMenu } from "@/components/actionMenu";

type InfoProps = {
  name: string;
  online?: boolean;
  euid: string;
  lastSeen: string;
};

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
