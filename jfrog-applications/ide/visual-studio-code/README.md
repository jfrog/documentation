# Visual Studio Code

![](../../.gitbook/assets/vscode-header.png)

### About the JFrog Extension;

The cost of remediating a vulnerability is akin to the cost of fixing a bug. The earlier you remediate a vulnerability in the release cycle, the lower the cost. The extension allows developers to find and fix security vulnerabilities in their projects and to see valuable information about the status of their code by continuously scanning it locally with [JFrog Xray](https://jfrog.com/xray/).

#### What security capabilities do we provide?

**Software Composition Analysis (SCA)**

Scan your project dependencies for security issues. For selected security issues, get leverage-enhanced CVE data that is provided by our JFrog Security Research team. To learn more about enriched CVEs, see [here](https://jfrog.com/help/r/jfrog-security-documentation/jfrog-security-cve-research-and-enrichment).

#### Advanced Security

Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced Security.

**Contextual Analysis**

With advanced Contextual Analysis, understand the applicability of CVEs in your application and utilize JFrog Security scanners to analyze the way you use 3rd party packages in your projects. Automatically validate some high-impact vulnerabilities, such as vulnerabilities that have prerequisites for exploitations, and reduce false positives and vulnerability noise with smart CVE analysis. To learn more, see [here](https://jfrog.com/help/r/jfrog-security-documentation/vulnerability-contextual-analysis).

**Infrastructure as Code (IaC) Scan**

* Analyze Infrastructure as Code (IaC) files, such as Terraform, to identify security vulnerabilities and misconfigurations before deploying your cloud infrastructure.
* Get actionable insights and recommendations for securing your IaC configurations.

**Secrets Detection**

Detect and prevent the inclusion of sensitive information, such as credentials and API keys, in your codebase.

**Supported Packages**

<table><thead><tr><th width="257">Feature</th><th>Go</th><th width="100">Maven</th><th>npm</th><th>Yarn v1</th><th>Yarn v2</th><th>Pip</th><th>Pipenv</th><th>Poetry</th><th>.NET CLI</th><th>NuGet</th><th>Terraform</th></tr></thead><tbody><tr><td>SCA</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>❌</td></tr><tr><td>Contextual Analysis</td><td>❌</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>❌</td><td>❌</td><td>❌</td></tr><tr><td>Secrets Detection</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td></tr><tr><td>SAST</td><td>❌</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>❌</td><td>❌</td><td>❌</td></tr><tr><td>Exclude dev dependencies</td><td>❌</td><td>❌</td><td>✅</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td></tr><tr><td>Infrastructure as Code (IaC)</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>❌</td><td>✅</td></tr><tr><td>Autofix for direct dep.</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>❌</td></tr><tr><td>License Violations</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td><td>✅</td></tr></tbody></table>

**Additional Perks**

* Security issues are easily visible inline.
* The results show issues with context, impact, and remediation.
* View all security issues in one place, in the JFrog tab.
* For Security issues with an available fixed version, you can upgrade to the fixed version within the plugin.
* Track the status of the code while it is being built, tested, and scanned on the CI server.

The extension also applies [JFrog File Spec JSON schema](https://raw.githubusercontent.com/jfrog/jfrog-cli/master/schema/filespec-schema.json) on the following file patterns: `**/filespecs/*.json`, `*filespec*.json` and `*.filespec`. Read more about JFrog File specs [here](https://www.jfrog.com/confluence/display/JFROG/FileSpec).
