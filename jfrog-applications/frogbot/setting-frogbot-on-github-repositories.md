# Setting Frogbot on GitHub Repositories

### Github Prerequisites

* Go to your repository's **settings** tab and save the JFrog connection details as repository secrets with the following names:
  * **JF\_URL** (JFrog Platform URL - Example: `https://acme.jfrog.io`)
  * **JF\_ACCESS\_TOKEN** (JFrog access token)

> You can also use **JF\_XRAY\_URL** and **JF\_ARTIFACTORY\_URL** instead of **JF\_URL**, and **JF\_USER** + **JF\_PASSWORD** instead of **JF\_ACCESS\_TOKEN**

![](../../../.gitbook/assets/github-repository-secrets.png)

* Under **Actions** > **General**, check the **Allow GitHub Actions to create and approve pull requests** check box.

![](../../../.gitbook/assets/github-pr-permissions.png)

* For open-source projects: Create a new [GitHub environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#creating-an-environment) called **frogbot** and add people or public teams as reviewers. The chosen reviewers can trigger Frogbot scans on pull requests.

![](../../../.gitbook/assets/github-environment.png)

### Frogbot GitHub Action Templates

* Begin by cloning the GitHub repository to your local environment.

* Switch to the target branch where you'd like the pull requests to be scanned.

* Create a file named **frogbot-scan-pull-request.yml**. Fill it with the template provided below, and then push it into the **.github/workflows** directory at the root of your GitHub repository.

<details>

<summary>frogbot-scan-pull-request.yml</summary>

```
name: "Frogbot Scan Pull Request"
on:
  pull_request_target:
    types: [opened, synchronize]
permissions:
  pull-requests: write
  contents: read
jobs:
  scan-pull-request:
    runs-on: ubuntu-latest
    # A pull request needs to be approved before Frogbot scans it. Any GitHub user who is associated with the
    # "frogbot" GitHub environment can approve the pull request to be scanned.
    environment: frogbot
    steps:
      - uses: jfrog/frogbot@v2
        env:
          # [Mandatory]
          # JFrog platform URL
          JF_URL: ${{ secrets.JF_URL }}

          # [Mandatory if JF_USER and JF_PASSWORD are not provided]
          # JFrog access token with 'read' permissions on Xray service
          JF_ACCESS_TOKEN: ${{ secrets.JF_ACCESS_TOKEN }}

          # [Mandatory if JF_ACCESS_TOKEN is not provided]
          # JFrog username with 'read' permissions for Xray. Must be provided with JF_PASSWORD
          # JF_USER: ${{ secrets.JF_USER }}

          # [Mandatory if JF_ACCESS_TOKEN is not provided]
          # JFrog password. Must be provided with JF_USER
          # JF_PASSWORD: ${{ secrets.JF_PASSWORD }}

          # [Mandatory]
          # The GitHub token is automatically generated for the job
          JF_GIT_TOKEN: ${{ secrets.GITHUB_TOKEN }}

          # [Optional, default: https://api.github.com]
          # API endpoint to GitHub
          # JF_GIT_API_ENDPOINT: https://github.example.com

          # [Optional]
          # By default, the Frogbot workflows download the Frogbot executable as well as other tools
          # needed from https://releases.jfrog.io
          # If the machine that runs Frogbot has no access to the internet, follow these steps to allow the
          # executable to be downloaded from an Artifactory instance, which the machine has access to:
          #
          # 1. Login to the Artifactory UI, with a user who has admin credentials.
          # 2. Create a Remote Repository with the following properties set.
          #    Under the 'Basic' tab:
          #       Package Type: Generic
          #       URL: https://releases.jfrog.io
          #    Under the 'Advanced' tab:
          #       Uncheck the 'Store Artifacts Locally' option
          # 3. Set the value of the 'JF_RELEASES_REPO' variable with the Repository Key you created.
          # JF_RELEASES_REPO: ""

          # [Optional]
          # Configure the SMTP server to enable Frogbot to send emails with detected secrets in pull request scans.
          # SMTP server URL including should the relevant port: (Example: smtp.server.com:8080)
          # JF_SMTP_SERVER: ""

          # [Mandatory if JF_SMTP_SERVER is set]
          # The username required for authenticating with the SMTP server.
          # JF_SMTP_USER: ""

          # [Mandatory if JF_SMTP_SERVER is set]
          # The password associated with the username required for authentication with the SMTP server.
          # JF_SMTP_PASSWORD: ""

          ##########################################################################
          ##   If your project uses a 'frogbot-config.yml' file, you can define   ##
          ##   the following variables inside the file, instead of here.          ##
          ##########################################################################

          # [Mandatory if the two conditions below are met]
          # 1. The project uses yarn 2, NuGet or .NET Core to download its dependencies
          # 2. The `installCommand` variable isn't set in your frogbot-config.yml file.
          #
          # The command that installs the project dependencies (e.g "nuget restore")
          # JF_INSTALL_DEPS_CMD: ""

          # [Optional, default: "."]
          # Relative path to the root of the project in the Git repository
          # JF_WORKING_DIR: path/to/project/dir

          # [Optional]
          # Xray Watches. Learn more about them here: https://www.jfrog.com/confluence/display/JFROG/Configuring+Xray+Watches
          # JF_WATCHES: <watch-1>,<watch-2>...<watch-n>

          # [Optional]
          # JFrog project. Learn more about it here: https://www.jfrog.com/confluence/display/JFROG/Projects
          # JF_PROJECT: <project-key>

          # [Optional, default: "FALSE"]
          # Displays all existing vulnerabilities, including the ones that were added by the pull request.
          # JF_INCLUDE_ALL_VULNERABILITIES: "TRUE"

          # [Optional, default: "FALSE"]
          # When adding new comments on pull requests, keep old comments that were added by previous scans.
          # JF_AVOID_PREVIOUS_PR_COMMENTS_DELETION: "TRUE"

          # [Optional, default: "TRUE"]
          # Fails the Frogbot task if any security issue is found.
          # JF_FAIL: "FALSE"

          # [Optional]
          # Frogbot will download the project dependencies if they're not cached locally. To download the
          # dependencies from a virtual repository in Artifactory, set the name of the repository. There's no
          # need to set this value, if it is set in the frogbot-config.yml file.
          # JF_DEPS_REPO: ""

          # [Optional, Default: "FALSE"]
          # If TRUE, Frogbot creates a single pull request with all the fixes.
          # If false, Frogbot creates a separate pull request for each fix.
          # JF_GIT_AGGREGATE_FIXES: "FALSE"

          # [Optional, Default: "FALSE"]
          # Handle vulnerabilities with fix versions only
          # JF_FIXABLE_ONLY: "TRUE"

          # [Optional]
          # Set the minimum severity for vulnerabilities that should be fixed and commented on in pull requests
          # The following values are accepted: Low, Medium, High or Critical
          # JF_MIN_SEVERITY: ""

          # [Optional]
          # List of comma separated email addresses to receive email notifications about secrets
          # detected during pull request scanning. The notification is also sent to the email set
          # in the committer git profile regardless of whether this variable is set or not.
          # JF_EMAIL_RECEIVERS: ""

          # [Optional]
          # Set the list of allowed licenses
          # The full list of licenses can be found in:
          # https://github.com/jfrog/frogbot/blob/master/docs/licenses.md
          # JF_ALLOWED_LICENSES: "MIT, Apache-2.0"

          # [Optional]
          # Avoid adding extra info to pull request comments. that isn't related to the scan findings.
          # JF_AVOID_EXTRA_MESSAGES: "TRUE"

          # [Optional]
          # Add a title to pull request comments generated by Frogbot.
          # JF_PR_COMMENT_TITLE: ""
```

</details>

* Return to the default branch.

* Now, create a file named **frogbot-scan-repository.yml**. Again, populate it with the provided template and push it into the **.github/workflows** directory at the root of your GitHub repository.

<details>

<summary>frogbot-scan-repository.yml</summary>

```
name: "Frogbot Scan Repository"
on:
  workflow_dispatch:
  schedule:
    # The repository will be scanned once a day at 00:00 GMT.
    - cron: "0 0 * * *"
permissions:
  contents: write
  pull-requests: write
  security-events: write
jobs:
  scan-repository:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        # The repository scanning will be triggered periodically on the following branches.
        branch: [ "dev" ]
    steps:
      - uses: jfrog/frogbot@v2
        env:
          # [Mandatory]
          # JFrog platform URL
          JF_URL: ${{ secrets.JF_URL }}

          # [Mandatory if JF_USER and JF_PASSWORD are not provided]
          # JFrog access token with 'read' permissions on Xray service
          JF_ACCESS_TOKEN: ${{ secrets.JF_ACCESS_TOKEN }}

          # [Mandatory if JF_ACCESS_TOKEN is not provided]
          # JFrog username with 'read' permissions for Xray. Must be provided with JF_PASSWORD
          # JF_USER: ${{ secrets.JF_USER }}

          # [Mandatory if JF_ACCESS_TOKEN is not provided]
          # JFrog password. Must be provided with JF_USER
          # JF_PASSWORD: ${{ secrets.JF_PASSWORD }}

          # [Mandatory]
          # The GitHub token is automatically generated for the job
          JF_GIT_TOKEN: ${{ secrets.GITHUB_TOKEN }}

          # [Mandatory]
          # The name of the branch on which Frogbot will perform the scan
          JF_GIT_BASE_BRANCH: ${{ matrix.branch }}

          # [Optional, default: https://api.github.com]
          # API endpoint to GitHub
          # JF_GIT_API_ENDPOINT: https://github.example.com

          # [Optional]
          # By default, the Frogbot workflows download the Frogbot executable as well as other tools
          # needed from https://releases.jfrog.io
          # If the machine that runs Frogbot has no access to the internet, follow these steps to allow the
          # executable to be downloaded from an Artifactory instance, which the machine has access to:
          #
          # 1. Login to the Artifactory UI, with a user who has admin credentials.
          # 2. Create a Remote Repository with the following properties set.
          #    Under the 'Basic' tab:
          #       Package Type: Generic
          #       URL: https://releases.jfrog.io
          #    Under the 'Advanced' tab:
          #       Uncheck the 'Store Artifacts Locally' option
          # 3. Set the value of the 'JF_RELEASES_REPO' variable with the Repository Key you created.
          # JF_RELEASES_REPO: ""

          ##########################################################################
          ##   If your project uses a 'frogbot-config.yml' file, you can define   ##
          ##   the following variables inside the file, instead of here.          ##
          ##########################################################################

          # [Optional, default: "."]
          # Relative path to the root of the project in the Git repository
          # JF_WORKING_DIR: path/to/project/dir

          # [Optional]
          # Xray Watches. Learn more about them here: https://www.jfrog.com/confluence/display/JFROG/Configuring+Xray+Watches
          # JF_WATCHES: <watch-1>,<watch-2>...<watch-n>

          # [Optional]
          # JFrog project. Learn more about it here: https://www.jfrog.com/confluence/display/JFROG/Projects
          # JF_PROJECT: <project-key>

          # [Optional, default: "TRUE"]
          # Fails the Frogbot task if any security issue is found.
          # JF_FAIL: "FALSE"

          # [Optional]
          # Frogbot will download the project dependencies, if they're not cached locally. To download the
          # dependencies from a virtual repository in Artifactory, set the name of the repository. There's no
          # need to set this value, if it is set in the frogbot-config.yml file.
          # JF_DEPS_REPO: ""

          # [Optional]
          # Template for the branch name generated by Frogbot when creating pull requests with fixes.
          # The template must include {BRANCH_NAME_HASH}, to ensure that the generated branch name is unique.
          # The template can optionally include the {IMPACTED_PACKAGE} and {FIX_VERSION} variables.
          # JF_BRANCH_NAME_TEMPLATE: "frogbot-{IMPACTED_PACKAGE}-{BRANCH_NAME_HASH}"

          # [Optional]
          # Template for the commit message generated by Frogbot when creating pull requests with fixes
          # The template can optionally include the {IMPACTED_PACKAGE} and {FIX_VERSION} variables.
          # JF_COMMIT_MESSAGE_TEMPLATE: "Upgrade {IMPACTED_PACKAGE} to {FIX_VERSION}"

          # [Optional]
          # Template for the pull request title generated by Frogbot when creating pull requests with fixes.
          # The template can optionally include the {IMPACTED_PACKAGE} and {FIX_VERSION} variables.
          # JF_PULL_REQUEST_TITLE_TEMPLATE: "[🐸 Frogbot] Upgrade {IMPACTED_PACKAGE} to {FIX_VERSION}"

          # [Optional, Default: "FALSE"]
          # If TRUE, Frogbot creates a single pull request with all the fixes.
          # If FALSE, Frogbot creates a separate pull request for each fix.
          # JF_GIT_AGGREGATE_FIXES: "FALSE"

          # [Optional, Default: "FALSE"]
          # Handle vulnerabilities with fix versions only
          # JF_FIXABLE_ONLY: "TRUE"

          # [Optional]
          # Set the minimum severity for vulnerabilities that should be fixed and commented on in pull requests
          # The following values are accepted: Low, Medium, High or Critical
          # JF_MIN_SEVERITY: ""

          # [Optional, Default: eco-system+frogbot@jfrog.com]
          # Set the email of the commit author
          # JF_GIT_EMAIL_AUTHOR: ""

          # [Optional]
          # Avoid adding extra info to pull request comments. that isn't related to the scan findings.
          # JF_AVOID_EXTRA_MESSAGES: "TRUE"
```

</details>