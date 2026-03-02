import { JokeViewer } from "@/API/getJoke";
import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";

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
