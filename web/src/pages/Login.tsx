import { useContext, useState } from "react";

import innLogoDark from "@assets/innoveriaDark.png";
import innLogoLight from "@assets/innoveriaLight.png";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { green } from "@mui/material/colors";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { ThemeContext } from "@shared/config/theme/themeContext";
import { useNavigate } from "react-router";

const textFieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
  "& .MuiInputLabel-root": { color: "primary.main" },
  "& .MuiInputLabel-root.Mui-focused": { color: "primary.main" },
};

const loginFields = ["email", "password"];
const newPassFields = ["newPassword", "repeatPass"];

/**
 * Login page with email and password fields and a link to account registration.
 * @returns The rendered Login page
 */
export default function Base() {
  const { mode } = useContext(ThemeContext);

  // Field values
  const [loginValues, setLoginValues] = useState<Record<string, string>>({});
  const [newValues, setNewValues] = useState<Record<string, string>>({});

  // Error types
  const [fillError, setFillError] = useState(false);
  const [equalError, setEqualError] = useState(false);

  const [firstLogin, setFirstLogin] = useState(false);
  const navigate = useNavigate();

  const hasLowercase = (str: string) => /[a-z]/.test(str);
  const hasUppercase = (str: string) => /[A-Z]/.test(str);
  const hasNumber = (str: string) => /[0-9]/.test(str);
  const hasSymbol = (str: string) => /[!@#$%^&*()_\-+=]/.test(str);

  const [passErrors, setPassErrors] = useState({
    length: false,
    lowercase: false,
    uppercase: false,
    number: false,
    symbol: false,
  });

  const validatePassword = (password: string) => ({
    length: password.length < 8,
    lowercase: !hasLowercase(password),
    uppercase: !hasUppercase(password),
    number: !hasNumber(password),
    symbol: !hasSymbol(password),
  });

  const handleLogin = () => {
    const allFilled = loginFields.every(
      (field) => (loginValues[field] ?? "").trim() !== "",
    );
    const newCreated = newPassFields.every(
      (field) => (newValues[field] ?? "").trim() !== "",
    );

    const equalPass = newValues.newPassword === newValues.repeatPass;

    if (!allFilled) {
      setFillError(true);
      return;
    }

    if (!equalPass) {
      setEqualError(true);
      return;
    }

    if (!firstLogin) {
      setFirstLogin(true);
      return;
    } else {
      if (!newCreated) {
        setFillError(true);
        return;
      }

      const password = newValues.newPassword ?? "";

      const errors = validatePassword(password);

      setPassErrors(errors);

      const hasAnyError = Object.values(errors).some(Boolean);

      if (hasAnyError) {
        return;
      }
    }

    setFillError(false);
    setEqualError(false);
    navigate("/");
  };

  return (
    <Box
      className="flex h-screen"
      sx={{
        backgroundColor: "primary.dark",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <Box
        sx={{
          backgroundColor: "secondary.light",
          borderRadius: 3,
          width: 400,
          p: 4,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: 2,
        }}
      >
        <img
          src={mode ? innLogoDark : innLogoLight}
          alt="Innoveria logo"
          style={{ width: 310 }}
        />
        <Typography variant="h5">
          {firstLogin ? "Create a new password" : "Log in"}
        </Typography>
        {!firstLogin && (
          <TextField
            label="Email"
            fullWidth
            sx={textFieldSx}
            value={loginValues.email ?? ""}
            onChange={(e) =>
              setLoginValues((prev) => ({ ...prev, email: e.target.value }))
            }
          />
        )}
        {!firstLogin && (
          <TextField
            label="Password"
            type="password"
            fullWidth
            sx={textFieldSx}
            value={loginValues.password ?? ""}
            onChange={(e) =>
              setLoginValues((prev) => ({ ...prev, password: e.target.value }))
            }
          />
        )}
        {firstLogin && (
          <TextField
            label="New password"
            type="password"
            fullWidth
            sx={textFieldSx}
            value={newValues.newPassword ?? ""}
            onChange={(e) => {
              const value = e.target.value;

              setNewValues((prev) => ({
                ...prev,
                newPassword: value,
              }));

              setPassErrors(validatePassword(value));
            }}
          />
        )}
        {firstLogin && (
          <TextField
            label="Repeat password"
            type="password"
            fullWidth
            sx={textFieldSx}
            value={newValues.repeatPass ?? ""}
            onChange={(e) =>
              setNewValues((prev) => ({ ...prev, repeatPass: e.target.value }))
            }
          />
        )}
        <Button
          onClick={handleLogin}
          fullWidth
          sx={{
            backgroundColor: green[500],
            color: "white",
            "&:hover": { backgroundColor: green[800] },
          }}
        >
          Log in
        </Button>
        {fillError && (
          <Typography color="error">All fields must be filled</Typography>
        )}
        {equalError && (
          <Typography color="error">
            The repeated password is not the same
            <br />
            as the new password
          </Typography>
        )}
        {passErrors.length && (
          <Typography color="error">Must at least 8 characters</Typography>
        )}
        {passErrors.lowercase && (
          <Typography color="error">Must contain a lowercase letter</Typography>
        )}
        {passErrors.uppercase && (
          <Typography color="error">
            Must contain an uppercase letter
          </Typography>
        )}
        {passErrors.number && (
          <Typography color="error">Must contain a number</Typography>
        )}
        {passErrors.symbol && (
          <Typography color="error">Must contain a symbol (!@#$...)</Typography>
        )}
      </Box>
    </Box>
  );
}
