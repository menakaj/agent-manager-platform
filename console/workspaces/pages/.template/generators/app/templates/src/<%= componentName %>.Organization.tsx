import React from 'react';
import { Box, Typography } from '@mui/material';

export interface <%= componentName %>OrganizationProps {
  title?: string;
  description?: string;
}

export const <%= componentName %>Organization: React.FC<<%= componentName %>OrganizationProps> = ({
  title = '<%= title %> - Organization Level',
  description = '<%= description %>',
}) => {
  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        {title}
      </Typography>
      <Typography variant="body1" color="text.secondary">
        {description}
      </Typography>
      <Typography variant="caption" display="block" sx={{ mt: 2 }}>
        Organization Level View
      </Typography>
    </Box>
  );
};

export default <%= componentName %>Organization;
