import { AddNewAgent } from './AddNewAgent';
import { Dashboard } from '@mui/icons-material';
import { absoluteRouteMap } from '@agent-management-platform/types';

export const metaData = {
  title: 'Add New Agent',
  description: 'A page component for Add New Agent',
  icon: Dashboard,
  path: absoluteRouteMap.children.org.children.projects.children.newAgent.path,
  component: AddNewAgent,
};

export { AddNewAgent };
export default AddNewAgent;
