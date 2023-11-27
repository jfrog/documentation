# Building and Testing the Sources

#### Preconditions

* npm 7 and above
* JFrog CLI's `jf` executable - required for tests

To build the extension from sources, please follow these steps:

1. Clone the code from Github.
2. Update submodules:

```
git submodule init
git submodule update
```

3. Build and create the VS-Code extension vsix file by running the following npm command:

```bash
    npm i
    npm run package
```

After the build finishes, you'll find the vsix file in the _jfrog-vscode-extension_ directory. The vsix file can be loaded into VS-Code

####  Test Instructions:
* Set the environment variables JFROG_IDE_PLATFORM_URL and JFROG_IDE_ACCESS_TOKEN to establish a connection with a valid JFrog instance.
* create a remote repository named `releases-proxy` designed to function as a  [proxy jfrog releases](https://docs.jfrog-applications.jfrog.io/jfrog-applications/ide/visual-studio-code/extension-settings#downloading-external-resources-through-artifactory).

To run the test, execute the following command:
```bash
npm t
```