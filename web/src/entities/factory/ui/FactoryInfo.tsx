import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

type InfoProps = {
  id: string;
  name: string;
  address: string;
  createdAt: string;
  updatedAt: string;
  addFactory: () => void;
};

/**
 * Displays a single factory row's data: id, name, address, date created, and date updated.
 * Includes a button for adding the first factory admin user.
 * @param props - Component props
 * @param props.id - Unique identifier of the factory
 * @param props.name - Display name of the factory
 * @param props.address - Main address of factory office
 * @param props.createdAt - Formatted date the factory was added
 * @param props.updatedAt - Formatted date of most recent change
 * @param props.addFactory - Callback to add a new factory
 * @returns A set of grid-aligned cells with an add user button
 */
export function FactoryInfo({
  id,
  name,
  address,
  createdAt,
  updatedAt,
  addFactory,
}: InfoProps) {
  return (
    <>
      <Typography>{id}</Typography>
      <Typography>{name}</Typography>
      <Typography>{address}</Typography>
      <Typography>{createdAt}</Typography>
      <Typography>{updatedAt}</Typography>
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
        onClick={addFactory}
      >
        Add first factory admin
      </Button>
    </>
  );
}
