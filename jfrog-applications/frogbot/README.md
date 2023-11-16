# JFrog Frogbot

![](https://raw.github.com/jfrog/frogbot/master/images/frogbot-intro.png)

### Overview

JFrog Frogbot is a Git bot that scans your Git repositories for security vulnerabilities.

1. It scans pull requests immediately after they are opened but before they are merged. This process notifies you if the pull request is about to introduce new vulnerabilities to your code. This unique capability ensures the code is scanned and can be fixed even before vulnerabilities are introduced into the codebase.
2. It scans the Git repository periodically and creates pull requests with fixes for detected vulnerabilities.

#### Frogbot supports the following Git providers:

| ![GitHub](https://raw.githubusercontent.com/jfrog/frogbot/master/images/github-icon.png) GitHub | ![GitLab](https://raw.githubusercontent.com/jfrog/frogbot/master/images/gitlab-icon.png) GitLab | ![Azure](https://raw.githubusercontent.com/jfrog/frogbot/master/images/azure-devops-icon.png) Azure Repos | ![Bitbucket](https://raw.githubusercontent.com/jfrog/frogbot/master/images/bitbucket-icon.png) Bitbucket Server |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |

#### For SCA Scanning. Frogbot supports projects that use the following package managers:

| ![Go](https://raw.githubusercontent.com/jfrog/frogbot/master/images/go-icon.png) Go         | ![Gradle](https://raw.githubusercontent.com/jfrog/frogbot/master/images/gradle-icon.png) Gradle | ![Maven](https://raw.githubusercontent.com/jfrog/frogbot/master/images/maven-icon.png) Maven | ![npm](https://raw.githubusercontent.com/jfrog/frogbot/master/images/npm-icon.png) npm       | ![Yarn](https://raw.githubusercontent.com/jfrog/frogbot/master/images/yarn-icon.png) Yarn       |
| ------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| ![.NET](https://raw.githubusercontent.com/jfrog/frogbot/master/images/dotnet-icon.png) .NET | ![NuGet](https://raw.githubusercontent.com/jfrog/frogbot/master/images/nuget-icon.png) NuGet    | ![Pip](https://raw.githubusercontent.com/jfrog/frogbot/master/images/pip-icon.png) Pip       | ![Pipenv](https://raw.githubusercontent.com/jfrog/frogbot/master/images/pip-icon.png) Pipenv | ![Poetry](https://raw.githubusercontent.com/jfrog/frogbot/master/images/poetry-icon.png) Poetry |

### Why use JFrog Frogbot?

* **Software Composition Analysis (SCA)**: Scan your project dependencies for security issues. For selected security issues, get leverage-enhanced CVE data from our JFrog Security Research team. Frogbot uses JFrog's vast vulnerabilities database, to which we continuously add new component vulnerability data.&#x20;
* **Validate Dependency Licenses**: Ensure that the licenses for the project's dependencies are in compliance with a predefined list of approved licenses.
* **Static Application Security Testing (SAST)**: Provides fast and accurate security-focused engines that detect zero-day security vulnerabilities on your source code sensitive operations, while minimizing false positives.
* **CVE Vulnerability Contextual Analysis**: This feature uses the code context to eliminate false positive reports on vulnerable dependencies that are not applicable to the code. For CVE vulnerabilities that are applicable to your code, Frogbot will create pull request comments on the relevant code lines with full descriptions regarding the security issues caused by the CVE. Vulnerability Contextual Analysis is currently supported for Python, JavaScript, and Java code.
* **Secrets Detection**: Detect any secrets left exposed inside the code. to stop any accidental leak of internal tokens or credentials.
* **Infrastructure as Code scans (IaC)**: Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

> _**NOTE:**_ **SAST**, **Vulnerability Contextual Analysis**, **Secrets Detection** and **Infrastructure as Code scans** require the [JFrog Advanced Security Package](https://jfrog.com/xray/).
