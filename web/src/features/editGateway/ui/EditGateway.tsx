import { useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import type {
  GatewayApiResponse,
  UpdateGatewayRequest,
} from "@entities/gateway";
import { DropDownSelect } from "@shared/ui/DropDownSelect";
import { fieldSx } from "@shared/ui/fieldSx";

export interface EditGatewayProps {
  open: boolean;
  gateway: GatewayApiResponse;
  onClose: () => void;
  onEdit: (id: string, payload: UpdateGatewayRequest) => Promise<void>;
  factoryOptions: { id: string; name: string }[];
  factoryAreaOptions: { id: string; name: string }[];
  onFactoryChange: (factoryId: string) => void;
  onErrorClear?: () => void;
  submitError?: string | null;
}

/**
 * Modal dialog for editing an existing gateway's details.
 * Pre-filled with the gateway's current name, description, factory, and factory area.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.gateway - The gateway being edited, used to pre-fill the form
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.onEdit - Called with the gateway ID and updated fields when the user confirms
 * @param props.factoryOptions - Available factories for the dropdown
 * @param props.factoryAreaOptions - Available factory areas for the dropdown (filtered by selected factory)
 * @param props.onFactoryChange - Called when the factory selection changes, to reload factory areas
 * @param props.onErrorClear - Called on any field change to clear a previously displayed server-side error
 * @param props.submitError - Optional server-side error message to display
 * @returns The rendered edit-gateway dialog
 */
export function EditGateway({
  open,
  gateway,
  onClose,
  onEdit,
  factoryOptions,
  factoryAreaOptions,
  onFactoryChange,
  onErrorClear,
  submitError,
}: EditGatewayProps) {
  const [name, setName] = useState(gateway.name);
  const [description, setDescription] = useState(gateway.description ?? "");
  const [factoryId, setFactoryId] = useState(gateway.factoryId);
  const [factoryAreaId, setFactoryAreaId] = useState(gateway.factoryAreaId);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [nameError, setNameError] = useState(false);

  const handleClose = () => {
    setName(gateway.name);
    setDescription(gateway.description ?? "");
    setFactoryId(gateway.factoryId);
    setFactoryAreaId(gateway.factoryAreaId);
    setNameError(false);
    onClose();
  };

  const handleFactoryChange = (newFactoryId: string) => {
    setFactoryId(newFactoryId);
    setFactoryAreaId("");
    onFactoryChange(newFactoryId);
    onErrorClear?.();
  };

  const handleSave = async () => {
    if (name.trim() === "") {
      setNameError(true);
      return;
    }
    const payload: UpdateGatewayRequest = {};
    if (name.trim() !== gateway.name) payload.name = name.trim();
    if (description.trim() !== (gateway.description ?? ""))
      payload.description = description.trim();
    if (factoryId !== gateway.factoryId) payload.factoryId = factoryId;
    if (factoryAreaId !== gateway.factoryAreaId)
      payload.factoryAreaId = factoryAreaId;

    if (Object.keys(payload).length === 0) {
      onClose();
      return;
    }

    setIsSubmitting(true);
    try {
      await onEdit(gateway.id, payload);
    } catch {
      // parent already sets submitError and shows snackbar
    } finally {
      setIsSubmitting(false);
    }
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
      <DialogTitle sx={{ color: "primary.main" }}>Edit gateway</DialogTitle>
      <DialogContent>
        <Box display="flex" flexDirection="column" gap={2}>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Name
            </Typography>
            <TextField
              value={name}
              onChange={(e) => {
                setName(e.target.value);
                if (nameError) setNameError(false);
                onErrorClear?.();
              }}
              fullWidth
              error={nameError}
              helperText={nameError ? "Name is required" : undefined}
              slotProps={{ htmlInput: { maxLength: 50 } }}
              sx={fieldSx}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Description
            </Typography>
            <TextField
              value={description}
              onChange={(e) => {
                setDescription(e.target.value);
                onErrorClear?.();
              }}
              fullWidth
              multiline
              maxRows={6}
              slotProps={{ htmlInput: { maxLength: 500 } }}
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
              onChange={(v) => {
                setFactoryAreaId(v);
                onErrorClear?.();
              }}
            />
          </Box>
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
          disabled={isSubmitting}
          autoFocus
        >
          Save
        </Button>
      </DialogActions>
    </Dialog>
  );
}
