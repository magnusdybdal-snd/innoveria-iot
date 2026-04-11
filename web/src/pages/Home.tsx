import Typography from "@mui/material/Typography";

import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

/**
 * Home page that displays a placholder text, wrapped in the main Menu layout.
 * This is the default landing page after login.
 * @returns The rendered Home page
 */
export default function Home() {
  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader />
        <Typography variant="h5" sx={{ mt: 2 }}>
          Home Page content placeholder
        </Typography>
      </PageContent>
    </div>
  );
}
