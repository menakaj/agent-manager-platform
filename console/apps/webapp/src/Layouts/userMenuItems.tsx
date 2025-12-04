import { Logout, Settings } from '@mui/icons-material';

export const createUserMenuItems = (orgId: string, logout: () => void) => [
  {
    label: 'Settings',
    href: "/unknown" + orgId,
    icon: <Settings fontSize='inherit' />,
  },
  {
    label: 'Logout',
    onClick: () => logout(),
    icon: <Logout fontSize='inherit' />,
  },
];
