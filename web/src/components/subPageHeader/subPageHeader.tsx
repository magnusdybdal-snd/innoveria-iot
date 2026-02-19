import type { ReactNode } from "react";

import Typography from "@mui/material/Typography";

import { Path } from "@/components/path";

type SubPageHeaderProps = {
  title?: string;
  action?: ReactNode;
};

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
