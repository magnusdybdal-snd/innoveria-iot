import {
  useContext,
  useState,
  type ComponentType,
  type ReactNode,
} from "react";

import AdminPanelSettingsIcon from "@mui/icons-material/AdminPanelSettings";
import AssignmentIcon from "@mui/icons-material/Assignment";
import BusinessIcon from "@mui/icons-material/Business";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import FactoryIcon from "@mui/icons-material/Factory";
import HomeIcon from "@mui/icons-material/Home";
import LightModeIcon from "@mui/icons-material/LightMode";
import MemoryIcon from "@mui/icons-material/Memory";
import NoteAltIcon from "@mui/icons-material/NoteAlt";
import PeopleIcon from "@mui/icons-material/People";
import RouterIcon from "@mui/icons-material/Router";
import RuleIcon from "@mui/icons-material/Rule";
import SensorsIcon from "@mui/icons-material/Sensors";
import SettingsRemoteIcon from "@mui/icons-material/SettingsRemote";
import SpaceDashboardIcon from "@mui/icons-material/SpaceDashboard";
import StraightenIcon from "@mui/icons-material/Straighten";
import SummarizeIcon from "@mui/icons-material/Summarize";
import Accordion from "@mui/material/Accordion";
import AccordionDetails from "@mui/material/AccordionDetails";
import AccordionSummary from "@mui/material/AccordionSummary";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Divider from "@mui/material/Divider";
import Drawer from "@mui/material/Drawer";
import IconButton from "@mui/material/IconButton";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import type { SvgIconProps } from "@mui/material/SvgIcon";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { Link as RouterLink, useLocation } from "react-router";

import { useCurrentUser } from "@app/providers/useCurrentUser";
import innLogoDark from "@assets/innoveriaDark.png";
import innLogoLight from "@assets/innoveriaLight.png";
import { postLogout } from "@entities/user";
import MainPages from "@shared/config/navigation/mainPageList";
import SubPages from "@shared/config/navigation/subPageList";
import { ThemeContext } from "@shared/config/theme/themeContext";

interface MenuProps {
  children: ReactNode;
}

const pageSymbol: Map<string, ComponentType<SvgIconProps>> = new Map([
  ["Home", HomeIcon],
  ["Admin", AdminPanelSettingsIcon],
  ["Dashboard views", SpaceDashboardIcon],
  ["Devices", MemoryIcon],
  ["Gateways", RouterIcon],
  ["Sensors", SettingsRemoteIcon],
  ["Reports", SummarizeIcon],
  ["Companies", BusinessIcon],
  ["Factories", FactoryIcon],
  ["Measurement types", StraightenIcon],
  ["Payload schema", NoteAltIcon],
  ["Users", PeopleIcon],
  ["Context rules", RuleIcon],
  ["Sensor Data", SensorsIcon],
  ["Context", AssignmentIcon],
]);

const roles: Map<string, string> = new Map([
  ["FACTORY_WORKER", "Factory worker"],
  ["PLATFORM_ADMIN", "Platform admin"],
]);

/**
 * Persistent sidebar navigation with logo, user info, page links, theme toggle, and logout.
 * @param menuProps props of menu
 * @param menuProps.children the page
 * @returns The rendered sidebar menu component
 */
