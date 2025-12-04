import React from 'react';
import { Box, useTheme } from '@mui/material';
import { NoDataFound } from '../../NoDataFound/NoDataFound';

export const EmptyState: React.FC = () => {
  const theme = useTheme();
  return (
    <Box 
      display="flex" 
      flexDirection="column"
      justifyContent="center" 
      alignItems="center" 
      minHeight={200}
      gap={theme.spacing(2)}
      padding={theme.spacing(4)}
      sx={{
        boxShadow: theme.shadows[1],
        backgroundColor: theme.palette.background.paper,
        borderRadius: theme.shape.borderRadius,
      }}
    >
      <NoDataFound
        message="No data found"
      />
    </Box>
  );
};
