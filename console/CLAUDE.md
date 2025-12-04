# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a React/TypeScript Agent Management Platform console built as a **Rush monorepo**. The project uses React 19, Vite for building, and pnpm as the package manager (managed by Rush). The architecture follows a clear separation between apps, shared libraries, and page components.

## Technology Stack

- **Monorepo**: Rush (v5.157.0) with pnpm (v9.12.3)
- **Runtime**: Node.js 18.20.3+ or 20.14.0+
- **Frontend**: React 19, TypeScript 5.9.3, Vite 7.1.7
- **UI**: Material-UI (MUI) v7 with Emotion for styling
- **Routing**: React Router v7.9.4
- **State Management**: TanStack React Query v5.90.5
- **Form Handling**: react-hook-form with yup validation
- **Authentication**: Asgardeo (can be disabled via config)
- **Testing**: Vitest with Testing Library
- **Component Development**: Storybook

## Essential Rush Commands

```bash
# Install all dependencies (first step for any new clone)
rush install

# Build all projects
rush build

# Build specific project and its dependencies
rush build --to @agent-management-platform/webapp

# Run linting across all projects
rush lint

# Run tests for all projects
rush test

# Clean all build outputs
rush purge

# Update dependencies after modifying package.json
rush update

# Create a new page component (uses Yeoman generator)
rush create-page
```

## Development Workflow

### Starting Development

```bash
# From console/ directory
rush install
rush build --to @agent-management-platform/webapp
cd apps/webapp
rushx dev
```

The dev server runs on `http://localhost:5173`.

### Working on Individual Packages

Navigate to any package and use `rushx` commands:

```bash
cd workspaces/libs/views
rushx build              # Build the package
rushx dev                # Watch mode for development
rushx test               # Run tests
rushx test-watch         # Run tests in watch mode
rushx lint               # Run ESLint
rushx lint:fix           # Fix linting issues
rushx storybook          # Launch Storybook (pages only)
```

### Running Tests

```bash
# All projects
rush test

# Specific package
cd workspaces/libs/views
rushx test
rushx test-watch
rushx test:ui  # Vitest UI
```

## Architecture

### Project Structure

The monorepo is organized into three main categories:

**apps/** - Main applications
- `webapp/` - Main React application (@agent-management-platform/webapp)

**workspaces/libs/** - Shared libraries
- `auth/` - Authentication provider and hooks (supports Asgardeo or no-auth mode)
- `types/` - Shared TypeScript types and route definitions
- `api-client/` - API client utilities with React Query hooks
- `views/` - Shared UI components, layouts, and MUI themes
- `eslint-config/` - Shared ESLint configuration

**workspaces/pages/** - Page components (federated modules)
- `agent-list/` - Agents list page component
- `add-new-agent/` - Add new agent page component
- `agent-view/` - Agent detail view page component
- `overview/` - Dashboard overview page component
- `.template/` - Yeoman generator for creating new pages

### Key Architectural Patterns

**Authentication Abstraction**
- The `auth` library exports `AuthProvider` and `useAuthHooks` that dynamically switch between Asgardeo and no-auth implementations based on `globalConfig.disableAuth`
- Location: `workspaces/libs/auth/src/index.ts`

**Route Management**
- Routes are defined in `workspaces/libs/types/src/routes/routes.map.ts`
- A script generates absolute paths: `cd workspaces/libs/types && rushx generate-route`
- Both relative and absolute route maps are exported from `@agent-management-platform/types`

**Runtime Configuration**
- App reads config from `window.__RUNTIME_CONFIG__` set in `apps/webapp/public/config.js`
- Template at `apps/webapp/public/config.template.js` supports environment variable substitution
- Config includes auth settings, API base URL, and auth enable/disable flag

**Provider Hierarchy** (apps/webapp/src/Providers/GlobalProviders/GlobalProviders.tsx)
```
ThemeProvider (custom context)
  └─ MuiThemeProvider (theme switches based on context)
      └─ AuthProvider (Asgardeo or no-auth)
          └─ ClientProvider (API client with React Query)
```

**Page Component Pattern**
- Each page is a separate package in `workspaces/pages/`
- Pages export metadata including component, path, and display info
- Pages are imported and registered in `apps/webapp/src/Route/Route.tsx`
- Pages include Storybook for isolated development
- Build output is library format (dist/index.js)

## Creating New Page Components

```bash
cd console
rush create-page
# Follow prompts for package name, title, description, route path
```

After generation:
1. Add the new package to `rush.json` projects array
2. Run `rush update` to install dependencies
3. Run `rush build --to @agent-management-platform/<package-name>`
4. Import and register the page in `apps/webapp/src/Route/Route.tsx`

## Configuration

### Environment Setup

Copy and customize the configuration:
```bash
cp apps/webapp/public/config.js.template apps/webapp/public/config.js
```

Edit `apps/webapp/public/config.js` with your settings:
- `authConfig`: Asgardeo authentication configuration
- `disableAuth`: Set to `true` to disable authentication
- `apiBaseUrl`: Backend API URL (e.g., 'http://localhost:8080')

### TypeScript Configuration

Projects use TypeScript 5.9.3. Each package has:
- `tsconfig.json` - General TypeScript config
- `tsconfig.lib.json` - Library build config (for libs and pages)

### Building Library Packages

Libraries and pages must be built before the webapp can use them:
```bash
# Build all dependencies for webapp
rush build --to @agent-management-platform/webapp

# Build individual library
cd workspaces/libs/views
rushx build
```

## Common Tasks

### Adding a New Route

1. Edit `workspaces/libs/types/src/routes/routes.map.ts` to add route definition
2. Run route generation: `cd workspaces/libs/types && rushx generate-route`
3. Update `apps/webapp/src/Route/Route.tsx` to add the route component

### Updating Shared Types

1. Edit types in `workspaces/libs/types/src/`
2. Build types: `cd workspaces/libs/types && rushx build`
3. Rebuild dependent packages: `rush build`

### Working with Themes

- Light and dark themes: `workspaces/libs/views/src/theme/`
- Theme selection: `apps/webapp/src/contexts/ThemeContext.tsx`
- Export both `aiAgentTheme` and `aiAgentDarkTheme` from `@agent-management-platform/views`

### API Client Usage

```typescript
import { useGetAgents } from '@agent-management-platform/api-client';

const { data, isLoading, error } = useGetAgents();
```

API client is configured in `workspaces/libs/api-client/` with React Query hooks.

## Dependency Management

- Rush uses pnpm workspaces
- Local packages use `workspace:*` protocol in package.json
- After modifying any package.json: `rush update`
- Rush maintains lockfile at `common/config/rush/pnpm-lock.yaml`

## Linting and Code Quality

```bash
# Lint all projects
rush lint

# Lint specific project
cd workspaces/libs/views
rushx lint
rushx lint:fix
```

ESLint config is shared from `@agent-management-platform/eslint-config`.

## Important Notes

- Always run `rush install` after pulling changes that modify package.json files
- Build libraries before building webapp: `rush build --to @agent-management-platform/webapp`
- Rush commands must be run from the `console/` directory
- Individual package scripts use `rushx` instead of `npm run`
- The monorepo supports project depth 1-3 (category/package or category/subcategory/package)
- Page components must be added to `rush.json` projects array manually after generation
