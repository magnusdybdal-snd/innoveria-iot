import { OutlinedButton } from "@/shared/ui/Button";
import Typography from "@mui/material/Typography";

type InfoProps = {
  name: string;
  address: string;
  created_at: string;
  updated_at: string;
  onViewUsers?: () => void;
};

/**
 * Displays a single company row's data: name, address, date created, date updated, and a link to its users.
 * @param props - Component props
 * @param props.name - Display name of the company
 * @param props.address - Main address of company office
 * @param props.created_at - Date the company was added
 * @param props.updated_at - Date of most recent change/addition/deletion related to company
 * @param props.onViewUsers - Optional callback to drill down into the company's users; omit to hide the button
 * @returns A set of grid-aligned cells with an optional users button
 */
export function CompanyInfo({
  name,
  address,
  created_at,
  updated_at,
  onViewUsers,
}: InfoProps) {
  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{created_at}</Typography>
      <Typography>{updated_at}</Typography>
      {onViewUsers && (
        <OutlinedButton
          onClick={(e) => {
            e.stopPropagation();
            onViewUsers();
          }}
        >
          Users
        </OutlinedButton>
      )}
    </>
  );
}
