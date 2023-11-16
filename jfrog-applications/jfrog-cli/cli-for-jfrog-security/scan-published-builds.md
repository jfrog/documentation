# Scan Published Builds

## Scanning Published Builds

JFrog CLI is integrated with JFrog Xray and JFrog Artifactory, allowing you to have your build artifacts and dependencies scanned for vulnerabilities and license violations. Please notice that the build in the below example had already been published to Artifactory using the [build-publish command](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#publishing-build-info).

### Commands Params

|                   |                                                                                                                                                                                                                                            |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command name      | build-scan                                                                                                                                                                                                                                 |
| Abbreviation      | bs                                                                                                                                                                                                                                         |
| Command options   |                                                                                                                                                                                                                                            |
| --server-id       | <p>[Optional]<br><br>Server ID configured by the <em>jf c add</em> command. If not specified, the default configured server is used.</p>                                                                                                   |
| --vuln            | <p>[Optional]<br><br>Set if you'd like to receive all vulnerabilities, regardless of the policy configured in Xray.</p>                                                                                                                    |
| --fail            | <p>[Default: true]<br><br>When set, the command returns exit code 3 if a 'Fail Build' rule is matched by Xray.<br>Set to false if you do not wish the command to return exit code 3 in such case, and an exit code 0 will be returned.</p> |
| --format          | <p>[Default: table]<br><br>Defines the output format of the command. The accepted values are: <em><strong>table</strong></em> and <em><strong>json</strong></em>.</p>                                                                      |
| --project         | <p>[Optional]<br><br>JFrog project key</p>                                                                                                                                                                                                 |
| --rescan          | <p>[Default: false]<br><br>Set to true when scanning an already successfully scanned build, for example after adding an ignore rule.</p>                                                                                                   |
| Command arguments | The command accepts two arguments.                                                                                                                                                                                                         |
| Build name        | Build name to be scanned.                                                                                                                                                                                                                  |
| Build number      | Build number to be scanned.                                                                                                                                                                                                                |

### Example

Scan build number 18, corresponding to the following build name: 'my-build-name'.
```
jf bs my-build-name 18
```
