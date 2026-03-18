import { useContext, useState } from "react";

import innLogoDark from "@assets/innoveriaDark.png";
import innLogoLight from "@assets/innoveriaLight.png";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { green } from "@mui/material/colors";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
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

/**
 * Login page with username and password fields and a link to account registration.
 * @returns The rendered Login page
 */
export default function Base() {
  const { mode } = useContext(ThemeContext);
  const [values, setValues] = useState<Record<string, string>>({});
  const [error, setError] = useState(false);

  const loginFields = ["Username", "Password"];

  const handleLogin = () => {
    const allFilled = loginFields.every(
      (field) => (values[field] ?? "").trim() !== "",
    );

    if (!allFilled) {
      setError(true);
      return;
    }

    setError(false);
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
        <Typography variant="h5">Log in</Typography>
        <TextField
          label="Username"
          fullWidth
          sx={textFieldSx}
          value={values["Username"] ?? ""}
          onChange={(e) =>
            setValues((prev) => ({ ...prev, Username: e.target.value }))
          }
        />
        <TextField
          label="Password"
          type="password"
          fullWidth
          sx={textFieldSx}
          value={values["Password"] ?? ""}
          onChange={(e) =>
            setValues((prev) => ({ ...prev, Password: e.target.value }))
          }
        />
        <Button
          onClick={handleLogin}
          {...(!error ? { href: "/" } : {})} // Go to home page if no error
          fullWidth
          sx={{
            backgroundColor: green[500],
            color: "white",
            "&:hover": { backgroundColor: green[800] },
          }}
        >
          Log in
        </Button>
        {error && (
          <Typography color="error">All fields must be filled</Typography>
        )}
      </Box>
    </Box>
  );
}
