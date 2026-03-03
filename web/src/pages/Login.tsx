import { useContext } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { green } from "@mui/material/colors";
import Link from "@mui/material/Link";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import innLogoDark from "@/assets/innoveriaDark.png";
import innLogoLight from "@/assets/innoveriaLight.png";
import { ThemeContext } from "@/theme/color/themeContext";

/**
 * Login page with username and password fields and a link to account registration.
 * @returns The rendered Login page
 */
export default function Base() {
  const { mode } = useContext(ThemeContext);

  return (
    <div className="flex h-screen">
      <Box
        sx={{
          backgroundColor: "primary.dark",
          color: "primary.main",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          textAlign: "center",
        }}
        className="flex-1 overflow-auto"
      >
        <Box
          sx={{
            backgroundColor: "secondary.light",
            color: "primary.main",
            margin: 1,
            borderRadius: 3,
            border: "3px grey",
            width: 400,
            height: 400,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            textAlign: "center",
          }}
        >
          <div className="w-64 text-white flex flex-col">
            <img
              src={mode ? innLogoDark : innLogoLight}
              alt="Innoveria logo"
              style={{
                width: "310px",
                height: "auto",
                marginBottom: 20,
              }}
            />
            <Typography className="font-normal text-3xl">Log in:</Typography>
            <TextField
              sx={{
                "& .MuiOutlinedInput-root": {
                  color: "primary.main", // input text color
                  "& fieldset": {
                    borderColor: "primary.main", // default border
                  },
                  "&:hover fieldset": {
                    borderColor: "primary.main", // hover border
                  },
                  "&.Mui-focused fieldset": {
                    borderColor: "primary.main", // focused border
                  },
                },
                "& .MuiInputLabel-root": {
                  color: "primary.main", // default label
                },
                "& .MuiInputLabel-root.Mui-focused": {
                  color: "primary.main", // focused label
                },
              }}
              id="username"
              label="Username"
              variant="outlined"
            />
            <br />
            <TextField
              sx={{
                "& .MuiOutlinedInput-root": {
                  color: "primary.main", // input text color
                  "& fieldset": {
                    borderColor: "primary.main", // default border
                  },
                  "&:hover fieldset": {
                    borderColor: "primary.main", // hover border
                  },
                  "&.Mui-focused fieldset": {
                    borderColor: "primary.main", // focused border
                  },
                },
                "& .MuiInputLabel-root": {
                  color: "primary.main", // default label
                },
                "& .MuiInputLabel-root.Mui-focused": {
                  color: "primary.main", // focused label
                },
              }}
              id="password"
              label="Password"
              type="password"
              variant="outlined"
            />
            <br />
            <Button
              href={"/"}
              variant="outlined"
              sx={{
                backgroundColor: green[500],
                color: "white",
                "&:hover": {
                  backgroundColor: green[800],
                },
                borderRadius: 2,
              }}
            >
              Log in
            </Button>
            <br />
            <Link href={"/"} fontSize={12} color={"primary"}>
              New user? Make an account here!
            </Link>
          </div>
        </Box>
      </Box>
    </div>
  );
}
