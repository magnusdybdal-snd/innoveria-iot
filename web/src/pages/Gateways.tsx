import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu";
import { PageDivider } from "@/components/pageDivider";
import { PageContent } from "@/components/pageContent";

export default function Gateways() {
  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader title="Gateways" />
        <PageDivider />
      </PageContent>
    </div>
  );
}
