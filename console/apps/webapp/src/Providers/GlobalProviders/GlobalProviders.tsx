import { AuthProvider } from "@agent-management-platform/auth";
import { ClientProvider } from "@agent-management-platform/api-client";
import { ThemeProvider as MuiThemeProvider } from "@mui/material";
import { aiAgentTheme, aiAgentDarkTheme } from "@agent-management-platform/views";
import { ThemeProvider, useTheme } from "../../contexts/ThemeContext";

const MuiThemeWrapper = ({ children }: { children: React.ReactNode }) => {
  const { actualTheme } = useTheme();
  const theme = actualTheme === 'dark' ? aiAgentDarkTheme : aiAgentTheme;

  return <MuiThemeProvider theme={theme}>{children}</MuiThemeProvider>;
};

export const GlobalProviders = ({ children }: { children: React.ReactNode }) => {
  return (
    <ThemeProvider>
      <MuiThemeWrapper>
        <AuthProvider>
          <ClientProvider>
            {children}
          </ClientProvider>
        </AuthProvider>
      </MuiThemeWrapper>
    </ThemeProvider>
  );
};
