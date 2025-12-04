import React from 'react';
import { Box, Typography } from '@mui/material';

export interface <%= componentName %>ProjectProps {
  title?: string;
  description?: string;
}

export const <%= componentName %>Project: React.FC<<%= componentName %>ProjectProps> = ({
  title = '<%= title %> - Project Level',
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
        Project Level View
      </Typography>
    </Box>
  );
};

export default <%= componentName %>Project;
