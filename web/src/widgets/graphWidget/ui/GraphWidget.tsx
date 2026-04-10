import { useCallback, useEffect, useState } from "react";

import {
  getContextData,
  getRules,
  toMinutes,
  type BucketUnit,
  type ContextDataResponse,
} from "@entities/context";
import { getSensors } from "@entities/sensor";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";
import SettingsIcon from "@mui/icons-material/Settings";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import IconButton from "@mui/material/IconButton";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import Typography from "@mui/material/Typography";
import { toLocalDateTimeString } from "@shared/lib";
import {
  CHART_TYPE,
  type ChartType,
  type GraphWidgetConfig,
} from "@widgets/graphWidget/model/types";
import { GraphWidgetChart } from "@widgets/graphWidget/ui/GraphWidgetChart";
import { GraphWidgetSettings } from "@widgets/graphWidget/ui/GraphWidgetSettings";

// TODO: Replace with auth context when backend auth is wired
const COMPANY_ID = "a0000000-0000-0000-0000-000000000001";

const INITIAL_CONFIG: GraphWidgetConfig = {
  title: "",
  deviceEui: "",
  ruleId: "",
  from: toLocalDateTimeString(new Date(Date.now() - 60 * 60 * 1000)),
  to: toLocalDateTimeString(new Date()),
  useCurrentTime: true,
  bucketValue: "1",
  bucketUnit: "hours",
};

interface GraphWidgetProps {
  defaultConfig?: Partial<GraphWidgetConfig>;
  onDelete?: () => void;
}

/**
 * Self-contained graph widget card. Owns all state for data fetching, settings,
 * and chart type selection. Multiple instances can be rendered independently.
 * @param props - Component props
 * @param props.defaultConfig - Optional partial config to pre-seed the widget
 * @param props.onDelete - Optional callback invoked when the user removes the widget
 * @returns The rendered widget card
 */
export function GraphWidget({ defaultConfig, onDelete }: GraphWidgetProps) {
  const [config, setConfig] = useState<GraphWidgetConfig>({
    ...INITIAL_CONFIG,
    ...defaultConfig,
  });
  const [chartType, setChartType] = useState<ChartType>(CHART_TYPE.bar);
  const [draft, setDraft] = useState<GraphWidgetConfig>(config);
  const [settingsOpen, setSettingsOpen] = useState(false);

  const [data, setData] = useState<ContextDataResponse[] | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [fetchError, setFetchError] = useState<string | null>(null);

  const [deviceOptions, setDeviceOptions] = useState<
    { id: string; name: string }[]
  >([]);
  const [ruleOptions, setRuleOptions] = useState<
    { id: string; name: string }[]
  >([]);
  const [optionsError, setOptionsError] = useState<string | null>(null);

  useEffect(() => {
    Promise.all([getSensors(), getRules(COMPANY_ID)])
      .then(([sensors, rules]) => {
        setDeviceOptions(
          sensors.map((s) => ({ id: s.deviceEui, name: s.name })),
        );
        setRuleOptions(rules.map((r) => ({ id: r.id, name: r.name })));
      })
      .catch((err: unknown) => {
        setOptionsError(
          err instanceof Error ? err.message : "Failed to load options.",
        );
      });
  }, []);

  const fetchData = useCallback((cfg: GraphWidgetConfig) => {
    const parsed = parseInt(cfg.bucketValue, 10);
    const bucketMinutes =
      !isNaN(parsed) && parsed > 0
        ? toMinutes(parsed, cfg.bucketUnit as BucketUnit)
        : undefined;

    setIsLoading(true);
    setFetchError(null);
    getContextData(
      COMPANY_ID,
      [cfg.deviceEui],
      cfg.ruleId,
      new Date(cfg.from).toISOString(),
      new Date(cfg.to).toISOString(),
      bucketMinutes,
    )
      .then(setData)
      .catch((err: unknown) => {
        setData(null);
        setFetchError(err instanceof Error ? err.message : String(err));
      })
      .finally(() => setIsLoading(false));
  }, []);

  const handleSettingsOpen = () => {
    setDraft(config);
    setSettingsOpen(true);
  };

  const handleSettingsCancel = () => {
    setSettingsOpen(false);
  };

  const handleSettingsConfirm = (newConfig: GraphWidgetConfig) => {
    const resolved = newConfig.useCurrentTime
      ? { ...newConfig, to: toLocalDateTimeString(new Date()) }
      : newConfig;
    setConfig(resolved);
    setSettingsOpen(false);
    fetchData(resolved);
  };

  const isConfigured = Boolean(config.deviceEui && config.ruleId);

  return (
    <Card
      sx={{
        backgroundColor: "secondary.light",
        borderRadius: 3,
        p: 2,
        color: "primary.main",
      }}
    >
      {/* Header */}
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          mb: 1,
        }}
      >
        <Typography variant="subtitle2">
          {config.title || "Graph Widget"}
        </Typography>
        <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
          <Select
            size="small"
            value={chartType}
            onChange={(e) => setChartType(e.target.value as ChartType)}
            sx={{ fontSize: "0.75rem" }}
          >
            <MenuItem value={CHART_TYPE.bar}>Bar</MenuItem>
            <MenuItem value={CHART_TYPE.line}>Line</MenuItem>
          </Select>
          <IconButton size="small" onClick={handleSettingsOpen}>
            <SettingsIcon fontSize="small" />
          </IconButton>
          {onDelete && (
            <IconButton size="small" onClick={onDelete}>
              <DeleteOutlineIcon fontSize="small" />
            </IconButton>
          )}
        </Box>
      </Box>

      {/* Chart body */}
      <GraphWidgetChart
        data={data}
        chartType={chartType}
        isLoading={isLoading}
        fetchError={fetchError}
        isConfigured={isConfigured}
      />

      {/* Settings dialog */}
      <GraphWidgetSettings
        open={settingsOpen}
        draft={draft}
        onChange={setDraft}
        onCancel={handleSettingsCancel}
        onConfirm={handleSettingsConfirm}
        deviceOptions={deviceOptions}
        ruleOptions={ruleOptions}
        optionsError={optionsError}
      />
    </Card>
  );
}
