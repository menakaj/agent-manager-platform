import { Box, Typography, useTheme } from "@mui/material";
import { ReactNode } from "react";

interface InfoFieldProps {
    label: string;
    value: ReactNode;
    isMonospace?: boolean;
}

export function InfoField({ label, value, isMonospace = false }: InfoFieldProps) {
    const theme = useTheme();

    return (
        <Box>
            <Typography 
                variant="caption" 
                fontWeight="600" 
                sx={{ 
                    color: theme.palette.text.secondary, 
                    display: 'block', 
                    mb: theme.spacing(0.5) 
                }}
            >
                {label}
            </Typography>
            <Typography 
                variant="body2" 
                sx={{ 
                    fontFamily: isMonospace ? 'monospace' : 'inherit',
                    fontSize: isMonospace ? theme.typography.caption.fontSize : 'inherit',
                    color: theme.palette.text.primary 
                }}
            >
                {value}
            </Typography>
        </Box>
    );
}

