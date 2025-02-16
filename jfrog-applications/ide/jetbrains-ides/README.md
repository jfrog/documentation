# JetBrains IDEs

![](../../.gitbook/assets/jetbrains-ides-header.png)

## About the JFrog Plugin

The plugin allows developers to find and fix security vulnerabilities in their projects and to see valuable information about the status of their code by continuously scanning it locally with [JFrog Security](https://jfrog.com/xray/).

#### What security capabilities do we provide?

**Basic**

**Software Composition Analysis (SCA)**

Scans your project dependencies for security issues and shows you which dependencies are vulnerable. If the vulnerabilities have a fix, you can upgrade to the version with the fix in a click of a button.

**CVE Research and Enrichment**

or selected security issues, get leverage-enhanced CVE data that is provided by our JFrog Security Research team. Prioritize the CVEs based on:

* **JFrog Severity**: The severity given by the JFrog Security Research team after the manual analysis of the CVE by the team. CVEs with the highest JFrog security severity are the most likely to be used by real-world attackers. This means that you should put effort into fixing them as soon as possible.
* **Research Summary**: The summary that is based on JFrog's security analysis of the security issue provides detailed technical information on the specific conditions for the CVE to be applicable.
* **Remediation**: Detailed fix and mitigation options for the CVEs

You can learn more about enriched CVEs [here](https://jfrog.com/help/r/jfrog-security-documentation/jfrog-security-cve-research-and-enrichment).

Check out what our research team is up to and stay updated on newly discovered issues by clicking on this link: [https://research.jfrog.com](https://research.jfrog.com)

**Advanced**

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with_ [_Advanced DevSecOps_](https://jfrog.com/xray/#xray-advanced)_._&#x20;

**CVEs Contextual Analysis**

Uses the code context to eliminate false positive reports on vulnerable dependencies that are not applicable to the code. CVEs Contextual Analysis is currently supported for Python, Java and JavaScript code.

**Secrets Detection**

Prevents the exposure of keys or credentials that are stored in your source code.

**Infrastructure as Code (IaC) Scan**

Secures your IaC files. Critical to keeping your cloud deployment safe and secure.

**Additional Perks**

* Security issues are easily visible inline.
* The results show issues with context, impact, and remediation.
* View all security issues in one place, in the JFrog tab.
* For Security issues with an available fixed version, you can upgrade to the fixed version within the plugin.
* Track the status of the code while it is being built, tested, and scanned on the CI server.

The JFrog Plugin supports the following IDEs:

* IntelliJ IDEA
* WebStorm
* PyCharm
* Android Studio
* GoLand
* Rider
* CLion

1. Install the JFrog Plugin via the Plugins tab in the IDE settings, or in [JetBrains Marketplace](https://plugins.jetbrains.com/plugin/9834-jfrog).
2. [Connect IDEA to your JFrog environment](connect-the-jfrog-plugin-to-the-jfrog-platform.md).
3. [Start](using-the-jfrog-plugin-in-the-jetbrains-ides.md) using the plugin.
