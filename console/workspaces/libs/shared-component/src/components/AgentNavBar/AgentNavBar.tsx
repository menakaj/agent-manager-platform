import { generatePath, Outlet, useParams } from "react-router-dom";
import { EnvBuildSelector } from "./EnvBuildSelector";
import { PageLayout } from "@agent-management-platform/views";
import { useGetAgent } from "@agent-management-platform/api-client";
import { absoluteRouteMap } from "@agent-management-platform/types";

export function AgentNavBar() {
    const { orgId, agentId, projectId } = useParams();
    const { data: agent } = useGetAgent({
        orgName: orgId ?? 'default',
        projName: projectId ?? 'default',
        agentName: agentId ?? '',
    });
    return (
        <PageLayout
            title={agent?.displayName ?? 'Agent'}
            description={agent?.description}
            backHref={
                generatePath(absoluteRouteMap.children.org.children.projects.path, { orgId: orgId ?? 'default', projectId: projectId ?? 'default'})
            }
            backLabel="Agents"
            actions={
                <EnvBuildSelector />
            }>
            <Outlet />
        </PageLayout>
    );
}
