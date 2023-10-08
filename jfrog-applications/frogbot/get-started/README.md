# Get Started

JFrog Frogbot is a Git bot that scans your git repositories for security vulnerabilities through the following key functions:

1. Scans repository selected branches periodically and creates pull requests with fixes for vulnerabilities that are detected.
2. Scans newly opened pull requests to ensure no new vulnerabilities are introduced to the codebase.


#### Why use JFrog Frogbot?
* **Software Composition Analysis (SCA)**: Scan your project dependencies for security issues. For selected security issues, get leverage-enhanced CVE data from our JFrog Security Research team. Frogbot uses JFrog's vast vulnerabilities database, to which we continuously add new component vulnerability data. Also included is VulnDB, the industry's most comprehensive security database, to further extend the range of vulnerabilities detected and fixed by Frogbot.
* **Static Application Security Testing (SAST)**: Provides fast and accurate security-focused engines that detect zero-day security vulnerabilities on your source code sensitive operations, while minimizing false positives.
* **CVE Vulnerability Contextual Analysis**: This feature uses the code context to eliminate false positive reports on vulnerable dependencies that are **not applicable** to the code. For CVE vulnerabilities that are **applicable** to your code, Frogbot will create pull request comments on the relevant code lines with full descriptions regarding the security issues caused by the CVE. Vulnerability Contextual Analysis is currently supported for Python, JavaScript, and Java code.
* **Secrets Detection**: Detect any secrets left exposed inside the code. to stop any accidental leak of internal tokens or credentials.
* **Infrastructure as Code scans (IaC)**: Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

> _**NOTE:**_ **SAST**, **Vulnerability Contextual Analysis**, **Secrets Detection** and **Infrastructure as Code scans** require the [JFrog Advanced Security Package](https://jfrog.com/xray/).

#### It supports the following Git providers:

| ![GitHub](../../.gitbook/assets/integrations/GitHub.png) GitHub | ![GitLab](../../.gitbook/assets/integrations/GitLab.png) GitLab | ![Azure](../../.gitbook/assets/integrations/AzureDevops.png) Azure Repos | ![Bitbucket](../../.gitbook/assets/integrations/Bitbucket.png) Bitbucket Server |
|----------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|


#### It supports the following package managers:

| ![Go](../../.gitbook/assets/integrations/Go.png) Go | ![Gradle](../../.gitbook/assets/integrations/Gradle.png) Gradle | ![Maven](../../.gitbook/assets/integrations/ApacheMaven.png) Maven | ![npm](../../.gitbook/assets/integrations/npm.png) npm | ![Yarn](../../.gitbook/assets/integrations/Yarn.png) Yarn |
|:----|:----|:----|:----|:----|
| ![.NET](../../.gitbook/assets/integrations/NET.png) .NET | ![NuGet](../../.gitbook/assets/integrations/NuGet.png) NuGet | ![Pip](../../.gitbook/assets/integrations/Python.png) Pip | ![Pipenv](../../.gitbook/assets/integrations/Python.png) Pipenv | ![Poetry](../../.gitbook/assets/integrations/Poetry.png) Poetry |

 # Setup Frogbot

#### What's needed for the setup?

* **JFrog Platform** server. (If you don't have a JFrog Platform, you can set up one for free)
  
* **CI server** to run the scan tasks.

#### Select your preferred CI server:

![GitHubActions](../../.gitbook/assets/integrations/GitHubActions.png) [GitHub Actions](../../jfrog-applications/frogbot/get-started/setting-frogbot-on-github-repositories.md)

![Jenkins](../../.gitbook/assets/integrations/Jenkins.png) [Jenkins](../../jfrog-applications/frogbot/get-started/set-up-frogbot-using-jenkins.md)

![JfrogPipelines](../../.gitbook/assets/integrations/JfrogPipelines.png) [JFrog Pipelines](../../jfrog-applications/frogbot/get-started/installing-frogbot-on-jfrog-pipelines.md)

![Gitlab](../../.gitbook/assets/integrations/Gitlab.png) [GitLab CI](../../jfrog-applications/frogbot/get-started/installing-frogbot-on-gitlab-repositories.md)

![AzurePipelines](../../.gitbook/assets/integrations/AzurePipelines.png) [Azure Pipelines](../../jfrog-applications/frogbot/get-started/installing-frogbot-on-azure-repositories.md)

<details>
  
<summary>Optionally - set up a FREE JFrog Platform in the Cloud</summary>

Frogbot requires a JFrog environment to scan your projects. If you don't have an environment, we can set up a free environment in the cloud for you. Just run one of the following commands in your terminal to set up an environment in less than a minute.

The commands will do the following:

1. Install [JFrog CLI](https://www.jfrog.com/confluence/display/CLI/JFrog+CLI) on your machine.
2. Create a FREE JFrog environment in the cloud for you.

**For macOS and Linux, use curl**

```
curl -fL "https://getcli.jfrog.io?setup" | sh
```

**For Windows, use PowerShell**

```
powershell "Start-Process -Wait -Verb RunAs powershell '-NoProfile iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/[RELEASE]/jfrog-cli-windows-amd64/jf.exe -OutFile $env:SYSTEMROOT\system32\jf.exe'" ; jf setup
```

After the setup is complete, you'll receive an email with your JFrog environment connection details, which can be stored as secrets in Git.

</details>

<details>

<summary>Advanced - Customize advanced settings with frogbot-config.yml</summary>

* [Creating the frogbot-config.yml file](../../jfrog-applications/frogbot/get-started/frogbot-configuration.md)

</details>

