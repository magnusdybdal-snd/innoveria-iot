import Button from "@mui/material/Button";

interface CustomButtonProps {
  onClick: () => void;
  children: React.ReactNode;
}
/**
 * A reusable styled button wrapping MUI Button.
 * @param props - Component props.
 * @param props.onClick - Called when the button is clicked.
 * @param props.children - Content displayed inside the button.
 * @returns The rendered button element.
 */
export function CustomButton({ onClick, children }: CustomButtonProps) {
  return (
    <Button
      variant="outlined"
      sx={{
        backgroundColor: "primary.main",
        color: "primary.dark",
        "&:hover": { backgroundColor: "primary.main" },
        borderRadius: 2,
        textTransform: "none",
        fontSize: 20,
      }}
      onClick={onClick}
    >
      {children}
    </Button>
  );
}
