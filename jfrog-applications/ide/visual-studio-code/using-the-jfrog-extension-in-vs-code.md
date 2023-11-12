# Using the JFrog Extension in VS Code

The extension offers two modes, **Local** and **CI**. The two modes can be toggled by pressing on their respective buttons that will appear next to the components tree.

* The **Local** view displays information about the local code as it is being developed in VS Code. The developer can scan their local workspace continuously. The information is displayed in the **Local** view.
* The **CI** view allows the tracking of the code as it is built, tested and scanned by the CI server. It displays information about the status of the build and includes a link to the build log on the CI server.

#### Severity Icons

The icon demonstrates the top severity issue of a selected component and its transitive dependencies. The following table describes the severities from highest to lowest:

| Icon                                                                                                                                                                                                                                                     | Severity       | Description                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------- | ---------------------------------------------------- |
| ![](../../.gitbook/assets/Critical.png)                                                                                                                                                                                                                  | Critical       | Issue with critical severity                         |
| ![](../../.gitbook/assets/High.png)                                                                                                                                                                                                                      | High           | Issue with high severity                             |
| ![](../../.gitbook/assets/Medium.png)                                                                                                                                                                                                                    | Medium         | Issue with medium severity                           |
| ![](../../.gitbook/assets/Low.png)                                                                                                                                                                                                                       | Low            | Issue with low severity                              |
| ![](../../.gitbook/assets/Unknown.png)                                                                                                                                                                                                                   | Unknown        | Issue with unknown severity                          |
| ![](../../.gitbook/assets/notApplicableCritical.png)![](../../.gitbook/assets/notApplicableHigh.png)![](../../.gitbook/assets/notApplicableMedium.png)![](../../.gitbook/assets/notApplicableLow.png)![](../../.gitbook/assets/notApplicableUnknown.png) | Not Applicable | CVE issue that is not applicable to your source code |
| ![](../../.gitbook/assets/Normal.png)                                                                                                                                                                                                                    | Normal         | No issues (Used only in CI view)                     |

## The Local View

### General

The JFrog VS Code Extension enables continuous scans of your project with the JFrog Platform. The security related information will be displayed under the Local view. It allows developers to view vulnerability information about their dependencies and source code in their IDE. With this information, you can make an informed decision on whether to use a component or not before it gets entrenched into the organization’s product.

scan your workspace by clicking the Scan/Rescan button, the ![](../../.gitbook/assets/vscode/refreshBtn.png) icon at the extension tab or click on Start Xray Scan from within the editor. The scan will create a list of files with vulnerabilities in the workspace.

![Refresh](../../.gitbook/assets/vscode/refresh.png)

## The CI View

The CI view of the extension allows you to view information about your builds directly from your CI system. This allows developers to keep track of the status of their code, while it is being built, tested and scanned as part of the CI pipeline, regardless of the CI provider used.

This information can be viewed inside JFrog VS Code Extension, from the JFrog Panel, after switching to CI mode.

The following details can be made available in the CI view.

* Status of the build run (passed or failed)
* Build run start time
* Git branch and latest commit message
* Link to the CI run log
* Security information about the build artifacts and dependencies

### How Does It Work?

The CI information displayed in VS Code is pulled by the JFrog Extension directly from JFrog Artifactory. This information is stored in Artifactory as part of the build-info, which is published to Artifactory by the CI server.

Read more about build-info in the [Build Integration](https://www.jfrog.com/confluence/display/JFROG/Build+Integration) documentation page. If the CI pipeline is also configured to scan the build-info by JFrog Xray, the JFrog VS Code Extension will pull the results of the scan from JFrog Xray and display them in the CI view as well.

### Setting Up Your CI Pipeline

Before VS Code can display information from your CI in the CI View, your CI pipeline needs to be configured to expose this data. Read [this guide](https://www.jfrog.com/confluence/display/JFROG/Setting+Up+CI+Integration) which describes how to configure your CI pipeline.

### Setting Up the CI View

Set your CI build name in the Build name pattern field at the [Extension Settings](using-the-jfrog-extension-in-vs-code.md#extension-settings). This is the name of the build published to Artifactory by your CI pipeline. You have the option of setting \* to view all the builds published to Artifactory.

After your builds were fetched from Artifactory, press on the Builds ![Builds](../../.gitbook/assets/vscode/build.png) button to choose what build to display.

![CI](../../.gitbook/assets/vscode/ci.gif)
