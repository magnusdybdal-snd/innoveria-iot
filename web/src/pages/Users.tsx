import { useState } from "react";

import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";

import {
  CompanyInfo,
  useCompanies,
  type CompanyApiResponse,
} from "@entities/company";
import {
  deleteUser,
  postUser,
  sortUsers,
  UserInfo,
  useUsers,
  type CreateUserRequest,
  type SortDirection,
  type UserSortKey,
} from "@entities/user";
import { AddUser } from "@features/addUser";
import { formatTimestamp } from "@shared/lib";
import { CustomButton } from "@shared/ui/Button";
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

const companyColumns = ["Name", "Address", "Created at", "Updated at"];
const userColumns = ["Name", "Email", "Role", "Created at"];
const sortableUserColumns: UserSortKey[] = [
  "Name",
  "Email",
  "Role",
  "Created at",
];

/**
 * Admin user management page.
 *
 * Default view lists all companies. Clicking a company drills down into
 * that company's users where admins can add or delete users.
 * @returns The rendered Users page
 */
export default function Users() {
  const [selectedCompany, setSelectedCompany] =
    useState<CompanyApiResponse | null>(null);

  const { companies, isLoading: companiesLoading } = useCompanies();
  const {
    users,
    isLoading: usersLoading,
    refetch,
  } = useUsers(selectedCompany?.companyId);

  const [openAdd, setOpenAdd] = useState(false);
  const [addError, setAddError] = useState<string | null>(null);
  const [sortConfig, setSortConfig] = useState<{
    key: UserSortKey | null;
    direction: SortDirection;
  }>({ key: null, direction: "asc" });

  const { show, hide, snackbar } = useSnackbar();

  const handleSelectCompany = (company: CompanyApiResponse) => {
    setSelectedCompany(company);
    setSortConfig({ key: null, direction: "asc" });
  };

  const handleBack = () => {
    setSelectedCompany(null);
  };

  const handleSort = (column: string) => {
    const col = column as UserSortKey;
    setSortConfig((prev) => {
      if (prev.key !== col) return { key: col, direction: "asc" };
      if (prev.direction === "asc") return { key: col, direction: "desc" };
      return { key: null, direction: "asc" };
    });
  };

  const handleAddUser = (userData: CreateUserRequest) => {
    setAddError(null);
    postUser({ ...userData, companyId: selectedCompany!.companyId })
      .then(() => {
        refetch();
        setOpenAdd(false);
        show("User added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        setAddError("Failed to add user. The email may already be registered.");
        show("Failed to add user", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const handleDeleteUser = (userId: string) => {
    deleteUser(userId)
      .then(() => {
        refetch();
        show("User deleted successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch(() => {
        show("Failed to delete user", SNACKBAR_SEVERITY.ERROR);
      });
  };

  const sorted = sortUsers(users, sortConfig.key, sortConfig.direction);

  // — Companies view —
  if (!selectedCompany) {
    return (
      <div className="flex h-screen">
        <PageContent>
          <SubPageHeader title="Users" />
          <PageDivider />
          <CategoryHeader
            categories={companyColumns}
            columns={companyColumns.length}
          >
            {companiesLoading && <p>Loading...</p>}
            {companies.map((company) => (
              <DeviceRow
                key={company.companyId}
                onClick={() => handleSelectCompany(company)}
              >
                <CompanyInfo
                  name={company.name}
                  address={company.address}
                  created_at={formatTimestamp(company.createdAt)}
                  updated_at={formatTimestamp(company.updatedAt)}
                  addUser={() => {}}
                />
              </DeviceRow>
            ))}
          </CategoryHeader>
          {!companiesLoading && companies.length === 0 && (
            <NotFoundCard page="companies" isEmpty={true} />
          )}
        </PageContent>
        <AppSnackbar
          open={snackbar?.open ?? false}
          message={snackbar?.message ?? ""}
          severity={snackbar?.severity}
          onClose={hide}
        />
      </div>
    );
  }

  // — Users view —
  const addButton = (
    <CustomButton onClick={() => setOpenAdd(true)}>Add user</CustomButton>
  );

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader
          title={`Users — ${selectedCompany.name}`}
          action={addButton}
        />
        <div className="flex items-center gap-2 mt-1 mb-2">
          <IconButton
            onClick={handleBack}
            size="small"
            sx={{ color: "primary.main" }}
            aria-label="back to companies"
          >
            <ArrowBackIcon />
          </IconButton>
          <Typography variant="body2" sx={{ color: "primary.main" }}>
            Back to companies
          </Typography>
        </div>
        <PageDivider />
        <CategoryHeader
          categories={userColumns}
          columns={userColumns.length + 1}
          sortableColumns={sortableUserColumns}
          sortConfig={sortConfig}
          onSort={handleSort}
        >
          {usersLoading && <p>Loading...</p>}
          {sorted.map((user) => (
            <DeviceRow key={user.id}>
              <UserInfo
                id={user.id}
                name={user.name}
                email={user.email}
                role={user.role}
                createdAt={formatTimestamp(user.createdAt)}
                onDelete={handleDeleteUser}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!usersLoading && sorted.length === 0 && (
          <NotFoundCard page="users" isEmpty={true} />
        )}
      </PageContent>
      <AddUser
        open={openAdd}
        companyId={selectedCompany.companyId}
        onClose={() => {
          setOpenAdd(false);
          setAddError(null);
        }}
        onAdd={handleAddUser}
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
