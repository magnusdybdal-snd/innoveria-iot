import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu";
import { PageDivider } from "@/components/pageDivider";
import { PageContent } from "@/components/pageContent";
import { CategoryHeader } from "@/components/CategoryHeader";
import { GatewayInfo } from "@/components/gatewayInfo";

{
  /* Categoriesdisplayed in header of gateways */
}
const gatewayDetails: string[] = ["Status", "Name", "EUI", "Last seen"];

export default function Gateways() {
  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader title="Gateways" />
        <PageDivider />
        <CategoryHeader
          categories={gatewayDetails}
          columns={gatewayDetails.length}
        >
          <GatewayInfo
            name="Gateway 1"
            euid="A1B2C3D4E5F6"
            lastSeen="2 min ago"
          />
        </CategoryHeader>
      </PageContent>
    </div>
  );
}
