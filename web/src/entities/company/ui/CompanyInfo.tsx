import Typography from "@mui/material/Typography";

import { ActionMenu } from "@shared/ui/actionMenu";

type InfoProps = {
  name: string;
  address: string;
  created_at: string;
  updated_at: string;
  addUser: () => void;
  onGenerateERPToken: () => void;
};

/**
 * Displays a single company row's data: name, address, date created, and date updated.
 * Includes an ActionMenu for company actions.
 * @param props - Component props
 * @param props.name - Display name of the company
 * @param props.address - Main address of company office
 * @param props.created_at - Date the company was added
 * @param props.updated_at - Date of most recent change/addition/deletion related to company
 * @param props.addUser - Button to add a new company admin user
 * @param props.onGenerateERPToken - Callback to issue ERP token for the company
 * @returns A set of grid-aligned cells with an action menu and rename dialog
 */
export function CompanyInfo({
  name,
  address,
  created_at,
  updated_at,
  addUser,
  onGenerateERPToken,
}: InfoProps) {
  const menuItems = [
    {
      label: "Generate ERP token",
      onClick: onGenerateERPToken,
    },
    {
      label: "Add first company admin",
      onClick: addUser,
      disabled: true,
    },
  ];

  return (
    <>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{created_at}</Typography>
      <Typography>{updated_at}</Typography>
      <ActionMenu items={menuItems} />
    </>
  );
}
