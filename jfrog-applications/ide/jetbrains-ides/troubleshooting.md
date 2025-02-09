# Troubleshooting

The JFrog Plugin uses the IDE log files. By default, the log level used by the plugin is INFO.

You have the option of increasing the log level to DEBUG. Here's how to do it:

1. Go to **Help** | **Diagnostic Tools** | **Debug Log Settings...**
2. Inside the **Custom Debug Log Configuration** window add the following line:

```java
#com.jfrog.ide.idea.log.Logger
```

To see the IDE log file, depending on the IDE version and OS as described [here](https://intellij-support.jetbrains.com/hc/en-us/articles/207241085-Locating-IDE-log-files), go to **Help** | **Show/reveal Log in Explorer/finder/Konqueror/Nautilus**.

## Reporting Issues

Please report issues by opening an issue on [GitHub](https://github.com/jfrog/jfrog-idea-plugin/issues).

## Contributions

We welcome community contributions through pull requests. To help us improve this project, please read our Contribution guide.

### Android Studio Support for JCEF

The JFrog Plugin uses JCEF (Java Chromium Embedded Framework) to create a webview component in the plugin's tool window.

Most IntelliJ-based IDEs use a boot runtime that contains JCEF by default.

Android Studio and some older versions of other IntelliJ-based IDEs use a boot runtime that doesn't contain JCEF by default, and therefore the plugin can't be loaded in them.

To solve this issue, open the ["Choose Boot Runtime for the IDE"](https://www.jetbrains.com/help/idea/switching-boot-jdk.html) dialog where you can change the boot runtime to one that contains JCEF.

## Release Notes

The release notes are available on [Marketplace](https://plugins.jetbrains.com/plugin/9834-jfrog/versions).
