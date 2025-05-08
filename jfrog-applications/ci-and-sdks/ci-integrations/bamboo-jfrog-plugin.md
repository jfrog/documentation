# Bamboo JFrog Plugin

### Overview

The **Bamboo JFrog Plugin** is designed to provide an easy integration between Bamboo and the [JFrog Platform](https://jfrog.com/solution-sheet/jfrog-platform/).

Unlike the legacy [Bamboo Artifactory Plugin](https://plugins.atlassian.com/plugin/details/27818), the new **Bamboo JFrog Plugin** focuses on a single task that runs [JFrog CLI](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli) commands. Worth mentioning that both JFrog plugins, can work side by side.

The advantage of this approach is that JFrog CLI is a powerful and versatile tool that integrates with all JFrog capabilities. It offers extensive features and functionalities, and it is constantly improved and updated with the latest enhancements from JFrog. This ensures that the Bamboo JFrog Plugin is always up-to-date with the newest features and improvements provided by JFrog.

With the Bamboo JFrog Plugin, you can easily deploy artifacts, resolve dependencies, and link them to the build jobs that created them. Additionally, you can scan your artifacts and builds for vulnerabilities using [JFrog Xray](https://jfrog.com/xray/) and distribute your software packages to remote locations using [JFrog Distribution](https://jfrog.com/distribution/).

### Key Features

* **Artifact Management**: Manage build artifacts with Artifactory.
* **Dependency Resolution**: Resolve dependencies from Artifactory for reliable builds.
* **Build Traceability**: Link artifacts to their corresponding build jobs for better traceability.
* **Security Scanning**: Scan artifacts and builds with JFrog Xray for vulnerabilities.
* **Software Distribution**: Distribute software packages to remote locations using JFrog Distribution.

### Installation and Configuration

1. Download the plugin (upto version 1.0.5) from the [Bamboo Marketplace](https://marketplace.atlassian.com/).
   > **Note: For the latest version 1.0.6, download source code of the plugin and build it until it is available in the Bamboo Marketplace**.
   > 1. Download the source code from the [GitHub Release Page](https://github.com/jfrog/bamboo-jfrog-plugin/releases/tag/1.0.6).
   > 2. From the source code root path, build the jar by running the command `mvn clean install -U -DskipTests`.
3. Install the plugin on your Bamboo server.
4. In the _Bamboo Administration_ section, go to _Manage Apps_ and select _JFrog Configuration_.
5. Click on _New JFrog Platform Configuration_.
6. Configure your credentials details and run a _Test Connection_, then click _Save_.


### Important Upgrade Instructions 

#### For Bamboo 10.0.0 and Above
- **Upgrade to Version 1.0.6 or Higher**: If you are using Bamboo version **10.0.0** or later, please upgrade your JFrog Bamboo Plugin to version **1.0.6** or above.
- **Post-Upgrade Configuration**: After upgrading to version **1.0.6**, any additions or modifications to your configurations will require a **restart of Bamboo** to take effect. 

    **We will resolve this issue in future releases**.

#### For Bamboo 9.x.x and Below
- **Do Not Upgrade to Version 1.0.6**: If your Bamboo version is **9.x.x** or lower, please do not upgrade to version **1.0.6**, as it is not compatible with these versions.


### JFrog CLI Settings

1. By default, latest JFrog CLI will be installed and used when the JFrog CLI task runs. You can specify a specific version to be used.
2.  If your Bamboo agents have access to the internet, you can set the JFrog Plugin to download JFrog CLI directly from https://releases.jfrog.io. If not, you can set the plugin to download JFrog CLI through the configured Artifactory instance.

    Set the Repository Name field value to the name of a Remote or Virtual repository in your Artifactory instance which proxies https://releases.jfrog.io/.

### Usage

Once installed and configured, you can use the JFrog CLI task in your Bamboo build plans. Follow these steps:

1. Go to the _Tasks_ section of your build plan.
2. Add the _JFrog CLI task_ to your plan.
3. Configure the JFrog CLI task by selecting the appropriate Server ID.
