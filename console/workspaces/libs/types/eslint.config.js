import sharedConfig from '@agent-management-platform/eslint-config'

export default [
  ...sharedConfig,
  {
    ignores: ['dist/**', 'node_modules/**'],
  },
]