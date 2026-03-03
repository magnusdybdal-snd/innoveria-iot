import { JokeViewer } from "@/API/getJoke";
import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu2 from "@/Menu2.tsx";

export default function Home() {
  return (
    <Menu2>
      <div className="flex h-screen">
        <PageContent>
          <SubPageHeader />
          <JokeViewer />
        </PageContent>
      </div>
    </Menu2>
  );
}
