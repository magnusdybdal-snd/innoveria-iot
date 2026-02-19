import { BrowserRouter, Routes, Route } from "react-router";
import { createRoot } from "react-dom/client";
import CssBaseline from "@mui/material/CssBaseline";
import { StyledEngineProvider, ThemeProvider } from "@mui/material/styles";
import "@/index.css";
import Home from "@/pages/Home.tsx";
import Login from "@/pages/Login.tsx";
import Dashboard from "@/pages/Dashboard.tsx";
import Sensors from "@/pages/Sensors.tsx";
import { DarkMode } from "@/theme/color/darkMode";
import { LightMode } from "@/theme/color/lightMode";
import { ThemeContext } from "@/theme/color/themeContext";
import { useState } from "react";
import Gateways from "@/pages/Gateways";

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

            <Routes>
              <Route path="/Login" element={<Login />} />
              <Route path="/" element={<Home />} />
              <Route path="/Dashboard" element={<Dashboard />} />
              <Route path="/Devices/Sensors" element={<Sensors />} />
              <Route path="/Devices/Gateways" element={<Gateways />} />
              <Route path="/Reports" element={<Home />} />
            </Routes>
          </ThemeProvider>
        </StyledEngineProvider>
      </BrowserRouter>
    </ThemeContext.Provider>
  );
}

createRoot(document.getElementById("root")!).render(<Root />);
