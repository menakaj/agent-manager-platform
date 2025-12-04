import { BuildComponent } from './Build.Component';
import { BuildProject } from './Build.Project';
import { BuildOrganization } from './Build.Organization';
import { Dashboard } from '@mui/icons-material';

export const metaData = {
  title: 'Build',
  description: 'A page component for Build',
  icon: Dashboard,
  path: '/build',
  component: BuildComponent,
  levels: {
    component: BuildComponent,
    project: BuildProject,
    organization: BuildOrganization,
  },
};

export { 
  BuildComponent,
  BuildProject,
  BuildOrganization,
};

export default BuildComponent;
