import Button from "@mui/material/Button";

interface OutlinedButtonProps {
  onClick: (e: React.MouseEvent) => void;
  children: React.ReactNode;
}

/**
 * A reusable small outlined button wrapping MUI Button.
 * @param props - Component props.
 * @param props.onClick - Called when the button is clicked.
 * @param props.children - Content displayed inside the button.
 * @returns The rendered button element.
 */
export function OutlinedButton({ onClick, children }: OutlinedButtonProps) {
  return (
    <Button
      size="small"
      variant="outlined"
      onClick={onClick}
      sx={{
        color: "primary.main",
        borderColor: "primary.main",
        justifySelf: "start",
        alignSelf: "center",
      }}
    >
      {children}
    </Button>
  );
}
