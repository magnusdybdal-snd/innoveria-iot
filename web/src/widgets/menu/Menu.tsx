import { useContext } from "react";

import innLogoDark from "@assets/innoveriaDark.png";
import innLogoLight from "@assets/innoveriaLight.png";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import LightModeIcon from "@mui/icons-material/LightMode";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Divider from "@mui/material/Divider";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import MainPages from "@shared/config/navigation/mainPageList";
import SubPages from "@shared/config/navigation/subPageList";
import { ThemeContext } from "@shared/config/theme/themeContext";
import { MenuBox } from "@shared/ui/MenuBox";
import { Link as RouterLink } from "react-router";

import viteLogo from "/vite.svg";

/**
 * Persistent sidebar navigation with logo, user info, page links, theme toggle, and logout.
 * @returns The rendered sidebar menu component
 */
export default function Menu() {
  const { mode, toggle } = useContext(ThemeContext);

  return (
    <Box
      sx={{
        backgroundColor: "primary.light",
        color: "primary.main",
      }}
    >
      <div className="w-64 bg-gray-900 text-white flex flex-col h-screen">
        {/* Top menu logo */}
        <a href="/">
          <img
            src={mode ? innLogoDark : innLogoLight}
            alt="Innoveria logo"
            style={{
              width: "310px",
              height: "auto",
            }}
          />
        </a>
        {/* User info */}
        <div className="flex items-center gap-3 p-4">
          <Avatar
            alt="User"
            src={viteLogo}
            style={{
              width: "60px",
              height: "auto",
            }}
          />
          <div>
            <Typography variant="h3">Username</Typography>
            <Typography variant="h6">email@email.com</Typography>
          </div>
          <Divider
            sx={{
              backgroundColor: "primary.main",
            }}
            variant="middle"
          />
        </div>
        {/* Menu navigation */}
        <nav className="flex-1 flex flex-col">
          {Array.from(MainPages.entries()).map(([category, page]) => (
            <MenuBox
              key={category}
              title={category}
              mainPage={page}
              subPages={SubPages.get(category) ?? []}
              add={category !== "Devices"}
            />
          ))}
        </nav>

        {/* Logout pinned to bottom */}
        <div className="p-4 flex justify-between">
          <Button
            component={RouterLink}
            to="/Login"
            variant="outlined"
            sx={{
              backgroundColor: "secondary.main",
              color: "white",
              "&:hover": {
                backgroundColor: "secondary.dark",
              },
              borderRadius: 2,
              margin: 2,
            }}
          >
            Log out
          </Button>
          <IconButton aria-label="delete" size="large" onClick={toggle}>
            {mode ? (
              <DarkModeIcon sx={{ fontSize: 25, color: "primary.main" }} />
            ) : (
              <LightModeIcon sx={{ fontSize: 25, color: "primary.main" }} />
            )}
          </IconButton>
        </div>
      </div>
    </Box>
  );
}
