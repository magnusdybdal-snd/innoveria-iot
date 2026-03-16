import ErrorIcon from "@mui/icons-material/Error";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import { yellow } from "@mui/material/colors";

type NotFoundProps = {
  page: string;
};

/**
 * NotFoundCard
 * @param NotFoundProps - Component props
 * @param NotFoundProps.children - Name of page not found in
 * @param NotFoundProps.page
 * @returns card with message that no object has been found
 */
export function NotFoundCard({ page }: NotFoundProps) {
  return (
    <>
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
            padding: 2,
          }}
        >
          <ErrorIcon sx={{ margin: 1, color: yellow[500] }} />
          <p>No {page} found</p>
        </Box>
      </Card>
    </>
  );
}
