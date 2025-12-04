import React from 'react';
import { Decorator } from '@storybook/react';
import { useDarkMode } from 'storybook-dark-mode';
import { BrowserRouter } from 'react-router-dom';
import { IntlProvider } from 'react-intl';
import { ThemeProvider } from '@mui/material/styles';
import { CssBaseline } from '@mui/material';
import relativeTime from 'dayjs/plugin/relativeTime';
import dayjs from 'dayjs';
import { aiAgentTheme, aiAgentDarkTheme } from '../src/theme';

dayjs.extend(relativeTime);

export const withTheme: Decorator = (Story) => {
  const isDark = useDarkMode();
  const theme = isDark ? aiAgentDarkTheme : aiAgentTheme;

  return (
    <BrowserRouter>
      <IntlProvider locale="en">
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Story />
        </ThemeProvider>
      </IntlProvider>
    </BrowserRouter>
  );
}; 
