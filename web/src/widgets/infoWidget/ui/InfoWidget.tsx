import Card from "@mui/material/Card";
import Typography from "@mui/material/Typography";

/** Props for the `InfoWidget` component. */
interface InfoWidgetProps {
  /** Short label displayed above the value (e.g. "Status", "Product"). */
  label: string;
  /** Main content displayed prominently in the center of the widget. */
  value: string;
  /** Unit of measurement for the value (e.g. "kg", "m"). */
  unit?: string;
}

/**
 * Square info card that displays a label and a single value.
 * Styling matches the graphWidget card appearance.
 * @param props - Component props
 * @param props.label - Descriptive label shown above the value
 * @param props.value - Main text value shown in the center of the card
 * @param props.unit - Unit of measurement for the value, shown below the value
 * @returns The rendered info widget card
 */
export function InfoWidget({ label, value, unit }: InfoWidgetProps) {
  return (
    <Card
      sx={{
        backgroundColor: "secondary.light",
        borderRadius: 3,
        p: 2,
        color: "primary.main",
        aspectRatio: "1",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        gap: 1,
        maxWidth: 200,
      }}
    >
      <Typography variant="subtitle2" sx={{ opacity: 0.7 }}>
        {label}
      </Typography>
      <Typography variant="h2" sx={{ fontWeight: 600, textAlign: "center" }}>
        {value}
      </Typography>
      <Typography variant="subtitle2" sx={{ opacity: 0.7 }}>
        {unit}
      </Typography>
    </Card>
  );
}
