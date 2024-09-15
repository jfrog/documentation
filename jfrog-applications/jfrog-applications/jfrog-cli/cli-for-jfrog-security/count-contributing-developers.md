# Count Contributing Developers

{% hint style="info" %}
This feature is supported in JFrog CLI version 2.60.0
{% endhint %}

The `git count-contributors` command allows JFrog users to easily determine the number of Git developers contributing to their code. The counts indicate the number of contributing developers to the **default branch**.&#x20;

The command works by counting the contributing developers for all commits performed within a time range you specify. The results are based on email addresses, thus giving you the specific number of unique developers.&#x20;

We provide several options to obtain the developer count:

* **A single repository**: Analyze a single Git repository by providing the repository name.
* **Across a project/group**: Analyze multiple repositories organized under a project/group by providing the owner command option.
* **Across multiple Git servers**: Analyze repositories across various Git servers by providing a YAML file as an input file with the required parameters outlined below.

This information can be helpful when purchasing an Advanced Security subscription, as the number of developers is often a key factor in pricing.

Supported Git providers:

* GitHub
* GitLab
* Bitbucket

> This command is intended to assist you in gaining insights to your contributors count. However, we recommend that in addition to utilizing this command that you confirm the data.&#x20;

### Usage

The `git count-contributors` command can be run from the JFrog CLI with the following syntax:

```
git count-contributors [command options]
```

| Command Option     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| --scm-type         | <p>(optional) The type of SCM to use for the analysis. </p><p>Supported Values: <code>github, gitlab, bitbucket</code> Example: <code>--scm-type=github</code></p>                                                                                                                                                                                                                                                                                                                                                          |
| --scm-api-url      | <p>(optional) The base URL of the SCM system's API endpoint. </p><p>Format: The full URL, including the protocol Example: <code>--scm-api-url=https://api.github.com</code></p>                                                                                                                                                                                                                                                                                                                                             |
| --token            | (optional) The authentication token required to access the SCM system's API. In the absence of a flag, tokens should be passed in the JF\_GIT\_TOKEN environment variable, or the corresponding environment variables 'JFROG\_CLI\_GITLAB\_TOKEN, JFROG\_CLI\_GITHUB\_TOKEN or JFROG\_CLI\_BITBUCKET\_TOKEN' Example: `--token:your_access_token`                                                                                                                                                                           |
| --owner            | (optional) The owner or organization of the repositories to be analyzed. Format: Depending on the Git provider. On GitHub and GitLab, the owner is typically an individual or an organization, On Bitbucket, the owner can also be a project. In the case of a private instance on Bitbucket, the individual or organization name should be prefixed with '\~'. When using this option without a specific repository name, all repositories will be analyzed at the group/project level. Example: `owner=your-organization` |
| --months           | (optional) The number of months to analyze for developer activity. Default: `1` Example: `--months=6`                                                                                                                                                                                                                                                                                                                                                                                                                       |
| --detailed-summary | (optional) Generates a more detailed summary of the contributors. Default: `false` Example: `--detailed-summary=true`                                                                                                                                                                                                                                                                                                                                                                                                       |
| --repo-name        | (optional) List of semicolon-separated(;) repositories names to analyze, If not provided all repositories related to the provided owner will be analyzed. Example: --repo-name=repo1;repo2                                                                                                                                                                                                                                                                                                                                  |
| --input-file       | (optional) The path to an input file in YAML format that contains multiple git providers. Example: `--input-file="/Users/path/to/file/input.yaml"`                                                                                                                                                                                                                                                                                                                                                                          |
| --verbose          | <p>(optional) </p><p>Enables verbose output, providing more detailed information.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                       |

#### Example Commands <a href="#example-commands" id="example-commands"></a>

**Single Repository**

```javascript
git cc --scm-type=github --scm-api-url=https://api.github.com --token=<token> --owner=jfrog --months=4 --detailed-summary=false --repo-name=cli-core
```

Required Parameters:

* \--scm-type
* \--scm-api-url
* \--token
* \--repo-name

**Group/Project**

```javascript
git cc --scm-type=gitlab --scm-api-url=https://git.vdoo.io --token=<token> --owner=vdoo --months=3 --detailed-summary
```

Required Parameters:

* \--scm-type
* \--scm-api-url
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

