export type SubPage = {
  name: string;
  path: string;
};

const subPages: Map<string, SubPage[]> = new Map([
  [
    "Dashboard views",
    [
      { name: "Sensor Data", path: "/Dashboard/Context/SensorData" },
      { name: "Context rules", path: "/Dashboard/Rules" },
      { name: "Context", path: "/Dashboard/Context/OrderList" },
    ],
  ],
  [
    "Devices",
    [
      { name: "Gateways", path: "/Devices/Gateways" },
      { name: "Sensors", path: "/Devices/Sensors" },
    ],
  ],
  ["Organization", [{ name: "Factories", path: "/Organization/Factories" }]],
  [
    "Admin",
    [
      { name: "Companies", path: "/Admin/Companies" },
      { name: "Measurement types", path: "/Admin/Measurement" },
      { name: "Payload schema", path: "/Admin/PayloadSchema" },
      { name: "Users", path: "/Admin/Users" },
    ],
  ],
]);

export default subPages;
