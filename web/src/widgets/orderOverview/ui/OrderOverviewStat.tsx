import WarningAmberIcon from "@mui/icons-material/WarningAmber";
import Card from "@mui/material/Card";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

/** Props for the `OrderOverviewStat` component. */
interface OrderOverviewStatProps {
  /** Short label displayed above the value. */
  label: string;
  /** Main value to display. Use `"—"` for unavailable data. */
  value: string;
  /** Unit of measurement shown below the value. */
  unit?: string;
  /** MUI system color applied to the value (e.g. `"warning.main"`). */
  valueColor?: string;
  /** If set, wraps the value in a Tooltip — used for placeholder `—` values. */
  tooltip?: string;
  /** If set, shows a warning triangle with this text as the Tooltip. */
  warningTooltip?: string;
}

/**
 * Stat tile used in the order overview row.
 * Supports colored values, placeholder tooltips, and an inline warning indicator.
 * @param props - Component props
 * @param props.label - Descriptive label shown above the value
 * @param props.value - Main value displayed prominently
 * @param props.unit - Unit of measurement shown below the value
 * @param props.valueColor - MUI color token applied to the value typography
 * @param props.tooltip - Tooltip text wrapping the value for unavailable data
 * @param props.warningTooltip - If set, renders a WarningAmberIcon with this tooltip
 * @returns The rendered stat tile card
 */
export function OrderOverviewStat({
  label,
  value,
  unit,
  valueColor,
  tooltip,
  warningTooltip,
}: OrderOverviewStatProps) {
  const valueNode = (
    <Typography
      variant="h4"
      sx={{
        fontWeight: 600,
        textAlign: "center",
        whiteSpace: "nowrap",
        color: valueColor,
      }}
    >
      {value}
    </Typography>
  );

  return (
    <Card
      sx={{
        backgroundColor: "secondary.light",
        borderRadius: 3,
        p: 2,
        color: "primary.main",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        gap: 0.5,
        minWidth: 130,
        flex: "1 1 130px",
      }}
    >
      <Typography variant="subtitle2" sx={{ opacity: 0.7 }}>
        {label}
      </Typography>

      {tooltip ? <Tooltip title={tooltip}>{valueNode}</Tooltip> : valueNode}

      {unit && (
        <Typography variant="subtitle2" sx={{ opacity: 0.7 }}>
          {unit}
        </Typography>
      )}

      {warningTooltip && (
        <Tooltip title={warningTooltip}>
          <WarningAmberIcon
            sx={{ color: "warning.main", fontSize: 18, mt: 0.5 }}
          />
        </Tooltip>
      )}
    </Card>
  );
}
