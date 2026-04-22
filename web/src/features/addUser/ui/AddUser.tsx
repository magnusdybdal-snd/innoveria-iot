import { useState } from "react";

import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import MenuItem from "@mui/material/MenuItem";
import TextField from "@mui/material/TextField";

import type { CreateUserRequest } from "@entities/user";
import { validatePassword, type PasswordErrors } from "@shared/lib";

const ROLES = [
  { label: "User", value: "FACTORY_WORKER" },
  { label: "Admin", value: "PLATFORM_ADMIN" },
];

const inputSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
  "& .MuiInputLabel-root": { color: "primary.main" },
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
  const [role, setRole] = useState("FACTORY_WORKER");
  const [fillError, setFillError] = useState(false);
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
    setName("");
    setEmail("");
    setPassword("");
    setRole("FACTORY_WORKER");
    setFillError(false);
    setEmailError(false);
    setPasswordErrors(noPasswordErrors);
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
    setPasswordErrors(noPasswordErrors);
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
      <DialogContent className="flex flex-col gap-4 pt-2">
        <TextField
          label="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          fullWidth
          sx={inputSx}
        />
        <TextField
          label="Email"
          type="email"
          value={email}
          onChange={(e) => {
            setEmail(e.target.value.toLowerCase());
            setEmailError(false);
          }}
          fullWidth
          sx={inputSx}
        />
        <TextField
          label="Password"
          type="password"
          value={password}
          onChange={(e) => {
            setPassword(e.target.value);
            setPasswordErrors(noPasswordErrors);
          }}
          fullWidth
          sx={inputSx}
        />
        <TextField
          label="Role"
          select
          value={role}
          onChange={(e) => setRole(e.target.value)}
          fullWidth
          sx={inputSx}
        >
          {ROLES.map((r) => (
            <MenuItem key={r.value} value={r.value}>
              {r.label}
            </MenuItem>
          ))}
        </TextField>
        {fillError && (
          <div style={{ color: "red" }}>
            Name, email and password are required
          </div>
        )}
        {emailError && (
          <div style={{ color: "red" }}>Please enter a valid email address</div>
        )}
        {passwordErrors.length && (
          <div style={{ color: "red" }}>
            Password must be at least 8 characters
          </div>
        )}
        {passwordErrors.lowercase && (
          <div style={{ color: "red" }}>
            Password must contain a lowercase letter
          </div>
        )}
        {passwordErrors.uppercase && (
          <div style={{ color: "red" }}>
            Password must contain an uppercase letter
          </div>
        )}
        {passwordErrors.number && (
          <div style={{ color: "red" }}>Password must contain a number</div>
        )}
        {passwordErrors.symbol && (
          <div style={{ color: "red" }}>
            Password must contain a symbol (!@#$...)
          </div>
        )}
        {submitError && <div style={{ color: "red" }}>{submitError}</div>}
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
