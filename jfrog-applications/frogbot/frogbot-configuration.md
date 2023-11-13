# Frogbot Configuration

## Creating the frogbot-config.yml file

### What is the frogbot-config.yml file?

The frogbot-config.yml file includes configuration related to your projects, to help Frogbot scan your Git repositories.

### Is the frogbot-config.yml file mandatory?

Not all projects require the **frogbot-config.yml** file, but any project can use it.

If your project doesn't use a **frogbot-config.yml** file, all of the configuration Frogbot requires\
should be provided as variables as part of the Frogbot workflows.

### How does the frogbot-config.yml file helps Frogbot scan the repository?

When your Git repository includes multiple subprojects, and each subproject has its own descriptor file (package.json in the case of npm), Frogbot will attempt to detect these projects recursively and scan them. In cases where the detection is not accurate, you can use the **frogbot-config.yml** file to include the relative paths to the subprojects. Frogbot uses this configuration to scan each subproject separately. In the following example, there are two subprojects under `path/to/project-1` and `path/to/project-2`.

```yaml
- params:
    git:
      repoName: my-git-repo-name
      branches:
        - master
    scan:
      projects:
        - workingDirs:
            - path/to/npm/project-1
            - path/to/npm/project-2
```

Here's another example. Notice that projects which use the nuget client to download the dependencies, the download command needs to be specified.

```yaml
- params:
    git:
      repoName: my-git-repo-name
      branches:
        - master
    scan:
      projects:
        - workingDirs:
            - path/to/node/project
        - installCommand: nuget restore
          workingDirs:
            - path/to/.net/project
```

See the full **frogbot-config.yml** structure here.

### Can one frogbot-config.yml file be used for multiple Git repositories?

You have the option of using a single **frogbot-config.yml** file for scanning multiple Git repositories in the same organization if one of the following platforms is used.

* GitHub with Jenkins or JFrog Pipelines
* Bitbucket Server
* Azure Repos

The file can be placed in any repository if it's in the same organization as all the repositories referenced in the file. Here's an example of a **frogbot-config.yml** referencing multiple repositories.

```yaml
- params:
    git:
      repoName: repo-1
      branches:
        - master
- params:
    git:
      repoName: repo-2
      branches:
        - master
        - dev
- params:
    git:
      repoName: repo-3
      branches:
        - master
    scan:
      projects:
        - pipRequirementsFile: requirements.txt
```

If however you're using one of the following platforms, each repository that needs to be scanned by Frogbot should include its own **frogbot-config.yml** file.

* GitHub with GitHub actions
* GitLab

### Where should the frogbot-config.yml file be placed in the repository?

Frogbot expects the frogbot-config.yml file to be in the following path from the root of the Git repository: `.frogbot/frogbot-config.yml`.

**IMPORTANT**: The `frogbot-config.yml` file must be pushed to the target branch before it can be used by Frogbot. This means that if, for example, a pull request includes the `frogbot-config.yml` and the target branch doesn't, the file will be ignored.

### The frogbot-config.yml file structure

See the complete content and structure of the **frogbot-config.yml** file [here](templates/.frogbot/frogbot-config.yml).
