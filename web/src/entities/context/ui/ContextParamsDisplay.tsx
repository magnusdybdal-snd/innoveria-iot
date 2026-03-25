import type { ContextQueryParams } from "@entities/context/model/contextSchema";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

interface ContextParamsDisplayProps {
  params: ContextQueryParams;
}

/**
 * Displays the current context query parameters as formatted JSON.
 * @param props - Component props
 * @param props.params - The context query parameters to display
 * @returns The rendered parameters display block
 */
export function ContextParamsDisplay({ params }: ContextParamsDisplayProps) {
  return (
    <Box>
      <Typography variant="body2" sx={{ mb: 1 }}>
        Request parameters
      </Typography>
      <Box
        component="pre"
        sx={{
          border: "1px solid",
          borderColor: "primary.main",
          borderRadius: 1,
          p: 2,
          fontSize: "0.75rem",
          overflow: "auto",
          whiteSpace: "pre-wrap",
          wordBreak: "break-all",
        }}
      >
        {JSON.stringify(params, null, 2)}
      </Box>
    </Box>
  );
}
