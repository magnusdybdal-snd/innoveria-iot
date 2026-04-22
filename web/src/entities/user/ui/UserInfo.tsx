import DeleteIcon from "@mui/icons-material/Delete";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";

type InfoProps = {
  id: string;
  name: string;
  email: string;
  role: string;
  createdAt: string;
  onDelete: (id: string) => void;
};

/**
 * Displays a single user row: name, email, role, created date, and a delete action.
 * @param props - Component props
 * @param props.id - User ID used for delete
 * @param props.name - Display name of the user
 * @param props.email - Email address of the user
 * @param props.role - Role assigned to the user
 * @param props.createdAt - Formatted creation timestamp
 * @param props.onDelete - Called with the user ID when the delete button is clicked
 * @returns A set of grid-aligned cells with a delete action
 */
export function UserInfo({
  id,
  name,
  email,
  role,
  createdAt,
  onDelete,
}: InfoProps) {
  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{email}</Typography>
      <Typography>{role}</Typography>
      <Typography>{createdAt}</Typography>
      <IconButton
        size="small"
        onClick={() => onDelete(id)}
        aria-label="delete user"
        sx={{ color: "primary.main" }}
      >
        <DeleteIcon fontSize="small" />
      </IconButton>
    </>
  );
}
