# Audit your Source Code

The _**jf audit**_ command allows scanning your source code dependencies to find security vulnerabilities and licenses violations, with the ability to scan against your Xray policies. The command builds a deep dependencies graph for your project, scans it with Xray, and displays the results. It uses the package manager used by the project to build the dependencies graph. Currently, the following package managers are supported.

* Maven (mvn) - Version 3.1.0 or above of Maven is supported.
* Gradle (gradle)
* Npm (npm)
* Pnpm (pnpm)
* Yarn (yarn)
* Pip (pip)
* Pipenv (pipenv)
* Poetry (poetry)
* Go Modules (go)
* NuGet (nuget)
* .NET Core CLI (dotnet)

The command will detect the package manager used by the project automatically. It requires version 3.29.0 or above of Xray and also version 2.13.0 or above of JFrog CLI.

### Advanced Scans

This command also supports the following Advanced Scans with the **Advanced Security Package** enabled on the JFrog Platform instance. To enable the Advanced Security Package, contact us using [this](https://jfrog.com/advanced-security-contact-us/) form.

* **Vulnerability Contextual Analysis**: This feature uses the code context to eliminate false positive reports on vulnerable dependencies that are not applicable to the code. Vulnerability Contextual Analysis is currently supported for Python and JavaScript code.
* **Secrets Detection**: Detect any secrets left exposed inside the code. to stop any accidental leak of internal tokens or credentials.
* **Infrastructure as Code scans (IaC)**: Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

***

**Note**

* The **jf audit** command does not extract the internal content of the scanned dependencies. This means that if a package includes other vulnerable components bundled inside the binary, they may not be shown as part of the results. This is contrary to the **jf scan** command, which drills down into the package content.
* To generate the dependency tree for scanning purposes, the system will execute an 'install' command on the project if it hasn't been executed previously.

***

#### Commands Params

|                       |                                                                                                                                                                                                                                                                                                                                                               |
|-----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Command name**      | audit                                                                                                                                                                                                                                                                                                                                                         |
| **Abbreviation**      | aud                                                                                                                                                                                                                                                                                                                                                           |
| **Command options**   |                                                                                                                                                                                                                                                                                                                                                               |
| --server-id           | <p>[Optional]<br>Server ID configured using the <em>jf c add</em> command. If not specified, the default configured server is used.</p>                                                                                                                                                                                                                       |
| --project             | <p>[Optional]<br>JFrog project key, to enable Xray to determine security violations accordingly. The command accepts this option only if the --repo-path and --watches options are not provided. If none of the three options are provided, the command will show all known vulnerabilities</p>                                                               |
| --repo-path           | <p>[Optional]<br>Artifactory repository path in the form of &#x3C;repository>/&#x3C;path in the repository>, to enable Xray to determine violations accordingly. The command accepts this option only if the --project and --watches options are not provided. If none of the three options are provided, the command will show all known vulnerabilities</p> |
| --watches             | <p>[Optional]<br>A comma-separated list of Xray watches, to enable Xray to determine violations accordingly. The command accepts this option only if the --repo-path and --repo-path options are not provided. If none of the three options are provided, the command will show all known vulnerabilities</p>                                                 |
| --licenses            | <p>[Default: false]<br>Set if you'd also like the list of licenses to be displayed.</p>                                                                                                                                                                                                                                                                       |
| --format              | <p>[Default: table]<br>Defines the output format of the command. Acceptable values are: table and json.</p>                                                                                                                                                                                                                                                   |
| --fail                | <p>[Default: true]<br>Set to false if you do not wish the command to return exit code 3, even if the 'Fail Build' rule is matched by Xray.</p>                                                                                                                                                                                                                |
| --use-wrapper         | <p>[Default: false] [Gradle]<br>Set to true if you'd like to use the Gradle wrapper.</p>                                                                                                                                                                                                                                                                      |
| --dep-type            | <p>[Default: all] [npm]<br>Defines npm dependencies type. Possible values are: all, devOnly and prodOnly</p>                                                                                                                                                                                                                                                  |
| --exclude-test-deps   | <p>[Default: false] [Gradle]<br>Set to true if you'd like to exclude Gradle test dependencies from Xray scanning.</p>                                                                                                                                                                                                                                         |
| --requirements-file   | <p>[Optional] [Pip]<br>Defines pip requirements file name. For example: 'requirements.txt'</p>                                                                                                                                                                                                                                                                |
| --working-dirs        | <p>[Optional]<br>A comma-separated list of relative working directories, to determine the audit targets locations.</p> If flag isn't provided, a recursive scan is triggered from the root directory of the project.                                                                                                                                          |
| --exclusions          | <p>[Default: *.git*;*node_modules*;*target*;*venv*;*test*]<br>List of Semicolon-separated exclusions, utilized to skip sub-projects from undergoing an audit. These exclusions may incorporate the * and ? wildcards.</p>                                                                                                                                     |
| --fixable-only        | <p>[Optional]<br>Set to true if you wish to display issues that have a fix version only.</p>                                                                                                                                                                                                                                                                  |
| --min-severity        | <p>[Optional]<br>Set the minimum severity of issues to display. The following values are accepted: Low, Medium, High or Critical</p>                                                                                                                                                                                                                          |
| --go                  | <p>[Default: false]<br>Set to true to request audit for a Go project.</p>                                                                                                                                                                                                                                                                                     |
| --gradle              | <p>[Default: false]<br>Set to true to request audit for a Gradle project.</p>                                                                                                                                                                                                                                                                                 |
| --mvn                 | <p>[Default: false]<br>Set to true to request audit for a Maven project.</p>                                                                                                                                                                                                                                                                                  |
| --npm                 | <p>[Default: false]<br>Set to true to request audit for a npm project.</p>                                                                                                                                                                                                                                                                                    |
| --pnpm                | <p>[Default: false]<br>Set to true to request audit for a pnpm project.</p>                                                                                                                                                                                                                                                                                   |
| --nuget               | <p>[Default: false]<br>Set to true to request audit for a .Net project.</p>                                                                                                                                                                                                                                                                                   |
| --pip                 | <p>[Default: false]<br>Set to true to request audit for a Pip project.</p>                                                                                                                                                                                                                                                                                    |
| --pipenv              | <p>[Default: false]<br>Set to true to request audit for a Pipenv project.</p>                                                                                                                                                                                                                                                                                 |
| --yarn                | <p>[Default: false]<br>Set to true to request audit for a Yarn project.</p>                                                                                                                                                                                                                                                                                   |
| **Command arguments** | The command accepts no arguments                                                                                                                                                                                                                                                                                                                              |

#### **Output Example**

#### Example 1

Audit the project at the current directory. Show all known vulnerabilities, regardless of the policies defined in Xray.

```
jf audit
```

#### Example 2

Audit the project at the current directory. Show all known vulnerabilities, regardless of the policies defined in Xray. Show only maven and npm vulnerabilities.

```
jf audit --mvn --npm
```

#### Example 3

Audit the project at the current directory using a watch named _watch1_ watch defined in Xray.

```
jf audit --watches "watch1"
```

#### Example 4

Audit the project at the current directory using _watch1_ and \_watch2\_defined in Xray.

```
jf audit --watches "watch1,watch2"
```

#### Example 5

Audit the project at the current directory using the policies defined for project-1.

```
jf audit --project "project-1"
```

#### Example 6

Audit the project at the current directory using the policies defined for the _libs-local/release-artifacts/_ path in Artifactory.

```
jf audit --repo-path "libs-local/release-artifacts/"
```

#### Example 7

Audit the project in the current directory, excluding all files inside the _node\_modules_ directory and files with the _to\_exclude_ suffix.

```
jf audit --exclusions "*node_modules*;*to_exclude"
```