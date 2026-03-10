import Dashboard from "@pages/Dashboard.tsx";
import Gateways from "@pages/Gateways";
import Home from "@pages/Home.tsx";
import Login from "@pages/Login.tsx";
import Sensors from "@pages/Sensors.tsx";
import { Route, Routes } from "react-router";

import Layout from "./Layout";

/**
 * AppRoutes defines the routing structure of the application, mapping URL paths to their corresponding page components.
 * @returns The rendered Routes component containing all defined routes
 */
export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/Login" element={<Login />} />
      <Route element={<Layout />}>
        <Route path="/" element={<Home />} />
        <Route path="/Dashboard" element={<Dashboard />} />
        <Route path="/Devices/Sensors" element={<Sensors />} />
        <Route path="/Devices/Gateways" element={<Gateways />} />
        <Route path="/Reports" element={<Home />} />
      </Route>
    </Routes>
  );
}
