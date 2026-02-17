import { Path } from "@/components/path";
import Typography from "@mui/material/Typography";
import type { ReactNode } from "react";

type SubPageHeaderProps = {
  title: string;
  action?: ReactNode;
};

export function SubPageHeader({ title, action }: SubPageHeaderProps) {
  return (
    <>
      <Path />
      <div className="flex justify-between flex-wrap">
        <Typography variant="h4">{title}</Typography>
        {action}
      </div>
    </>
  );
}
