import React, { useState } from 'react';
import { IconButton, Menu, MenuItem, useTheme } from '@mui/material';
import { MoreVert } from '@mui/icons-material';

export interface ActionItem {
  label: string;
  value: string;
  onClick?: (row: any) => void;
}

export interface ActionMenuProps<T = any> {
  row: T;
  actions: ActionItem[];
  onActionClick: (action: string, row: T) => void;
}

export const ActionMenu = <T extends Record<string, any>>({
  row,
  actions,
  onActionClick,
}: ActionMenuProps<T>) => {
  const theme = useTheme();
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);

  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleActionClick = (action: ActionItem) => {
    handleMenuClose();
    onActionClick(action.value, row);
    action.onClick?.(row);
  };

  if (actions.length === 0) return null;

  return (
    <>
      <IconButton
        onClick={handleMenuOpen}
        size="small"
        aria-label="actions"
        sx={{
          color: theme.palette.text.secondary,
          '&:hover': {
            backgroundColor: theme.palette.action.hover,
            color: theme.palette.text.primary,
          },
        }}
      >
        <MoreVert />
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleMenuClose}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'right',
        }}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        PaperProps={{
          sx: {
            backgroundColor: theme.palette.background.paper,
            border: `1px solid ${theme.palette.divider}`,
            boxShadow: theme.shadows[3],
            minWidth: 160,
          },
        }}
      >
        {actions.map((action) => (
          <MenuItem
            key={action.value}
            onClick={() => handleActionClick(action)}
            sx={{
              fontSize: theme.typography.body2.fontSize,
              color: theme.palette.text.primary,
              padding: theme.spacing(1, 2),
              '&:hover': {
                backgroundColor: theme.palette.action.hover,
              },
            }}
          >
            {action.label}
          </MenuItem>
        ))}
      </Menu>
    </>
  );
};
