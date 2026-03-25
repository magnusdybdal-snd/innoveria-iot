import { useState } from "react";

import {
  BUCKET_UNIT_OPTIONS,
  ContextParamsDisplay,
  toMinutes,
  type BucketUnit,
} from "@entities/context";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { DropDownSelect } from "@shared/ui/DropDownSelect";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

const now = new Date();
const minus1h = new Date(now.getTime() - 60 * 60 * 1000);
const minus6h = new Date(now.getTime() - 6 * 60 * 60 * 1000);
const minus24h = new Date(now.getTime() - 24 * 60 * 60 * 1000);
const minus7d = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);
const minus30d = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000);

// Hardcoded until auth context provides the active company
const companyOptions = [
  { id: "a0000000-0000-0000-0000-000000000001", name: "Innoveria" },
];

// TODO: fetch device EUIs from the device/sensor API (same pattern as Sensors page)
const deviceEuiOptions = [
  { id: "0000000000000001", name: "0000000000000001" },
  { id: "0000000000000002", name: "0000000000000002" },
];

// TODO: fetch aggregation rules from the context-service rules API
const ruleOptions = [
  { id: "rule-uuid-1", name: "Temperature average" },
  { id: "rule-uuid-2", name: "Humidity average" },
];

const fromOptions = [
  { id: minus1h.toISOString(), name: "1 hour ago" },
  { id: minus6h.toISOString(), name: "6 hours ago" },
  { id: minus24h.toISOString(), name: "24 hours ago" },
  { id: minus7d.toISOString(), name: "7 days ago" },
  { id: minus30d.toISOString(), name: "30 days ago" },
];

const toOptions = [{ id: now.toISOString(), name: "Now" }];

const textFieldSx = {
  "& .MuiOutlinedInput-root": {
    color: "primary.main",
    "& .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&:hover .MuiOutlinedInput-notchedOutline": { borderColor: "primary.main" },
    "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
      borderColor: "primary.main",
    },
  },
};

/**
 * Context page component for displaying and managing context parameters.
 * @returns The rendered Context page
 */
export default function Context() {
  const [companyId, setCompanyId] = useState(companyOptions[0].id);
  const [deviceEui, setDeviceEui] = useState(deviceEuiOptions[0].id);
  const [ruleId, setRuleId] = useState(ruleOptions[0].id);
  const [from, setFrom] = useState(fromOptions[0].id);
  const [to, setTo] = useState(toOptions[0].id);
  const [bucketValue, setBucketValue] = useState("1");
  const [bucketUnit, setBucketUnit] = useState<BucketUnit>("hours");

  const parsedBucketValue = parseInt(bucketValue, 10);
  const bucketMinutes =
    !isNaN(parsedBucketValue) && parsedBucketValue > 0
      ? toMinutes(parsedBucketValue, bucketUnit)
      : undefined;

  const params = {
    companyId,
    deviceEui,
    ruleId,
    from,
    to,
    bucketMinutes,
  };

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Context" />
        <PageDivider />
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: 2,
            maxWidth: 600,
            mt: 2,
          }}
        >
          <Box>
            <Typography variant="body2" sx={{ mb: 0.5 }}>
              Company
            </Typography>
            <DropDownSelect
              options={companyOptions}
              value={companyId}
              onChange={setCompanyId}
            />
          </Box>

          <Box>
            <Typography variant="body2" sx={{ mb: 0.5 }}>
              Device EUI
            </Typography>
            <DropDownSelect
              options={deviceEuiOptions}
              value={deviceEui}
              onChange={setDeviceEui}
            />
          </Box>

          <Box>
            <Typography variant="body2" sx={{ mb: 0.5 }}>
              Aggregation Rule
            </Typography>
            <DropDownSelect
              options={ruleOptions}
              value={ruleId}
              onChange={setRuleId}
            />
          </Box>

          <Box>
            <Typography variant="body2" sx={{ mb: 0.5 }}>
              From
            </Typography>
            <DropDownSelect
              options={fromOptions}
              value={from}
              onChange={setFrom}
            />
          </Box>

          <Box>
            <Typography variant="body2" sx={{ mb: 0.5 }}>
              To
            </Typography>
            <DropDownSelect options={toOptions} value={to} onChange={setTo} />
          </Box>

          <Box>
            <Typography variant="body2" sx={{ mb: 0.5 }}>
              Time interval
            </Typography>
            <Box sx={{ display: "flex", gap: 1 }}>
              <TextField
                type="number"
                value={bucketValue}
                onChange={(e) => setBucketValue(e.target.value)}
                slotProps={{ htmlInput: { min: 1 } }}
                sx={{ ...textFieldSx, width: 100 }}
              />
              <Box sx={{ flex: 1 }}>
                <DropDownSelect
                  options={BUCKET_UNIT_OPTIONS}
                  value={bucketUnit}
                  onChange={(v) => setBucketUnit(v as BucketUnit)}
                />
              </Box>
            </Box>
          </Box>

          <PageDivider />

          <ContextParamsDisplay params={params} />
        </Box>
      </PageContent>
    </div>
  );
}
