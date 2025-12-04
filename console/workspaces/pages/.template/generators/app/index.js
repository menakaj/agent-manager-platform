const Generator = require('yeoman-generator');

module.exports = class extends Generator {
  async prompting() {
    const prompts = [
      {
        type: 'input',
        name: 'name',
        message: 'What is the name of your page package?',
        default: 'my-page',
        validate: (input) => {
          if (!input || input.trim() === '') {
            return 'Package name is required';
          }
          if (!/^[a-z0-9-]+$/.test(input)) {
            return 'Package name should only contain lowercase letters, numbers, and hyphens';
          }
          return true;
        }
      },
      {
        type: 'input',
        name: 'title',
        message: 'What is the display title for your page?',
        default: (answers) => {
          return answers.name
            .split('-')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1))
            .join(' ');
        }
      },
      {
        type: 'input',
        name: 'description',
        message: 'What is the description for your page?',
        default: (answers) => `A page component for ${answers.title}`
      },
      {
        type: 'input',
        name: 'routePath',
        message: 'What is the route path for your page?',
        default: (answers) => `/${answers.name.replace(/-/g, '/')}`
      }
    ];

    const answers = await this.prompt(prompts);
    
    // Generate component name from package name
    answers.componentName = answers.name
      .split('-')
      .map(word => word.charAt(0).toUpperCase() + word.slice(1))
      .join('');

    this.props = answers;
  }

  writing() {
    // Copy all template files
    this.fs.copyTpl(
      this.templatePath('**/*'),
      this.destinationPath(this.props.name),
      this.props
    );
  }

  install() {
    // This will be handled by the monorepo's package manager
    this.log('Template generated successfully!');
    this.log(`Next steps:`);
    this.log(`1. cd ${this.props.name}`);
    this.log(`2. Run rush update to install dependencies`);
    this.log(`3. Run rushx build to build the package`);
    this.log(`4. Run rushx storybook to view the component in Storybook`);
  }
};