export default function Menu(menuProps: MenuProps) {
  const { children } = menuProps;
  const { mode, toggle } = useContext(ThemeContext);
  const location = useLocation();
  const [expanded, setExpanded] = useState<string[]>([]);
  const { user } = useCurrentUser();
  const isAdmin = user?.role === "PLATFORM_ADMIN";

  const handleAccordionChange =
    (category: string) =>
    (_event: React.SyntheticEvent, isExpanded: boolean) => {
      setExpanded(
        (prev) =>
          isExpanded
            ? [...prev, category] // open category
            : prev.filter((c) => c !== category), // close category
      );
    };

  return (
    <Box sx={{ display: "flex" }}>
      <Drawer
        sx={{
          width: 270,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: 270,
            boxSizing: "border-box",
            display: "flex",
            flexDirection: "column",

            // Hide scrollbar
            "&::-webkit-scrollbar": {
              display: "none", // Chrome, Safari, Edge
            },
            scrollbarWidth: "none", // Firefox
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <Toolbar>
          {/* Top menu logo */}
          <a href="/">
            <img
              src={mode ? innLogoDark : innLogoLight}
              alt="Innoveria logo"
              style={{
                width: "220px",
                height: "auto",
              }}
            />
          </a>
        </Toolbar>
        <Divider />
        <List sx={{ flexGrow: 1 }}>
          {/* User info */}
          <ListItem>
            <div className="flex items-center gap-3 p-4">
              <div>
                {user ? (
                  <>
                    <Typography variant="h6">{user.email}</Typography>
                    <Typography variant="h6">{roles.get(user.role)}</Typography>
                  </>
                ) : (
                  <Typography variant="h6">Loading...</Typography>
                )}
              </div>
              <Divider
                sx={{
                  backgroundColor: "primary.main",
                }}
                variant="middle"
              />
            </div>
          </ListItem>
          {/* Menu navigation */}
          {Array.from(MainPages.entries()).map(([category, page]) => {
            if (category === "Admin" && !isAdmin) return null;
            const subPagesForCategory = SubPages.get(category) ?? [];
            const isCategoryActive = subPagesForCategory.some(
              (sub) => location.pathname === sub.path,
            );
            const MainIcon = pageSymbol.get(category);

            // If no subpages → normal link
            if (subPagesForCategory.length === 0) {
              return (
                <ListItem key={category} disablePadding>
                  <ListItemButton
                    component={RouterLink}
                    to={page}
                    selected={location.pathname === page}
                  >
                    {/*Add symbol if page has it*/}
                    {MainIcon && (
                      <ListItemIcon sx={{ minWidth: 28 }}>
                        <MainIcon fontSize="small" />
                      </ListItemIcon>
                    )}
                    <ListItemText primary={category} />
                  </ListItemButton>
                </ListItem>
              );
            }

            return (
              <ListItem key={category} disablePadding sx={{ display: "block" }}>
                <Accordion
                  expanded={expanded.includes(category)}
                  onChange={handleAccordionChange(category)}
                  disableGutters
                  elevation={0}
                  sx={{
                    backgroundColor: "transparent",
                    "&:before": { display: "none" },
                  }}
                >
                  <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    sx={{
                      px: 2,
                      backgroundColor: isCategoryActive
                        ? "action.selected"
                        : "transparent",
                      "&:hover": {
                        backgroundColor: isCategoryActive
                          ? "action.selected"
                          : "action.hover",
                      },
                    }}
                  >
                    {MainIcon && (
                      <ListItemIcon sx={{ minWidth: 28 }}>
                        <MainIcon fontSize="small" />
                      </ListItemIcon>
                    )}
                    <Typography>{category}</Typography>
                  </AccordionSummary>

                  <AccordionDetails sx={{ p: 0 }}>
                    <List component="div" disablePadding>
                      {subPagesForCategory.map((subPage) => {
                        const SubIcon = pageSymbol.get(subPage.name);

                        return (
                          <ListItem
                            key={subPage.path}
                            disablePadding
                            component="div"
                          >
                            <ListItemButton
                              component={RouterLink}
                              to={subPage.path}
                              sx={{ pl: 4 }}
                              selected={location.pathname === subPage.path}
                            >
                              {/*Add symbol if page has it*/}
                              {SubIcon && (
                                <ListItemIcon sx={{ minWidth: 28 }}>
                                  <SubIcon fontSize="small" />
                                </ListItemIcon>
                              )}
                              <ListItemText primary={subPage.name} />
                            </ListItemButton>
                          </ListItem>
                        );
                      })}
                    </List>
                  </AccordionDetails>
                </Accordion>
              </ListItem>
            );
          })}

          <ListItem></ListItem>
        </List>
        <Box sx={{ p: 2, display: "flex", justifyContent: "space-between" }}>
          <Button
            onClick={() => {
              postLogout().finally(() => {
                localStorage.removeItem("access_token");
                window.location.replace("/Login");
              });
            }}
            variant="outlined"
            sx={{
              backgroundColor: "secondary.main",
              color: "white",
              "&:hover": {
                backgroundColor: "secondary.dark",
                color: "white",
              },
              borderRadius: 2,
              margin: 0,
            }}
          >
            Log out
          </Button>
          <IconButton aria-label="delete" size="large" onClick={toggle}>
            {mode ? (
              <DarkModeIcon sx={{ fontSize: 25, color: "primary.main" }} />
            ) : (
              <LightModeIcon sx={{ fontSize: 25, color: "primary.main" }} />
            )}
          </IconButton>
        </Box>
      </Drawer>

      <Box component="main" sx={{ flexGrow: 1 }}>
        {children}
      </Box>
    </Box>
  );
}
