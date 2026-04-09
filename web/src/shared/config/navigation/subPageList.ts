export type SubPage = {
  name: string;
  path: string;
};

const subPages: Map<string, SubPage[]> = new Map([
  [
    "Dashboard views",
    [
      //{ name: "KPI dashboard", path: "/Dashboard/KPI_dashboard" },
      //{ name: "Building dashboard", path: "/Dashboard/Building_dashboard" },
      //{ name: "Prod.line 5", path: "/Dashboard/Prod.line_5" },
      { name: "Context dashboard", path: "/Dashboard/Context" },
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
      // Rules are supposed to be viewed by normal users. NOT JUST innoveria
      { name: "Context rules", path: "/Admin/Rules" },
    ],
  ],
]);

export default subPages;
