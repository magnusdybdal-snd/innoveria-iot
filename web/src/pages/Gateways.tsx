import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu";
import Divider from "@mui/material/Divider";
import { PageContent } from "@/components/pageContent";

export default function Gateways() {
  return (
    <div className="flex h-screen">
      <Menu />
      <PageContent>
        <SubPageHeader title="Gateways" />
        <Divider sx={{ backgroundColor: "primary.main" }} />
      </PageContent>
    </div>
  );
}
