import { useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import type { SensorApiResponse, UpdateSensorRequest } from "@entities/sensor";
import { DropDownSelect } from "@shared/ui/DropDownSelect";

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
};

export interface EditSensorProps {
  open: boolean;
  sensor: SensorApiResponse;
  onClose: () => void;
  onEdit: (id: string, payload: UpdateSensorRequest) => Promise<void>;
  factoryOptions: { id: string; name: string }[];
  factoryAreaOptions: { id: string; name: string }[];
  sensorProfileOptions: { id: string; name: string }[];
  productionResourceOptions: { id: string; name: string }[];
  voltageOptions: { id: string; name: string }[];
  onFactoryChange: (factoryId: string) => void;
  submitError?: string | null;
}

/**
 * Modal dialog for editing an existing sensor's details.
 * Pre-filled with the sensor's current values.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.sensor - The sensor being edited, used to pre-fill the form
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.onEdit - Called with the sensor ID and updated fields when the user confirms
 * @param props.factoryOptions - Available factories for the dropdown
 * @param props.factoryAreaOptions - Available factory areas filtered by selected factory
 * @param props.sensorProfileOptions - Available sensor profiles for the dropdown
 * @param props.productionResourceOptions - Available production resources for the dropdown
 * @param props.voltageOptions - Available voltage options shown when electricity sensor is enabled
 * @param props.onFactoryChange - Called when the factory selection changes, to reload factory areas
 * @param props.submitError - Optional server-side error message to display
 * @returns The rendered edit-sensor dialog
 */
export function EditSensor({
  open,
  sensor,
  onClose,
  onEdit,
  factoryOptions,
  factoryAreaOptions,
  sensorProfileOptions,
  productionResourceOptions,
  voltageOptions,
  onFactoryChange,
  submitError,
}: EditSensorProps) {
  const [name, setName] = useState(sensor.name);
  const [description, setDescription] = useState(sensor.description ?? "");
  const [factoryId, setFactoryId] = useState(sensor.factory);
  const [factoryAreaId, setFactoryAreaId] = useState(sensor.factoryAreaId);
  const [sensorProfileId, setSensorProfileId] = useState(
    sensor.sensorProfileId,
  );
  const [electricitySensor, setElectricitySensor] = useState(
    sensor.electricitySensor,
  );
  const [voltage, setVoltage] = useState(
    sensor.voltage != null ? String(sensor.voltage) : "",
  );
  const [productionResource, setProductionResource] = useState(
    sensor.productionResource != null ? String(sensor.productionResource) : "",
  );

  const handleClose = () => {
    setName(sensor.name);
    setDescription(sensor.description ?? "");
    setFactoryId(sensor.factory);
    setFactoryAreaId(sensor.factoryAreaId);
    setSensorProfileId(sensor.sensorProfileId);
    setElectricitySensor(sensor.electricitySensor);
    setVoltage(sensor.voltage != null ? String(sensor.voltage) : "");
    setProductionResource(
      sensor.productionResource != null
        ? String(sensor.productionResource)
        : "",
    );
    onClose();
  };

  const handleFactoryChange = (newFactoryId: string) => {
    setFactoryId(newFactoryId);
    setFactoryAreaId("");
    onFactoryChange(newFactoryId);
  };

  const handleSave = () => {
    const payload: UpdateSensorRequest = {};

    if (name.trim() !== sensor.name) payload.name = name.trim();
    if (description.trim() !== (sensor.description ?? ""))
      payload.description = description.trim();
    if (factoryId !== sensor.factory) payload.factoryId = factoryId;
    if (factoryAreaId !== sensor.factoryAreaId)
      payload.factoryAreaId = factoryAreaId;
    if (sensorProfileId !== sensor.sensorProfileId)
      payload.sensorProfileId = sensorProfileId;
    if (electricitySensor !== sensor.electricitySensor)
      payload.electricitySensor = electricitySensor;

    const parsedVoltage = voltage !== "" ? Number(voltage) : null;
    if (parsedVoltage !== sensor.voltage) payload.voltage = parsedVoltage;

    const parsedProductionResource =
      productionResource !== "" ? Number(productionResource) : null;
    if (parsedProductionResource !== sensor.productionResource)
      payload.productionResource = parsedProductionResource;

    if (Object.keys(payload).length === 0) {
      onClose();
      return;
    }

    onEdit(sensor.id, payload).catch(() => {});
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      maxWidth="sm"
      fullWidth
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.contrastText",
        },
      }}
    >
      <DialogTitle sx={{ color: "primary.main" }}>Edit sensor</DialogTitle>
      <DialogContent>
        <Box display="flex" flexDirection="column" gap={2}>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Name
            </Typography>
            <TextField
              value={name}
              onChange={(e) => setName(e.target.value)}
              fullWidth
              sx={fieldSx}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Description
            </Typography>
            <TextField
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              fullWidth
              multiline
              maxRows={6}
              sx={fieldSx}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Factory
            </Typography>
            <DropDownSelect
              options={factoryOptions}
              value={factoryId}
              onChange={handleFactoryChange}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Factory area
            </Typography>
            <DropDownSelect
              options={factoryAreaOptions}
              value={factoryAreaId}
              onChange={setFactoryAreaId}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Sensor profile
            </Typography>
            <DropDownSelect
              options={sensorProfileOptions}
              value={sensorProfileId}
              onChange={setSensorProfileId}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Production resource
            </Typography>
            <DropDownSelect
              options={productionResourceOptions}
              value={productionResource}
              onChange={setProductionResource}
            />
          </Box>
          <Box display="flex" alignItems="center" gap={1}>
            <Typography variant="body2" color="primary.main">
              Electricity sensor
            </Typography>
            <Checkbox
              checked={electricitySensor}
              onChange={(_, checked) => {
                setElectricitySensor(checked);
                if (!checked) setVoltage("");
              }}
              slotProps={{ input: { "aria-label": "electricity sensor" } }}
            />
          </Box>
          {electricitySensor && (
            <Box>
              <Typography variant="body2" color="primary.main" mb={0.5}>
                Voltage
              </Typography>
              <DropDownSelect
                options={voltageOptions}
                value={voltage}
                onChange={setVoltage}
              />
            </Box>
          )}
          {submitError && <Typography color="error">{submitError}</Typography>}
        </Box>
      </DialogContent>
      <DialogActions>
        <Button
          sx={{
            backgroundColor: "primary.main",
            color: "primary.contrastText",
          }}
          onClick={handleClose}
        >
          Close
        </Button>
        <Button
          sx={{
            backgroundColor: "primary.main",
            color: "primary.contrastText",
          }}
          onClick={handleSave}
          autoFocus
        >
          Save
        </Button>
      </DialogActions>
    </Dialog>
  );
}
