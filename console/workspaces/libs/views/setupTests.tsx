import React from 'react';
import { expect } from 'vitest';
import * as matchers from '@testing-library/jest-dom/matchers';
import { TestingLibraryMatchers } from '@testing-library/jest-dom/matchers';

// Extend vitest expect with jest-dom matchers
declare module 'vitest' {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  interface Assertion<T = any>
    extends jest.Matchers<void, T>,
      TestingLibraryMatchers<T, void> {}
}

expect.extend(matchers);

// Mock MUI theme for testing
import { ThemeProvider } from '@mui/material/styles';
import { aiAgentTheme } from './src/theme';

export const testTheme = aiAgentTheme;

export const TestWrapper = ({ children }: { children: React.ReactNode }) => (
  <ThemeProvider theme={testTheme}>
    {children}
  </ThemeProvider>
);


