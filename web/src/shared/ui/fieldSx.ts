/**
 * Shared MUI TextField sx prop that overrides the outlined input border and text
 * color to use the theme's primary color in all states (default, hover, focused).
 * Apply to any TextField that should match the dark-background dialog style.
 */
export const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
};
