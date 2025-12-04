import { TestComponent } from './Test.Component';
import { TestProject } from './Test.Project';
import { TestOrganization } from './Test.Organization';
import { Dashboard } from '@mui/icons-material';

export const metaData = {
  title: 'Test',
  description: 'A page component for Test',
  icon: Dashboard,
  path: '/test',
  component: TestComponent,
  levels: {
    component: TestComponent,
    project: TestProject,
    organization: TestOrganization,
  },
};

export { 
  TestComponent,
  TestProject,
  TestOrganization,
};

export default TestComponent;
