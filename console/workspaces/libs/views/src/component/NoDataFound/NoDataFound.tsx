import { Box, Typography, useTheme } from "@mui/material";
import { SearchOffOutlined } from "@mui/icons-material";
import { FadeIn } from "../FadeIn/FadeIn";
import { ReactNode } from "react";

interface NoDataFoundProps {
    message?: string;
    action?: ReactNode;
    icon?: ReactNode;
    subtitle?: string;
}

export function NoDataFound({ 
    message = "No data found", 
    action,
    icon,
    subtitle
}: NoDataFoundProps) {
    const theme = useTheme();
    return (
        <FadeIn>
            <Box sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                height: '100%',
                color: 'text.secondary',
                p: theme.spacing(2)
            }}>
                {icon || <SearchOffOutlined sx={{ fontSize: 100, mb: 2, opacity: 0.2 }} color="inherit" />}
                <Typography variant="h6" align="center" color="textSecondary" sx={{ mb: subtitle ? 1 : 2 }}>
                    {message}
                </Typography>
                {subtitle && (
                    <Typography variant="body2" align="center" color="textSecondary" sx={{ mb: 2, opacity: 0.7 }}>
                        {subtitle}
                    </Typography>
                )}
                {action && (
                    <Box sx={{ mt: 2 }}>
                        {action}
                    </Box>
                )}
            </Box>
        </FadeIn>
    );
}
