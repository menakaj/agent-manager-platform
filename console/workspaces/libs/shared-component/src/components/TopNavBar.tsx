import { Box } from "@mui/material";
import { useParams, useSearchParams } from "react-router-dom";
import { TabStatus, TopNavBarGroup } from "./LinkTab";
import { useGetAgent, useListAgentDeployments, useListEnvironments } from "@agent-management-platform/api-client";
import { useEffect, useMemo } from "react";


export const TopNavBar: React.FC = () => {
    const { orgId, agentId, projectId } = useParams();
    const [searchParams, setSearchParams] = useSearchParams();
    const { data: agent } = useGetAgent({
        orgName: orgId ?? 'default',
        projName: projectId ?? 'default',
        agentName: agentId ?? ''
    });
    const { data: environments } = useListEnvironments({ orgName: orgId ?? 'default' });
    const { data: deployments } = useListAgentDeployments({
        orgName: orgId || '',
        projName: projectId || '',
        agentName: agentId || '',
    }, {
        enabled: agent?.provisioning.type === 'internal',
    });

    const sortedEnvironments = useMemo(() =>
        environments?.sort((a) => a.isProduction ? 1 : -1) ?? [], [environments]);

    const selectedEnvironment = searchParams.get('environment');

    // Set first environment as default if no environment is selected
    useEffect(() => {
        if (!selectedEnvironment && sortedEnvironments.length > 0 && agent?.provisioning.type === 'internal') {
            setSearchParams(prev => {
                const newSearchParams = new URLSearchParams(prev);
                newSearchParams.set('environment', sortedEnvironments[0].name);
                return newSearchParams; 
            }, { replace: true });
        }
    }, [selectedEnvironment, sortedEnvironments, setSearchParams, agent?.provisioning.type]);

    if (agent?.provisioning.type === 'external' || !agent) {
        return null;
    }

    return (
        <Box display="flex" gap={1}>
            <TopNavBarGroup
                // eslint-disable-next-line max-len
                tabs={sortedEnvironments.map((env) => {
                    const tabSearchParams = new URLSearchParams(searchParams);
                    tabSearchParams.set('environment', env.name);
                    return {
                        to: `?${tabSearchParams.toString()}`,
                        label: env.displayName ?? env.name,
                        status: deployments?.[env.name]?.status as TabStatus,
                        isProduction: env.isProduction,
                        id: env.name,
                    };
                })}
                selectedId={selectedEnvironment || sortedEnvironments[0]?.name}
            />
        </Box>
    );
};
