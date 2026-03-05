import { JokeViewer } from "@/entities/joke";
import { PageContent } from "@/shared/ui/PageContent";
import { SubPageHeader } from "@/shared/ui/SubPageHeader";
import { Menu } from "@/widgets/menu";

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
