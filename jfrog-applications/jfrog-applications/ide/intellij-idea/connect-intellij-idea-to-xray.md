# Connect the JFrog Plugin to the JFrog Platform

<details>

<summary>Optionally set up a free JFrog Environment in the Cloud</summary>

Need a free JFrog environment in the Cloud, so that JFrog IntelliJ IDEA Plugin can connect to it? Just run one of the following commands in your terminal. The commands will do the following:

1. Install JFrog CLI on your machine.
2. Create a FREE JFrog environment in the Cloud for you.
3. Configure IntelliJ IDEA to connect to your new environment.

**MacOS and Linux using cURL**

```bash
curl -fL https://getcli.jfrog.io?setup | sh
```

**Windows using PowerShell**

```powershell
powershell "Start-Process -Wait -Verb RunAs powershell '-NoProfile iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/[RELEASE]/jfrog-cli-windows-amd64/jf.exe -OutFile $env:SYSTEMROOT\system32\jf.exe'" ; jf setup
```

</details>

<details>

<summary>Connect the JFrog Plugin to an existing JFrog Environment</summary>

You can connect the plugin to your JFrog environment using one of the following methods:

</details>

> **NOTES:**
>
> * If your JFrog Platform instance uses a domain with a self-signed certificate, add the certificate to IDEA as described [here](https://www.jetbrains.com/help/idea/settings-tools-server-certificates.html).
> * From JFrog Xray version **1.9** to **2.x**, IntelliJ IDEA users connecting to Xray from IntelliJ are required to be granted the ‘View Components’ action in Xray.
> * From JFrog Xray version **3.x**, as part of the JFrog Platform, IntelliJ IDEA users connecting to Xray from IntelliJ require ‘Read’ permission. For more information, see [here](https://jfrog.com/help/r/jfrog-platform-administration-documentation/permissions).
