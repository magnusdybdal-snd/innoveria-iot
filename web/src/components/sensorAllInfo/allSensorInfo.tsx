import { useEffect, useState } from "react";

import CircleIcon from "@mui/icons-material/Circle";
import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";

import type { SensorReadingApiResponse } from "@/API/apiClient";
import { fetchSensorReading } from "@/API/fetchSensorReading";
import { CategoryHeader } from "@/components/CategoryHeader";
import { DeviceRow } from "@/components/gatewayRow";
import type { Sensor } from "@/mocks/sensors.ts";

export interface AddDeviceProps {
  open: boolean;
  onClose: () => void;
  sensor: Sensor;
}

type InfoAllProps = {
  name: string;
  status: number;
  euid: string;
  machine: string;
  lastReading: string;
  appKey: string;
  devProf: string;
};

{
  /*Format for all single sensor info*/
}
function SensorAllInfo({
  name,
  status,
  euid,
  machine,
  lastReading,
  appKey,
  devProf,
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
      <Typography>{lastReading}</Typography>
      <Typography>{appKey}</Typography>
      <Typography>{devProf}</Typography>
    </>
  );
}

export function SensorAllInfoPopUp(props: AddDeviceProps) {
  const { onClose, open, sensor } = props;
  const [reading, setReading] = useState<SensorReadingApiResponse | null>(null);

  useEffect(() => {
    if (!open) return;
    fetchSensorReading("e112f47bdd366c03").then(setReading); // TODO: replace hardcoded value with 'sensor.euid'
  }, [open, sensor.euid]);

  const handleClose = () => {
    onClose();
  };

  const sensorAllDetails: string[] = [
    "Status",
    "Name",
    "DeviceEUI",
    "Machine",
    "Last reading",
    "Application key",
    "Device profile",
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
              euid={sensor.euid}
              machine={sensor.machine}
              appKey={sensor.appKey}
              devProf={sensor.devProf}
            />
          </DeviceRow>
        </CategoryHeader>

        {reading && (
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
              {new Date(reading.timestamp).toLocaleString()}
            </Typography>
            <CategoryHeader
              categories={Object.keys(reading.payload)}
              columns={Object.keys(reading.payload).length}
            >
              <DeviceRow key="reading">
                {Object.values(reading.payload).map((value, i) => (
                  <Typography key={i}>{String(value)}</Typography>
                ))}
              </DeviceRow>
            </CategoryHeader>
          </>
        )}

        {reading === null && (
          <Typography sx={{ mt: 2, opacity: 0.5, color: "primary.main" }}>
            No reading available
          </Typography>
        )}
      </DialogContent>
    </Dialog>
  );
}
