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
import { green, grey, red, yellow } from "@mui/material/colors";
type InfoProps = {
  name: string;
  status: number;
  euid: string;
  machine: string;
};

{
  /*Format for single sensor info*/
}
export function SensorInfo({ name, status, euid, machine }: InfoProps) {
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

  const getColor = (status: number) => {
    switch (status) {
      case 0:
        return green[500];
      case 1:
        return yellow[500];
      case 2:
        return red[500];
      default:
        return grey[500];
    }
  };
  return (
    <>
      <CircleIcon
        sx={{
          color: getColor(Number(status)),
          fontSize: 14,
          alignSelf: "center",
          filter: "drop-shadow(0 0 1px grey)",
        }}
      ></CircleIcon>
      <Typography>{currentName}</Typography>
      <Typography>{euid}</Typography>
      <Typography>{machine}</Typography>
      <ActionMenu items={menuItems} />
      <Dialog open={editOpen} onClose={() => setEditOpen(false)}>
        <DialogTitle>Rename sensor</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            value={editValue}
            onChange={(e) => setEditValue(e.target.value)}
            label="Sensor name"
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
