import { useEffect, useRef, useState } from "react";

import { CustomButton } from "@/shared/ui/Button";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import {
  CompanyInfo,
  postCompany,
  postCompanyERPAgentToken,
  sortCompanies,
  useCompanies,
  type CompanySortKey,
  type SortDirection,
} from "@entities/company";
import { AddCompany } from "@features/addCompany";
import { CompanyERPTokenDialog } from "@features/companyERPToken";
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
 * column sort state. Delegates row rendering to CompanyRow/UserInfo.
 * @returns The rendered Companies page
 */
export default function Companies() {
  const { companies, isLoading, error, refetch } = useCompanies();
  const [addError, setAddError] = useState<string | null>(null);
  const [erpTokenError, setERPTokenError] = useState<string | null>(null);
  const [erpTokenLoading, setERPTokenLoading] = useState(false);
  const [erpToken, setERPToken] = useState<string | null>(null);
  const erpTokenCancelledRef = useRef(false);
  const erpTokenInFlightRef = useRef(false);
  const erpTokenRequestIdRef = useRef(0);
  const [selectedCompany, setSelectedCompany] = useState<{
    companyId: string;
    name: string;
  } | null>(null);

  useEffect(() => {
    return () => {
      erpTokenCancelledRef.current = true;
      erpTokenRequestIdRef.current += 1;
      erpTokenInFlightRef.current = false;
    };
  }, []);

  // State for controlling success snackbar
  const { show, hide, snackbar } = useSnackbar();

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
  const closeERPTokenDialog = () => {
    erpTokenCancelledRef.current = true;
    erpTokenRequestIdRef.current += 1;
    erpTokenInFlightRef.current = false;
    setSelectedCompany(null);
    setERPToken(null);
    setERPTokenError(null);
    setERPTokenLoading(false);
  };

  const generateERPToken = (companyId: string) => {
    if (erpTokenInFlightRef.current) {
      return;
    }

    erpTokenInFlightRef.current = true;
    erpTokenCancelledRef.current = false;
    const requestId = ++erpTokenRequestIdRef.current;
    setERPTokenLoading(true);
    setERPTokenError(null);

    postCompanyERPAgentToken(companyId)
      .then((token) => {
        if (
          erpTokenCancelledRef.current ||
          requestId !== erpTokenRequestIdRef.current
        ) {
          return;
        }
        setERPToken(token);
        show("ERP token generated", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        if (
          erpTokenCancelledRef.current ||
          requestId !== erpTokenRequestIdRef.current
        ) {
          return;
        }
        setERPTokenError("Failed to generate ERP token.");
        show("Failed to generate ERP token", SNACKBAR_SEVERITY.ERROR);
      })
      .finally(() => {
        erpTokenInFlightRef.current = false;
        if (
          erpTokenCancelledRef.current ||
          requestId !== erpTokenRequestIdRef.current
        ) {
          return;
        }
        setERPTokenLoading(false);
      });
  };

  const handleOpenERPTokenDialog = (companyId: string, companyName: string) => {
    if (!companyId) {
      show("Company ID is missing", SNACKBAR_SEVERITY.ERROR);
      return;
    }

    setSelectedCompany({ companyId, name: companyName });
    setERPToken(null);
    setERPTokenError(null);
    setERPTokenLoading(false);
  };

  const handleCopyERPToken = () => {
    if (!erpToken) {
      return;
    }

    if (window.isSecureContext && navigator.clipboard) {
      navigator.clipboard
        .writeText(erpToken)
        .then(() => {
          show("ERP token copied", SNACKBAR_SEVERITY.SUCCESS);
        })
        .catch(() => {
          show("Failed to copy ERP token", SNACKBAR_SEVERITY.ERROR);
        });
      return;
    }

    const textArea = document.createElement("textarea");
    textArea.value = erpToken;
    textArea.style.position = "fixed";
    textArea.style.opacity = "0";
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();

    try {
      const copied = document.execCommand("copy");
      if (copied) {
        show("ERP token copied", SNACKBAR_SEVERITY.SUCCESS);
      } else {
        show("Failed to copy ERP token", SNACKBAR_SEVERITY.ERROR);
      }
    } catch {
      show("Failed to copy ERP token", SNACKBAR_SEVERITY.ERROR);
    } finally {
      document.body.removeChild(textArea);
    }
  };

  const handleAddCompany = (companyData: { name: string; address: string }) => {
    setAddError(null);
    postCompany({
      name: companyData.name,
      address: companyData.address,
    })
      .then(() => {
        refetch();
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
          {error ? (
            <Box sx={{ p: 2 }}>
              <Typography variant="body2" sx={{ color: "error.main", mb: 1 }}>
                Failed to load companies. {error.message}
              </Typography>
              <Button variant="outlined" size="small" onClick={refetch}>
                Retry
              </Button>
            </Box>
          ) : (
            sorted.map((company) => (
              <DeviceRow key={company.companyId}>
                <CompanyInfo
                  name={company.name}
                  address={company.address}
                  created_at={formatTimestamp(company.createdAt)}
                  updated_at={formatTimestamp(company.updatedAt)}
                  onShowApiToken={() =>
                    handleOpenERPTokenDialog(company.companyId, company.name)
                  }
                />
              </DeviceRow>
            ))
          )}
        </CategoryHeader>
        {!isLoading && !error && sorted.length === 0 && (
          <NotFoundCard page="companies" action="registered" />
        )}
      </PageContent>
      <AddCompany
        open={openAdd}
        onClose={handleCloseAdd}
        addOptions={addCompanyDetails}
        onAdd={handleAddCompany}
        submitError={addError}
      />
      {selectedCompany && (
        <CompanyERPTokenDialog
          open={true}
          companyName={selectedCompany.name}
          token={erpToken}
          isLoading={erpTokenLoading}
          error={erpTokenError}
          onClose={closeERPTokenDialog}
          onGenerate={() => {
            generateERPToken(selectedCompany.companyId);
          }}
          onCopy={handleCopyERPToken}
        />
      )}
      <AppSnackbar
        open={snackbar?.open ?? false}
        message={snackbar?.message ?? ""}
        severity={snackbar?.severity}
        onClose={hide}
      />
    </div>
  );
}
