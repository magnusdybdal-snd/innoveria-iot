import { useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import type { UpdateUserRequest } from "@entities/user";
import { validatePassword, type PasswordErrors } from "@shared/lib";

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
  "& .MuiSelect-icon": { color: "primary.main" },
};

export interface EditUserProps {
  open: boolean;
  user: { id: string; name: string; email: string };
  onClose: () => void;
  onEdit: (userId: string, payload: UpdateUserRequest) => void;
  submitError?: string | null;
}

/**
 * Modal dialog for editing an existing user's details.
 * Pre-filled with the user's current name and email.
 * Password is optional — leave blank to keep the existing password.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.user - The user being edited, used to pre-fill the form
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.onEdit - Called with the user ID and updated fields when the user confirms
 * @param props.submitError - Optional server-side error message to display
 * @returns The rendered edit-user dialog
 */
export function EditUser({
  open,
  user,
  onClose,
  onEdit,
  submitError,
}: EditUserProps) {
  const [name, setName] = useState(user.name);
  const [email, setEmail] = useState(user.email);
  const [password, setPassword] = useState("");
  const [emailError, setEmailError] = useState(false);
  const [passwordErrors, setPasswordErrors] = useState<PasswordErrors>({
    length: false,
    lowercase: false,
    uppercase: false,
    number: false,
    symbol: false,
  });

  const isValidEmail = (value: string) =>
    /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value);

  const noPasswordErrors = {
    length: false,
    lowercase: false,
    uppercase: false,
    number: false,
    symbol: false,
  };

  const handleClose = () => {
    setName(user.name);
    setEmail(user.email);
    setPassword("");
    setEmailError(false);
    setPasswordErrors(noPasswordErrors);
    onClose();
  };

  const handleSave = () => {
    if (!isValidEmail(email)) {
      setEmailError(true);
      return;
    }
    if (password) {
      const pwErrors = validatePassword(password);
      if (Object.values(pwErrors).some(Boolean)) {
        setPasswordErrors(pwErrors);
        return;
      }
    }
    setEmailError(false);
    setPasswordErrors(noPasswordErrors);

    const payload: UpdateUserRequest = {};
    if (name.trim() !== user.name) payload.name = name.trim();
    if (email !== user.email) payload.email = email;
    if (password) payload.password = password;

    onEdit(user.id, payload);
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
      <DialogTitle sx={{ color: "primary.main" }}>Edit user</DialogTitle>
      <DialogContent>
        <Box display="flex" flexDirection="column" gap={2}>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Name
            </Typography>
            <TextField
              placeholder="Enter name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              fullWidth
              sx={fieldSx}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Email
            </Typography>
            <TextField
              placeholder="Enter email"
              type="email"
              value={email}
              onChange={(e) => {
                setEmail(e.target.value.toLowerCase());
                setEmailError(false);
              }}
              fullWidth
              sx={fieldSx}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              New password{" "}
              <Typography
                component="span"
                variant="body2"
                color="text.secondary"
              >
                (leave blank to keep current password)
              </Typography>
            </Typography>
            <TextField
              placeholder="Enter new password"
              type="password"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
                setPasswordErrors(noPasswordErrors);
              }}
              fullWidth
              sx={fieldSx}
            />
          </Box>
          {emailError && (
            <Typography color="error">
              Please enter a valid email address
            </Typography>
          )}
          {passwordErrors.length && (
            <Typography color="error">
              Password must be at least 8 characters
            </Typography>
          )}
          {passwordErrors.lowercase && (
            <Typography color="error">
              Password must contain a lowercase letter
            </Typography>
          )}
          {passwordErrors.uppercase && (
            <Typography color="error">
              Password must contain an uppercase letter
            </Typography>
          )}
          {passwordErrors.number && (
            <Typography color="error">
              Password must contain a number
            </Typography>
          )}
          {passwordErrors.symbol && (
            <Typography color="error">
              Password must contain a symbol (!@#$...)
            </Typography>
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
