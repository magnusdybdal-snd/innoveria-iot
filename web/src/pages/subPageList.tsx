export type SubPage = {
    name: string;
    path: string;
};

const subPages: Map<string, SubPage[]> = new Map([
    [
        "Dashboard views",
        [
            {name: "KPI dashboard",      path: "/Dashboard/KPI_dashboard"},
            {name: "Building dashboard", path: "/Dashboard/Building_dashboard"},
            {name: "Prod.line 5",        path: "/Dashboard/Prod.line_5"},
        ],
    ],
    [
        "Devices",
        [
            {name: "Gateways",           path: "/Devices/Gateways"},
            {name: "Sensors",            path: "/Devices/Sensors"},
        ],
    ],
    [
        "Reports",
        [
        ],
    ],
]);

export default subPages;