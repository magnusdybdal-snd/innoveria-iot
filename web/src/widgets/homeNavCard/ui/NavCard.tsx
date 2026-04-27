import { type ReactNode } from "react";

import ArrowForwardIcon from "@mui/icons-material/ArrowForward";
import Box from "@mui/material/Box";
import Divider from "@mui/material/Divider";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";
import { Link as RouterLink } from "react-router";

/** A navigable sub-page entry within a nav card. */
export type NavPage = {
  name: string;
  path: string;
};

/** Props for the NavCard widget. */
export type NavCardProps = {
  /** Section label shown as the card header. */
  label: string;
  /** Short description of what this section contains. */
  description: string;
  /** Icon displayed next to the section label. */
  icon: ReactNode;
  /** List of pages to render as clickable links inside the card. */
  pages: NavPage[];
};

/**
 * A bordered navigation card linking to a group of related pages.
 * Renders an icon, label, description, and a list of page links with hover animations.
 * @param props - NavCardProps
 * @param props.label
 * @param props.description
 * @param props.icon
 * @param props.pages
 * @returns The rendered NavCard widget
 */
export function NavCard({ label, description, icon, pages }: NavCardProps) {
  const theme = useTheme();

  const isDark = theme.palette.mode === "dark";
  const borderColor = theme.palette.secondary.light;
  const labelColor = theme.palette.primary.main;
  const hoverBg = isDark ? "rgba(255,255,255,0.04)" : "rgba(0,0,0,0.03)";

  return (
    <Box
      sx={{
        border: `2px solid ${borderColor}`,
        borderRadius: 4,
        p: 3,
        transition: "border-color 0.15s ease",
        "&:hover": { borderColor: "text.primary" },
      }}
    >
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          gap: 1,
          mb: 0.75,
          color: labelColor,
        }}
      >
        {icon}
        <Typography
          variant="overline"
          sx={{
            letterSpacing: 2,
            fontWeight: 700,
            color: labelColor,
            lineHeight: 1,
            fontSize: "0.65rem",
          }}
        >
          {label}
        </Typography>
      </Box>

      <Typography
        variant="body2"
        sx={{
          color: "text.secondary",
          mb: 2.5,
          fontSize: "0.8rem",
          lineHeight: 1.5,
        }}
      >
        {description}
      </Typography>

      <Divider sx={{ mb: 1.5 }} />

      <Box sx={{ display: "flex", flexDirection: "column" }}>
        {pages.map((page) => (
          <Box
            key={page.path}
            component={RouterLink}
            to={page.path}
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              py: 1,
              px: 1.5,
              textDecoration: "none",
              color: "inherit",
              transition: "background-color 0.15s ease",
              "&:hover": { backgroundColor: hoverBg },
              "&:hover .arrow-icon": { transform: "translateX(3px)" },
            }}
          >
            <Typography variant="body2" sx={{ fontWeight: 500 }}>
              {page.name}
            </Typography>
            <ArrowForwardIcon
              className="arrow-icon"
              sx={{
                fontSize: 15,
                color: "text.secondary",
                transition: "transform 0.15s ease",
              }}
            />
          </Box>
        ))}
      </Box>
    </Box>
  );
}
