import { Box, Typography, IconButton } from "@mui/material";
import { Close, Timeline } from "@mui/icons-material";

interface SpanPanelHeaderProps {
    onClose: () => void;
}

export function SpanPanelHeader({ onClose }: SpanPanelHeaderProps) {
    return (
        <Box 
            sx={{ 
                display: 'flex', 
                justifyContent: 'space-between', 
                alignItems: 'center' 
            }}
        >
            <Typography variant="h4">
                <Timeline fontSize="inherit" />
                &nbsp;
                Span Details
            </Typography>
            <IconButton color="error" size="small" onClick={onClose}>
                <Close />
            </IconButton>
        </Box>
    );
}

