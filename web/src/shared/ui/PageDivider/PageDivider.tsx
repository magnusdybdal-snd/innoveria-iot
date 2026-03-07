import Divider from "@mui/material/Divider";

/**
 * PageDivider
 * @returns Styled horizontal ruler used to visually separate page sections.
 */
export function PageDivider() {
  return (
    <Divider
      sx={{
        backgroundColor: "primary.main",
        marginTop: 2,
        marginBottom: 5,
      }}
    />
  );
}
