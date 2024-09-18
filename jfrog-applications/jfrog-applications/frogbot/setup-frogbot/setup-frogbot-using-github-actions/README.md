---
description: Install Frogbot on GitHub using GitHub Actions
---

# Setup Frogbot Using GitHub Actions

## Prepare GitHub to work with Frogbot

Perform the following steps to allow GitHub and Frogbot to work together:

### Step 1: Provide connection details

<details>

<summary>Set Frogbot's connection details as GitHub secrets</summary>

Go to your repository's **settings** tab and save the JFrog connection details as repository secrets with the following names:

1. **JF\_URL**  (JFrog Platform URL)\
   Example: `https://acme.jfrog.io`\
   You can also use **JF\_XRAY\_URL** and **JF\_ARTIFACTORY\_URL** instead of **JF\_URL.**
2. **JF\_ACCESS\_TOKEN** (JFrog access token)\
   You can also use **JF\_USER** + **JF\_PASSWORD** instead of **JF\_ACCESS\_TOKEN**.\
   Instead of using **JF\_ACCESS\_TOKEN** and providing an access token as a GitHub secret, you can utilize the [GitHub OpenID Connect (OIDC)](oidc-authentication.md) authentication protocol.
3. **JF\_GIT\_TOKEN** (GitHub token)\
   You can utilize [$\{{secrets.GITHUB\_TOKEN\}}](https://docs.github.com/en/actions/security-guides/automatic-token-authentication) for **JF\_GIT\_TOKEN**, which is an automatically generated token by GitHub. However, this option comes with a limitation: a workflow, such as Frogbot itself, cannot trigger another workflow. Consequently, if you have additional workflows intended to activate upon the creation of a new pull request, they might not be initiated. To resolve this issue, you can generate a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) and use it as JF\_GIT\_TOKEN.

<img src="../../../../.gitbook/assets/github-repository-secrets.png" alt="" data-size="original">

</details>

### Step 2: Allow Frogbot to open Pull Requests

<details>

<summary>Allow Pull Requests</summary>

Under **Actions** > **General**, check the **Allow GitHub Actions to create and approve pull requests** check box.

<img src="../../../../.gitbook/assets/github-pr-permissions.png" alt="" data-size="original">

</details>

### Step 3: Create execution Environment (open source projects only)

<details>

<summary>Create a dedicated execution environment for Frogbot</summary>

Create a new [GitHub environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#creating-an-environment) called **frogbot** and add people or public teams as reviewers. \
The chosen reviewers can trigger Frogbot scans on pull requests.

![](<../../../../.gitbook/assets/image (1).png>)

</details>

## Create the required GitHub Actions templates

### Step 1: Navigate to the project you wish to scan&#x20;

Clone the GitHub repository you wish to scan to your local environment:

```shell-session
> git clone <my-project.git>
> cd <my-project>
```

Switch to the branch you'd like to scan with Frogbot:

```shell-session
> git checkout <branch-to-scan>
```

### Step 2: Set up Repository Scan

In the branch you'd like to scan, create a file named `frogbot-scan-repository.yml`. Fill it with the provided [template](./#basic-frogbot-scan-repository.yml-template) and push it into the `.github/workflows` directory at the root of your GitHub repository.\
You can see more advanced options in the [full scan repository template](scan-git-repository-full-template.md).

<details>

<summary>frogbot-scan-repository.yml template</summary>

```yaml
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
  # [Mandatory If using OIDC authentication protocol instead of JF_ACCESS_TOKEN]
  # id-token: write
jobs:
  scan-repository:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        # The repository scanning will be triggered periodically on the following branches.
        branch: ["dev"]
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

        # [Mandatory if using OIDC authentication protocol instead of JF_ACCESS_TOKEN]
        # Insert to oidc-provider-name the 'Provider Name' defined in the OIDC integration configured in the JPD
        # with:
        #   oidc-provider-name: ""
```

</details>

### Step 3: Set up Pull Request Scan

Create a file named `frogbot-scan-pull-request.yml`. Fill it with the provided [template](./#basic-frogbot-scan-pull-request.yml-template), and then push it into the `.github/workflows` directory at the root of your GitHub repository.\
You can see more advanced options in the [full scan pull request template](scan-pull-request-full-template.md).

<details>

<summary>frogbot-scan-pull-request.yml template</summary>

{% code fullWidth="true" %}
```yaml
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
  # [Mandatory If using OIDC authentication protocol instead of JF_ACCESS_TOKEN]
  # id-token: write
jobs:
  scan-repository:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        # The repository scanning will be triggered periodically on the following branches.
        branch: ["dev"]
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

      # [Mandatory if using OIDC authentication protocol instead of JF_ACCESS_TOKEN]
      # Insert to oidc-provider-name the 'Provider Name' defined in the OIDC integration configured in the JPD
      # with:
      #   oidc-provider-name: ""
```
{% endcode %}

</details>
