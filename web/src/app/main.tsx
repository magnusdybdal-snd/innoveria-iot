import CssBaseline from "@mui/material/CssBaseline";
import { StyledEngineProvider, ThemeProvider } from "@mui/material/styles";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router";

import "@/app/providers/styles/index.css";

import { useEffect, useState } from "react";

import Box from "@mui/material/Box";

import { UserContext } from "@app/providers/UserContext";
import AppRoutes from "@app/routes/index.tsx";
import { getUser, type CurrentUserApiResponse } from "@entities/user";
import { DarkMode } from "@shared/config/theme/darkMode";
import { LightMode } from "@shared/config/theme/lightMode";
import { ThemeContext } from "@shared/config/theme/themeContext";

/**
 * Application root that wires up current user, theme persistence, MUI ThemeProvider, and the React Router route tree.
 * @returns The rendered application root with all providers and routes
 */
export default function Root() {
  // Fetch the current user once on mount and share via UserContext
  const [user, setUser] = useState<CurrentUserApiResponse | null>(null);
  const [userLoading, setUserLoading] = useState(
    () => !!localStorage.getItem("access_token"),
  );
  const [userError, setUserError] = useState<Error | null>(null);

  useEffect(() => {
    if (!localStorage.getItem("access_token")) {
      return;
    }
    getUser()
      .then(setUser)
      .catch((err: unknown) => {
        console.error("Failed to fetch current user:", err);
        setUserError(err instanceof Error ? err : new Error(String(err)));
      })
      .finally(() => setUserLoading(false));
  }, []);

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
    <UserContext.Provider
      value={{ user, isLoading: userLoading, error: userError }}
    >
      <ThemeContext.Provider value={{ mode, toggle }}>
        <BrowserRouter>
          <StyledEngineProvider injectFirst>
            <ThemeProvider theme={mode ? DarkMode : LightMode}>
              <CssBaseline />
              <Box // Hide scrollbar for whole page
                sx={{
                  height: "100vh",
                  overflowY: "auto",
                  "&::-webkit-scrollbar": {
                    display: "none",
                  },
                  scrollbarWidth: "none",
                }}
              >
                <AppRoutes />
              </Box>
            </ThemeProvider>
          </StyledEngineProvider>
        </BrowserRouter>
      </ThemeContext.Provider>
    </UserContext.Provider>
  );
}

createRoot(document.getElementById("root")!).render(<Root />);
