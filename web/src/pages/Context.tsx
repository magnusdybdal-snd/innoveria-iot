import { useEffect, useState } from "react";

import {
  BucketIntervalField,
  ContextParamsDisplay,
  ContextResultDisplay,
  DateTimeField,
  getContextData,
  getRules,
  LabeledSelect,
  toMinutes,
  type BucketUnit,
  type ContextDataResponse,
} from "@entities/context";
import { getSensors } from "@entities/sensor";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { CustomButton } from "@shared/ui/Button";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

// TODO: Hardcoded until auth context provides the active company
const companyOptions = [
  { id: "a0000000-0000-0000-0000-000000000001", name: "Innoveria" },
];

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

  const [from, setFrom] = useState(() =>
    new Date(Date.now() - 60 * 60 * 1000).toISOString(),
  );
  const [toNow, setToNow] = useState(true);
  const [to, setTo] = useState(() => new Date().toISOString());
  const [bucketValue, setBucketValue] = useState("1");
  const [bucketUnit, setBucketUnit] = useState<BucketUnit>("hours");
  const [result, setResult] = useState<ContextDataResponse[] | null>(null);
  const [fetchError, setFetchError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
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

  const resolvedTo = toNow ? new Date().toISOString() : to;
  const params = {
    companyId,
    deviceEui,
    ruleId,
    from,
    to: resolvedTo,
    bucketMinutes,
  };

  const handleFetch = () => {
    const errors: { deviceEui?: string; ruleId?: string } = {};
    if (!deviceEui) errors.deviceEui = "Device EUI is required";
    if (!ruleId) errors.ruleId = "Aggregation Rule is required";
    if (Object.keys(errors).length > 0) {
      setFieldErrors(errors);
      return;
    }
    setFieldErrors({});
    setIsLoading(true);
    setFetchError(null);
    getContextData(
      companyId,
      [deviceEui],
      ruleId,
      from,
      resolvedTo,
      bucketMinutes,
    )
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
        <Box sx={{ display: "flex", gap: 4, mt: 2, alignItems: "flex-start" }}>
          {/* Left: filters */}
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              gap: 2,
              width: 640,
              flexShrink: 0,
              border: "2px dashed",
              borderColor: "primary.main",
              borderRadius: 1,
              p: 2,
            }}
          >
            <Typography variant="body2" sx={{ fontWeight: 600 }}>
              Filters
            </Typography>
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
              <DateTimeField label="From" value={from} onChange={setFrom} />
              <DateTimeField
                label="To"
                value={to}
                onChange={setTo}
                useNow={toNow}
                onUseNowChange={setToNow}
              />
            </Box>
            <BucketIntervalField
              value={bucketValue}
              onValueChange={setBucketValue}
              unit={bucketUnit}
              onUnitChange={setBucketUnit}
            />
            <PageDivider />
            <ContextParamsDisplay params={params} />
            <CustomButton onClick={handleFetch}>
              {isLoading ? "Fetching..." : "Fetch context data"}
            </CustomButton>
          </Box>

          {/* Right: results */}
          <Box
            sx={{ flex: 1, display: "flex", flexDirection: "column", gap: 2 }}
          >
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
        </Box>
      </PageContent>
    </div>
  );
}
