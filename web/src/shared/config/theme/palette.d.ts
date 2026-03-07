// TypeScript declaration file
// Should only contain type information - no runtime code!

import "@mui/material/styles";

declare module "@mui/material/styles" {
  // what useTheme() returns - after theme is built
  interface Palette {
    status: {
      online: string;
      warning: string;
      offline: string;
      unknown: string;
    };
  }
  // Accepted input in createTheme({palette{ ... }})
  interface PaletteOptions {
    status?: {
      online?: string;
      warning?: string;
      offline?: string;
      unknown?: string;
    };
  }
}
