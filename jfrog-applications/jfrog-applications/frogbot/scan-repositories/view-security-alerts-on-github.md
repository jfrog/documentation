# View Security Alerts on GitHub

For GitHub repositories, issues that are found during Frogbot's periodic scans are also added to the [Security Alerts](https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/managing-code-scanning-alerts-for-your-repository) view in the UI.

![](../../../.gitbook/assets/github-code-scanning.png)

The following alert types are supported:

**1. CVEs on vulnerable dependencies**

![](../../../.gitbook/assets/github-code-scanning-content.png)

**2. Secrets that are exposed in the code**

![](../../../.gitbook/assets/github-code-scanning-secrets-content.png)

**3. Infrastructure as Code (Iac) issues on Terraform packages**

![](../../../.gitbook/assets/github-code-scanning-iac-content.png)

**4. Static Application Security Testing (Sast) vulnerabilities**

![](../../../.gitbook/assets/github-code-scanning-sast-content.png)

**5. Validate Allowed Licenses**

When Frogbot scans the repository periodically, it checks the licenses of any project dependencies. If Frogbot identifies licenses that are not listed in a predefined set of approved licenses, it adds an alert. The list of allowed licenses is set up as a variable within the Frogbot workflow.

![](../../../.gitbook/assets/github-code-scanning-license-violation-content.png)
