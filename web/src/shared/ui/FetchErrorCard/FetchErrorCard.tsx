import Box from "@mui/material/Box";
import Card from "@mui/material/Card";

import { CustomButton } from "@shared/ui/Button";

type FetchErrorCardProps = {
  message: string;
  onRetry: () => void;
};

/**
 * Displays a styled error card with a retry button for failed data fetches.
 * @param props - Component props
 * @param props.message - The error message to display
 * @param props.onRetry - Called when the user clicks Retry
 * @returns A card with the error message and a retry button
 */
export function FetchErrorCard({ message, onRetry }: FetchErrorCardProps) {
  return (
    <Card
      sx={{
        marginTop: 2,
        color: "primary.main",
        backgroundColor: "primary.light",
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          justifyContent: "center",
          gap: 2,
          padding: 2,
        }}
      >
        <p>{message}</p>
        <CustomButton onClick={onRetry}>Retry</CustomButton>
      </Box>
    </Card>
  );
}
