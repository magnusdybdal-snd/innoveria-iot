import Typography from "@mui/material/Typography";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { Menu } from "@widgets/menu";

/**
 * Home page that displays a random Chuck Norris joke alongside the sub-page header.
 * @returns The rendered Home page
 */
export default function Home() {
  return (
    <Menu>
      <div className="flex h-screen">
        <PageContent>
          <SubPageHeader />
          <Typography variant="h5" sx={{ mt: 2 }}>
            Home Page content placeholder
          </Typography>
        </PageContent>
      </div>
    </Menu>
  );
}
