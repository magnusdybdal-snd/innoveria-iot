import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";
import { JokeViewer } from "@/API/getJoke";
import { ApiViewer } from "@/API/getData/";

export default function Home() {
  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader />
        <JokeViewer />
        <ApiViewer />
      </PageContent>
    </div>
  );
}
