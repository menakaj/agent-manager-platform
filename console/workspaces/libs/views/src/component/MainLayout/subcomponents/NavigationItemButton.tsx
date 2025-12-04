import {
    ListItemButton,
    ListItemIcon,
    ListItemText,
    useTheme,
    Typography,
    alpha,
    Tooltip,
} from '@mui/material';
import { Link as RouterLink } from "react-router-dom";
import { NavigationItem } from './Sidebar';

export interface NavigationItemButtonProps {
    item: NavigationItem;
    sidebarOpen: boolean;
    isMobile?: boolean;
    onNavigationClick?: () => void;
    subButton?: boolean;
}

export function NavigationItemButton({
    item,
    sidebarOpen,
    isMobile = false,
    onNavigationClick,
    subButton = false,
}: NavigationItemButtonProps) {
    const theme = useTheme();

    return (
        <Tooltip title={item.label} placement='right' disableHoverListener={sidebarOpen}>
            <ListItemButton
                onClick={() => {
                    if ('onClick' in item && item.onClick) {
                        item.onClick();
                    }
                    if (isMobile) {
                        onNavigationClick?.();
                    }
                }}
                component={item.href ? RouterLink : 'div'}
                to={item.href ?? ''}
                selected={item.isActive}
                sx={{
                    justifyContent: 'start',
                    alignItems: 'center',
                    height: theme.spacing(5.5),
                    //   p: theme.spacing(1),
                    // p:0,
                    m:0,
                    pl: (subButton && sidebarOpen) ? theme.spacing(5) : theme.spacing(1.75),
                    transition: theme.transitions.create('all', { duration: theme.transitions.duration.short }),
                    '&.Mui-selected': {
                        backgroundColor: alpha(theme.palette.secondary.light, 0.4),
                        '&:hover': {
                            backgroundColor: alpha(theme.palette.secondary.light, 0.3),
                            opacity: 1,
                        },
                    },
                    '&:hover': {
                        backgroundColor: alpha(theme.palette.secondary.light, 0.2),
                        opacity: 1,
                    },
                }}
            >
                {item.icon && (
                    <ListItemIcon
                        sx={{
                            minWidth: theme.spacing(3),
                            color: item.isActive ?
                                theme.palette.primary.contrastText :
                                theme.palette.secondary.light,
                        }}
                    >
                        {item.icon}
                    </ListItemIcon>
                )}
                {sidebarOpen && (
                    <ListItemText
                        primary={
                            <Typography
                                variant="body2"
                                noWrap
                                sx={{
                                    color: item.isActive
                                        ? theme.palette.primary.contrastText :
                                        theme.palette.secondary.light,
                                }}
                            >
                                {item.label}
                            </Typography>
                        }
                    />
                )}
            </ListItemButton>
        </Tooltip>
    );
}

