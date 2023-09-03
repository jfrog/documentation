# Analyze your Results

#### Viewing vulnerabilities

The JFrog extension incorporates a file tree displaying all the vulnerabilities within the project. Each file that is infected with a vulnerability appears as a tree node.

Descriptor file (e.g., pom.xml in Maven, go.mod in Go, etc.) has a special meaning that outlines the available direct dependencies for the project. The tree will show these descriptor files containing vulnerable dependencies. In cases where a direct dependency contains vulnerable child dependencies, the tree will show the vulnerable child dependencies instead, denoting them with a '(indirect)' postfix.

Furthermore, various types of vulnerability nodes, such as Contextual Analysis Vulnerabilities or hard-coded secrets, may be present in other source code files.

Each file node in the tree is interactive, click and expand it to view its children node and navigate to the corresponding file in the IDE for better visibility. Upon navigating to a file, the extension will highlight the vulnerable line, making it easier to locate the specific issue

In addition, the locations with vulnerabilities will be marked in the editor. By clicking on the light bulb icon next to a vulnerable location in the editor, we can instantly jump to the corresponding entry in the tree view.

<figure><img src="../../../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

Clicking on a CVE in the list will open the location with the issue in the editor and a vulnerability details view. This view contains information about the vulnerability, the vulnerable component, fixed versions, impact paths, and much more.

<figure><img src="../../../.gitbook/assets/image (4).png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (5).png" alt=""><figcaption></figcaption></figure>

<details>

<summary>CVE Research and Enrichment</summary>

For selected security issues, get leverage-enhanced CVE data that is provided by our JFrog Security Research team. Prioritize the CVEs based on:

* JFrog Severity: The severity given by the JFrog Security Research team after the manual analysis of the CVE by the team. CVEs with the highest JFrog security severity are the most likely to be used by real-world attackers. This means that you should put effort into fixing them as soon as possible.
* Research Summary: The summary that is based on JFrog's security analysis of the security issue provides detailed technical information on the specific conditions for the CVE to be applicable. Remediation: Detailed fix and mitigation options for the CVEs

Check out what our research team is up to and stay updated on newly discovered issues by clicking on this [link](https://research.jfrog.com/).



</details>

<details>

<summary>Contextual Analysis</summary>

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced Security._

Xray automatically validates some high and very high-impact vulnerabilities, such as vulnerabilities that have prerequisites for exploitations, and provides contextual analysis information for these vulnerabilities, to assist you in figuring out which vulnerabilities need to be fixed. Contextual Analysis data includes:

* Contextual analysis status: Contextual analysis results indicate if a CVE was found applicable in your application or not applicable.
* Contextual Analysis breakdown: An explanation provided by our research team as to why the CVE was found applicable or not applicable.
* Remediation: Contextual mitigation steps and options provided by our research team that assist you with remediating the issues.



</details>

<details>

<summary>Secrets Detection</summary>

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced Security._

Detect any secrets left exposed inside the code to prevent any accidental leak of internal tokens or credentials.

![](broken-reference)

</details>

<details>

<summary>Infrastructure as Code (IaC) Scan</summary>

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced Security._

Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

![](broken-reference)

</details>
