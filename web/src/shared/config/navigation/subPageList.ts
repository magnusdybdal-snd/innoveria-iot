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
    ],
  ],
  [
    "Devices",
    [
      { name: "Gateways", path: "/Devices/Gateways" },
      { name: "Sensors", path: "/Devices/Sensors" },
    ],
  ],
  ["Reports", []],
  [
    "Admin",
    [
      { name: "Companies", path: "/Admin/Companies" },
      { name: "Factories", path: "/Admin/Factories" },
      { name: "Users", path: "/Admin/Users" },
    ],
  ],
]);

export default subPages;
