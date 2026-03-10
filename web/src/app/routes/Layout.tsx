import { Menu } from "@widgets/menu";
import { Outlet } from "react-router";

/**
 *  Layout component that wraps all pages with a common Menu and renders the current route's content via Outlet.
 * @returns The rendered layout with menu and page content
 */
export default function Layout() {
  return (
    <Menu>
      <Outlet />
    </Menu>
  );
}
