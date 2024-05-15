# Scan Published Builds

## Scanning Published Builds

JFrog CLI is integrated with JFrog Xray and JFrog Artifactory, allowing you to have your build artifacts and dependencies scanned for vulnerabilities and license violations. Please notice that the build in the below example had already been published to Artifactory using the [build-publish command](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#publishing-build-info).

### Commands Params

|                   |                                                                                                                                                                                                                                |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name      | build-scan                                                                                                                                                                                                                     |
| Abbreviation      | bs                                                                                                                                                                                                                             |
| **Command options:**   |                                                                                                                                                                                                                                |
| `--server-id` | <p>[Optional]<br>Server ID configured by the <em>jf c add</em> command. If not specified, the default configured server is used.</p>                                                                                           |
| `--vuln` | <p>[Optional]<br>Set if you'd like to receive all vulnerabilities, regardless of the policy configured in Xray.</p>                                                                                                            |
| `--fail` | <p>[Default: true]<br>When using one of the flags --watches, --project or --repo-path and a Fail build rule is matched the command will return exit code 3. Set to false if you'd like to see violations with exit code 0.</p> |
| `--format` | <p>[Default: table]<br>Defines the output format of the command. The accepted values are: <em><strong>table</strong></em> and <em><strong>json</strong></em>.</p>                                                              |
| `--project` | <p>[Optional]<br>JFrog project key</p>                                                                                                                                                                                         |
| `--rescan` | <p>[Default: false]<br>Set to true when scanning an already successfully scanned build, for example after adding an ignore rule.</p>                                                                                           |
| **Command arguments:** | The command accepts two arguments.                                                                                                                                                                                             |
| Build name        | Build name to be scanned.                                                                                                                                                                                                      |
| Build number      | Build number to be scanned.                                                                                                                                                                                                    |

### Example

Scan build number 18, corresponding to the following build name: 'my-build-name'.
```
jf bs my-build-name 18
```