Sample Output:

<pre><code>{
  "total_unique_contributors": 4,	-	        <a data-footnote-ref href="#user-content-fn-1">The number of unique developers (dedup)</a>
  "total_commits": 4,				-	<a data-footnote-ref href="#user-content-fn-2">The number of commits examined on the default branch (total)</a>
  "scanned_repos": [				-	<a data-footnote-ref href="#user-content-fn-3">The repositories that were scanned</a>
    "test-go",
    "test-cli-core"
  ],
  "report_date": "2024-07-22T12:08:04+03:00",	-	<a data-footnote-ref href="#user-content-fn-4">The report date</a>
  "number_of_months": "5",			-	<a data-footnote-ref href="#user-content-fn-5">The time range specified</a>
  "unique_contributors_list": [			-	<a data-footnote-ref href="#user-content-fn-6">The evidence of the last seen developer</a>
    {
      "email": "dev1@users.noreply.github.com",
      "name": "`Developer 1",
      "last_commit": {
        "repo": "test-cli-core",
        "date": "2024-02-22T14:21:55Z",
        "hash": "3463b55aa453fb5dd3d5e7c6ebf45a3e33710e72"
      }
    },
    {
      "email": "dev2@users.noreply.github.com",
      "name": "Developer 2",
      "last_commit": {
        "repo": "test-go",
        "date": "2024-03-24T07:56:17Z",
        "hash": "8b102603458044b434689fc3832e12d30af12d15"
      }
    },
    {
      "email": "dev3@jfrog.com",
      "name": "Developer 3",
      "last_commit": {
        "repo": "test-cli-core",
        "date": "2024-02-25T15:15:19Z",
        "hash": "de88b95a38242b9984877a8e928ceafedb147843"
      }
    }
  ],
  "detailed_contributors_list": {		-	<a data-footnote-ref href="#user-content-fn-7">With verbose; detailed evidence per user</a>
    "dev2@users.noreply.github.com": [
      {
        "repo_path": "test-go",
        "last_commit": {
          "date": "2024-03-24T07:56:17Z",
          "hash": "8b102603458044b434689fc3832e12d30af12d15"
        }
      },
      {
        "repo_path": "test-cli-core",
        "last_commit": {
          "date": "2024-02-25T12:40:40Z",
          "hash": "0941c5ce1007501c2793efa0e09b0e9531b8d503"
        }
      }
    ],
    "dev3@jfrog.com": [
      {
        "repo_path": "test-cli-core",
        "last_commit": {
          "date": "2024-02-25T15:15:19Z",
          "hash": "de88b95a38242b9984877a8e928ceafedb147843"
        }
      }
    ],
    "dev1@users.noreply.github.com": [
      {
        "repo_path": "test-cli-core",
        "last_commit": {
          "date": "2024-02-22T14:21:55Z",
          "hash": "3463b55aa453fb5dd3d5e7c6ebf45a3e33710e72"
        }
      }
    ]
  },
  "detailed_repos_list": { -	<a data-footnote-ref href="#user-content-fn-8">With verbose; detailed evidence per repo</a>
    "test-go": [
      {
        "email": "dev2@users.noreply.github.com",
        "last_commit": {
          "date": "2024-03-24T07:56:17Z",
          "hash": "8b102603458044b434689fc3832e12d30af12d15"
        }
      }
    ],
    "test-cli-core": [
      {
        "email": "dev3@jfrog.com",
        "last_commit": {
          "date": "2024-02-25T15:15:19Z",
          "hash": "de88b95a38242b9984877a8e928ceafedb147843"
        }
      },
      {
        "email": "dev2@users.noreply.github.com",
        "last_commit": {
          "date": "2024-02-25T12:40:40Z",
          "hash": "0941c5ce1007501c2793efa0e09b0e9531b8d503"
        }
      },
      {
        "email": "dev1@users.noreply.github.com",
        "last_commit": {
          "date": "2024-02-22T14:21:55Z",
          "hash": "3463b55aa453fb5dd3d5e7c6ebf45a3e33710e72"
        }
      }
    ]
  }
}
</code></pre>

[^1]: 

[^2]: 

[^3]: 

[^4]: 

[^5]: 

[^6]: 

[^7]: 

[^8]: 
