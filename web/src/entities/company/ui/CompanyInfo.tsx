import Typography from "@mui/material/Typography";

type InfoProps = {
  name: string;
  address: string;
  created_at: string;
  updated_at: string;
};

/**
 * Displays a single company row's data: name, address, date created, and date updated.
 * @param props - Component props
 * @param props.name - Display name of the company
 * @param props.address - Main address of company office
 * @param props.created_at - Date the company was added
 * @param props.updated_at - Date of most recent change/addition/deletion related to company
 * @returns A set of grid-aligned cells
 */
export function CompanyInfo({
  name,
  address,
  created_at,
  updated_at,
}: InfoProps) {
  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{created_at}</Typography>
      <Typography>{updated_at}</Typography>
    </>
  );
}
