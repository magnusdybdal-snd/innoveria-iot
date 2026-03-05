import {
  useContext,
  useState,
  type ComponentType,
  type ReactNode,
} from "react";

import DarkModeIcon from "@mui/icons-material/DarkMode";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import LightModeIcon from "@mui/icons-material/LightMode";
import MemoryIcon from "@mui/icons-material/Memory";
import RouterIcon from "@mui/icons-material/Router";
import SettingsRemoteIcon from "@mui/icons-material/SettingsRemote";
import SpaceDashboardIcon from "@mui/icons-material/SpaceDashboard";
import SummarizeIcon from "@mui/icons-material/Summarize";
import Accordion from "@mui/material/Accordion";
import AccordionDetails from "@mui/material/AccordionDetails";
import AccordionSummary from "@mui/material/AccordionSummary";
import Avatar from "@mui/material/Avatar";
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

import innLogoDark from "@/assets/innoveriaDark.png";
import innLogoLight from "@/assets/innoveriaLight.png";
import MainPages from "@/pages/mainPageList.tsx";
import SubPages from "@/pages/subPageList.tsx";
import { ThemeContext } from "@/theme/color/themeContext";

import viteLogo from "/vite.svg";

interface MenuProps {
  children: ReactNode;
}

const pageSymbol: Map<string, ComponentType<SvgIconProps>> = new Map([
  ["Dashboard views", SpaceDashboardIcon],
  ["Devices", MemoryIcon],
  ["Gateways", RouterIcon],
  ["Sensors", SettingsRemoteIcon],
  ["Reports", SummarizeIcon],
]);

export default function Menu({ children }: MenuProps) {
  const { mode, toggle } = useContext(ThemeContext);
  const location = useLocation();
  const [expanded, setExpanded] = useState<string | false>(false);

  const handleAccordionChange =
    (category: string) =>
    (_event: React.SyntheticEvent, isExpanded: boolean) => {
      setExpanded(isExpanded ? category : false);
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
              <Avatar
                alt="User"
                src={viteLogo}
                style={{
                  width: "60px",
                  height: "auto",
                }}
              />
              <div>
                <Typography variant="h4">Username</Typography>
                <Typography variant="h6">email@email.com</Typography>
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
            const subPagesForCategory = SubPages.get(category) ?? [];
            const MainIcon = pageSymbol.get(category);
            const isCategoryActive = subPagesForCategory.some(
              // Keep open category dropdown
              (sub) => location.pathname === sub.path,
            );

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
                  expanded={expanded === category || isCategoryActive}
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
                    sx={{ px: 2 }}
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
                          <ListItem key={subPage.path} disablePadding>
                            <ListItemButton
                              component={RouterLink}
                              to={subPage.path}
                              sx={{ pl: 4 }}
                              selected={location.pathname === page}
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
            component={RouterLink}
            to="/Login"
            variant="outlined"
            sx={{
              backgroundColor: "secondary.main",
              color: "white",
              "&:hover": {
                backgroundColor: "secondary.dark",
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
