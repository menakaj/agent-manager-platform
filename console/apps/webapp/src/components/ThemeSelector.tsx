import React from 'react';
import { IconButton, Tooltip } from '@mui/material';
import { Brightness4, Brightness7, BrightnessAuto } from '@mui/icons-material';
import { useTheme } from '../contexts/ThemeContext';

/**
 * Optional component that provides explicit theme mode selection
 * Allows users to toggle between light, dark, or system theme
 */
export const ThemeSelector: React.FC = () => {
  const { mode, setMode } = useTheme();

  const handleToggle = () => {
    const modes: Array<'light' | 'dark' | 'system'> = ['light', 'dark', 'system'];
    const currentIndex = modes.indexOf(mode);
    const nextIndex = (currentIndex + 1) % modes.length;
    setMode(modes[nextIndex]);
  };

  const getIcon = () => {
    switch (mode) {
      case 'light':
        return <Brightness7 color="inherit" />;
      case 'dark':
        return <Brightness4 color="inherit" />;
      case 'system':
        return <BrightnessAuto color="inherit"/>;
      default:
        return <BrightnessAuto color="inherit"/>;
    }
  };

  const getTooltipText = () => {
    switch (mode) {
      case 'light':
        return 'Light mode (click to switch to dark)';
      case 'dark':
        return 'Dark mode (click to switch to system)';
      case 'system':
        return 'System mode (click to switch to light)';
      default:
        return 'Toggle theme';
    }
  };

  return (
    <Tooltip title={getTooltipText()}>
      <IconButton onClick={handleToggle} color="default">
        {getIcon()}
      </IconButton>
    </Tooltip>
  );
};

