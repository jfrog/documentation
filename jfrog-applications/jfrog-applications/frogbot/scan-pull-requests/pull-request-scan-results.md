# Pull Request Scan Results

### Scan results

Frogbot adds the scan results to the pull request in the following format:

#### 👍 No issues

If no new vulnerabilities are found, Frogbot automatically adds the following comment to the pull request:

[![](https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/noVulnerabilityBannerPR.png)](pull-request-scan-results.md#-no-issues)

#### 👎 Issues were found

**Software Composition Analysis (SCA)**

If new vulnerabilities are found, Frogbot adds them as a comment on the pull request. For example:

[![](https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/vulnerabilitiesBannerPR.png)](pull-request-scan-results.md#-issues)

**VULNERABLE DEPENDENCIES**

|                                                               SEVERITY                                                              | CONTEXTUAL ANALYSIS | DIRECT DEPENDENCIES | IMPACTED DEPENDENCY |       FIXED VERSIONS      |
| :---------------------------------------------------------------------------------------------------------------------------------: | :-----------------: | :-----------------: | :-----------------: | :-----------------------: |
| <p><img src="https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/notApplicableCritical.png" alt=""><br>Critical</p> |    Not Applicable   |    minimist:1.2.5   |    minimist:1.2.5   | <p>[0.2.4]<br>[1.2.6]</p> |
|   <p><img src="https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/applicableHighSeverity.png" alt=""><br>High</p>  |      Applicable     |  protobufjs:6.11.2  |  protobufjs:6.11.2  |         \[6.11.3]         |
|     <p><img src="https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/notApplicableHigh.png" alt=""><br>High</p>     |    Not Applicable   |    lodash:4.17.19   |    lodash:4.17.19   |         \[4.17.21]        |

**Vulnerability Contextual Analysis**

![](https://raw.githubusercontent.com/jfrog/frogbot/master/images/pr-vuln-contextual-analysis.png)

**Static Application Security Testing (SAST)**

![](https://raw.githubusercontent.com/jfrog/frogbot/master/images/pr-sast.png)

**Infrastructure as Code scans (IaC)**

![](https://raw.githubusercontent.com/jfrog/frogbot/master/images/pr-iac.png)

**Validate Allowed Licenses**

When Frogbot scans newly opened pull requests, it checks the licenses of any new direct project dependencies introduced by the pull request. If Frogbot identifies licenses that are not listed in a predefined set of approved licenses, it appends a comment to the pull request providing this information. The list of allowed licenses is set up as a variable within the Frogbot workflow.

![](https://raw.githubusercontent.com/jfrog/frogbot/master/images/violated-licenses.png)

#### Secrets Detection

When Frogbot detects secrets that have been inadvertently exposed within the code of a pull request, it promptly triggers an email notification to the user who pushed the corresponding commit. The email address utilized for this notification is sourced from the committer's Git profile configuration. Moreover, Frogbot offers the flexibility to direct the email notification to an extra email address if desired. To activate email notifications, it is necessary to configure your SMTP server details as variables within your Frogbot workflows.

![](https://raw.githubusercontent.com/jfrog/frogbot/master/images/secrets-email.png)

xx
