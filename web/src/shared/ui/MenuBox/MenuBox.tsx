import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import { Link as RouterLink } from "react-router";

import type { SubPage } from "@/shared/config/navigation/subPageList";

type MenuBoxProps = {
  title: string;
  mainPage: string;
  subPages: SubPage[];
  add: boolean;
};

{
  /*Box for menu components*/
}
/**
 * Sidebar navigation card that renders a primary page link and its optional sub-page links.
 * @param root0 - Component props
 * @param root0.title - Display label for the primary navigation link
 * @param root0.mainPage - Route path for the primary link
 * @param root0.subPages - List of sub-page routes rendered below the primary link
 * @param root0.add - Whether to show an "+ Add" link at the bottom of the sub-page list
 * @returns The rendered menu box
 */
export function MenuBox({ title, mainPage, subPages, add }: MenuBoxProps) {
  return (
    <Box
      sx={{
        backgroundColor: "secondary.light",
        color: "primary.main",
        margin: 1,
        borderRadius: 3,
        width: 300,
      }}
    >
      <div className="w-64 p-[10px]">
        {/*Main page*/}
        <h2 className="font-normal text-3xl" style={{ margin: 0 }}>
          <Link
            key={mainPage}
            component={RouterLink}
            to={mainPage}
            underline="none"
            sx={{ color: "primary.main" }}
          >
            {title}
          </Link>
        </h2>

        <nav className="flex flex-col mt-2 space-y-1">
          {subPages.map((page) => (
            <Link
              key={page.path}
              component={RouterLink}
              to={page.path}
              underline="none"
              sx={{ color: "primary.main" }}
            >
              | {page.name}
            </Link>
          ))}

          {add && (
            <Link underline="none" sx={{ color: "primary.main" }}>
              + Add
            </Link>
          )}
        </nav>
      </div>
    </Box>
  );
}
