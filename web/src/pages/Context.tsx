import { useEffect, useRef, useState } from "react";

import {
  BucketIntervalField,
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
  const [loadError, setLoadError] = useState<string | null>(null);

  useEffect(() => {
    getSensors()
      .then((sensors) => {
        const options = sensors.map((s) => ({
          id: s.deviceEui,
          name: s.deviceEui,
        }));
        setDeviceEuiOptions(options);
        if (options.length > 0) setDeviceEui(options[0].id);
      })
      .catch((err: unknown) => {
        setLoadError(
          err instanceof Error ? err.message : "Failed to load sensors.",
        );
      });
  }, []);

  useEffect(() => {
    getRules(companyId)
      .then((rules) => {
        const options = rules.map((r) => ({ id: r.id, name: r.name }));
        setRuleOptions(options);
        if (options.length > 0) setRuleId(options[0].id);
      })
      .catch((err: unknown) => {
        setLoadError(
          err instanceof Error
            ? err.message
            : "Failed to load aggregation rules.",
        );
      });
  }, [companyId]);

  const [from, setFrom] = useState(() => {
    const d = new Date(Date.now() - 60 * 60 * 1000);
    return d.toISOString().slice(0, 16);
  });
  const [toNow, setToNow] = useState(true);
  const [to, setTo] = useState(() => new Date().toISOString().slice(0, 16));
  const [bucketValue, setBucketValue] = useState("1");
  const [bucketUnit, setBucketUnit] = useState<BucketUnit>("hours");
  const [result, setResult] = useState<ContextDataResponse[] | null>(null);
  const [fetchError, setFetchError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const hasFetchedOnce = useRef(false);
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

  const fromIso = new Date(from).toISOString();
  const resolvedTo = toNow
    ? new Date().toISOString()
    : new Date(to).toISOString();

  const runFetch = (minutes: number | undefined) => {
    setIsLoading(true);
    setFetchError(null);
    getContextData(companyId, [deviceEui], ruleId, fromIso, resolvedTo, minutes)
      .then((data) => setResult(data))
      .catch((err: unknown) => {
        setResult(null);
        setFetchError(err instanceof Error ? err.message : String(err));
      })
      .finally(() => setIsLoading(false));
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
    hasFetchedOnce.current = true;
    runFetch(bucketMinutes);
  };

  // Re-fetch when bucket interval changes after the first manual fetch.
  // setState calls here are intentional — syncing UI state with an async API response
  // is a valid useEffect use case. Filter deps are intentionally excluded: this effect
  // should only re-run on bucket changes, not every filter change.
  useEffect(() => {
    if (!hasFetchedOnce.current) return;
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setIsLoading(true);

    setFetchError(null);
    getContextData(
      companyId,
      [deviceEui],
      ruleId,
      fromIso,
      resolvedTo,
      bucketMinutes,
    )
      .then((data) => setResult(data))
      .catch((err: unknown) => {
        setResult(null);
        setFetchError(err instanceof Error ? err.message : String(err));
      })
      .finally(() => setIsLoading(false));
  }, [bucketMinutes]); // eslint-disable-line react-hooks/exhaustive-deps

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
            {/*TODO: remove company option, companies should not have the option to select a company. This is there for testing purposes only */}
            <LabeledSelect
              label="Company"
              options={companyOptions}
              value={companyId}
              onChange={setCompanyId}
            />
            {loadError && (
              <Typography variant="body2" color="error">
                {loadError}
              </Typography>
            )}
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
            <CustomButton onClick={handleFetch}>
              {isLoading ? "Fetching..." : "Fetch context data"}
            </CustomButton>
          </Box>

          {/* Right: results */}
          <Box
            sx={{ flex: 1, display: "flex", flexDirection: "column", gap: 2 }}
          >
            <BucketIntervalField
              value={bucketValue}
              onValueChange={setBucketValue}
              unit={bucketUnit}
              onUnitChange={setBucketUnit}
            />
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
