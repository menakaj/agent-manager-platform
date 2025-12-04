import { Box } from "@mui/material";
import { ChatBubbleOutline as ChatBubbleIcon, AnalyticsOutlined } from '@mui/icons-material';
import { GroupNavLinks, SubTopNavBar } from "./SubTopNavBar";
import { generatePath, matchPath, Outlet, useLocation, useParams } from "react-router-dom";
import { absoluteRouteMap } from "@agent-management-platform/types";

export const EnvSubNavBar = () => {
    const { orgId, projectId, agentId, envId } = useParams();
    const { pathname } = useLocation();
    const navLinks: GroupNavLinks[] = [
        {
            id: 'overview',
            navLinks: [
                {
                    id: 'try-out',
                    label: 'Try Out',
                    icon: <ChatBubbleIcon />,
                    isActive: !!matchPath(
                        absoluteRouteMap.children.org.children.projects.children.
                            agents.children.environment.path, pathname),
                    path: generatePath(absoluteRouteMap.children.org.children.projects.children.agents.children.environment.path, { orgId: orgId ?? 'default', projectId: projectId ?? 'default', agentId: agentId ?? 'default', envId: envId ?? 'default' })
                },
                {
                    id: 'observe',
                    label: 'Observe',
                    icon: <AnalyticsOutlined />,
                    isActive: !!matchPath(
                        absoluteRouteMap.children.org.children.projects.children.
                            agents.children.environment.children.observability.wildPath, pathname),
                    path: generatePath(absoluteRouteMap.children.org.children.projects.children.agents.children.environment.children.observability.path, { orgId: orgId ?? 'default', projectId: projectId ?? 'default', agentId: agentId ?? 'default', envId: envId ?? 'default' })
                }
            ]
        }
    ];
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <SubTopNavBar navLinks={navLinks} />
            <Outlet />
        </Box>
    );
};
