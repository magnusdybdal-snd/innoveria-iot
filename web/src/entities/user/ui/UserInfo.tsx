import Typography from "@mui/material/Typography";

type InfoProps = {
  company_id: string;
  email: string;
  name: string;
  role: string;
};

/**
 * Displays a single user row's data: company id, email, name, role, and user id.
 * @param props - Component props
 * @param props.company_id - Company user is part of
 * @param props.email - Email of user
 * @param props.name - Name of user
 * @param props.role - Functional role in the company
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function UserInfo({ company_id, email, name, role }: InfoProps) {
  return (
    <>
      <Typography>{company_id}</Typography>
      <Typography>{email}</Typography>
      <Typography>{name}</Typography>
      <Typography>{role}</Typography>
    </>
  );
}
