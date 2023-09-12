# Visual Studio Code

### About the JFrog Extension;

The cost of remediating a vulnerability is akin to the cost of fixing a bug. The earlier you remediate a vulnerability in the release cycle, the lower the cost. The extension allows developers to find and fix security vulnerabilities in their projects and to see valuable information about the status of their code by continuously scanning it locally with [JFrog Xray](https://jfrog.com/xray/).&#x20;

#### What security capabilities do we provide?&#x20;

**Software Composition Analysis (SCA)**

Scan your project dependencies for security issues. For selected security issues, get leverage-enhanced CVE data that is provided by our JFrog Security Research team. To learn more about enriched CVEs, see [here](https://jfrog.com/help/r/jfrog-security-documentation/jfrog-security-cve-research-and-enrichment).&#x20;

#### Advanced Security&#x20;

Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced Security.&#x20;

**Contextual Analysis**&#x20;

With advanced Contextual Analysis, understand the applicability of CVEs in your application and utilize JFrog Security scanners to analyze the way you use 3rd party packages in your projects. Automatically validate some high-impact vulnerabilities, such as vulnerabilities that have prerequisites for exploitations, and reduce false positives and vulnerability noise with smart CVE analysis. To learn more, see [here](https://jfrog.com/help/r/jfrog-security-documentation/vulnerability-contextual-analysis).&#x20;

**Infrastructure as Code (IaC) Scan**&#x20;

* Analyze Infrastructure as Code (IaC) files, such as Terraform, to identify security vulnerabilities and misconfigurations before deploying your cloud infrastructure.
* Get actionable insights and recommendations for securing your IaC configurations.&#x20;

**Secrets Detection**&#x20;

Detect and prevent the inclusion of sensitive information, such as credentials and API keys, in your codebase.&#x20;

**Supported Packages**

| Feature                                           | Go | Maven | npm | Yarn v1 | PypI | .NET | Terraform |
| ------------------------------------------------- | -- | ----- | --- | ------- | ---- | ---- | --------- |
| SCA                                               | ✅  | ✅     | ✅   | ✅       | ✅    | ✅    | ❌         |
| Upgrade vulnerable dependencies to fixed versions | ✅  | ✅     | ✅   | ✅       | ✅    | ✅    | ❌         |
| Contextual Analysis                               | ❌  | ❌     | ✅   | ✅       | ✅    | ❌    | ❌         |
| Secrets Detection                                 | ✅  | ✅     | ✅   | ✅       | ✅    | ✅    | ✅         |
| Exclude dev dependencies                          | ❌  | ❌     | ✅   | ❌       | ❌    | ❌    | ❌         |
| Infrastructure as Code (IaC)                      | ❌  | ❌     | ❌   | ❌       | ❌    | ❌    | ✅         |

**Additional Perks**

* Security issues are easily visible inline.
* The results show issues with context, impact, and remediation.
* View all security issues in one place, in the JFrog tab.
* For Security issues with an available fixed version, you can upgrade to the fixed version within the plugin.
* Track the status of the code while it is being built, tested, and scanned on the CI server.

The extension also applies [JFrog File Spec JSON schema](https://raw.githubusercontent.com/jfrog/jfrog-cli/master/schema/filespec-schema.json) on the following file patterns: `**/filespecs/*.json`, `*filespec*.json` and `*.filespec`. Read more about JFrog File specs [here](https://www.jfrog.com/confluence/display/JFROG/FileSpec).
