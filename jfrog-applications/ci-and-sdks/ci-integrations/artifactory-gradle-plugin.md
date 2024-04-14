# Artifactory Gradle Plugin

### Overview

The Gradle Artifactory Plugin provides tight integration with Gradle. All that is needed is a simple modification of your `build.gradle` script file with a few configuration parameters, and you can deploy your build artifacts and build information to Artifactory.

The plugin adds the `artifactoryPublish` task for each project, in the 'publishing' group. The task performs the following actions on the project and its submodules:

1. Extracting the [build-info](https://www.buildinfo.org/) file located in the root project. This file contains comprehensive information about the build, such as its configuration, dependencies, and other relevant details.
2. Deploying both the generated artifacts and the build-info file to your Artifactory repository. This ensures that the artifacts, which are the output of the build process, and the accompanying build-info file are stored and organized in your Artifactory repository for easy access and management.

> _**NOTE:**_ The minimum supported Gradle version to use this plugin is v6.9

<details>

<summary>🚚 Migrating from Version 4 to Version 5 of the Plugin</summary>

***

**Version 5 of the Gradle Artifactory Plugin includes the following breaking changes compared to version 4**

* The minimum version of Gradle required to use this plugin has been upgraded to version 6.9.
*   The below convention attributes have been removed:

    | Attribute | Migration action                                                                                                                                                                                                                              |
    |:---------:|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |  parent   | No longer supported.                                                                                                                                                                                                                          |
    |  resolve  | To define the Artifactory resolution repositories for your build, declare the repositories under the repositories section as described [here](https://docs.gradle.org/current/userguide/declaring\_repositories.html#declaring-repositories). |

</details>

***

### 📦 Installation

<details>

<summary>Step 1 - Add the plugin to your project</summary>

***

Add the following snippet to your build script:

***

</details>

<details>

<summary>Step 2 - Configure the plugin with your Artifactory</summary>

***

To configure the plugin with your Artifactory, add the following basic snippet to your project root build script, and make the necessary adjustments based on your platform information:

#### ⚙️ Advance Configurations

For advanced configurations and finer control over the plugin's operations, refer to the following documentation that outlines all the available configuration options. These options allow you to customize the behavior of the plugin according to your specific needs.

</details>

***

### 🚀 Usage

To deploy the project artifacts and build info to Artifactory, execute the following Gradle task

```bash
./gradlew artifactoryPublish
```

***

### 💡 Examples

The following are links to the build scripts of different types of projects that are configured to use the plugin.

**Multi Modules Project (Groovy)**

Sample project that uses the Gradle Artifactory Plugin with Gradle Publications.

**Multi Modules Project (Kotlin)**

Sample project that configures the Gradle Artifactory Plugin with the Gradle Kotlin DSL.

We highly recommend also using our [gradle project examples](https://github.com/JFrog/project-examples/tree/master/gradle-examples?\_gl=1\*pgsvlz\*\_ga\*MTc3OTI0ODE4NS4xNjYyMjgxMjI1\*\_ga\_SQ1NR9VTFJ\*MTY4NTM2OTcwMC4yNi4wLjE2ODUzNjk3MDAuNjAuMC4w) as a reference when configuring your build scripts.

***

### 🐞 Reporting Issues

We highly recommend running Gradle with the `-d` option to get useful and readable debug information if something goes wrong with your build.

Please help us improve the plugin by [reporting any issues](https://github.com/jfrog/artifactory-gradle-plugin/issues/new/choose) you encounter.

***

### 🫱🏻‍🫲🏼 Contributions

We welcome pull requests from the community. To help us improve this project, please read our Contribution guide.
