import ErrorIcon from "@mui/icons-material/Error";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import { yellow } from "@mui/material/colors";

type NotFoundProps = {
  page: string;
  action?: string;
  extra?: string;
};

/**
 * NotFoundCard
 * @param NotFoundProps - Component props
 * @param NotFoundProps.page - Name of page not found in
 * @param NotFoundProps.action - Action that is not, default is 'found'
 * @param NotFoundProps.extra - Extra info at the end of the message
 * @returns card with message that no object has been found. Distinguishes between error or none registered
 */
export function NotFoundCard({ page, action = "found", extra }: NotFoundProps) {
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
          {action == "registered" && (
            <ErrorIcon sx={{ margin: 1, color: yellow[500] }} />
          )}
          <p>
            No {page} {action}
            {extra?.length ? extra : null}
          </p>
        </Box>
      </Card>
    </>
  );
}
