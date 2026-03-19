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

/**
 * Login page with email and password fields and a link to account registration.
 * @returns The rendered Login page
 */
export default function Base() {
  const { mode } = useContext(ThemeContext);
  const [loginValues, setLoginValues] = useState<Record<string, string>>({});
  const [newValues, setNewValues] = useState<Record<string, string>>({});
  const [fillError, setFillError] = useState(false);
  const [equalError, setEqualError] = useState(false);
  const [firstLogin, setFirstLogin] = useState(false);
  const navigate = useNavigate();

  const loginFields = ["email", "password"];
  const newPassFields = ["newPassword", "repeatPass"];

  const handleLogin = () => {
    const allFilled = loginFields.every(
      (field) => (loginValues[field] ?? "").trim() !== "",
    );
    const equalPass = newValues.newPassword === newValues.repeatPass;

    const newCreated = newPassFields.every(
      (field) => (newValues[field] ?? "").trim() !== "",
    );

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
    }

    setFillError(false);
    setFillError(false);
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
            onChange={(e) =>
              setNewValues((prev) => ({ ...prev, newPassword: e.target.value }))
            }
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
          //{...(!error ? { href: "/" } : {})} // Go to home page if no error
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
      </Box>
    </Box>
  );
}
