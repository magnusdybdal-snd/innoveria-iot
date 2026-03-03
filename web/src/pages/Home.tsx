import { JokeViewer } from "@/API/getJoke";
import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";

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
