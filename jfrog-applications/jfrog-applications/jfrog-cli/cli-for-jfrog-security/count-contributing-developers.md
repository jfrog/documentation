# Count Contributing Developers

{% hint style="info" %}
This feature is supported in JFrog CLI version 2.60.0
{% endhint %}

The `git count-contributors` command allows JFrog users to easily determine the number of Git developers contributing to their code. The counts indicate the number of contributing developers to the **default branch**.&#x20;

This information can be helpful when purchasing an Advanced Security subscription, as the number of developers is often a key factor in pricing.&#x20;

We provide several options to obtain the developer count:

* **A single repository**: Analyze a single Git repository by providing the repository name.
* **Across a project/group**: Analyze multiple repositories organized under a project/group by providing the owner command option.
* **Across multiple Git servers**: Analyze repositories across various Git servers by providing a YAML file as an input file with the required parameters outlined below. &#x20;

Supported Git providers:

* GitHub
* GitLab
* Bitbucket

### Usage

The git count-contributors command can be run from the JFrog CLI with the following syntax:

```
git count-contributors [command optios]
```

| Command Option     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| --scm-type         | <p>(optional)<br>The type of SCM to use for the analysis. <br>Supported Values: <code>github, gitlab, bitbucket</code><br>Example: <code>--scm-type=github</code></p>                                                                                                                                                                                                                                                                                                                                                                                      |
| --scm-api-url      | <p>(optional)<br>The base URL of the SCM system's API endpoint.<br>Format: The full URL, including the protocol<br>Example: <code>--scm-api-url=https://api.github.com</code></p>                                                                                                                                                                                                                                                                                                                                                                          |
| --token            | <p>(optional)<br>The authentication token required to access the SCM system's API.<br>In the absence of a flag, tokens should be passed in the JF_GIT_TOKEN environment variable, or the corresponding environment variables 'JFROG_CLI_GITLAB_TOKEN, JFROG_CLI_GITHUB_TOKEN or JFROG_CLI_BITBUCKET_TOKEN'<br>Example: <code>--token:your_access_token</code></p>                                                                                                                                                                                          |
| --owner            | <p>(optional)<br>The owner or organization of the repositories to be analyzed. <br>Format: Depending on the Git provider. On GitHub and GitLab, the owner is typically an individual or an organization, On Bitbucket, the owner can also be a project. In the case of a private instance on Bitbucket, the individual or organization name should be prefixed with '~'.<br>When using this option without a specific repository name, all repositories will be analyzed at the group/project level. <br>Example: <code>owner=your-organization</code></p> |
| --months           | <p>(optional)<br>The number of months to analyze for developer activity. <br>Default: <code>1</code><br>Example: <code>--months=6</code></p>                                                                                                                                                                                                                                                                                                                                                                                                               |
| --detailed-summary | <p>(optional)<br>Generates a more detailed summary of the contributors. <br>Default: <code>false</code><br>Example: <code>--detailed-summary=true</code></p>                                                                                                                                                                                                                                                                                                                                                                                               |
| --repo-name        | <p>(optional)<br>List of semicolon-separated(;) repositories names to analyze, If not provided all repositories related to the provided owner will be analyzed.<br>Example: --repo-name=repo1;repo2</p>                                                                                                                                                                                                                                                                                                                                                    |
| --input-file       | <p>(optional)<br>The path to an input file in YAML format that contains multiple git providers.<br>Example: <code>--input-file="/Users/path/to/file/input.yaml"</code></p>                                                                                                                                                                                                                                                                                                                                                                                 |
| --verbose          | <p>(optional)<br>Enables verbose output, providing more detailed information.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |

### Example Commands

**Single Repository**

```javascript
git cc --scm-type=github --scm-api-url=https://api.github.com --token=<token> --owner=jfrog --months=4 --detailed-summary=false --repo-name=cli-core
```

Required Parameters:

* \--scm-type
* &#x20;\--scm-api-url
* \--token
* \--repo-name

Output:

```
// Some code
```

**Group/Project**

```javascript
git cc --scm-type=gitlab --scm-api-url=https://git.vdoo.io --token=<token> --owner=vdoo --months=3 --detailed-summary
```

Required Parameters:

* \--scm-type
* &#x20;\--scm-api-url
* \--token
* \--owner

**Multiple Git Servers- YAML File**

```javascript
git-servers-list:
  - scm-type: bitbucket
    scm-api-url: "https://api.bitbucket.url"
    token: "token"
    owner: "owner"
    repositories:
      - "repo1"
      - "repo2"
  - scm-type: gitlab
    scm-api-url: "https://api.github.com"
    token: "token"
    owner: "owner"
```
