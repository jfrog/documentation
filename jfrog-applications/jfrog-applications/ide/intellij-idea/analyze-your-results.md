# Analyze your Results

<details>

<summary>Software Composition Analysis (SCA)</summary>



Each descriptor file (like pom.xml in Maven, go.mod in Go, etc.) displayed in the JFrog Panel contains vulnerable dependencies, and each dependency contains the vulnerabilities themselves.

By right-clicking on a dependency line, you can jump to the dependency's declaration in the descriptor file or have the depedency upgraded to a version with a fix.

You can also create an [Ignore Rule](https://jfrog.com/help/r/jfrog-security-documentation/ignore-rules) in Xray.

_Creating Ignore Rules is only available_ [_when a JFrog Project or Watch is defined_](broken-reference)_._

Clicking a vulnerability in the list will open the vulnerability details view. This view contains information about the vulnerability, the vulnerable component, fixed versions, impact paths and much more.

</details>

<details>

<summary>CVEs Contextual Analysis</summary>

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced DevSecOps._

Xray automatically validates some high and very high impact vulnerabilities, such as vulnerabilities that have prerequisites for exploitations, and provides contextual analysis information for these vulnerabilities, to assist you in figuring out which vulnerabilities need to be fixed.

CVEs Contextual Analysis data includes:

* **Contextual Analysis status**: Contextual Analysis results indicating if a CVE was found applicable in your application or not applicable.
* **Contextual Analysis breakdown**: An explanation provided by our research team as to why the CVE was found applicable or not applicable.
* **Remediation**: Contextual mitigation steps and options provided by our research team that assist you with remediating the issues.

</details>

<details>

<summary>Secrets Detection</summary>

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced DevSecOps._

Detect any secrets left exposed inside the code. to prevent any accidental leak of internal tokens or credentials.

**NOTE:** To ignore detected secrets, you can add a comment which includes the phrase **jfrog-ignore** above the line with the secret.

</details>

<details>

<summary>Infrastructure as Code (IaC) Scan</summary>

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced DevSecOps._

Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

</details>

**Severity Icons**

| Icon |              Severity              |
| :--: | :--------------------------------: |
|      |              Critical              |
|      |                High                |
|      |               Medium               |
|      |                 Low                |
|      |               Unknown              |
|      | [Not Applicable](broken-reference) |

###
