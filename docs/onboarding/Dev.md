# Setting up your Laptop for Development

## Tools and CLIs

You'll need to install the following tools/binaries, as we depend on them for local development. 

### Go

- Install [Go](https://golang.org/dl/).
- Install [Air](https://github.com/cosmtrek/air), live-reload tool necessary for some Go projects.

### Node, nvm, npm

- Install [nvm](https://github.com/nvm-sh/nvm#installing-and-updating) (node version manager)
- Run `nvm install node` to install the latest node version directly
- Set up [nvm autoload](#nvm-autoload)
- Install [Yarn](https://classic.yarnpkg.com/en/docs/install) package manager

## `nvm` autoload

We have a `.nvmrc` file which ensures the correct Node version is used in every project. 
You can configure your terminal to load the correct node version when you `cd` into it. Follow the [Deeper Shell Integration](https://github.com/nvm-sh/nvm/blob/master/README.md#deeper-shell-integration) guide for the shell of your choice (bash, zsh, fish…).
