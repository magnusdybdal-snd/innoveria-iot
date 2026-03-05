import { JokeViewer } from "@entities/joke";
import { PageContent } from "@shared/ui/PageContent";
import { SubPageHeader } from "@shared/ui/SubPageHeader";
import { Menu } from "@widgets/menu";

/**
 * Home page that displays a random Chuck Norris joke alongside the sub-page header.
 * @returns The rendered Home page
 */
export default function Home() {
  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader />
        <JokeViewer />
      </PageContent>
    </div>
  );
}
