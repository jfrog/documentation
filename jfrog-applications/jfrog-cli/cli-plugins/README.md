# JFrog CLI Plugins

![](../../../.gitbook/assets/jfrog-cli-header.png)

## General

JFrog CLI plugins enhance the functionality of JFrog CLI to meet specific user and organization needs. The source code of a plugin is maintained as an open-source Go project on GitHub. All public plugins are registered in [JFrog CLI's Plugins Registry](https://github.com/jfrog/jfrog-cli-plugins-reg). We encourage developers to create plugins and share them publicly with the community. When a plugin is included in the registry, it becomes publicly available and can be installed using JFrog CLI. For more information, see the [JFrog CLI Plugins Developer Guide](developer-guide.md) if you want to create and publish your own plugins.

## Installing Plugins

A plugin which is included [JFrog CLI's Plugins Registry](https://github.com/jfrog/jfrog-cli-plugins-reg) can be installed using the following command.

```
$ jf plugin install the-plugin-name
```

This command will install the plugin from the official public registry by default. You can also install a plugin from a private JFrog CLI Plugin registry, as described in the _Private Plugins Registries_ section.

## Private Plugins Registries

In addition to the public official JFrog CLI Plugins Registry, JFrog CLI supports publishing and installing plugins to and from private JFrog CLI Plugins Registries. A private registry can be hosted on any Artifactory server. It uses a local generic Artifactory repository for storing the plugins.

To create your own private plugins registry, follow these steps.

* On your Artifactory server, create a local generic repository named _jfrog-cli-plugins_.
* Make sure your Artifactory server is included in JFrog CLI's configuration, by running the _jf c show_ command.
* If needed, configure your Artifactory instance using the _jf c add_ command.
* Set the ID of the configured server as the value of the JFROG\_CLI\_PLUGINS\_SERVER environment variable.
* If you wish the name of the plugins repository to be different from jfrog-cli-plugins, set this name as the value of the JFROG\_CLI\_PLUGINS\_REPO environment variable.

The **jf plugin install** command will now install plugins stored in your private registry.

To publish a plugin to the private registry, run the following command, while inside the root of the plugin's sources directory. This command will build the sources of the plugin for all the supported operating systems. All binaries will be uploaded to the configured registry.

```
jf plugin publish the-plugin-name the-plugin-version
```
