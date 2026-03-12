import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

type InfoProps = {
  name: string;
  address: string;
  created_at: string;
  updated_at: string;
  addUser: () => void;
};

/**
 * Displays a single company row's data: name, address, date created, and date updated.
 * Includes an ActionMenu for adding FACTORY_SUPERUSER to the company.
 * @param props - Component props
 * @param props.name - Display name of the company
 * @param props.address - Main address of company office
 * @param props.created_at - Date the company was added
 * @param props.updated_at - Date of most recent change/addition/deletion related to company
 * @param props.addUser - Button to add a new company admin user
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function CompanyInfo({
  name,
  address,
  created_at,
  updated_at,
  addUser,
}: InfoProps) {
  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{created_at}</Typography>
      <Typography>{updated_at}</Typography>
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
        onClick={addUser}
      >
        {" "}
        {/* TODO: say first only when no company admin has been added. */}
        Add first company admin
      </Button>
    </>
  );
}
