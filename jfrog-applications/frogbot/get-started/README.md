# Get Started

JFrog Frogbot is a Git bot that scans your git repositories for security vulnerabilities through the following key functions:

1. Scans repository selected branches periodically and creates pull requests with fixes for vulnerabilities that are detected.
2. Scans newly opened pull requests to ensure no new vulnerabilities are introduced to the codebase.


#### Why use JFrog Frogbot?
* **Software Composition Analysis (SCA)**: Scan your project dependencies for security issues. For selected security issues, get leverage-enhanced CVE data from our JFrog Security Research team. Frogbot uses JFrog's vast vulnerabilities database, to which we continuously add new component vulnerability data. Also included is VulnDB, the industry's most comprehensive security database, to further extend the range of vulnerabilities detected and fixed by Frogbot.
* **Static Application Security Testing (SAST)**: Provides fast and accurate security-focused engines that detect zero-day security vulnerabilities on your source code sensitive operations, while minimizing false positives.
* **CVE Vulnerability Contextual Analysis**: This feature uses the code context to eliminate false positive reports on vulnerable dependencies that are not *applicable* to the code. Vulnerability Contextual Analysis is currently supported for Python and JavaScript code.
* **Secrets Detection**: Detect any secrets left exposed inside the code. to stop any accidental leak of internal tokens or credentials.
* **Infrastructure as Code scans (IaC)**: Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

> _**NOTE:**_ **SAST**, **Vulnerability Contextual Analysis**, **Secrets Detection** and **Infrastructure as Code scans** require the [JFrog Advanced Security Package](https://jfrog.com/xray/).

#### What's needed for the setup?

* JFrog Platform to scan your Git repositories. If you don't have a JFrog Platform, you can set up one for free.
  
* CI server to run the scan tasks. The following environments are supported:
 
<img height="20" width="20"  src="https://cdn.simpleicons.org/GitHubActions" alt="GitHubActions" /> GitHub Actions

<img height="20" width="20"  src="https://cdn.simpleicons.org/Jenkins" alt="Jenkins" /> Jenkins

<img height="20" width="20"  src="https://cdn.simpleicons.org/JfrogPipelines" alt="jfrogpipelines" /> JFrog Pipelines

<img height="20" width="20"  src="https://cdn.simpleicons.org/Gitlab" alt="Gitlab" /> GitLab CI

<img height="20" width="20"  src="https://cdn.simpleicons.org/AzurePipelines" alt="AzurePipelines" /> Azure Pipelines
