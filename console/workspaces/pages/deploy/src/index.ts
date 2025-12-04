import { DeployComponent } from './Deploy.Component';
import { DeployProject } from './Deploy.Project';
import { DeployOrganization } from './Deploy.Organization';
import { Dashboard } from '@mui/icons-material';

export const metaData = {
  title: 'Deploy',
  description: 'A page component for Deploy',
  icon: Dashboard,
  path: '/deploy',
  component: DeployComponent,
  levels: {
    component: DeployComponent,
    project: DeployProject,
    organization: DeployOrganization,
  },
};

export { 
  DeployComponent,
  DeployProject,
  DeployOrganization,
};

export default DeployComponent;
