import type { ReactNode } from "react";

import HelpOutlineIcon from "@mui/icons-material/HelpOutline";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import KeyboardArrowUpIcon from "@mui/icons-material/KeyboardArrowUp";
import SwapVertIcon from "@mui/icons-material/SwapVert";
import Box from "@mui/material/Box";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

type SortDirection = "asc" | "desc";

type CategoryHeaderProps = {
  categories: string[];
  columns?: number;
  children?: ReactNode;
  sortableColumns?: string[];
  sortConfig?: { key: string | null; direction: SortDirection };
  onSort?: (column: string) => void;
  fit?: boolean;
};

/**
 * Renders a grid header row with column labels and optional sort controls.
 * Sortable columns show a pointer cursor and display an arrow (↑/↓) when active.
 * Child elements (e.g. data rows) are rendered inside the same grid below the header.
 * @param props - Component props
 * @param props.categories - Ordered list of column label strings
 * @param props.columns - Total column count for the grid template; defaults to categories.length if omitted
 * @param props.children - Data rows to render inside the grid
 * @param props.sortableColumns - Subset of category labels that are clickable for sorting
 * @param props.sortConfig - Currently active sort key and direction
 * @param props.onSort - Callback invoked with the column label when a sortable header is clicked
 * @param props.fit - Whether the header stretches the widt of the content
 * @returns A full-width grid box with header labels and child content
 */
export function CategoryHeader({
  categories,
  columns,
  children,
  sortableColumns = [],
  sortConfig,
  onSort,
  fit,
}: CategoryHeaderProps) {
  // Creates number of columns based on string[] passed as parameter
  // First category is made to fit the object through "auto"
  const templateColumns = fit
    ? `repeat(${columns ?? categories.length}, max-content)`
    : columns
      ? `auto ${Array(columns - 1)
          .fill("1fr")
          .join(" ")}`
      : `repeat(${categories.length}, 1fr)`;

  return (
    <Box
      sx={{
        display: "grid",
        gridTemplateColumns: templateColumns,
        columnGap: 4,
        borderRadius: 2,
        backgroundColor: "primary.light",
        color: "primary.main",
        padding: 2,
        width: fit ? "fit-content" : "100%",
      }}
    >
      {/* Map every category to display as text */}
      {categories.map((category) => {
        const isSortable = sortableColumns.includes(category);
        const isActive = sortConfig?.key === category;
        const arrow = isActive ? (
          sortConfig.direction === "asc" ? (
            <KeyboardArrowDownIcon />
          ) : (
            <KeyboardArrowUpIcon />
          )
        ) : sortConfig?.key == null ? (
          <SwapVertIcon />
        ) : (
          <SwapVertIcon style={{ visibility: "hidden" }} />
        );

        return (
          <Typography
            key={category}
            fontSize={18}
            fontWeight={"bold"}
            className="px-4 py-2"
            onClick={isSortable && onSort ? () => onSort(category) : undefined}
            sx={{
              whiteSpace: "nowrap",
              ...(isSortable && { cursor: "pointer", userSelect: "none" }),
            }}
          >
            <span style={{ display: "flex", alignItems: "center", gap: 4 }}>
              {category}
              {isSortable && arrow}

              {/* Adds info on status color of column is status*/}
              {category === "Status" && (
                <Tooltip
                  title={
                    <>
                      Green = Online
                      <br />
                      Yellow = Never connect
                      <br />
                      Red = Offline
                    </>
                  }
                >
                  <HelpOutlineIcon sx={{ fontSize: 15 }} />
                </Tooltip>
              )}
            </span>
          </Typography>
        );
      })}
      {children}
    </Box>
  );
}
