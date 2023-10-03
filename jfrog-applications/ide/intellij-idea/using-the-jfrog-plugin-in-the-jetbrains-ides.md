# Using the JFrog Plugin in the JetBrains IDEs

After the JFrog Plugin is installed, a new JFrog panel is added at the bottom of the screen. Opening the JFrog panel displays two views:

* The **Local** view displays information about the local code as it is being developed in the IDE. You can continuously scan your project locally. The information is displayed in the **Local** view.
* The **CI** view allows the tracking of the code as it is built, tested and scanned by the CI server. It displays information about the status of the build and includes a link to the build log on the CI server.

### The Local View

#### General

The JFrog Plugin enables continuous scans of your project with the JFrog Platform. The security-related information will be displayed under the Local view. It allows developers to view vulnerability information about their dependencies and source code in their IDE. With this information, a developer can make an informed decision on whether to use a component or not before it gets entrenched into the organization’s product.

Scan your project by clicking the Run Scan button. After the scan is done, a list of vulnerable files will appear.

#### Software Composition Analysis (SCA)

Each descriptor file (like pom.xml in Maven, go.mod in Go, etc.) displayed in the JFrog Panel contains vulnerable dependencies, and each dependency contains the vulnerabilities themselves.

By right-clicking on a dependency line, you can jump to the dependency's declaration in the descriptor file or have the dependency upgraded to a version with a fix.

<figure><img src="../../.gitbook/assets/jump-to-descriptor.png" alt=""><figcaption></figcaption></figure>

You can also create an [Ignore Rule](https://jfrog.com/help/r/jfrog-security-documentation/ignore-rules) in Xray.

> _**NOTE:**_ Creating Ignore Rules is only available\_when a JFrog Project or Watch is defined.

<div align="left">

<figure><img src="../../.gitbook/assets/create-ignore-rule.png" alt=""><figcaption></figcaption></figure>

</div>

Clicking a vulnerability in the list will open the vulnerability details view. This view contains information about the vulnerability, the vulnerable component, fixed versions, impact paths and much more.

<figure><img src="../../.gitbook/assets/vuln-details.png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../.gitbook/assets/vuln-impact-graph.png" alt=""><figcaption></figcaption></figure>

#### CVEs Contextual Analysis

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced DevSecOps._

Xray automatically validates some high and very high-impact vulnerabilities, such as vulnerabilities that have prerequisites for exploitations, and provides contextual analysis information for these vulnerabilities, to assist you in figuring out which vulnerabilities need to be fixed.

CVEs Contextual Analysis data includes:

* **Contextual Analysis status**: Contextual Analysis results indicate if a CVE was found applicable in your application or not applicable.
* **Contextual Analysis breakdown**: An explanation provided by our research team as to why the CVE was found applicable or not applicable.
* **Remediation**: Contextual mitigation steps and options provided by our research team that assist you with remediating the issues.

<figure><img src="../../.gitbook/assets/not-applicable.png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../.gitbook/assets/applicable.png" alt=""><figcaption></figcaption></figure>

#### Secrets Detection

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced DevSecOps._

Detect any secrets left exposed inside the code. to prevent any accidental leak of internal tokens or credentials.

> **NOTE:** To ignore detected secrets, you can add a comment which includes the phrase **jfrog-ignore** above the line with the secret.

<figure><img src="../../.gitbook/assets/secrets (1).png" alt=""><figcaption></figcaption></figure>

#### Infrastructure as Code (IaC) Scan

_Requires Xray version 3.66.5 or above and Enterprise X / Enterprise+ subscription with Advanced DevSecOps._

Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

<figure><img src="../../.gitbook/assets/iac (1).png" alt=""><figcaption></figcaption></figure>

#### Severity Icons

The icon demonstrates the top severity issue of a selected component and its transitive dependencies. The following table describes the severities from highest to lowest:

|                                                                                                                                   Icon                                                                                                                                  |    Severity    |
| :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: | :------------: |
|                                                                                                                ![](../../.gitbook/assets/Critical.png)                                                                                                               |    Critical    |
|                                                                                                                  ![](../../.gitbook/assets/High.png)                                                                                                                 |      High      |
|                                                                                                                 ![](../../.gitbook/assets/Medium.png)                                                                                                                |     Medium     |
|                                                                                                                  ![](../../.gitbook/assets/Low.png)                                                                                                                  |       Low      |
|                                                                                                                ![](../../.gitbook/assets/Unknown.png)                                                                                                                |     Unknown    |
| ![](../../.gitbook/assets/notApplicableCritical.png)![](../../.gitbook/assets/notApplicableHigh.png)![](../../.gitbook/assets/notApplicableLow.png)![](../../.gitbook/assets/notApplicableMedium.png)![](../../.gitbook/assets/notApplicableUnknown.png) | Not Applicable |

### The CI View

The JFrog Plugin allows you to view information about your builds directly from your CI system. This allows developers to keep track of the status of their code, while it is being built, tested, and scanned as part of the CI pipeline, regardless of the CI provider used.

This information can be viewed inside a Jetbrains IDE, from the **JFrog** Panel, under the **CI** tab.

The following details can be made available in the CI view:

* Status of the build run (passed or failed)
* Build run start time
* Git branch and latest commit message
* Link to the CI run log
* Security information about the build artifacts and dependencies

<figure><img src="../../.gitbook/assets/ci-view.png" alt=""><figcaption></figcaption></figure>

#### How Does It Work?

The CI information displayed in IDEA is pulled by the JFrog IDEA Plugin directly from JFrog Artifactory. This information is stored in Artifactory as part of the build-info, which is published to Artifactory by the CI server. Read more about build-info in the [Build Integration](https://jfrog.com/help/r/jfrog-integrations-documentation/build-integration) documentation page. If the CI pipeline is also configured to scan the build-info by JFrog Xray, the JFrog IDEA Plugin will pull the results of the scan from JFrog Xray and display them in the CI view as well.

#### Setting Up CI Integration

Set up your CI pipeline to expose information, so that it is visible in IDEA as described [here](https://jfrog.com/help/r/jfrog-integrations-documentation/setting-up-ci-integration).

Next, follow these steps:

1. Under **Settings (Preferences)** | **Other Settings**, click **JFrog Global Configuration**. configure the JFrog Platform URL and the user you created.
2. Under **Settings (Preferences)** | **Other Settings**, click **JFrog CI Integration**. Set your CI build name in the **Build name pattern** field. This is the name of the build published to Artifactory by your CI pipeline. You have the option of setting \* to view all the builds published to Artifactory.\
   ![](../../.gitbook/assets/ci-settings.png)
3. Click **Apply** and open the **CI** tab under the **JFrog** panel at the bottom of the screen and click the **Refresh** button.
