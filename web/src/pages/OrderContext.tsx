import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

/**
 * OrderContext page for displaying ERP order data enriched with sensor context.
 * @returns The rendered OrderContext page
 */
export default function OrderContext() {
  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Order Context" />
        <PageDivider />
      </PageContent>
    </div>
  );
}
