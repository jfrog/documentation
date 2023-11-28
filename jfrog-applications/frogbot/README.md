# JFrog Frogbot

![](https://raw.github.com/jfrog/frogbot/master/images/frogbot-intro.png)

### Overview

JFrog Frogbot is a Git bot that scans your Git repositories for security vulnerabilities.

1. It scans pull requests immediately after they are opened but before they are merged. This process notifies you if the pull request is about to introduce new vulnerabilities to your code. This unique capability ensures the code is scanned and can be fixed even before vulnerabilities are introduced into the codebase.
2. It scans the Git repository periodically and creates pull requests with fixes for detected vulnerabilities.

#### Frogbot supports the following Git providers:

- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/github-icon.png" width="15"> GitHub
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/gitlab-icon.png" width="15"> GitLab
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/azure-devops-icon.png" width="15"> Azure Repos
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/bitbucket-icon.png" width="15"> Bitbucket Server

#### For SCA Scanning. Frogbot supports projects that use the following package managers:

- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/go-icon.png" width="15"> Go
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/gradle-icon.png" width="15"> Gradle
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/maven-icon.png" width="15"> Maven
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/npm-icon.png" width="15"> npm
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/yarn-icon.png" width="15"> Yarn
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/dotnet-icon.png" width="15"> .NET
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/nuget-icon.png" width="15"> NuGet
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/pip-icon.png" width="15"> Pip
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/pip-icon.png" width="15"> Pipenv
- <img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/poetry-icon.png" width="15"> Poetry

### Why use JFrog Frogbot?

* **Software Composition Analysis (SCA)**: Scan your project dependencies for security issues. For selected security issues, get leverage-enhanced CVE data from our JFrog Security Research team. Frogbot uses JFrog's vast vulnerabilities database, to which we continuously add new component vulnerability data.&#x20;
* **Validate Dependency Licenses**: Ensure that the licenses for the project's dependencies are in compliance with a predefined list of approved licenses.
* **Static Application Security Testing (SAST)**: Provides fast and accurate security-focused engines that detect zero-day security vulnerabilities on your source code sensitive operations, while minimizing false positives.
* **CVE Vulnerability Contextual Analysis**: This feature uses the code context to eliminate false positive reports on vulnerable dependencies that are not applicable to the code. For CVE vulnerabilities that are applicable to your code, Frogbot will create pull request comments on the relevant code lines with full descriptions regarding the security issues caused by the CVE. Vulnerability Contextual Analysis is currently supported for Python, JavaScript, and Java code.
* **Secrets Detection**: Detect any secrets left exposed inside the code. to stop any accidental leak of internal tokens or credentials.
* **Infrastructure as Code scans (IaC)**: Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

> _**NOTE:**_ **SAST**, **Vulnerability Contextual Analysis**, **Secrets Detection** and **Infrastructure as Code scans** require the [JFrog Advanced Security Package](https://jfrog.com/xray/).
