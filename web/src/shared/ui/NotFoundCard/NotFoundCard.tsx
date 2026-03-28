import ErrorIcon from "@mui/icons-material/Error";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import { yellow } from "@mui/material/colors";

type NotFoundProps = {
  page: string;
  isEmpty: boolean;
};

/**
 * NotFoundCard
 * @param NotFoundProps - Component props
 * @param NotFoundProps.page - Name of page not found in
 * @param NotFoundProps.isEmpty - Whether there are no objects at all
 * @returns card with message that no object has been found. Distinguishes between error or none registered
 */
export function NotFoundCard({ page, isEmpty }: NotFoundProps) {
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
          {isEmpty && <ErrorIcon sx={{ margin: 1, color: yellow[500] }} />}
          <p>{isEmpty ? `No ${page} found` : `No ${page} registered`}</p>
        </Box>
      </Card>
    </>
  );
}
