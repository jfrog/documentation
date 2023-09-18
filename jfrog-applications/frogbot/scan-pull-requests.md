# Scan Pull Requests

Frogbot uses [JFrog Xray](https://jfrog.com/xray/) (version 3.29.0 and above is required) to scan your pull requests. It adds the scan results as a comment on the pull request. If no new vulnerabilities are found, Frogbot will also add a comment, confirming this.

The following features use the package manager used for building the project:

* Software Composition Analysis (SCA)
* Vulnerability Contextual Analysis


#### How to use Pull Request scanning?

<details>

<summary>Azure Repos</summary>



After you create a new pull request, Frogbot will automatically scan it.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the pull request. Vulnerabilities that aren't new, and existed in the code before the pull request was created, will not be included in the report. In order to include all the vulnerabilities in the report, including older ones that weren't added by this PR, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot Azure Repos scan workflow is:

1. The developer opens a pull request.
2. Frogbot scans the pull request and adds a comment with the scan results.
3. Frogbot can be triggered again following new commits, by adding a comment with the `rescan` text.

</details>

<details>

<summary>Bitbucket Server</summary>

After you create a new pull request, Frogbot will automatically scan it.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the pull request. Vulnerabilities that aren't new, and existed in the code before the pull request was created, will not be included in the report. In order to include all of the vulnerabilities in the report, including older ones that weren't added by this PR, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot scan on Bitbucket Server workflow:

1. The developer opens a pull request.
2. Frogbot scans the pull request and adds a comment with the scan results.
3. Frogbot can be triggered again following new commits, by adding a comment with the `rescan` text.

</details>

<details>

<summary>GitHub</summary>

After you create a new pull request, the maintainer of the Git repository can trigger Frogbot to scan the pull request from the pull request UI.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the pull request. Vulnerabilities that aren't new, and existed in the code before the pull request was created, will not be included in the report. In order to include all of the vulnerabilities in the report, including older ones that weren't added by this PR, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot GitHub scan workflow is:

1. The developer opens a pull request.
2. The Frogbot workflow automatically gets triggered and a [GitHub environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#creating-an-environment) named `frogbot` becomes pending the maintainer's approval.\
   ![](../.gitbook/assets/github-pending-deployment.png)
3. The maintainer of the repository reviews the pull request and approves the scan
4. Frogbot can be triggered again following new commits, by repeating steps 2 and 3.

</details>

<details>

<summary>GitLab</summary>

After you create a new merge request, the maintainer of the Git repository can trigger Frogbot to scan the merge request from the merge request UI.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the merge request. Vulnerabilities that aren't new, and existed in the code before the merge request was created, will not be included in the report. In order to include all of the vulnerabilities in the report, including older ones that weren't added by this merge request, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot GitLab flow is as follows:

1. The developer opens a merge request.
2. The maintainer of the repository reviews the merge request and approves the scan by triggering the manual _frogbot-scan_ job.
3. Frogbot is then triggered by the job, it scans the merge request and adds a comment with the scan results.
4. Frogbot can be triggered again following new commits, by triggering the _frogbot-scan_ job again.\
   ![](../.gitbook/assets/gitlab-run-button.png)

</details>

#### 👮 Security note for pull requests scanning

When installing Frogbot using JFrog Pipelines, Jenkins, and Azure DevOps, Frogbot will not wait for a maintainer's approval before scanning newly opened pull requests. Using Frogbot with these platforms is therefore not recommended for open-source projects.

When installing Frogbot using GitHub Actions and GitLab however, Frogbot will initiate the scan only after a maintainer approves it of the project. The goal of this review is to ensure that external code contributors don't introduce malicious code as part of the pull request. Since this review step is enforced by Frogbot when used with GitHub Actions and GitLab, it is safe to be used for open-source projects.

#### Scan results

Frogbot adds the scan results to the pull request in the following format:

**👍 No issues**

If no new vulnerabilities are found, Frogbot automatically adds the following comment to the pull request:

[![](https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/noVulnerabilityBannerPR.png)](broken-reference)

**👎 Issues were found**

If new vulnerabilities are found, Frogbot adds them as a comment on the pull request. For example:

[![](https://raw.githubusercontent.com/jfrog/frogbot/master/resources/v2/vulnerabilitiesBannerPR.png)](broken-reference)

**VULNERABLE DEPENDENCIES**

| SEVERITY                                          | CONTEXTUAL ANAYSIS | DIRECT DEPENDENCY | IMPACTED DEPENDENCY | FIXED VERSIONS    |
| ------------------------------------------------- | ------------------ | ----------------- | ------------------- | ----------------- |
| ![](../.gitbook/assets/notApplicableCritical.png) | Not Applicable     | minimist:1.2.5    | minimist:1.2.5      | \[0.2.4] \[1.2.6] |
| ![](../.gitbook/assets/High.png)                  | Applicable         | protobufjs:6.11.2 | protobufjs:6.11.2   | \[6.11.3]         |
| ![](../.gitbook/assets/notApplicableHigh.png)     | Not Applicable     | lodash:4.17.19    | lodash:4.17.19      | \[4.17.21]        |


## For source code scanning Frogbot will add code review comments and attach them to the vulnerable code lines. For example:

**Static Application Security Testing (SAST)**

![image](https://github.com/jfrog/documentation/assets/29822394/9ed841b1-d0f1-46d6-b448-ece558ef0c77)


**CVE Vulnerability Contextual Analysis**

![image](https://github.com/jfrog/documentation/assets/29822394/7474abbd-e439-4f80-be03-9c77ea70d182)


**INFRASTRUCTURE AS CODE (IaC)**

| SEVERITY                                          | FILE    | LINE:COLUMN | FINDING                             |
| ------------------------------------------------- | ------- | ----------- | ----------------------------------- |
| ![](../.gitbook/assets/notApplicableCritical.png) | test.js | 1:20        | kms\_key\_id='' was detected        |
| ![](../.gitbook/assets/High.png)                  | mock.js | 4:30        | Deprecated TLS version was detected |


**Secrets Detection**

When Frogbot detects secrets that have been inadvertently exposed within the code of a pull request, it promptly triggers an email notification to the user who pushed the corresponding commit. The email address utilized for this notification is sourced from the committer's Git profile configuration. Moreover, Frogbot offers the flexibility to direct email notifications to an extra email address if desired. To activate email notifications, it is necessary to configure your SMTP server details as variables within your Frogbot workflows.

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/secrets-email.png" alt="" data-size="original">
