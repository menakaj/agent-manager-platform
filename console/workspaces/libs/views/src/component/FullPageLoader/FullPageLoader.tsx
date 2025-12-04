import { PersonOutlined } from "@mui/icons-material";
import { Box, LinearProgress, useTheme } from "@mui/material";

export function FullPageLoader() {
    const theme = useTheme();
    return (
        <Box sx={{
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'center',
            alignItems: 'center',
            height: '100vh',
            width: '100vw'
        }}>
            <Box sx={{ display: 'flex', flexDirection:'column', justifyContent: 'center', alignItems: 'center', gap: theme.spacing(2) }}>
                <PersonOutlined sx={{ fontSize: 100 }} color="primary" />
                <LinearProgress color="primary" value={50} sx={{ width: '100%' }} />
            </Box>
        </Box>
    );
}
