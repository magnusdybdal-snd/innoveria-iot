import { useState } from "react";

import {
  postRule,
  RuleRow,
  useRules,
  type CreateRuleRequest,
} from "@entities/context";
import { AddRule } from "@features/addRule";
import { CustomButton } from "@shared/ui/Button";
import { CategoryHeader } from "@shared/ui/CategoryHeader";
import { DeviceRow } from "@shared/ui/DeviceRow";
import { NotFoundCard } from "@shared/ui/NotFoundCard";
import { PageContent } from "@shared/ui/PageContent";
import {
  AppSnackbar,
  SNACKBAR_SEVERITY,
  useSnackbar,
} from "@shared/ui/snackbar";
import { SubPageHeader } from "@shared/ui/SubPageHeader";

// TODO: replace with company ID from auth context once auth context is wired up
const COMPANY_ID = "a0000000-0000-0000-0000-000000000001";

const RULE_COLUMNS = [
  "Name",
  "Measurement type",
  "Aggregation method",
  "Active",
  "", // Action column with menu, no header
];

/**
 * Page for listing and creating context aggregation rules for a company.
 * @returns The rendered rules page.
 */
export default function Rules() {
  const { rules, isLoading, refetch } = useRules(COMPANY_ID);
  const [openAdd, setOpenAdd] = useState(false);
  const [addError, setAddError] = useState<string | null>(null);
  const { show, hide, snackbar } = useSnackbar();

  const handleAdd = (rule: CreateRuleRequest): Promise<void> => {
    setAddError(null);
    return postRule(rule)
      .then(() => {
        refetch();
        setOpenAdd(false);
        show("Rule added successfully", SNACKBAR_SEVERITY.SUCCESS);
      })
      .catch((err: unknown) => {
        setAddError("Failed to add rule. Please try again.");
        show("Failed to add rule.", SNACKBAR_SEVERITY.ERROR);
        throw err;
      });
  };

  const addButton = (
    <CustomButton onClick={() => setOpenAdd(true)}>Add rule</CustomButton>
  );

  return (
    <div className="flex h-screen">
      <PageContent>
        <SubPageHeader title="Context rules" action={addButton} />
        <CategoryHeader categories={RULE_COLUMNS} columns={RULE_COLUMNS.length}>
          {rules.map((rule) => (
            <DeviceRow key={rule.id}>
              <RuleRow rule={rule} onDelete={refetch} />
            </DeviceRow>
          ))}
        </CategoryHeader>
        {!isLoading && rules.length === 0 && (
          <NotFoundCard page="Rules" isEmpty={true} />
        )}
      </PageContent>
      <AddRule
        open={openAdd}
        companyId={COMPANY_ID}
        onClose={() => {
          setOpenAdd(false);
          setAddError(null);
        }}
        onAdd={handleAdd}
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
