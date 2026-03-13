import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";

interface RenameDialogProps {
  open: boolean;
  onClose: () => void;
  value: string;
  onChange: (value: string) => void;
  onSave: () => void;
  label?: string;
  title?: string;
}
/**
 * A reusable styled dialog rename.
 * @param props - Component props.
 * @param props.open - UseState for opening dialog.
 * @param props.onClose - UseState for closing dialog.
 * @param props.value - The name to be changed.
 * @param props.onChange - Function for changing text field.
 * @param props.onSave - Function for saving change.
 * @param props.label - Text field label.
 * @param props.title - Title dialog.
 * @returns The rendered dialog element.
 */
export function RenameDialog({
  open,
  onClose,
  value,
  onChange,
  onSave,
  label = "Name",
  title = "Rename",
}: RenameDialogProps) {
  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle>{title}</DialogTitle>

      <DialogContent>
        <TextField
          autoFocus
          fullWidth
          label={label}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          sx={{ mt: 1 }}
        />
      </DialogContent>

      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button variant="contained" onClick={onSave}>
          Save
        </Button>
      </DialogActions>
    </Dialog>
  );
}
