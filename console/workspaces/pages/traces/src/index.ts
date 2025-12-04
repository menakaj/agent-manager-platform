import { TracesComponent } from './Traces.Component';
import { TracesProject } from './Traces.Project';
import { TracesOrganization } from './Traces.Organization';
import { Dashboard } from '@mui/icons-material';

export const metaData = {
  title: 'Traces',
  description: 'A page component for Traces',
  icon: Dashboard,
  path: '/traces',
  component: TracesComponent,
  levels: {
    component: TracesComponent,
    project: TracesProject,
    organization: TracesOrganization,
  },
};

export { 
  TracesComponent,
  TracesProject,
  TracesOrganization,
};

export default TracesComponent;
