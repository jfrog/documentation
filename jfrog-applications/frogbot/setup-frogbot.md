# Setup Frogbot

#### What's needed for the setup?

* **JFrog Platform** server. (If you don't have a JFrog Platform, you can set up one for free)
  
* **CI server** to run the scan tasks.

#### Select your preferred CI server:

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/github-actions-icon.png" alt="GitHubActions" /> [GitHub Actions](./setup-frogbot-using-github-actions.md)

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/jenkins-icon.png" alt="Jenkins" /> [Jenkins](./setup-frogbot-using-jenkins.md)

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/jfrog-pipelines-icon.png" alt="jfrogpipelines" /> [JFrog Pipelines](./setup-frogbot-using-jfrog-pipelines.md)

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/gitlab-icon.png" alt="Gitlab" /> [GitLab CI](./setup-frogbot-using-gitlab-ci.md)

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/azure-pipelines-icon.png" alt="AzurePipelines" /> [Azure Pipelines](./setup-frogbot-using-azure-pipelines.md)


<details>
  
<summary>Optionally - set up a FREE JFrog Platform in the Cloud</summary>

Frogbot requires a JFrog environment to scan your projects. If you don't have an environment, we can set up a free environment in the cloud for you. Just run one of the following commands in your terminal to set up an environment in less than a minute.

The commands will do the following:

1. Install [JFrog CLI](https://www.jfrog.com/confluence/display/CLI/JFrog+CLI) on your machine.
2. Create a FREE JFrog environment in the cloud for you.

**For macOS and Linux, use curl**

```
curl -fL "https://getcli.jfrog.io?setup" | sh
```

**For Windows, use PowerShell**

```
powershell "Start-Process -Wait -Verb RunAs powershell '-NoProfile iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/[RELEASE]/jfrog-cli-windows-amd64/jf.exe -OutFile $env:SYSTEMROOT\system32\jf.exe'" ; jf setup
```

After the setup is complete, you'll receive an email with your JFrog environment connection details, which can be stored as secrets in Git.

</details>

<details>

<summary>Advanced - Customize advanced settings with frogbot-config.yml</summary>

* [Creating the frogbot-config.yml file](./frogbot-configuration.md)

</details>

