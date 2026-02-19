import type { ReactNode } from "react";

import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import KeyboardArrowUpIcon from "@mui/icons-material/KeyboardArrowUp";
import SwapVertIcon from "@mui/icons-material/SwapVert";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

type SortDirection = "asc" | "desc";

type CategoryHeaderProps = {
  categories: string[];
  columns?: number;
  children?: ReactNode;
  sortableColumns?: string[];
  sortConfig?: { key: string | null; direction: SortDirection };
  onSort?: (column: string) => void;
};

export function CategoryHeader({
  categories,
  columns,
  children,
  sortableColumns = [],
  sortConfig,
  onSort,
}: CategoryHeaderProps) {
  // Creates number of columns based on string[] passed as parameter
  // First category is made to fit the object through "auto"
  const templateColumns = columns
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
      }}
      className="w-full"
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
            sx={
              isSortable ? { cursor: "pointer", userSelect: "none" } : undefined
            }
          >
            <span style={{ display: "flex", alignItems: "center", gap: 4 }}>
              {category}
              {isSortable && arrow}
            </span>
          </Typography>
        );
      })}
      {children}
    </Box>
  );
}
