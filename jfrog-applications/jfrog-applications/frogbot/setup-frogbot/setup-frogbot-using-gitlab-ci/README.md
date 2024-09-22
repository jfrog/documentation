---
description: Install Frogbot on GitLab repositories using GitLab CI
---

# Setup Frogbot Using GitLab CI

## Prepare GitLab to work with Frogbot

1. Make sure you have the connection details of your JFrog environment.
2. Go to your GitLab repository settings page and save the JFrog connection details as repository secrets with the following names - **JF\_URL**, **JF\_USER**, and **JF\_PASSWORD.**

> _**NOTE:**_
>
> * You can use **JF\_XRAY\_URL** and **JF\_ARTIFACTORY\_URL** instead of **JF\_URL**.
> * You can use **JF\_ACCESS\_TOKEN** instead of **JF\_USER** and **JF\_PASSWORD**.
> * Ensure not set these tokens as **protected** in Gitlab.

3. Add a job named **frogbot-scan** to your `.gitlab-ci.yml` file in your GitLab repository. Use the following for execution:

<details>

<summary>Frogbot template to scan Repository and Pull Request</summary>

```yml
frogbot-scan:
  rules:
    - if: $CI_PIPELINE_SOURCE == 'merge_request_event'
      when: manual
      variables:
        FROGBOT_CMD: "scan-pull-request"
        JF_GIT_BASE_BRANCH: $CI_MERGE_REQUEST_TARGET_BRANCH_NAME
      # Repository scanning is triggered by any push to the default branch.
      # If you'd like a different branch to be scanned, replace $CI_DEFAULT_BRANCH in the line below with the name of the branch, wrapped with quotes (").
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH || $CI_PIPELINE_SOURCE == "schedule"
      variables:
        FROGBOT_CMD: "scan-repository"
        JF_GIT_BASE_BRANCH: $CI_COMMIT_BRANCH
  variables:
    # [Mandatory]
    # JFrog platform URL (This functionality requires version 3.29.0 or above of Xray)
    JF_URL: $JF_URL

    # [Mandatory if JF_USER and JF_PASSWORD are not provided]
    # JFrog access token with 'read' permissions for Xray
    JF_ACCESS_TOKEN: $JF_ACCESS_TOKEN

    # [Mandatory if JF_ACCESS_TOKEN is not provided]
    # JFrog user and password with 'read' permissions for Xray
    # JF_USER: $JF_USER
    # JF_PASSWORD: $JF_PASSWORD

    # [Mandatory]
    # GitLab access token. Ensure the token has the following permissions, depedending on your GiLab deployment type:
    # Self hosted: api, read_api, read_user, read_repository.
    # Cloud: api, read_api, read_repository
    JF_GIT_TOKEN: $USER_TOKEN

    # Predefined GitLab variables. There's no need to set them.
    JF_GIT_PROVIDER: gitlab
    JF_GIT_OWNER: $CI_PROJECT_NAMESPACE
    JF_GIT_REPO: $CI_PROJECT_NAME
    JF_GIT_PULL_REQUEST_ID: $CI_MERGE_REQUEST_IID

    # [Mandatory if the two conditions below are met]
    # 1. The project uses yarn 2, NuGet, or .NET to download its dependencies
    # 2. The `installCommand` variable isn't set in your frogbot-config.yml file.
    #
    # The command that installs the project dependencies (e.g "nuget restore")
    JF_INSTALL_DEPS_CMD: ""
    
  script:
    # For Linux / MacOS runner:
    - |
      getFrogbotScriptPath=$(if [ -z "$JF_RELEASES_REPO" ]; then echo "https://releases.jfrog.io"; else echo "${JF_URL}/artifactory/${JF_RELEASES_REPO}"; fi)
      curl -fLg "$getFrogbotScriptPath/artifactory/frogbot/v2/[RELEASE]/getFrogbot.sh" | sh
      ./frogbot ${FROGBOT_CMD}

    # For Windows runner:
    # 
    # - $getFrogbotScriptPath = $(if ([string]::IsNullOrEmpty($env:JF_RELEASES_REPO)) { "https://releases.jfrog.io" } else { "$($env:JF_URL)/artifactory/$($env:JF_RELEASES_REPO)" })
    # - Invoke-WebRequest -Uri "$getFrogbotScriptPath/artifactory/frogbot/v2/[RELEASE]/getFrogbot.sh" -UseBasicParsing | ForEach-Object { & $_.Content }
    # - .\frogbot ${FROGBOT_CMD}
```

</details>

{% hint style="info" %}
In the `gitlab-ci.yml` file, Make sure that either **JF\_USER** and **JF\_PASSWORD** or **JF\_ACCESS\_TOKEN** are set, **but not both**.

For more advanced configuration, use [GitLab full template](gitlab-full-template.md) to see all available options.
{% endhint %}
