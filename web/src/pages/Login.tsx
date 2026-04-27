import { useContext, useState } from "react";

import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { green } from "@mui/material/colors";
import IconButton from "@mui/material/IconButton";
import InputAdornment from "@mui/material/InputAdornment";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import innLogoDark from "@assets/innoveriaDark.png";
import innLogoLight from "@assets/innoveriaLight.png";
import { postLogin } from "@entities/user";
import { ThemeContext } from "@shared/config/theme/themeContext";

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
  const [loginError, setLoginError] = useState(false);

  const [firstLogin, setFirstLogin] = useState(false);

  // Show password
  const [showPassword, setShowPassword] = useState({
    login: false,
    new: false,
  });
  const togglePassword = (key: "login" | "new") => {
    setShowPassword((prev) => ({
      ...prev,
      [key]: !prev[key],
    }));
  };

  const getPasswordAdornment = (key: "login" | "new") => (
    <InputAdornment position="end">
      <IconButton onClick={() => togglePassword(key)} edge="end">
        {showPassword[key] ? <VisibilityOffIcon /> : <VisibilityIcon />}
      </IconButton>
    </InputAdornment>
  );

  // Password validation helpers
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

  const handleLogin = async (loginData: {
    email: string;
    password: string;
  }) => {
    setFirstLogin(false);
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
      //setFirstLogin(true);
      //return;
    }

    if (!newCreated) {
      //setFillError(true);
      //return;
    }

    /*const password = newValues.newPassword ?? "";

    const errors = validatePassword(password);

    setPassErrors(errors);

    if (Object.values(errors).some(Boolean)) {
      return;
    }*/

    setFillError(false);
    setEqualError(false);
    setLoginError(false);

    try {
      const data = await postLogin(loginData);

      localStorage.setItem("access_token", data.accessToken);

      window.location.href = "/";
    } catch (err) {
      setLoginError(true);
      console.error(err);
    }
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
              setLoginValues((prev) => ({
                ...prev,
                email: e.target.value.toLowerCase(),
              }))
            }
          />
        )}
        {!firstLogin && (
          <TextField
            label="Password"
            type={showPassword.login ? "text" : "password"}
            fullWidth
            sx={textFieldSx}
            value={loginValues.password ?? ""}
            onChange={(e) =>
              setLoginValues((prev) => ({ ...prev, password: e.target.value }))
            }
            slotProps={{
              input: {
                endAdornment: getPasswordAdornment("login"),
              },
            }}
          />
        )}
        {firstLogin && (
          <TextField
            label="New password"
            type={showPassword.new ? "text" : "password"}
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
            slotProps={{
              input: {
                endAdornment: getPasswordAdornment("new"),
              },
            }}
          />
        )}
        {firstLogin && (
          <TextField
            label="Repeat password"
            type={showPassword.new ? "text" : "password"}
            fullWidth
            sx={textFieldSx}
            value={newValues.repeatPass ?? ""}
            onChange={(e) =>
              setNewValues((prev) => ({ ...prev, repeatPass: e.target.value }))
            }
            slotProps={{
              input: {
                endAdornment: getPasswordAdornment("new"),
              },
            }}
          />
        )}
        <Button
          onClick={() =>
            handleLogin({
              email: loginValues.email ?? "",
              password: loginValues.password ?? "",
            })
          }
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
        {loginError && (
          <Typography color="error">Email or password is wrong</Typography>
        )}
        {passErrors.length && (
          <Typography color="error">Must be at least 8 characters</Typography>
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
