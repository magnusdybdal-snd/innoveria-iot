import { useEffect, useState } from "react";

import {
  CompanyInfo,
  getCompanies,
  postCompany,
  sortCompanies,
  type CompanyApiResponse,
  type CompanySortKey,
  type SortDirection,
} from "@entities/company";
import { AddCompany } from "@features/addCompany";
import { formatTimestamp } from "@shared/lib";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { NotFoundCard } from "@shared/ui/NotFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import { PageDivider } from "@shared/ui/PageDivider";
import {
  AppSnackbar,
  SNACKBAR_SEVERITY,
  useSnackbar,
} from "@shared/ui/snackbar";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

import { CustomButton } from "@/shared/ui/Button";

// Column labels rendered by CategoryHeader; order determines grid layout
const companyDetails: string[] = [
  "Name",
  "Address",
  "Created at",
  "Updated at",
];

// EUI is display-only; excluded because the ChirpStack identifier is not a meaningful value to sort by.
const sortableColumns: CompanySortKey[] = [
  "Name",
  "Address",
  "Created at",
  "Updated at",
];

const addCompanyDetails: string[] = ["Name", "Address"];

/**
 * Full-page view listing all companies regisered on the site.
 *
 * Fetches data from the database and manages
 * column sort state. Delegates row rendering to CompanyRow/CompanyInfo.
 * @returns The rendered Companies page
 */
export default function Companies() {
  const [companies, setCompanies] = useState<CompanyApiResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [addError, setAddError] = useState<string | null>(null);

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

  const fetchCompanies = () => {
    getCompanies().then((data) => {
      setCompanies(data);
      setIsLoading(false);
    });
  };

  useEffect(() => {
    fetchCompanies();
  }, []);
  const [openAdd, setOpenAdd] = useState(false);
  const [sortConfig, setSortConfig] = useState<{
    key: CompanySortKey | null;
    direction: SortDirection;
  }>({
    key: null,
    direction: "asc",
  });

  // Handler for opening and closing add company pop-up
  const handleClickOpenAdd = () => {
    setOpenAdd(true);
  };
  const handleCloseAdd = () => {
    setOpenAdd(false);
    setAddError(null);
  };
  const addAdminUser = () => {};
  const handleAddCompany = (companyData: { name: string; address: string }) => {
    setAddError(null);
    postCompany({
      name: companyData.name,
      address: companyData.address,
    })
      .then(() => {
        fetchCompanies();
        setOpenAdd(false);
        show("Company added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        setAddError(
          "Failed to add company. The name may already be registered.", // TODO: throw non-hardcoded error messages - based on actual error
        );
        show("Failed to add company", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const addButton = (
    <CustomButton onClick={handleClickOpenAdd}>Add company</CustomButton>
  );

  /**
   * Updates sort state when a column header is clicked.
   * @param column - The column label passed up from CategoryHeader's onSort callback
   */
  function handleSort(column: string) {
    const col = column as CompanySortKey;
    setSortConfig((prev) => {
      if (prev.key !== col) return { key: col, direction: "asc" };
      if (prev.direction === "asc") return { key: col, direction: "desc" };
      return { key: null, direction: "asc" }; // third click resets to initial sorting (unsorted?) //TODO check if default is unsorted or sorted when connected to API.
    });
  }

  // Derive sorted list on every render; sortCompanies returns a new array and does not mutate companies
  const sorted = sortCompanies(companies, sortConfig.key, sortConfig.direction);

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Companies" action={addButton} />
        <PageDivider />
        <CategoryHeader
          categories={companyDetails}
          columns={companyDetails.length + 1}
          sortableColumns={sortableColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {isLoading && <p>Loading...</p>}{" "}
          {/*TODO: make a better looking loading indicator */}
          {sorted.map((company) => (
            <DeviceRow key={company.companyId}>
              <CompanyInfo
                name={company.name}
                address={company.address}
                created_at={formatTimestamp(company.createdAt)}
                updated_at={formatTimestamp(company.updatedAt)}
                addUser={addAdminUser}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && sorted.length === 0 && <NotFoundCard page="companies" />}
      </PageContent>
      <AddCompany
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addCompanyDetails}
        onAdd={handleAddCompany}
        submitError={addError}
      />
      <AppSnackbar
        open={snackbar?.open ?? false}
        message={snackbar?.message ?? ""}
        severity={snackbar?.severity}
        onClose={hide}
      />
    </div>
  );
}
