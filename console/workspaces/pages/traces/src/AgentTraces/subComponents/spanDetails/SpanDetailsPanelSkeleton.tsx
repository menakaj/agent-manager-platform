import { Box, Divider, Skeleton, useTheme } from "@mui/material";

export function SpanDetailsPanelSkeleton() {
    const theme = useTheme();

    return (
        <Box
            sx={{
                width: theme.spacing(80),
                p: theme.spacing(2),
                height: '100%',
                display: 'flex',
                flexDirection: 'column',
                gap: theme.spacing(2),
                bgcolor: theme.palette.background.paper
            }}
        >
            {/* Header */}
            <Box 
                sx={{ 
                    display: 'flex', 
                    justifyContent: 'space-between', 
                    alignItems: 'center' 
                }}
            >
                <Skeleton variant="text" width={theme.spacing(20)} height={theme.spacing(5)} />
                <Skeleton variant="circular" width={theme.spacing(4)} height={theme.spacing(4)} />
            </Box>
            <Divider />
            
            {/* Content */}
            <Box 
                sx={{ 
                    display: 'flex', 
                    flexDirection: 'column', 
                    gap: theme.spacing(2), 
                    overflow: 'auto', 
                    flex: 1 
                }}
            >
                {/* Basic Info Section */}
                <Box>
                    <Skeleton variant="text" width={theme.spacing(18)} height={theme.spacing(3)} />
                    <Skeleton 
                        variant="rectangular" 
                        width="100%" 
                        height={theme.spacing(30)} 
                        sx={{ mt: theme.spacing(1.5), borderRadius: 1 }} 
                    />
                </Box>

                <Divider />

                {/* Timing Section */}
                <Box>
                    <Skeleton variant="text" width={theme.spacing(12)} height={theme.spacing(3)} />
                    <Skeleton 
                        variant="rectangular" 
                        width="100%" 
                        height={theme.spacing(20)} 
                        sx={{ mt: theme.spacing(1.5), borderRadius: 1 }} 
                    />
                </Box>

                <Divider />

                {/* Status Section */}
                <Box>
                    <Skeleton variant="text" width={theme.spacing(10)} height={theme.spacing(3)} />
                    <Skeleton 
                        variant="rectangular" 
                        width="100%" 
                        height={theme.spacing(15)} 
                        sx={{ mt: theme.spacing(1.5), borderRadius: 1 }} 
                    />
                </Box>

                <Divider />

                {/* Attributes Section */}
                <Box>
                    <Skeleton variant="text" width={theme.spacing(14)} height={theme.spacing(3)} />
                    <Box sx={{ display: 'flex', flexDirection: 'column', gap: theme.spacing(2), mt: theme.spacing(1.5) }}>
                        {[...Array(3)].map((_, index) => (
                            <Box key={index}>
                                <Skeleton variant="text" width={theme.spacing(20)} height={theme.spacing(2.5)} />
                                <Skeleton 
                                    variant="rectangular" 
                                    width="100%" 
                                    height={theme.spacing(12)} 
                                    sx={{ mt: theme.spacing(0.75), borderRadius: 1 }} 
                                />
                            </Box>
                        ))}
                    </Box>
                </Box>
            </Box>
        </Box>
    );
}

