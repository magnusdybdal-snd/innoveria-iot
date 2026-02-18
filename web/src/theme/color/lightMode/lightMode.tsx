import { createTheme } from "@mui/material/styles";

export const LightMode = createTheme({
  palette: {
    primary: {
      main: "#000000",
      dark: "#fefefe", // main site background
      light: "#f1f1f1", // menu background
    },
    secondary: {
      main: "#391212",
      light: "#dcdcdc",
    },
  },
});
