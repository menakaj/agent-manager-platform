import type { Preview } from '@storybook/react'
import { withTheme } from './Decorator';

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
      expanded: true,
    },
    backgrounds: {
      values: [
        { name: 'Light', value: '#ffffff' },
        { name: 'Dark', value: '#121212' },
        { name: 'Paper', value: '#fafafa' },
        { name: 'Default', value: '#f5f5f5' },
      ],
      default: 'Light',
    },
    docs: {
      theme: {
        // Customize the docs theme to match MUI
        colorPrimary: '#1976d2',
        colorSecondary: '#dc004e',
        fontBase: '"Roboto", "Helvetica", "Arial", sans-serif',
        fontCode: '"Fira Code", "Monaco", "Consolas", monospace',
      },
    },
    viewport: {
      viewports: {
        mobile: {
          name: 'Mobile',
          styles: {
            width: '375px',
            height: '667px',
          },
        },
        tablet: {
          name: 'Tablet',
          styles: {
            width: '768px',
            height: '1024px',
          },
        },
        desktop: {
          name: 'Desktop',
          styles: {
            width: '1024px',
            height: '768px',
          },
        },
      },
    },
  },
  decorators: [withTheme],
  globalTypes: {
    theme: {
      description: 'Global theme for components',
      defaultValue: 'light',
      toolbar: {
        title: 'Theme',
        icon: 'circlehollow',
        items: [
          { value: 'light', title: 'Light' },
          { value: 'dark', title: 'Dark' },
        ],
        dynamicTitle: true,
      },
    },
  },
};

export default preview;
