import { useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import Stack from "@mui/material/Stack";
import { useTheme } from "@mui/material/styles";
import TextField from "@mui/material/TextField";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";

import { exportMeasurements } from "@entities/sensor/api";
import { ActionMenu } from "@shared/ui/actionMenu";
import { DeleteConfirmation } from "@shared/ui/DeleteConfirmation";

type InfoMainProps = {
  name: string;
  status: number;
  lastReading: string;
  description: string | null;
  deviceEui: string;
  onClick: () => void;
  onDelete: () => void;
  onEdit: () => void;
};

/**
 * Renders the primary sensor row cells: status indicator, name, last reading, and an action menu.
 * @param props - Component props
 * @param props.name - Display name of the sensor
 * @param props.status - Numeric status code: 0 = online, 1 = warning, 2 = offline
 * @param props.lastReading - Timestamp of the most recent sensor reading
 * @param props.description - Description shown in a tooltip on hover, or null if absent
 * @param props.deviceEui - LoRaWAN Device EUI used to identify the sensor
 * @param props.onClick - Called when the user clicks "Extra sensor info" to open the detail dialog
 * @param props.onDelete - Called when the user confirms deletion
 * @param props.onEdit - Called when the user clicks Edit in the action menu
 * @returns The rendered sensor row cells
 */
export function SensorMainInfo({
  name,
  status,
  lastReading,
  description,
  deviceEui,
  onClick,
  onDelete,
  onEdit,
}: InfoMainProps) {
  const theme = useTheme();
  const [deleteOpen, setDeleteOpen] = useState(false);
  const [exportOpen, setExportOpen] = useState(false);
  const [fromDate, setFromDate] = useState("");
  const [toDate, setToDate] = useState("");
  const [exporting, setExporting] = useState(false);
  const [exportError, setExportError] = useState<string | null>(null);

  const handleDeleteConfirm = () => {
    onDelete();
    setDeleteOpen(false);
  };

  const handleExport = async () => {
    if (!fromDate || !toDate) return;
    setExporting(true);
    setExportError(null);
    try {
      const from = new Date(`${fromDate}T00:00:00`);
      const to = new Date(`${toDate}T23:59:59.999`);
      await exportMeasurements(deviceEui, from.toISOString(), to.toISOString());
      setExportOpen(false);
    } catch (err) {
      if (!(err instanceof Error))
        console.error("Unexpected export error:", err);
      setExportError(err instanceof Error ? err.message : "Export failed");
    } finally {
      setExporting(false);
    }
  };

  const menuItems = [
    { label: "Edit", onClick: onEdit },
    { label: "Export CSV", onClick: () => setExportOpen(true) },
    { label: "Delete", onClick: () => setDeleteOpen(true) },
  ];

  const statusColor = (status: number) => {
    const s = theme.palette.status;
    switch (status) {
      case 0:
        return s.online;
      case 1:
        return s.warning;
      case 2:
        return s.offline;
      default:
        return s.unknown;
    }
  };

  return (
    <>
      <CircleIcon
        sx={{
          color: statusColor(Number(status)),
          fontSize: 14,
          alignSelf: "center",
          filter: "drop-shadow(0 0 1px grey)",
        }}
      />
      <Typography>{name}</Typography>
      <Typography>{lastReading}</Typography>
      {description ? (
        <Tooltip title={description} enterDelay={800} arrow>
          <InfoOutlinedIcon
            sx={{
              fontSize: 18,
              alignSelf: "center",
              color: "primary.main",
              cursor: "default",
            }}
          />
        </Tooltip>
      ) : (
        <span />
      )}
      <Button
        variant="outlined"
        sx={{
          backgroundColor: "primary.main",
          color: "primary.dark",
          "&:hover": { backgroundColor: "primary.main" },
          borderRadius: 2,
          textTransform: "none",
          fontSize: 15,
        }}
        onClick={onClick}
      >
        Extra sensor info
      </Button>
      <ActionMenu items={menuItems} />
      <DeleteConfirmation
        open={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDeleteConfirm}
      />
      <Dialog
        open={exportOpen}
        onClose={() => {
          setExportOpen(false);
          setExportError(null);
        }}
        fullWidth
        maxWidth="xs"
      >
        <DialogTitle>Export CSV</DialogTitle>
        <DialogContent>
          <Stack spacing={2} sx={{ mt: 1 }}>
            {exportError && <Alert severity="error">{exportError}</Alert>}
            <TextField
              label="From"
              type="date"
              value={fromDate}
              onChange={(e) => setFromDate(e.target.value)}
              slotProps={{ inputLabel: { shrink: true } }}
              fullWidth
            />
            <TextField
              label="To"
              type="date"
              value={toDate}
              onChange={(e) => setToDate(e.target.value)}
              slotProps={{ inputLabel: { shrink: true } }}
              fullWidth
            />
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => {
              setExportOpen(false);
              setExportError(null);
            }}
          >
            Cancel
          </Button>
          <Button
            onClick={handleExport}
            disabled={!fromDate || !toDate || exporting}
            variant="contained"
          >
            {exporting ? "Downloading…" : "Download"}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
