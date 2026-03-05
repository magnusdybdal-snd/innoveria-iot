import CssBaseline from "@mui/material/CssBaseline";
import { StyledEngineProvider, ThemeProvider } from "@mui/material/styles";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router";

import "@/app/providers/styles/index.css";

import { useState } from "react";

import { DarkMode } from "@/shared/config/theme/darkMode";
import { LightMode } from "@/shared/config/theme/lightMode";
import { ThemeContext } from "@/shared/config/theme/themeContext";

import AppRoutes from "./routes";

/**
 * Application root that wires up theme persistence, MUI ThemeProvider, and the React Router route tree.
 * @returns The rendered application root with all providers and routes
 */
export default function Root() {
  // read saved theme on first load
  const [mode, setDark] = useState(() => {
    const saved = localStorage.getItem("theme");
    return saved ? saved === "dark" : true; // default dark
  });

  // toggle + save
  const toggle = () => {
    setDark((prev) => {
      const next = !prev;
      localStorage.setItem("theme", next ? "dark" : "light");
      return next;
    });
  };

  return (
    <ThemeContext.Provider value={{ mode, toggle }}>
      <BrowserRouter>
        <StyledEngineProvider injectFirst>
          <ThemeProvider theme={mode ? DarkMode : LightMode}>
            <CssBaseline />
            <AppRoutes />
          </ThemeProvider>
        </StyledEngineProvider>
      </BrowserRouter>
    </ThemeContext.Provider>
  );
}

createRoot(document.getElementById("root")!).render(<Root />);
