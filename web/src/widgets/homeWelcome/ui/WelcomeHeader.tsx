import { useEffect, useState } from "react";

import Box from "@mui/material/Box";
import { useTheme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";

/**
 * Returns a time-of-day greeting based on the current hour.
 * @param hour - Current hour in 24h format (0–23)
 * @returns Greeting string
 */
function getGreeting(hour: number): string {
  if (hour < 12) return "Good morning";
  if (hour < 17) return "Good afternoon";
  return "Good evening";
}

/**
 * Live clock that re-renders every second in isolation.
 * @returns The rendered Clock component
 */
function Clock() {
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => setTime(new Date()), 1000);
    return () => clearInterval(timer);
  }, []);

  const timeStr = time.toLocaleTimeString(undefined, {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false,
  });

  return (
    <Typography
      sx={{
        fontFamily: "monospace",
        fontSize: "3rem",
        fontWeight: 200,
        letterSpacing: "0.04em",
        lineHeight: 1,
        color: "text.primary",
      }}
    >
      {timeStr}
    </Typography>
  );
}

/**
 * Displays a live clock alongside a contextual greeting and today's date.
 * Intended as the top welcome strip on the Home page.
 * @returns The rendered WelcomeHeader widget
 */
export function WelcomeHeader() {
  const theme = useTheme();
  const now = new Date();
  const greeting = getGreeting(now.getHours());
  const dateStr = now.toLocaleDateString("no-NO", {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
  });

  const borderColor = theme.palette.secondary.light;
  const labelColor = theme.palette.primary.main;

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "flex-end",
        justifyContent: "space-between",
        py: 4,
        borderBottom: `1px solid ${borderColor}`,
        mb: 4,
      }}
    >
      <Box>
        <Typography
          variant="overline"
          sx={{
            color: labelColor,
            letterSpacing: 4,
            fontWeight: 700,
            fontSize: "0.65rem",
            display: "block",
            mb: 0.5,
          }}
        >
          Factory pulse
        </Typography>
        <Typography
          variant="h4"
          sx={{ fontWeight: 300, lineHeight: 1.2, mb: 1 }}
        >
          {greeting}
        </Typography>
        <Typography
          variant="body2"
          sx={{ color: "text.secondary", textTransform: "capitalize" }}
        >
          {dateStr}
        </Typography>
      </Box>

      <Clock />
    </Box>
  );
}
