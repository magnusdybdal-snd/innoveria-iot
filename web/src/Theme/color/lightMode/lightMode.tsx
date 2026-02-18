import { createTheme } from "@mui/material/styles";

export const LightMode = createTheme({
  palette: {
    primary: {
      main: "#000000",
      dark: "#ececec", // main site background
      light: "#e3e3e3", // menu background
    },
    secondary: {
      main: "#391212",
      light: "#dcdcdc",
    },
  },
});
