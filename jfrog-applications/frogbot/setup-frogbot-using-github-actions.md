# Setup Frogbot Using GitHub Actions

### Github Prerequisites

* Go to your repository's **settings** tab and save the JFrog connection details as repository secrets with the following names:
  * **JF\_URL** (JFrog Platform URL - Example: `https://acme.jfrog.io`)
  * **JF\_ACCESS\_TOKEN** (JFrog access token)

> You can also use **JF\_XRAY\_URL** and **JF\_ARTIFACTORY\_URL** instead of **JF\_URL**, and **JF\_USER** + **JF\_PASSWORD** instead of **JF\_ACCESS\_TOKEN**

![](../.gitbook/assets/github-repository-secrets.png)

* Under **Actions** > **General**, check the **Allow GitHub Actions to create and approve pull requests** check box.

![](../.gitbook/assets/github-pr-permissions.png)

* For open-source projects: Create a new [GitHub environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#creating-an-environment) called **frogbot** and add people or public teams as reviewers. The chosen reviewers can trigger Frogbot scans on pull requests.

![](../../../.gitbook/assets/github-environment.png)

### Frogbot GitHub Action Templates

1. Begin by cloning the GitHub repository to your local environment.

2. Switch to the target branch where you'd like the pull requests to be scanned.

3. Create a file named **frogbot-scan-pull-request.yml**. Fill it with the provided [template](templates/github-actions/frogbot-scan-pull-request.yml), and then push it into the **.github/workflows** directory at the root of your GitHub repository.

4. Return to the default branch.

5. Now, create a file named **frogbot-scan-repository.yml**. Again, populate it with the provided [template](templates/github-actions/frogbot-scan-repository.yml) and push it into the **.github/workflows** directory at the root of your GitHub repository.
