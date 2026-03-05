import AddBoxOutlinedIcon from "@mui/icons-material/AddBoxOutlined";
import Box from "@mui/material/Box";
import Link from "@mui/material/Link";

{
  /*Button to add new components*/
}
/**
 * Placeholder card that lets users add a new dashboard component.
 * @returns The rendered add-component box
 */
export function AddBox() {
  return (
    <Box
      sx={{
        backgroundColor: "secondary.light",
        color: "primary.main",
        margin: 1,
        borderRadius: 3,
        border: "3px dashed grey",
        width: 250,
        height: 250,
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        textAlign: "center",
      }}
    >
      <Link href="#" underline="none">
        <div className="w-64 p-[10px]">
          <h2 className="font-normal text-3xl" style={{ margin: 0 }}>
            Add dashboard component
          </h2>
          <AddBoxOutlinedIcon fontSize="large" />
        </div>
      </Link>
    </Box>
  );
}
