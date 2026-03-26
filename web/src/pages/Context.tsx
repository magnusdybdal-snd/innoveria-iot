import { useEffect, useState } from "react";

import {
  BucketIntervalField,
  ContextParamsDisplay,
  ContextResultDisplay,
  getContextData,
  getRules,
  LabeledSelect,
  toMinutes,
  type BucketUnit,
  type ContextDataResponse,
} from "@entities/context";
import { getSensors } from "@entities/sensor";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Accordion from "@mui/material/Accordion";
import AccordionDetails from "@mui/material/AccordionDetails";
import AccordionSummary from "@mui/material/AccordionSummary";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { CustomButton } from "@shared/ui/Button";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

const now = new Date();
const minus1h = new Date(now.getTime() - 60 * 60 * 1000);
const minus6h = new Date(now.getTime() - 6 * 60 * 60 * 1000);
const minus24h = new Date(now.getTime() - 24 * 60 * 60 * 1000);
const minus7d = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);
const minus30d = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000);

// TODO: Hardcoded until auth context provides the active company
const companyOptions = [
  { id: "a0000000-0000-0000-0000-000000000001", name: "Innoveria" },
];

const fromOptions = [
  { id: minus1h.toISOString(), name: "1 hour ago" },
  { id: minus6h.toISOString(), name: "6 hours ago" },
  { id: minus24h.toISOString(), name: "24 hours ago" },
  { id: minus7d.toISOString(), name: "7 days ago" },
  { id: minus30d.toISOString(), name: "30 days ago" },
];

const toOptions = [{ id: now.toISOString(), name: "Now" }];

/**
 * Context page component for displaying and managing context parameters.
 * @returns The rendered Context page
 */
export default function Context() {
  const [companyId, setCompanyId] = useState(companyOptions[0].id);
  const [deviceEuiOptions, setDeviceEuiOptions] = useState<
    { id: string; name: string }[]
  >([]);
  const [ruleOptions, setRuleOptions] = useState<
    { id: string; name: string }[]
  >([]);
  const [deviceEui, setDeviceEui] = useState("");
  const [ruleId, setRuleId] = useState("");

  useEffect(() => {
    getSensors().then((sensors) => {
      const options = sensors.map((s) => ({
        id: s.deviceEui,
        name: s.deviceEui,
      }));
      setDeviceEuiOptions(options);
      if (options.length > 0) setDeviceEui(options[0].id);
    });
  }, []);

  useEffect(() => {
    getRules(companyId).then((rules) => {
      const options = rules.map((r) => ({ id: r.id, name: r.name }));
      setRuleOptions(options);
      if (options.length > 0) setRuleId(options[0].id);
    });
  }, [companyId]);

  const [from, setFrom] = useState(fromOptions[0].id);
  const [to, setTo] = useState(toOptions[0].id);
  const [bucketValue, setBucketValue] = useState("1");
  const [bucketUnit, setBucketUnit] = useState<BucketUnit>("hours");
  const [result, setResult] = useState<ContextDataResponse[] | null>(null);
  const [fetchError, setFetchError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [filtersOpen, setFiltersOpen] = useState(true);
  const [fieldErrors, setFieldErrors] = useState<{
    deviceEui?: string;
    ruleId?: string;
  }>({});
  const [selectedField, setSelectedField] =
    useState<keyof ContextDataResponse>("totalValue");

  const parsedBucketValue = parseInt(bucketValue, 10);
  const bucketMinutes =
    !isNaN(parsedBucketValue) && parsedBucketValue > 0
      ? toMinutes(parsedBucketValue, bucketUnit)
      : undefined;

  const params = { companyId, deviceEui, ruleId, from, to, bucketMinutes };

  const handleFetch = () => {
    const errors: { deviceEui?: string; ruleId?: string } = {};
    if (!deviceEui) errors.deviceEui = "Device EUI is required";
    if (!ruleId) errors.ruleId = "Aggregation Rule is required";
    if (Object.keys(errors).length > 0) {
      setFieldErrors(errors);
      setFiltersOpen(true);
      return;
    }
    setFieldErrors({});
    setIsLoading(true);
    setFetchError(null);
    setFiltersOpen(false);
    getContextData(companyId, [deviceEui], ruleId, from, to, bucketMinutes)
      .then((data) => setResult(data))
      .catch((err: unknown) => {
        setResult(null);
        setFetchError(err instanceof Error ? err.message : String(err));
      })
      .finally(() => setIsLoading(false));
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
          <Accordion
            expanded={filtersOpen}
            onChange={(_, expanded) => setFiltersOpen(expanded)}
            disableGutters
            elevation={0}
            sx={{ background: "transparent", "&:before": { display: "none" } }}
          >
            <AccordionSummary expandIcon={<ExpandMoreIcon />} sx={{ px: 0 }}>
              <Typography variant="body2">Filters</Typography>
            </AccordionSummary>
            <AccordionDetails
              sx={{ px: 0, display: "flex", flexDirection: "column", gap: 2 }}
            >
              <LabeledSelect
                label="Company"
                options={companyOptions}
                value={companyId}
                onChange={setCompanyId}
              />
              <LabeledSelect
                label="Device EUI"
                options={deviceEuiOptions}
                value={deviceEui}
                onChange={(v) => {
                  setDeviceEui(v);
                  setFieldErrors((e) => ({ ...e, deviceEui: undefined }));
                }}
                error={fieldErrors.deviceEui}
              />
              <LabeledSelect
                label="Aggregation Rule"
                options={ruleOptions}
                value={ruleId}
                onChange={(v) => {
                  setRuleId(v);
                  setFieldErrors((e) => ({ ...e, ruleId: undefined }));
                }}
                error={fieldErrors.ruleId}
              />
              <Box sx={{ display: "flex", gap: 2 }}>
                <LabeledSelect
                  label="From"
                  options={fromOptions}
                  value={from}
                  onChange={setFrom}
                  flex={1}
                />
                <LabeledSelect
                  label="To"
                  options={toOptions}
                  value={to}
                  onChange={setTo}
                  flex={1}
                />
              </Box>
              <BucketIntervalField
                value={bucketValue}
                onValueChange={setBucketValue}
                unit={bucketUnit}
                onUnitChange={setBucketUnit}
              />
              <PageDivider />
              {/* TODO: Replace with actual context data display once API integration is done */}
              <ContextParamsDisplay params={params} />
            </AccordionDetails>
          </Accordion>
          <CustomButton onClick={handleFetch}>
            {isLoading ? "Fetching..." : "Fetch context data"}
          </CustomButton>
          {fetchError && (
            <Typography variant="body2" color="error">
              {fetchError}
            </Typography>
          )}
          {result !== null && result.length === 0 && (
            <Typography variant="body2">No data returned.</Typography>
          )}
          {result !== null && result.length > 0 && (
            <ContextResultDisplay
              result={result}
              selectedField={selectedField}
              onFieldChange={setSelectedField}
            />
          )}
        </Box>
      </PageContent>
    </div>
  );
}
