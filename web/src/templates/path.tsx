import Divider from "@mui/material/Divider";
import Typography from "@mui/material/Typography";
import { useLocation } from "react-router";

{
  /*Path of current site + line divider*/
}
export default function Home() {
  const location = useLocation();
  const pathNames = location.pathname.split("/").filter((x) => x); // removes empty strings

  return (
    <div>
      <Typography variant="h6">
        Home
        {pathNames.map((name, index) => {
          const formatted =
            name.charAt(0).toUpperCase() + name.slice(1).toLowerCase();

          return (
            <span key={index}>
              {" > "}
              {formatted}
            </span>
          );
        })}
      </Typography>
      <Divider
        sx={{
          backgroundColor: "primary.main",
          marginTop: 2,
          marginBottom: 5,
        }}
      />
    </div>
  );
}
