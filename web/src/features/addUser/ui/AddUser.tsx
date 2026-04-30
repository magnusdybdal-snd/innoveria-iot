import { useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import MenuItem from "@mui/material/MenuItem";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import type { CreateUserRequest, UserRole } from "@entities/user";
import {
  emptyPasswordErrors,
  isValidEmail,
  validatePassword,
  type PasswordErrors,
} from "@shared/lib";

const ROLES = [
  { label: "User", value: "FACTORY_WORKER" },
  { label: "Admin", value: "PLATFORM_ADMIN" },
];

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
  "& .MuiSelect-icon": { color: "primary.main" },
};

export interface AddUserProps {
  open: boolean;
  companyId: string;
  onClose: () => void;
  onAdd: (user: CreateUserRequest) => void;
  submitError?: string | null;
}

/**
 * Modal dialog for creating a new user.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.companyId - Pre-filled company ID for the new user
 * @param props.onClose - Called when the dialog should close without submitting
 * @param props.onAdd - Called with the validated user data when the user confirms
 * @param props.submitError - Optional server-side error message to display
 * @returns The rendered add-user dialog
 */
export function AddUser({
  open,
  companyId,
  onClose,
  onAdd,
  submitError,
}: AddUserProps) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState<UserRole>("FACTORY_WORKER");
  const [fillError, setFillError] = useState(false);
  const [emailError, setEmailError] = useState(false);
  const [passwordErrors, setPasswordErrors] =
    useState<PasswordErrors>(emptyPasswordErrors);

  const handleClose = () => {
    setName("");
    setEmail("");
    setPassword("");
    setRole("FACTORY_WORKER");
    setFillError(false);
    setEmailError(false);
    setPasswordErrors(emptyPasswordErrors);
    onClose();
  };

  const handleAdd = () => {
    if (!name.trim() || !email.trim() || !password) {
      setFillError(true);
      return;
    }
    if (!isValidEmail(email)) {
      setFillError(false);
      setEmailError(true);
      return;
    }
    const pwErrors = validatePassword(password);
    if (Object.values(pwErrors).some(Boolean)) {
      setFillError(false);
      setEmailError(false);
      setPasswordErrors(pwErrors);
      return;
    }
    setFillError(false);
    setEmailError(false);
    setPasswordErrors(emptyPasswordErrors);
    onAdd({
      companyId,
      name: name.trim(),
      email: email.trim(),
      password,
      role,
    });
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
      <DialogTitle sx={{ color: "primary.main" }}>Add user</DialogTitle>
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
              Password
            </Typography>
            <TextField
              placeholder="Enter password"
              type="password"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
                setPasswordErrors(emptyPasswordErrors);
              }}
              fullWidth
              sx={fieldSx}
            />
          </Box>
          <Box>
            <Typography variant="body2" color="primary.main" mb={0.5}>
              Role
            </Typography>
            <TextField
              select
              value={role}
              onChange={(e) => setRole(e.target.value as UserRole)}
              fullWidth
              sx={fieldSx}
            >
              {ROLES.map((r) => (
                <MenuItem key={r.value} value={r.value}>
                  {r.label}
                </MenuItem>
              ))}
            </TextField>
          </Box>
          {fillError && (
            <Typography color="error">
              Name, email and password are required
            </Typography>
          )}
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
          onClick={handleAdd}
          autoFocus
        >
          Add
        </Button>
      </DialogActions>
    </Dialog>
  );
}
