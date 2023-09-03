# Scan Pull Requests

Frogbot uses [JFrog Xray](https://jfrog.com/xray/) (version 3.29.0 and above is required) to scan your pull requests. It adds the scan results as a comment on the pull request. If no new vulnerabilities are found, Frogbot will also add a comment, confirming this.

The following features use the package manager used for building the project:

* Software Composition Analysis (SCA)
* Vulnerability Contextual Analysis

The supported package managers are:

* Go
* Gradle
* Maven
* .NET
* npm
* NuGet
* Pip
* Pipenv
* Poetry
* Yarn

#### How to use Pull Request scanning?

#### 👮 Security note for pull requests scanning

When installing Frogbot using JFrog Pipelines, Jenkins, and Azure DevOps, Frogbot will not wait for a maintainer's approval before scanning newly opened pull requests. Using Frogbot with these platforms is therefore not recommended for open-source projects.

When installing Frogbot using GitHub Actions and GitLab however, Frogbot will initiate the scan only after it is approved by a maintainer of the project. The goal of this review is to ensure that external code contributors don't introduce malicious code as part of the pull request. Since this review step is enforced by Frogbot when used with GitHub Actions and GitLab, it is safe to be used for open-source projects.

#### Scan results

**Software Composition Analysis (SCA), Vulnerability Contextual Analysis and Infrastructure as Code scans (IaC)**

Frogbot adds the scan results to the pull request in the following format:

**👍 No issues**

If no new vulnerabilities are found, Frogbot automatically adds the following comment to the pull request:

[![](https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/noVulnerabilityBannerPR.png)](broken-reference)

**👎 Issues were found**

If new vulnerabilities are found, Frogbot adds them as a comment on the pull request. For example:

[![](https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/vulnerabilitiesBannerPR.png)](broken-reference)

**VULNERABLE DEPENDENCIES**

**INFRASTRUCTURE AS CODE**

**Secrets Detection**

When Frogbot detects secrets that have been inadvertently exposed within the code of a pull request, it promptly triggers an email notification to the user who pushed the corresponding commit. The email address utilized for this notification is sourced from the committer's Git profile configuration. Moreover, Frogbot offers the flexibility to direct email notification to an extra email address if desired. To activate email notifications, it is necessary to configure your SMTP server details as variables within your Frogbot workflows.

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/secrets-email.png" alt="" data-size="original">
