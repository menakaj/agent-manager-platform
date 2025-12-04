import { <%= componentName %>Component } from './<%= componentName %>.Component';
import { <%= componentName %>Project } from './<%= componentName %>.Project';
import { <%= componentName %>Organization } from './<%= componentName %>.Organization';
import { Dashboard } from '@mui/icons-material';

export const metaData = {
  title: '<%= title %>',
  description: '<%= description %>',
  icon: Dashboard,
  path: '<%= routePath %>',
  component: <%= componentName %>Component,
  levels: {
    component: <%= componentName %>Component,
    project: <%= componentName %>Project,
    organization: <%= componentName %>Organization,
  },
};

export { 
  <%= componentName %>Component,
  <%= componentName %>Project,
  <%= componentName %>Organization,
};

export default <%= componentName %>Component;
