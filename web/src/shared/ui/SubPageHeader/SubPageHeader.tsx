import type { ReactNode } from "react";

import Typography from "@mui/material/Typography";

import { Path } from "@/shared/ui/Path";

/**
 * Props for the SubPageHeader component.
 * @param title - Page heading displayed as an h4; omit to render no heading
 * @param action - Optional toolbar element (e.g. a button or menu) placed to the right of the title
 */
type SubPageHeaderProps = {
  title?: string;
  action?: ReactNode;
};

/**
 * SubPageHeader renders the top section of a sub-page, combining a breadcrumb
 * path with an optional title and action slot.
 * @param root0 - Component props
 * @param root0.title - Page heading text; conditionally rendered when provided
 * @param root0.action - Arbitrary action element aligned to the trailing edge of the header row
 * @returns The rendered sub-page header
 */
export function SubPageHeader({ title, action }: SubPageHeaderProps) {
  return (
    <>
      <Path />
      <div className="flex justify-between flex-wrap">
        {title && <Typography variant="h4">{title}</Typography>}
        {action}
      </div>
    </>
  );
}
