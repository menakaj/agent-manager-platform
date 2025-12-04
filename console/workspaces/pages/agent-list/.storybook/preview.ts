import type { Preview } from '@storybook/react';
import { ThemeProvider, CssBaseline } from '@mui/material';
import { aiAgentTheme, aiAgentDarkTheme } from '@agent-management-platform/views';
import React from 'react';

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
    backgrounds: {
      default: 'light',
      values: [
        {
          name: 'light',
          value: '#F8F9FB',
        },
        {
          name: 'dark',
          value: '#1A1A1A',
        },
      ],
    },
  },
  decorators: [
    (Story, context) => {
      const isDark = context.globals.backgrounds?.value === '#1A1A1A';
      const theme = isDark ? aiAgentDarkTheme : aiAgentTheme;
      
      return React.createElement(
        ThemeProvider,
        { theme },
        React.createElement(React.Fragment, null, [
          React.createElement(CssBaseline, { key: 'baseline' }),
          React.createElement(Story, { key: 'story' }),
        ])
      );
    },
  ],
};

export default preview;

