export type SubPage = {
  name: string;
  path: string;
};

const subPages: Map<string, SubPage[]> = new Map([
  [
    "Dashboard views",
    [
      { name: "Context dashboard", path: "/Dashboard/Context" },
      { name: "Order context", path: "/Dashboard/Context/Orders" },
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
