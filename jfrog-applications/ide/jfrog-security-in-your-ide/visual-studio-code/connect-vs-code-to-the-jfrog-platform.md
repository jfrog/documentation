# Connect VS Code to the JFrog Platform

## Optionally set up a free JFrog Platform in the cloud

<details>

<summary>Run one of the following commands in your terminal.</summary>

**MacOS and Linux using cUrl**

```
curl -fL "https://getcli.jfrog.io?setup" | sh
```

**Windows using PowerShell**

```
powershell "Start-Process -Wait -Verb RunAs powershell '-NoProfile iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/[RELEASE]/jfrog-cli-windows-amd64/jf.exe -OutFile $env:SYSTEMROOT\system32\jf.exe'" ; jf setup
```

The commands will do the following:

1. Install JFrog CLI on your machine.
2. Create a FREE JFrog environment in the cloud for you.
3. Configure VS Code to connect to your new environment.

</details>

## Connect VS Code to your JFrog Platform

Connect to your JFrog environment by clicking on the green Connect button or the provided button in the JFrog extension tab:&#x20;

<figure><img src="../../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

You can leave the platform URL empty, to enter the separate URLs for Artifactory and Xray.

**Note**: If the credentials are configured and used by JFrog CLI the extension will try to read and use them to connect.

The extension also supports connecting to your JFrog environment using environment variables. You may provide basic auth credentials or access tokens as follows:

**Note**: For security reasons, it is recommended to unset the environment variables after launching VS Code.

* `JFROG_IDE_URL` - JFrog URL
* `JFROG_IDE_USERNAME` - JFrog username
* `JFROG_IDE_PASSWORD` - JFrog password
* `JFROG_IDE_ACCESS_TOKEN` - JFrog access token
* `JFROG_IDE_STORE_CONNECTION` - Set the value of this environment variable to **true**, if you'd like VS Code to store the connection details after reading them from the environment variables.
