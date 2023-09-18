# Scan Repositories

#### Automatic pull requests creation

Frogbot scans your Git repositories periodically and automatically creates pull requests for upgrading vulnerable dependencies to a version with a fix.
> **_NOTE:_** Currently not supported in Gradle.
  
  ![](../.gitbook/assets/fix-pr.png)

#### Adding Security Alerts

For GitHub repositories, issues that are found during Frogbot's periodic scans are also added to the [Security Alerts](https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/managing-code-scanning-alerts-for-your-repository) view in the UI. The following alert types are supported:

**1. CVEs on vulnerable dependencies**

![](../.gitbook/assets/github-code-scanning.png)\
![](../.gitbook/assets/github-code-scanning-content.png)

**2. Secrets that are exposed in the code**

![](../.gitbook/assets/github-code-scanning-secrets-content.png)

**3. Infrastructure as Code (Iac) issues on Terraform packages**

![](../.gitbook/assets/github-code-scanning-iac-content.png)
