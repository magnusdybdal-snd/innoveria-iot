import Context from "@/pages/Context";
import Factories from "@/pages/Factories";
import OrderContext from "@/pages/OrderContext";
import OrderDetail from "@/pages/OrderDetail";
import OrderList from "@/pages/OrderList";
import Rules from "@/pages/Rules";
import StatusPage from "@/pages/StatusPage";
import { Route, Routes } from "react-router";

import Companies from "@pages/Companies.tsx";
import Gateways from "@pages/Gateways";
import Home from "@pages/Home.tsx";
import Login from "@pages/Login.tsx";
import MeasurementTypes from "@pages/MeasurementTypes.tsx";
import PayloadSchema from "@pages/PayloadSchema.tsx";
import Sensors from "@pages/Sensors.tsx";

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
        <Route path="/Dashboard/Context" element={<Context />} />
        <Route path="/Dashboard/Context/Orders" element={<OrderContext />} />
        <Route path="/Dashboard/Context/OrderList" element={<OrderList />} />
        <Route
          path="/Dashboard/Context/OrderList/:id"
          element={<OrderDetail />}
        />
        <Route path="/Admin/Rules" element={<Rules />} />
        <Route path="/Devices/Sensors" element={<Sensors />} />
        <Route path="/Devices/Gateways" element={<Gateways />} />
        <Route path="/Admin/Companies" element={<Companies />} />
        <Route path="/Admin/Factories" element={<Factories />} />
        <Route path="/Admin/Measurement" element={<MeasurementTypes />} />
        <Route path="/Admin/PayloadSchema" element={<PayloadSchema />} />
        <Route path="/Admin/Users" element={<Home />} />
        <Route
          path="*"
          element={<StatusPage code="404" message="Page not found" />}
        />
      </Route>
    </Routes>
  );
}
