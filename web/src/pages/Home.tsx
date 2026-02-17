import { AddBox } from "@/components/addCompBox";
import { PageContent } from "@/components/pageContent";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";

export default function Home() {
  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader />
        <AddBox />
      </PageContent>
    </div>
  );
}
