import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

interface GraphWidgetPlaceholderProps {
  onAdd: () => void;
}

/**
 * Empty state placeholder shown when no graph widgets have been added yet.
 * Clicking it adds the first widget.
 * @param props - Component props
 * @param props.onAdd - Called when the user clicks the placeholder to add a widget
 * @returns The rendered placeholder card
 */
export function GraphWidgetPlaceholder({ onAdd }: GraphWidgetPlaceholderProps) {
  return (
    <Box
      onClick={onAdd}
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        gap: 1,
        backgroundColor: "secondary.light",
        borderRadius: 3,
        border: "2px dashed",
        borderColor: "primary.main",
        opacity: 0.5,
        cursor: "pointer",
        py: 4,
        "&:hover": { opacity: 0.8 },
      }}
    >
      <AddCircleOutlineIcon />
      <Typography variant="caption">No widgets added yet, add here</Typography>
    </Box>
  );
}
