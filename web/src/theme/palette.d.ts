// TypeScript declaration file
// Should only contain type information - no runtime code!

import "@mui/material/styles";

declare module "@mui/metarial/styles" {
  // what useTheme() returns - after theme is built
  interface Palette {
    status: {
      online: string;
      warning: string;
      offline: string;
      unknown: string;
    };
  }
  interface PaletteOptions {
    status?: {
      online?: string;
      warning?: string;
      offline?: string;
      unknown?: string;
    };
  }
}
