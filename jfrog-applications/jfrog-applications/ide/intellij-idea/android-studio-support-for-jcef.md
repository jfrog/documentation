# Android Studio Support for JCEF

The JFrog IDEA Plugin uses JCEF (Java Chromium Embedded Framework) to create a webview component in the plugin's tool window.

Most IntelliJ-based IDEs use a boot runtime that contains JCEF by default.

Android Studio and some older versions of other IntelliJ-based IDEs use a boot runtime that doesn't contain JCEF by default, and therefore the plugin can't be loaded in them.

To solve this issue, open the ["Choose Boot Runtime for the IDE"](https://www.jetbrains.com/help/idea/switching-boot-jdk.html) dialog where you can change the boot runtime to one that contains JCEF.
