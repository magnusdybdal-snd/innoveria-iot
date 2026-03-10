import { useState, type ReactNode } from "react";

import MoreVertIcon from "@mui/icons-material/MoreVert";
import IconButton from "@mui/material/IconButton";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";

export type ActionMenuItem = {
  label: string;
  onClick?: () => void;
  disabled?: boolean;
};

type ActionMenuProps = {
  items: ActionMenuItem[];
  icon?: ReactNode;
};

/**
 * A three-dot icon button that opens a dropdown context menu.
 * Accepts a list of labeled action items with optional disabled states.
 * The trigger icon defaults to MoreVertIcon but can be overridden.
 * @param props - Component props
 * @param props.items - List of menu actions, each with a label, handler, and optional disabled flag
 * @param props.icon - Custom trigger icon; defaults to a vertical three-dot icon
 * @returns An icon button paired with a MUI Menu dropdown
 */
export function ActionMenu({ items, icon }: ActionMenuProps) {
  const [anchor, setAnchor] = useState<null | HTMLElement>(null);

  const handleOpen = (e: React.MouseEvent<HTMLElement>) => {
    setAnchor(e.currentTarget);
  };

  const handleClose = () => {
    setAnchor(null);
  };

  return (
    <>
      <IconButton
        size="small"
        onClick={handleOpen}
        sx={{
          justifySelf: "end",
          alignSelf: "center",
        }}
      >
        {icon ?? <MoreVertIcon sx={{ color: "primary.main" }} />}
      </IconButton>
      <Menu
        anchorEl={anchor}
        open={Boolean(anchor)}
        onClose={handleClose}
        sx={{
          "& .MuiPaper-root": {
            backgroundColor: "primary.light",
            color: "primary.main",
          },
        }}
      >
        {items.map((item) => (
          <MenuItem
            key={item.label}
            disabled={item.disabled}
            onClick={() => {
              item.onClick?.();
              handleClose();
            }}
          >
            {item.label}
          </MenuItem>
        ))}
      </Menu>
    </>
  );
}
