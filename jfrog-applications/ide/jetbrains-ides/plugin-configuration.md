# Plugin Configuration

To access the plugin configuration, click on the gear icon:


![setting-button](../../.gitbook/assets/intellij/settingButton.png)

## Advanced

#### Downloading External Resources Through Artifactory

JFrog IntelliJ Plugin requires necessary resources for scanning your projects.
By default, the JFrog IntelliJ Plugin downloads the necessary resources needed from https://releases.jfrog.io. If the machine that runs JFrog IntelliJ Plugin has no access to it, follow these steps to allow the tools to be downloaded through an Artifactory instance, which the machine has access to:

1. Login to the JFrog Platform UI, with a user who has admin credentials.

2. Create a Remote Repository with the following properties set:
    - Under the 'Basic' tab:
        Package Type: Generic
        URL: https://releases.jfrog.io
    - Under the 'Advanced' tab:
        Uncheck the 'Store Artifacts Locally' option
3. Navigate to the Advanced tab within the JFrog IntelliJ Plugin, opt for the "Download resources through Artifactory", and input the Repository Key you generated into the Repository name text field:

![externalResourcesThroughArtifactory](../../.gitbook/assets/intellij/externalResourcesThroughArtifactory.png)
