import { createTheme } from "@mui/material/styles";

export const DarkMode = createTheme({
  palette: {
    primary: {
      main: "#ffffff",
      dark: "#131313", // main site background
      light: "#1c1c1c", // menu background
    },
    secondary: {
      main: "#391212",
      light: "#232323",
    },
  },
});
