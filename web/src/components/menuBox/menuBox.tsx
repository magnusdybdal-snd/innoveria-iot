import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import { Link as RouterLink } from "react-router";
import type { SubPage } from "@/pages/subPageList.tsx";

type MenuBoxProps = {
  title: string;
  mainPage: string;
  subPages: SubPage[];
  add: boolean;
};

{
  /*Box for menu components*/
}
export default function MenuBox({
  title,
  mainPage,
  subPages,
  add,
}: MenuBoxProps) {
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
