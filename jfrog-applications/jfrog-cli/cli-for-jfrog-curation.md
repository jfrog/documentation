# JFrog CLI curation-audit command (for JFrog Curation project audit)

## Overview

JFrog Curation defends your software supply chain, enabling early blocking of malicious or risky open-source packages before they even enter. Seamlessly identify harmful, vulnerable, or risky packages, ensuring increased security, compliance, and developer productivity.

For more information see: https://jfrog.com/curation/

"The 'curation-audit' is a JFrog CLI command designed for developers to scan their projects and identify third-party dependencies that violate the restrictions set by the Curation service. This command provides detailed insights into the specific package policies that are being violated, leading to their blockage by the Curation service. Additionally, when feasible, 'curation-audit' may suggest alternative versions of the packages that comply with the Curation policies."

## Supported package managers & build systems

For a full list of the package managers and build systems supported by the curation-audit command and the required Artifactory and Xray versions to use it please see: https://jfrog.com/help/r/jfrog-curation/jfrog-curation-support-matrix

curation-audit command supported package managers and build systems:
* Npm (npm)
* Maven (mvn) - Requires xray 3.92 and above, and Artifactory 7.82 and above
* Pip (pip) - Requires xray 3.92 and above, and Artifactory 7.82 and above
***

### Commands

Audit your Project with JFrog CLI curation-audit command

### Setup:

Prerequisites:

Make sure your JFrog Artifactory admin configured the curated remote repository you are using during your build process. For more information refer your Artifactory admin to:
https://jfrog.com/help/r/jfrog-curation/configure-curation-pass-through

1. **Connect JFrog CLI to JFrog Platform**

   Connect the JFrog CLI to your JFrog Platform instance by running the following command:

    ```
    jf c add
    ```

   - When prompted for the access token, use the token generated from Artifactory. For more details, refer to the [JFrog CLI documentation](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration#adding-and-editing-configured-servers).

    ```
    jf c show
    ```

   - It should present Artifactory server just added (with default true)
     </br></br>
2. **Configure JFrog CLI for Project**</br>
   Ensure your project is configured in the JFrog CLI with the repository you would like to resolve dependencies from. Here are details for each package manager:

   - **NPM:**

      - Set the resolved repository using the [**jf npmc**](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory/package-managers-integration#setting-npm-repositories) command inside the project directory.

   - **MAVEN:**

      - Set the resolved repository using the [**jf mvnc**](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory/package-managers-integration#setting-maven-repositories) command inside the project directory.
   
   - **PIP:**

       - Set the resolved repository using the [**jf pipc**](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory/package-managers-integration#setting-python-repository) command inside the project directory (The only package installer supported for now by Python is "pip").

#### Commands Params

|                     |                                                                                                                                        |
| ------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| **Command name**    | curation-audit                                                                                                                         |
| **Abbreviation**    | ca                                                                                                                                     |
| **Command options** |                                                                                                                                        |
| --format            | <p>[Default: table]<br><br>Defines the output format of the command. Acceptable values are: table and json.</p>                        |
| --working-dirs      | <p>[Optional]<br><br>A comma separated list of relative working directories, to determine the audit targets locations.</p>             |
| --threads           | <p>[Default: 10]<br><br>The number of parallel threads used to determine the curation status for each package in the project tree.</p> |
| --requirements-file | <p>[Optional] [Pip]<br><br>Defines pip requirements file name. For example: 'requirements.txt'</p>                                                                                                                                                                                                                                                                |


#### Example 1

Audit the project in the current directory. Displays all known packages that were blocked by Curation Policies.

```
jf curation-audit
```

#### Example 2

Audit the projects according to the specific paths defined in the "working-dirs" option. Displays all known packages that were blocked by Curation Policies for all projects. The data is displayed in separate tables.

```
jf curation-audit --working-dirs="/path/to/project/npm_project1,/path/to/project/npm_project2"
```

#### Example 3

Audit the project in the current directory using 5 threads to check the packages Curation status in parallel. Displays all known packages blocked by Curation Policies.

```
jf curation-audit --threads=5
```

