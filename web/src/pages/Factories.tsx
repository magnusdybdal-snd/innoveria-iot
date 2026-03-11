import { useEffect, useState } from "react";

import { getFactories } from "@entities/factory";

import { CategoryHeader } from "@/shared/ui/CategoryHeader";
import { PageContent } from "@/shared/ui/PageContent";
import { PageDivider } from "@/shared/ui/PageDivider";
import { SubPageHeader } from "@/shared/ui/SubPageHeader";

const factoryDetails: string[] = [
  "ID",
  "Name",
  "Address",
  "Created At",
  "Updated At",
];

/**
 * Full-page view listing all factories regisered on the site.
 * @returns The rendered Factories page
 */
export default function Factories() {
  const [isLoading, setIsLoading] = useState(true);

  const fetchFactories = () => {
    getFactories().then(() => {
      setIsLoading(false);
    });
  };

  useEffect(() => {
    fetchFactories();
  }, []);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Companies" />
        <PageDivider />
        <CategoryHeader
          categories={factoryDetails}
          columns={factoryDetails.length + 1}
          //sortableColumns={sortableColumns}
          //sortConfig={sortConfig}
          //onSort={handleSort}
        >
          {isLoading && <p>Loading...</p>}{" "}
          {/*TODO: make a better looking loading indicator */}
        </CategoryHeader>
      </PageContent>
    </div>
  );
}
