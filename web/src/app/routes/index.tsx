import { Route, Routes } from "react-router";

import Dashboard from "@/pages/Dashboard.tsx";
import Gateways from "@/pages/Gateways";
import Home from "@/pages/Home.tsx";
import Login from "@/pages/Login.tsx";
import Sensors from "@/pages/Sensors.tsx";

export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/Login" element={<Login />} />
      <Route path="/" element={<Home />} />
      <Route path="/Dashboard" element={<Dashboard />} />
      <Route path="/Devices/Sensors" element={<Sensors />} />
      <Route path="/Devices/Gateways" element={<Gateways />} />
      <Route path="/Reports" element={<Home />} />
    </Routes>
  );
}
