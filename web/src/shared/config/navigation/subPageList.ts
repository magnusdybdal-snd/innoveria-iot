export type SubPage = {
  name: string;
  path: string;
};

const subPages: Map<string, SubPage[]> = new Map([
  [
    "Dashboard views",
    [
      { name: "Context dashboard", path: "/Dashboard/Context" },
      { name: "Context rules", path: "/Dashboard/Rules" },
      { name: "Orders", path: "/Dashboard/Context/OrderList" },
    ],
  ],
  [
    "Devices",
    [
      { name: "Gateways", path: "/Devices/Gateways" },
      { name: "Sensors", path: "/Devices/Sensors" },
    ],
  ],
  [
    "Admin",
    [
      { name: "Companies", path: "/Admin/Companies" },
      { name: "Factories", path: "/Admin/Factories" },
      { name: "Measurement types", path: "/Admin/Measurement" },
      { name: "Payload schema", path: "/Admin/PayloadSchema" },
      { name: "Users", path: "/Admin/Users" },
    ],
  ],
]);

export default subPages;
