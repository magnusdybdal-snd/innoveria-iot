import { useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

const fieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& fieldset": { borderColor: "primary.main" },
    "&:hover fieldset": { borderColor: "primary.main" },
    "&.Mui-focused fieldset": { borderColor: "primary.main" },
  },
  "& .MuiInputBase-input": {
    fontSize: "0.72rem",
    fontFamily: "monospace",
    lineHeight: 1.2,
    whiteSpace: "nowrap",
    overflowX: "auto",
    textOverflow: "clip",
    scrollbarWidth: "thin",
  },
};

type CompanyERPTokenDialogProps = {
  open: boolean;
  companyName: string;
  token: string | null;
  isLoading: boolean;
  error?: string | null;
  onClose: () => void;
  onGenerate: () => void;
  onCopy: () => void;
};

/**
 * Dialog for generating, viewing, and copying a company's ERP agent token.
 * @param props - Component props
 * @param props.open - Whether the dialog is visible
 * @param props.companyName - Name shown in the dialog title
 * @param props.token - Current generated token value
 * @param props.isLoading - Whether token generation request is in progress
 * @param props.error - Optional generation error to display
 * @param props.onClose - Callback for closing the dialog
 * @param props.onGenerate - Callback for generating or regenerating a token
 * @param props.onCopy - Callback for copying token value
 * @returns Styled ERP token dialog
 */
export function CompanyERPTokenDialog({
  open,
  companyName,
  token,
  isLoading,
  error,
  onClose,
  onGenerate,
  onCopy,
}: CompanyERPTokenDialogProps) {
  const [isTokenVisible, setIsTokenVisible] = useState(false);

  const displayedToken =
    token && !isTokenVisible
      ? "Token hidden. Click Show token or use Copy token."
      : (token ?? "");

  return (
    <Dialog
      open={open}
      onClose={onClose}
      maxWidth="sm"
      fullWidth
      aria-labelledby="company-erp-token-dialog-title"
      sx={{
        "& .MuiPaper-root": {
          backgroundColor: "primary.dark",
          color: "primary.contrastText",
        },
      }}
    >
      <DialogTitle
        id="company-erp-token-dialog-title"
        sx={{ color: "primary.main" }}
      >
        ERP token for {companyName}
      </DialogTitle>
      <DialogContent>
        <Typography variant="caption" color="primary.main" mb={0.5}>
          ERP agent token
        </Typography>
        <TextField
          sx={fieldSx}
          fullWidth
          size="small"
          value={displayedToken}
          placeholder={
            isLoading ? "Generating token..." : "No token generated yet"
          }
          slotProps={{
            input: {
              readOnly: true,
            },
          }}
        />
        <Typography
          variant="caption"
          color="primary.main"
          mt={1}
          display="block"
        >
          Long tokens can be scrolled horizontally. Store securely after
          copying.
        </Typography>
        {error && (
          <Typography color="error" mt={1}>
            {error}
          </Typography>
        )}
      </DialogContent>
      <DialogActions
        sx={{
          px: 3,
          pb: 2,
          pt: 1,
          justifyContent: "space-between",
          alignItems: "center",
          gap: 1,
          flexWrap: "wrap",
        }}
      >
        <Button
          variant="outlined"
          sx={{
            color: "primary.main",
            borderColor: "primary.main",
            "&:hover": {
              borderColor: "primary.main",
              backgroundColor: "action.hover",
            },
          }}
          onClick={() => setIsTokenVisible((prev) => !prev)}
          disabled={!token || isLoading}
        >
          {isTokenVisible ? "Hide token" : "Show token"}
        </Button>
        <Box sx={{ display: "flex", gap: 1, flexWrap: "wrap" }}>
          <Button
            sx={{
              backgroundColor: "primary.main",
              color: "primary.contrastText",
            }}
            onClick={() => {
              setIsTokenVisible(false);
              onGenerate();
            }}
            disabled={isLoading}
          >
            {token ? "Regenerate" : "Generate"}
          </Button>
          <Button
            variant="outlined"
            sx={{
              color: "primary.main",
              borderColor: "primary.main",
              "&:hover": {
                borderColor: "primary.main",
                backgroundColor: "action.hover",
              },
            }}
            onClick={onCopy}
            disabled={!token || isLoading}
          >
            Copy token
          </Button>
          <Button
            variant="outlined"
            sx={{
              color: "primary.main",
              borderColor: "primary.main",
              "&:hover": {
                borderColor: "primary.main",
                backgroundColor: "action.hover",
              },
            }}
            onClick={onClose}
          >
            Close
          </Button>
        </Box>
      </DialogActions>
    </Dialog>
  );
}
