import React from 'react';
import { Box, CircularProgress, Typography, useTheme } from '@mui/material';

export interface LoadingStateProps {
  message?: string;
  minHeight?: number;
}

export const LoadingState: React.FC<LoadingStateProps> = ({
  message = 'Loading...',
  minHeight = 200,
}) => {
  const theme = useTheme();

  return (
    <Box 
      display="flex" 
      flexDirection="column"
      justifyContent="center" 
      alignItems="center" 
      minHeight={minHeight}
      gap={theme.spacing(2)}
      padding={theme.spacing(4)}
    >
      <CircularProgress 
        size={40}
        sx={{
          color: theme.palette.primary.main,
        }}
      />
      <Typography 
        variant="body2" 
        sx={{
          color: theme.palette.text.secondary,
          fontSize: theme.typography.body2.fontSize,
        }}
      >
        {message}
      </Typography>
    </Box>
  );
};
