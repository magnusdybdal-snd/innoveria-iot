import SettingsIcon from "@mui/icons-material/Settings";
import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";
import Typography from "@mui/material/Typography";

import {
  BucketBarChart,
  BucketLineChart,
  type ContextDataResponse,
} from "@entities/context";
import { CHART_TYPE, type ChartType } from "@widgets/graphWidget/model/types";

interface GraphWidgetChartProps {
  data: ContextDataResponse[] | null;
  chartType: ChartType;
  isLoading: boolean;
  fetchError: string | null;
  isConfigured: boolean;
}

/**
 * Renders the chart body of a GraphWidget, handling all intermediate states
 * (unconfigured, loading, error, empty) before delegating to the appropriate chart.
 * @param props - Component props
 * @param props.data - Fetched context data responses, or null if not yet loaded
 * @param props.chartType - Which chart type to render when data is available
 * @param props.isLoading - Whether a fetch is currently in progress
 * @param props.fetchError - Error message from the last failed fetch, or null
 * @param props.isConfigured - Whether the widget has a device and rule selected
 * @returns The appropriate chart, state indicator, or empty state element
 */
export function GraphWidgetChart({
  data,
  chartType,
  isLoading,
  fetchError,
  isConfigured,
}: GraphWidgetChartProps) {
  if (!isConfigured) {
    return (
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          gap: 1,
          py: 6,
          color: "text.secondary",
        }}
      >
        <SettingsIcon sx={{ fontSize: 36, opacity: 0.4 }} />
        <Typography variant="body2" sx={{ opacity: 0.6 }}>
          No data configured
        </Typography>
        <Typography variant="caption" sx={{ opacity: 0.4 }}>
          Click the settings icon to get started
        </Typography>
      </Box>
    );
  }

  if (isLoading) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          py: 6,
        }}
      >
        <CircularProgress size={32} />
      </Box>
    );
  }

  if (fetchError) {
    return (
      <Typography variant="body2" color="error" sx={{ py: 2 }}>
        {fetchError}
      </Typography>
    );
  }

  if (!data || data.length === 0 || data[0].buckets.length === 0) {
    return (
      <Typography variant="body2" sx={{ py: 2, opacity: 0.6 }}>
        No data returned.
      </Typography>
    );
  }

  const buckets = data[0].buckets;
  const unit = data[0].unit;

  switch (chartType) {
    case CHART_TYPE.bar:
      return <BucketBarChart buckets={buckets} unit={unit} />;
    case CHART_TYPE.line:
      return <BucketLineChart buckets={buckets} unit={unit} />;
    case CHART_TYPE.lineFilled:
      return <BucketLineChart buckets={buckets} unit={unit} area />;
    default: {
      // Will throw a compile-time error if a new chart type is added to CHART_TYPE but not handled here.
      const _exhaustiveCheck: never = chartType;
      return _exhaustiveCheck;
    }
  }
}
