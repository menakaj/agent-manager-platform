import { OverviewComponent } from './Overview.Component';
import { OverviewProject } from './Overview.Project';
import { OverviewOrganization } from './Overview.Organization';
import { Dashboard } from '@mui/icons-material';

export const metaData = {
  title: 'Overview',
  description: 'A page component for Overview',
  icon: Dashboard,
  path: '/overview',
  component: OverviewComponent,
  levels: {
    component: OverviewComponent,
    project: OverviewProject,
    organization: OverviewOrganization,
  },
};

export { 
  OverviewComponent,
  OverviewProject,
  OverviewOrganization,
};

export default OverviewComponent;
