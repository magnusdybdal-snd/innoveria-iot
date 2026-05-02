import MemoryIcon from "@mui/icons-material/Memory";
import SpaceDashboardIcon from "@mui/icons-material/SpaceDashboard";
import Box from "@mui/material/Box";

import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { NavCard, type NavCardProps } from "@widgets/homeNavCard";
import { WelcomeHeader } from "@widgets/homeWelcome";

const NAV_SECTIONS: NavCardProps[] = [
  {
    label: "Dashboard",
    description: "Monitor live sensor context and production order status.",
    icon: <SpaceDashboardIcon fontSize="small" />,
    pages: [{ name: "Sensor Data", path: "/Dashboard/Context/SensorData" }],
  },
  {
    label: "Devices",
    description:
      "Manage connected LoRaWAN gateways and sensors on the factory floor.",
    icon: <MemoryIcon fontSize="small" />,
    pages: [
      { name: "Gateways", path: "/Devices/Gateways" },
      { name: "Sensors", path: "/Devices/Sensors" },
    ],
  },
];

/**
 * Home page displaying a live clock, welcome message, and navigation cards
 * to all main sections of the application (Admin excluded).
 * @returns The rendered Home page
 */
export default function Home() {
  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader />
        <WelcomeHeader />
        <Box
          sx={{
            display: "grid",
            gridTemplateColumns: `repeat(${NAV_SECTIONS.length}, 1fr)`,
            gap: 2,
            maxWidth: 800,
          }}
        >
          {NAV_SECTIONS.map((section) => (
            <NavCard key={section.label} {...section} />
          ))}
        </Box>
      </PageContent>
    </div>
  );
}
