import Divider from "@mui/material/Divider";
import Link from "@mui/material/Link";
import Toolbar from "@mui/material/Toolbar";
import { Link as RouterLink, useLocation } from "react-router";

{
  /*Path of current site + line divider*/
}
export function Path() {
  const location = useLocation();
  const pathNames = location.pathname.split("/").filter((x) => x); // removes empty strings

  return (
    <div>
      <Toolbar>
        <Link
          component={RouterLink}
          to="/"
          color="inherit"
          style={{
            textDecoration: "none",
          }}
        >
          Home
        </Link>

        {pathNames.map((name, index) => {
          const formatted =
            name.charAt(0).toUpperCase() + name.slice(1).toLowerCase();

          const to = "/" + pathNames.slice(0, index + 1).join("/");

          return (
            <span key={index}>
              {" > "}
              <Link
                component={RouterLink}
                to={to}
                color="inherit"
                style={{ textDecoration: "none" }}
              >
                {formatted}
              </Link>
            </span>
          );
        })}
      </Toolbar>
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
