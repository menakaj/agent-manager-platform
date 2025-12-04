#!/usr/bin/env node

const path = require('path');
const { execSync } = require('child_process');

console.log('🚀 Creating a new page component...\n');

try {
  // Change to the template directory
  const templateDir = path.join(__dirname, 'workspaces', 'pages', '.template');
  process.chdir(templateDir);
  
  // Link the generator globally
  console.log('📦 Linking generator globally...');
  execSync('npm link', { stdio: 'inherit' });
  
  // Change back to pages directory
  process.chdir(path.join(__dirname, 'workspaces', 'pages'));
  
  // Run the Yeoman generator
  console.log('🎯 Running generator...');
  execSync('npx yo agent-page', { stdio: 'inherit' });
  
  console.log('\n✅ Page created successfully!');
  console.log('\nNext steps:');
  console.log('1. Add the new page to rush.json projects list');
  console.log('2. Run "rush update" from the console root to install dependencies');
  console.log('3. Run "rushx build" to build the package');
  console.log('4. Run "rushx storybook" to view the component in Storybook');
  
} catch (error) {
  console.error('❌ Error creating page:', error.message);
  process.exit(1);
}
