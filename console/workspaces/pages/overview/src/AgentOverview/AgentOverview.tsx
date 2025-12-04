import { useGetAgent } from "@agent-management-platform/api-client";
import { InternalAgentOverview } from "./InternalAgentOverview";
import { ExternalAgentOverview } from "./ExternalAgentOverview";
import { metaData as tracesMetadata } from '@agent-management-platform/traces';
import { useParams } from "react-router-dom";

export function AgentOverview() {
    const { orgId, agentId, projectId } = useParams();
    const { data: agent } = useGetAgent({
        orgName: orgId ?? 'default',
        projName: projectId ?? 'default',
        agentName: agentId ?? ''
    });

    if (agent?.provisioning.type === 'internal') {
        return (
            <InternalAgentOverview />
        )
    }
    if (agent?.provisioning.type === 'external') {
        return (
            <>
                <ExternalAgentOverview />
                <tracesMetadata.levels.component />
            </>
        )
    }

    return null;
}
