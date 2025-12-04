import { useGetAgent, useListAgentDeployments } from "@agent-management-platform/api-client";
import { StatusCard, UnderDevelopment } from "@agent-management-platform/views";
import { BusinessSharp, CheckCircle, CircleOutlined, ErrorRounded } from "@mui/icons-material";
import { Box, CircularProgress } from "@mui/material";
import dayjs from "dayjs";
import { useParams } from "react-router-dom";
import { PromotionTargetEnvironment } from "@agent-management-platform/types";
import { TabStatus } from "../../components/LinkTab";

const getStatusValue = (status: string) => {
  switch (status) {
    case TabStatus.ACTIVE:
      return 'Active';
    case TabStatus.INACTIVE:
      return 'Not deployed';
    default:
      return 'In-progress';
   
  }
}
const getStatusSubtitle = (status: string,
  promotionTargetEnvironment?: PromotionTargetEnvironment) => {
  switch (status) {
    case TabStatus.ACTIVE:
      return `Ready for ${promotionTargetEnvironment?.displayName}`;
    case TabStatus.INACTIVE:
      return `Not ready for ${promotionTargetEnvironment?.displayName}`;
    case TabStatus.DEPLOYING:
      return `Deploying for ${promotionTargetEnvironment?.displayName}`;
    case TabStatus.ERROR:
      return `Failed to deploy for ${promotionTargetEnvironment?.displayName}`;
  }
}
const getStatusIcon = (status: TabStatus) => {
  switch (status) {
    case TabStatus.ACTIVE:
      return <CheckCircle color="success" />;
    case TabStatus.INACTIVE:
      return <CircleOutlined color="disabled" />;
    case TabStatus.DEPLOYING:
      return <CircularProgress size={20} color="warning" />;
    case TabStatus.ERROR:
      return <ErrorRounded color="error" />;
  }
}
const getStatusColor = (status: TabStatus) => {
  switch (status) {
    case TabStatus.ACTIVE:
      return "success";
    case TabStatus.INACTIVE:
      return "primary";
    case TabStatus.DEPLOYING:
      return "warning";
    case TabStatus.ERROR:
      return "error";
  }
}

export function Overview() {
  const { orgId, projectId, agentId, envId } = useParams();
  const { data: agent } = useGetAgent({ 
    orgName: orgId ?? 'default', 
    projName: projectId ?? 'default', 
    agentName: agentId ?? '' 
  });
  const { data: deployments } = useListAgentDeployments({
    orgName: orgId || '',
    projName: projectId || '',
    agentName: agentId || '',
  }, {
    enabled: agent?.provisioning.type === 'internal',
  });
  const currentDeployment = deployments && envId ? deployments[envId] : undefined;
  return (
    <Box display="flex" flexDirection="column" gap={1} pt={1}>
      <Box display="flex" gap={1}>
        <StatusCard
          icon={<BusinessSharp />}
          title={"Endpoints"}
          value={`${(currentDeployment?.endpoints.length || 0)}`}
          subtitle={`Last deployed: ${currentDeployment?.lastDeployed ?
            dayjs(currentDeployment?.lastDeployed).format('MM/DD/YYYY HH:mm') : ''}`}
        />
        <StatusCard
          title={"Endpoints"}
          iconVariant={getStatusColor(currentDeployment?.status as TabStatus ?? TabStatus.INACTIVE)}
          icon={getStatusIcon(currentDeployment?.status as TabStatus ?? TabStatus.INACTIVE)}
          value={getStatusValue(currentDeployment?.status ?? '') ?? ''}
          subtitle={getStatusSubtitle(currentDeployment?.status ?? '', currentDeployment?.promotionTargetEnvironment) ?? ''}
        />
      </Box>
      <UnderDevelopment />
    </Box>
  );
}


