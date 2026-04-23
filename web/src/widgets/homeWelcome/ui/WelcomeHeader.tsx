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
 * Displays a live updating clock alongside a contextual greeting and today's date.
 * Intended as the top welcome strip on the Home page.
 * @returns The rendered WelcomeHeader widget
 */
export function WelcomeHeader() {
  const theme = useTheme();
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => setTime(new Date()), 1000);
    return () => clearInterval(timer);
  }, []);

  const greeting = getGreeting(time.getHours());

  const timeStr = time.toLocaleTimeString("no-NO", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false,
  });

  const dateStr = time.toLocaleDateString("no-NO", {
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
    </Box>
  );
}
