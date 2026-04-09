import { useEffect, useState } from "react";

import { fetchSensorReading } from "@entities/sensor/api";
import type {
  SensorApiResponse,
  SensorReadingApiResponse,
} from "@entities/sensor/model/sensorSchema";
import CircleIcon from "@mui/icons-material/Circle";
import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";
import { formatReading, formatTimestamp } from "@shared/lib";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";

export interface AddDeviceProps {
  open: boolean;
  onClose: () => void;
  sensor: SensorApiResponse;
}

type InfoAllProps = {
  name: string;
  status: number;
  sensorEui: string;
  machine: string;
  lastReading: string;
  senProf: string;
};

function SensorAllInfo({
  name,
  status,
  sensorEui: euid,
  machine,
  lastReading,
  senProf,
}: InfoAllProps) {
  const theme = useTheme();

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
      <Typography>{euid}</Typography>
      <Typography>{machine}</Typography>
      <Typography>{formatTimestamp(lastReading)}</Typography>
      <Typography>{senProf}</Typography>
    </>
  );
}

/**
 * Modal dialog showing all fields and the latest sensor reading payload for a selected sensor.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.onClose - Called when the dialog should close
 * @param props.sensor - The sensor whose full details and latest reading are displayed
 * @returns The rendered sensor detail dialog
 */
export function SensorAllInfoPopUp(props: AddDeviceProps) {
  const { onClose, open, sensor } = props;
  const [reading, setReading] = useState<SensorReadingApiResponse | null>(null);

  useEffect(() => {
    if (!open) return;
    fetchSensorReading(sensor.deviceEui).then(setReading);
  }, [open, sensor.deviceEui]);

  const handleClose = () => {
    onClose();
  };

  const sensorAllDetails: string[] = [
    "Status",
    "Name",
    "DeviceEUI",
    "Machine",
    "Last reading",
    "Sensor profile",
  ];

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      maxWidth="lg"
      fullWidth
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.contrastText",
        },
      }}
    >
      <DialogTitle id="alert-dialog-title" sx={{ color: "primary.main" }}>
        {"All sensor info"}
      </DialogTitle>
      <DialogContent>
        <CategoryHeader
          categories={sensorAllDetails}
          columns={sensorAllDetails.length}
        >
          <DeviceRow key={sensor.id}>
            <SensorAllInfo
              name={sensor.name}
              status={sensor.status}
              lastReading={sensor.lastReading}
              sensorEui={sensor.deviceEui}
              machine={sensor.machine}
              senProf={sensor.sensorProfileId}
            />
          </DeviceRow>
        </CategoryHeader>
        {/* Show latest reading if available, otherwise show a message indicating no data */}
        {reading &&
        reading.payload &&
        Object.keys(reading.payload).length > 0 ? (
          <>
            <Typography
              sx={{ color: "primary.main", mt: 3, mb: 1, fontWeight: "bold" }}
            >
              Latest reading
            </Typography>
            <Typography
              variant="body2"
              sx={{ mb: 1, opacity: 0.7, color: "primary.main" }}
            >
              {formatTimestamp(reading.timestamp)}
            </Typography>
            <CategoryHeader
              categories={Object.keys(reading.payload)}
              columns={Object.keys(reading.payload).length}
            >
              <DeviceRow key="reading">
                {Object.values(reading.payload).map((value, i) => (
                  <Typography key={i}>{formatReading(value)}</Typography>
                ))}
              </DeviceRow>
            </CategoryHeader>
          </>
        ) : (
          // Show message when no reading data is available
          <Typography sx={{ mt: 2, opacity: 0.5, color: "primary.main" }}>
            No sensor data available — device may need configuration or is
            waiting for its first reading.
          </Typography>
        )}
      </DialogContent>
    </Dialog>
  );
}
