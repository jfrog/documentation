# Scan Repositories

#### Automatic pull requests creation

Frogbot scans your Git repositories periodically and automatically creates pull requests for upgrading vulnerable dependencies to a version with a fix.
  
  ![](../.gitbook/assets/fix-pr.png)

#### Adding Security Alerts

For GitHub repositories, issues that are found during Frogbot's periodic scans are also added to the [Security Alerts](https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/managing-code-scanning-alerts-for-your-repository) view in the UI. 

![](../.gitbook/assets/github-code-scanning.png)

The following alert types are supported:

**1. CVEs on vulnerable dependencies**


![](../.gitbook/assets/github-code-scanning-content.png)

**2. Secrets that are exposed in the code**

![](../.gitbook/assets/github-code-scanning-secrets-content.png)

**3. Infrastructure as Code (Iac) issues on Terraform packages**

![](../.gitbook/assets/github-code-scanning-iac-content.png)

**4. Static Application Security Testing (Sast) vulnerabilities**

![](../.gitbook/assets/github-code-scanning-sast-content.png)

**5. Validate Dependency Licenses**

![](../.gitbook/assets/github-code-scanning-license-violation-content.png)