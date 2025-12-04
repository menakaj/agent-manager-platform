import { Box, Typography, Divider, Button, Drawer, useTheme, Skeleton } from "@mui/material";
import { useTrace } from "@agent-management-platform/api-client";
import { FadeIn, TraceExplorer } from "@agent-management-platform/views";
import { generatePath, Link, useParams } from "react-router-dom";
import { absoluteRouteMap, Span } from "@agent-management-platform/types";
import { ListAltOutlined } from "@mui/icons-material";
import { useState, useCallback } from "react";
import { SpanDetailsPanel } from "./SpanDetailsPanel";

function TraceDetailsSkeleton() {
    const theme = useTheme();

    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                gap: theme.spacing(2)
            }}
        >
            <Box
                sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center'
                }}
            >
                <Skeleton variant="rectangular" width={theme.spacing(15)} height={theme.spacing(4.5)} />
                <Skeleton variant="text" width={theme.spacing(20)} height={theme.spacing(4)} />
            </Box>

            <Divider />

            <Box
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    gap: theme.spacing(1.5)
                }}
            >
                <Skeleton variant="rectangular" width="100%" height={theme.spacing(8)} />
                {[...Array(8)].map((_, index) => (
                    <Skeleton
                        key={index}
                        variant="rectangular"
                        width="100%"
                        height={theme.spacing(6)}
                        sx={{
                            ml: theme.spacing(index % 3 * 2)
                        }}
                    />
                ))}
            </Box>
        </Box>
    );
}

export function TraceDetails() {
    const theme = useTheme();
    const { orgId = "default", projectId = "default", agentId = "default", envId, traceId = "default" } = useParams();
    const { data: traceDetails, isLoading } = useTrace(
        orgId,
        projectId,
        agentId,
        envId ?? '',
        traceId
    );

    const [selectedSpan, setSelectedSpan] = useState<Span | null>(null);

    const handleCloseSpan = useCallback(() => setSelectedSpan(null), []);

    if (isLoading) {
        return <TraceDetailsSkeleton />;
    }

    const spans = traceDetails?.spans ?? [];

    if (spans.length === 0) {
        return (
            <FadeIn>
                <Box
                    sx={{
                        display: 'flex',
                        justifyContent: 'space-between',
                        alignItems: 'center'
                    }}
                >
                    <Button
                        startIcon={<ListAltOutlined fontSize='inherit' />}
                        component={Link}
                        to={
                            envId ?
                                generatePath(absoluteRouteMap.children.org.
                                    children.projects.children.agents.children.environment.
                                    children.observability.children.traces.path,
                                    { orgId: orgId ?? '', projectId: projectId ?? '', agentId: agentId ?? '', envId: envId ?? '', traceId: traceId as string })
                                :
                                generatePath(absoluteRouteMap.children.org.
                                    children.projects.children.agents.children.traces.path,
                                    { orgId: orgId ?? '', projectId: projectId ?? '', agentId: agentId ?? '', traceId: traceId as string })
                        }
                    >
                        Trace List
                    </Button>
                    <Typography variant="h6" fontWeight="bold">
                        Trace Details
                    </Typography>
                </Box>
                <Box
                    sx={{
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        height: '100%',
                        padding: theme.spacing(10)
                    }}
                >
                    <Typography variant="body1" color="text.secondary">No spans found</Typography>
                </Box>
            </FadeIn>
        );
    }

    return (
        <FadeIn>
            <Box
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    gap: theme.spacing(2),
                    height: '100%'
                }}
            >
                <Box
                    sx={{
                        display: 'flex',
                        justifyContent: 'space-between',
                        alignItems: 'center'
                    }}
                >
                    <Button
                        startIcon={<ListAltOutlined fontSize='inherit' />}
                        component={Link}
                        to={
                            envId ?
                                generatePath(absoluteRouteMap.children.org.
                                    children.projects.children.agents.children.environment.
                                    children.observability.children.traces.path,
                                    { orgId: orgId ?? '', projectId: projectId ?? '', agentId: agentId ?? '', envId: envId ?? '', traceId: traceId as string })
                                :
                                generatePath(absoluteRouteMap.children.org.
                                    children.projects.children.agents.path,
                                    { orgId: orgId ?? '', projectId: projectId ?? '', agentId: agentId ?? '' })
                        }
                    >
                        Trace List
                    </Button>
                    <Typography variant="h6" fontWeight="bold">
                        Trace Details
                    </Typography>
                </Box>

                <Divider />

                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: 'column',
                        gap: theme.spacing(2)
                    }}
                >
                    {traceId && (
                        <TraceExplorer onOpenAtributesClick={setSelectedSpan} spans={spans} />
                    )}
                </Box>
                <Drawer
                    anchor="right"
                    open={!!selectedSpan}
                    onClose={handleCloseSpan}
                    sx={{ zIndex: 1300 }}
                >
                    <SpanDetailsPanel span={selectedSpan} onClose={handleCloseSpan} />
                </Drawer>
            </Box>
        </FadeIn>
    );
}

