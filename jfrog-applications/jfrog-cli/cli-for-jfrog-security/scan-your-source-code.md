# Scan your Source Code

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
* CocoaPods (pod)
* SwiftPM (swift)

The command will detect the package manager used by the project automatically. It requires version 3.29.0 or above of Xray and also version 2.13.0 or above of JFrog CLI.

### Advanced Scans

This command also supports the following Advanced Scans with the **Advanced Security Package** enabled on the JFrog Platform instance. To enable the Advanced Security Package, contact us using [this](https://jfrog.com/advanced-security-contact-us/) form.

* **Vulnerability Contextual Analysis**: This feature uses the code context to eliminate false positive reports on vulnerable dependencies that are not applicable to the code. Vulnerability Contextual Analysis is currently supported for Python, Go and JavaScript code.
* **Secrets Detection**: Detect any secrets left exposed inside the code. to stop any accidental leak of internal tokens or credentials.
* **Infrastructure as Code scans (IaC)**: Scan Infrastructure as Code (Terraform) files for early detection of cloud and infrastructure misconfigurations.

***

**Note**

* The **jf audit** command does not extract the internal content of the scanned dependencies. This means that if a package includes other vulnerable components bundled inside the binary, they may not be shown as part of the results. This is contrary to the **jf scan** command, which drills down into the package content.
* To generate the dependency tree for scanning purposes, the system will execute an `install` command on the project if it hasn't been executed previously.

***

#### Commands Params

|                                 |                                                                  |                                                                                                                                                                                                                                                                                                                                            |
| ------------------------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Command name**                | **Default**                                                      | **Description**                                                                                                                                                                                                                                                                                                                            |
| `audit, aud`                    |                                                                  | Executes all following sub-scans: SCA (Software Composition Analysis) + Contextual Analysis, SAST, IaC, and Secrets.                                                                                                                                                                                                                       |
| **Command options**             |                                                                  |                                                                                                                                                                                                                                                                                                                                            |
| `--help`                        |                                                                  | Displays information about each of the `jf audit` command's options.                                                                                                                                                                                                                                                                       |
| `--server-id`                   | Default server                                                   | <p>[Optional]<br>Server ID configured using the <em><code>jf c add</code></em> command.</p>                                                                                                                                                                                                                                                |
| `--project`                     |                                                                  | <p>[Optional]<br>JFrog project key, to enable Xray to determine security violations accordingly. The command accepts this option only if the <code>--repo-path</code> and <code>--watches</code> options are not provided. If none of the three options are provided, the command will show all known vulnerabilities</p>                  |
| `--repo-path`                   |                                                                  | <p>[Optional]<br>Artifactory repository path, to enable Xray to determine violations accordingly. The command accepts this option only if the <code>--project</code> and <code>--watches</code> options are not provided. If none of the three options are provided, the command will show all known vulnerabilities</p>                   |
| `--watches`                     |                                                                  | <p>[Optional]<br>A comma-separated(,) list of Xray watches, to enable Xray to determine violations accordingly. The command accepts this option only if the <code>--repo-path</code> and <code>--repo-path</code> options are not provided. If none of the three options are provided, the command will show all known vulnerabilities</p> |
| `--licenses`                    | `false`                                                          | Set if you'd also like the list of licenses to be displayed.                                                                                                                                                                                                                                                                               |
| `--format`                      | `table`                                                          | Defines the output format of the command. Acceptable values are: `table` and `json`.                                                                                                                                                                                                                                                       |
| `--fail`                        | `true`                                                           | When using one of the flags `--watches`, `--project,` or `--repo-path` and a Fail build rule is matched the command will return exit code 3. Set to `false` if you want to see violations with exit code `0`.                                                                                                                              |
| `--use-wrapper`                 | `false`                                                          | <p>[Gradle]<br>Set to <code>true</code> if you'd like to use the Gradle wrapper.</p>                                                                                                                                                                                                                                                       |
| `--dep-type`                    | `all`                                                            | <p>[npm]<br>Defines npm dependency types. Acceptable values are: <code>all</code>, <code>devOnly</code> and <code>prodOnly</code></p>                                                                                                                                                                                                      |
| `--exclude-test-deps`           | `false`                                                          | <p>[Gradle]<br>Set to <code>true</code> if you'd like to exclude Gradle test dependencies from Xray scanning.</p>                                                                                                                                                                                                                          |
| `--requirements-file`           |                                                                  | <p>[Optional] [Pip]<br>Defines pip requirements file name. For example: '<code>requirements.txt</code>'</p>                                                                                                                                                                                                                                |
| `--working-dirs`                |                                                                  | <p>[Optional]<br>A comma-separated(,) list of relative working directories, to determine the audit targets locations.</p><p>If flag isn't provided, a recursive scan is triggered from the root directory of the project.</p>                                                                                                              |
| `--exclusions`                  | `.`_`git`_`;`_`node_modules`_`;`_`target`_`;`_`venv`_`;`_`test`_ | List of semicolon-separated(;) exclusions, utilized to skip sub-projects from undergoing an audit. These exclusions may incorporate the `*` and `?` wildcards.                                                                                                                                                                             |
| `--fixable-only`                |                                                                  | <p>[Optional]<br>Set to true if you wish to display issues that have a fix version only.</p>                                                                                                                                                                                                                                               |
| `--min-severity`                |                                                                  | <p>[Optional]<br>Set the minimum severity of issues to display. Acceptable values: <code>Low</code>, <code>Medium</code>, <code>High</code>, or <code>Critical</code></p>                                                                                                                                                                  |
| `--threads`                     | `3`                                                              | The number of parallel threads used to scan the source code project.                                                                                                                                                                                                                                                                       |
| `--go`                          | `false`                                                          | Set to true to request audit for a Go project.                                                                                                                                                                                                                                                                                             |
| `--gradle`                      | `false`                                                          | Set to true to request audit for a Gradle project.                                                                                                                                                                                                                                                                                         |
| `--mvn`                         | `false`                                                          | Set to true to request audit for a Maven project.                                                                                                                                                                                                                                                                                          |
| `--npm`                         | `false`                                                          | Set to true to request audit for a npm project.                                                                                                                                                                                                                                                                                            |
| `--pnpm`                        | `false`                                                          | Set to true to request audit for a pnpm project.                                                                                                                                                                                                                                                                                           |
| `--nuget`                       | `false`                                                          | Set to true to request audit for a .Net project.                                                                                                                                                                                                                                                                                           |
| `--pip`                         | `false`                                                          | Set to true to request audit for a Pip project.                                                                                                                                                                                                                                                                                            |
| `--pipenv`                      | `false`                                                          | Set to true to request audit for a Pipenv project.                                                                                                                                                                                                                                                                                         |
| `--yarn`                        | `false`                                                          | Set to true to request audit for a Yarn project.                                                                                                                                                                                                                                                                                           |
| `--sca`                         | `false`                                                          | <p><strong>Selective scanners mode</strong> </p><p>Executes SCA (Software Composition Analysis) sub-scan. Use <code>--sca</code> to run both SCA and Contextual Analysis. Use <code>--sca --without-contextual-analysis</code> to to run SCA. Can be combined with <code>--secrets</code>, <code>--sast</code>, <code>--iac.</code></p>    |
| `--iac`                         | `false`                                                          | <p><strong>Selective scanners mode</strong></p><p>Executes IaC sub-scan. Can be combined with <code>--sca</code>, <code>--secrets</code> and <code>--sast</code>.</p>                                                                                                                                                                      |
| `--secrets`                     | `false`                                                          | <p><strong>Selective scanners mode</strong></p><p>Executes Secrets sub-scan. Can be combined with <code>--sca</code>, <code>--sast,</code> and <code>--iac</code>.</p>                                                                                                                                                                     |
| `--sast`                        | `false`                                                          | <p><strong>Selective scanners mode</strong></p><p>Executes SAST sub-scan. Can be combined with <code>--sca</code>, <code>--secrets,</code> and <code>--iac</code>.</p>                                                                                                                                                                     |
| `--without-contextual-analysis` | `false`                                                          | <p><strong>Selective scanners mode</strong></p><p>Disable Contextual Analysis scanner after SCA. Relevant only with <code>--sca</code> flag.</p>                                                                                                                                                                                           |
| `--vuln`                        |                                                                  | <p>[Optional]<br>Set if you'd like to receive all vulnerabilities, regardless of the policy configured in Xray.</p>                                                                                                                                                                                                                        |
| `--validate-secrets`            | `false`                                                          | Triggers token validation on found secrets                                                                                                                                                                                                                                                                                                 |
| **Command arguments**           |                                                                  | The command accepts no arguments                                                                                                                                                                                                                                                                                                           |

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

Audit the project at the current directory using _watch1_ and \_watch2\_ defined in Xray.

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
