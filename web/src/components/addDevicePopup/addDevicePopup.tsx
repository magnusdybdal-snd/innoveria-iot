import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";

import { CategoryHeader } from "@/components/CategoryHeader";
import { LightMode } from "@/theme/color/lightMode";

export interface AddDeviceProps {
  open: boolean;
  onClose: () => void;
  addOptions: string[];
}

export function AddDevice(props: AddDeviceProps) {
  const { onClose, open, addOptions } = props;

  const handleClose = () => {
    onClose();
  };

  const handleSafeClose = () => {
    onClose();
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">{"Insert device info"}</DialogTitle>
      <DialogContent>
        <CategoryHeader categories={addOptions} columns={addOptions.length}>
          {addOptions.map((option) => (
            <TextField
              sx={{
                "& .MuiOutlinedInput-root": {
                  color: LightMode.palette.primary.main, // input text color
                  "& fieldset": {
                    borderColor: LightMode.palette.primary.main, // default border
                  },
                  "&:hover fieldset": {
                    borderColor: LightMode.palette.primary.main, // hover border
                  },
                  "&.Mui-focused fieldset": {
                    borderColor: LightMode.palette.primary.main, // focused border
                  },
                },
                "& .MuiInputLabel-root": {
                  color: LightMode.palette.primary.main, // default label
                },
                "& .MuiInputLabel-root.Mui-focused": {
                  color: LightMode.palette.primary.main, // focused label
                },
              }}
              id={option}
              type={option}
            />
          ))}
        </CategoryHeader>
      </DialogContent>
      <DialogActions>
        <Button sx={{ color: "primary.dark" }} onClick={handleClose}>
          Close
        </Button>
        <Button
          sx={{ color: "primary.dark" }}
          onClick={handleSafeClose}
          autoFocus
        >
          Add
        </Button>
      </DialogActions>
    </Dialog>
  );
}
