import { CategoryHeader } from "@/components/CategoryHeader";
import { GatewayInfo } from "@/components/gatewayInfo";
import { GatewayRow } from "@/components/gatewayRow/gatewayRow";
import { PageContent } from "@/components/pageContent";
import { PageDivider } from "@/components/pageDivider";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu";
import { mockGateways } from "@/mocks/gateways";

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
          columns={gatewayDetails.length + 1}
        >
          {mockGateways.map((gateway) => (
            <GatewayRow key={gateway.id}>
              <GatewayInfo
                name={gateway.name}
                online={gateway.online}
                euid={gateway.euid}
                lastSeen={gateway.lastSeen}
              />
            </GatewayRow>
          ))}
        </CategoryHeader>
      </PageContent>
    </div>
  );
}
