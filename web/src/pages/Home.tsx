import { JokeViewer } from "@/API/getJoke";
import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu2 from "@/Menu2.tsx";

export default function Home() {
  return (
    <div className="flex h-screen">
      <Menu2>
        <PageContent>
          <SubPageHeader />
          <JokeViewer />
        </PageContent>
      </Menu2>
    </div>
  );
}
