# Setup Frogbot

### Setting up Frogbot

<details>

<summary>1. Set up Frogbot on your preferred CI server</summary>

* GitHub Actions
* Jenkins
* JFrog Pipelines
* GitLab Pipelines
* Azure Pipelines

</details>

<details>

<summary>2. Optionally set up a FREE JFrog Environment in the Cloud</summary>

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

<summary>3. Advanced - Customize advanced settings with frogbot-config.yml</summary>

* Creating the frogbot-config.yml File

</details>
