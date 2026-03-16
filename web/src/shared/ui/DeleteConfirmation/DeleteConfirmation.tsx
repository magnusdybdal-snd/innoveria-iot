import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogTitle from "@mui/material/DialogTitle";

interface DeleteConfirmationProps {
  open: boolean;
  onClose: () => void;
  onConfirm: () => void;
}
/**
 * A reusable styled dialog delete confirmation.
 * @param props - Component props.
 * @param props.open - UseState for opening dialog.
 * @param props.onClose - UseState for closing dialog.
 * @param props.onConfirm - Called when the confirm is clicked.
 * @returns The rendered dialog element.
 */
export function DeleteConfirmation({
  open,
  onClose,
  onConfirm,
}: DeleteConfirmationProps) {
  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle>Delete item?</DialogTitle>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button
          sx={{
            backgroundColor: "primary.main",
            color: "primary.dark",
            "&:hover": { backgroundColor: "primary.main" },
          }}
          onClick={onConfirm}
        >
          Confirm
        </Button>
      </DialogActions>
    </Dialog>
  );
}
