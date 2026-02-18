import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import type { ReactNode } from "react";

type CategoryHeaderProps = {
  categories: string[];
  columns?: number;
  children?: ReactNode;
};

export function CategoryHeader({
  categories,
  columns,
  children,
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
      {categories.map((category) => (
        <Typography key={category} className="px-4 py-2">
          {category}
        </Typography>
      ))}
      {children}
    </Box>
  );
}
