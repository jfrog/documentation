# Bamboo JFrog Plugin

### Overview

The **Bamboo JFrog Plugin** is designed to provide an easy integration between Bamboo and the [JFrog Platform](https://jfrog.com/solution-sheet/jfrog-platform/).

Unlike the legacy [Bamboo Artifactory Plugin](https://plugins.atlassian.com/plugin/details/27818), the new **Bamboo JFrog Plugin** focuses on a single task that runs [JFrog CLI](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli) commands. Worth mentioning that both JFrog plugins, can work side by side.

The advantage of this approach is that JFrog CLI is a powerful and versatile tool that integrates with all JFrog capabilities. It offers extensive features and is constantly updated with the latest enhancements from JFrog. This ensures that the Bamboo JFrog Plugin is always up-to-date with the newest features and improvements.

With the Bamboo JFrog Plugin, you can easily deploy artifacts, resolve dependencies, and link them to the build jobs that created them. Additionally, you can scan your artifacts and builds for vulnerabilities using [JFrog Xray](https://jfrog.com/xray/) and distribute your software packages to remote locations using [JFrog Distribution](https://jfrog.com/distribution/).

#### Key Features

* Artifact Management: Manage build artifacts with Artifactory.
* Dependency Resolution: Resolve dependencies from Artifactory for reliable builds.
* Build Traceability: Link artifacts to their corresponding build jobs for better traceability.
* Security Scanning: Scan artifacts and builds with JFrog Xray for vulnerabilities.
* Software Distribution: Distribute software packages to remote locations using JFrog Distribution.

### Installation and Configuration

1. Download the plugin from the [Bamboo Marketplace](https://marketplace.atlassian.com/).&#x20;
2. Install the plugin on your Bamboo server.
3. In the _Bamboo Administration_ section, go to _Manage Apps_ and select _JFrog Configuration_.
4. Click on _New JFrog Platform Configuration_.
5. Configure your credentials details and run a _Test Connection_, then click _Save_.

***

### Important Upgrade Instructions

#### For Bamboo 10.0.0 and Above

* Upgrade to Version 1.0.6 or Higher: If you are using Bamboo version 10.0.0 or later, upgrade your JFrog Bamboo Plugin to version 1.0.6 or above.
* Post-Upgrade Configuration: After upgrading, you must restart Bamboo for any new or modified configurations to take effect. This issue will be resolved in a future release.

#### For Bamboo 9.x.x and Below

Warning: Do not upgrade to version 1.0.6 if your Bamboo version is 9.x.x or lower, as it is not compatible.

### JFrog CLI Settings

1. By default, the JFrog CLI task installs and uses the latest version of JFrog CLI. You can specify a different version to be used.
2.  If your Bamboo agents have internet access, you can set the plugin to download JFrog CLI directly from `https://releases.jfrog.io`. If not, you can configure the plugin to download JFrog CLI through a configured Artifactory instance. To do this, set the Repository Name field to the name of a Remote or Virtual repository in your Artifactory instance that proxies `https://releases.jfrog.io/`.



### Usage

After the plugin is installed and configured, use the JFrog CLI task in your Bamboo build plans.

To add the JFrog CLI task:

1. Go to the Tasks section of your build plan.
2. Add the JFrog CLI task to your plan.
3. Configure the task by selecting the appropriate Server ID.
