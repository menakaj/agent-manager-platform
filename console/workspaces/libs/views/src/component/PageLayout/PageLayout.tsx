import { ArrowBack } from "@mui/icons-material";
import { Box, Button, Container, Divider, Typography, useTheme } from "@mui/material";
import { ReactNode } from "react";
import { Link } from "react-router-dom";

export interface PageLayoutProps {
    children: ReactNode;
    backHref?: string;
    title?: string;
    backLabel?: string;
    description?: string;
    titleTail?: ReactNode;
    actions?: ReactNode;
    disablePadding?: boolean;
}
export function PageLayout({ children,
    title, backHref,
    backLabel,
    description,
    titleTail,
    actions,
    disablePadding = false
}: PageLayoutProps) {
    const theme = useTheme();
    return (
        <Box
            display="flex"
            flexDirection="column"
            overflow="auto"
            py={disablePadding ? 0 : 1.5}
            px={disablePadding ? 0 : 3}
            gap={2}
            sx={{
                animation: 'fadeIn 0.3s ease-in-out',
                '@keyframes fadeIn': {
                    from: {
                        opacity: 0,
                    },
                    to: {
                        opacity: 1,
                    },
                },
            }}>
            <Box flexGrow={1} display="flex" justifyContent="space-between" flexDirection="row" gap={2}>
                <Box display="flex" alignItems="center" gap={2}>
                    {backHref && (
                        <>
                            <Button
                                variant="text"
                                size="small"
                                sx={{
                                    color: theme.palette.text.secondary,
                                }}
                                component={Link}
                                startIcon={<ArrowBack fontSize="inherit" />}
                                to={backHref}
                            >
                                {backLabel || "Back"}
                            </Button>
                            <Divider orientation="vertical" flexItem />
                        </>

                    )}
                    <Box display="flex" flexDirection="column" gap={theme.spacing(1)}>
                        <Box display="flex" alignItems="center" gap={theme.spacing(1)}>
                            <Typography variant={backHref ? "h4" : "h2"} color="text.primary">{title}</Typography>
                            {titleTail && (
                                <Box>{titleTail}</Box>
                            )}
                        </Box>

                        {description && (
                            <Typography variant="body2" color="text.secondary">{description}</Typography>
                        )}
                    </Box>
                </Box>
                <Box>
                    {actions && (
                        <Box>{actions}</Box>
                    )}
                </Box>
            </Box>
            <Box flexGrow={1} display="flex" flexDirection="column">
                <Divider />
                <Box gap={theme.spacing(2)}>
                    {children}
                </Box>
            </Box>
        </Box>
    );
}

export function PageLayoutContent(props: Omit<PageLayoutProps, 'disablePadding'>) {
    return (
        <Container maxWidth="lg" disableGutters>   
            <PageLayout disablePadding={true} {...props} />
        </Container>
    );
}
