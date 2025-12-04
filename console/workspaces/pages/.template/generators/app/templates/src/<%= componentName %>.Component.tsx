import React from 'react';
import { Box, Typography } from '@mui/material';

export interface <%= componentName %>ComponentProps {
  title?: string;
  description?: string;
}

export const <%= componentName %>Component: React.FC<<%= componentName %>ComponentProps> = ({
  title = '<%= title %> - Component Level',
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
        Component Level View
      </Typography>
    </Box>
  );
};

export default <%= componentName %>Component;
