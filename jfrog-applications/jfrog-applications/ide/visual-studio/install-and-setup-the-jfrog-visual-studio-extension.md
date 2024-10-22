# Install and Setup the JFrog Visual Studio Extension

This section reviews how to setup and install the JFrog Visual Studio Extension. It lists the supported visual studio versions, and instructions on installing and configuring the extension. It reviews the following:

* Supported Visual Studio Versions
* Install the JFrog Visual Studio Extension
* Configure the JFrog Visual Studio Extension to Connect to JFrog Xray
* Use the JFrog Visual Studio Extension

### Supported Visual Studio Versions

Two extensions are shared to the marketplace - each of them supports a different Visual Studio version:

* Visual Studio 2022 - [JFrogV2](https://marketplace.visualstudio.com/items?itemName=JFrog.JFrogV2)
* Visual Studio 2017 and 2019 - [JFrog](https://marketplace.visualstudio.com/items?itemName=JFrog.JFrog)

Prerequisites

Frog Xray version 2.5.0 and above.

### Install the JFrog Visual Studio Extension

To install and work with the extension:

1. Open the terminal window.
2. Run the **nuget** command. If it is not recognized as a command, please add **nuget.exe** to the **PATH** environment variable.
3. If your projects use **NPM**, Run the **npm** command. If it is not recognized as a command, please add **npm.exe** to the **PATH**environment variable.
4. Open Visual Studio
5. Go to Tools | Extensions and Updates
6. Search for **JFrog**.
7. Click on Download
8. Once the installation is completed, re-open Visual Studio.

### Configure the JFrog Visual Studio Extension to Connect to JFrog Xray

Once the extension is successfully installed, connect Visual Studio to your instance of JFrog Xray.

### Go to Tools | Options | JFrog | JFrog Xray

1. Set your JFrog Platform URL and login credentials.
2.  Test your connection to Xray using the **Test connection** button.\


    <div align="left">

    <figure><img src="../../../.gitbook/assets/spaces_HtpcI8sApaH537Ph5QxY_uploads_9WPR1FlGfqJNZbX6jYoR_0.png" alt="" width="375"><figcaption></figcaption></figure>

    </div>
