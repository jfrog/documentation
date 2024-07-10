# GitLab Templates for JFrog

This repository includes pipeline templates for GitLab CI, for a quick and easy integration with the [JFrog Platform](https://jfrog.com/platform/).

The templates use the [.setup-jfrog.yml](https://github.com/jfrog/jfrog-cli/blob/v2/build/gitlab/v2/) pipeline scripts. The script is included by each of the templates, and sets up the integration between the pipeline and the JFrog Platform.

The script does the following:

* Installs [JFrog CLI](https://www.jfrog.com/confluence/display/CLI/JFrog+CLI)
* Configures JFrog CLI to work with the JFrog Platform
* Sets the build name and build number values with the values of `$CI_PROJECT_PATH_SLUG-$CI_COMMIT_REF_NAME` and `$CI_PIPELINE_ID` respectively, to allow publishing [build-info](https://www.buildinfo.org/) to Artifactory
* Optionally replaces the default Docker Registry with an [Artifactory Docker Registry](https://www.jfrog.com/confluence/display/JFROG/Docker+Registry)

### Installation

1. Ensure you have the connection details for the JFrog Platform.&#x20;
2. [Store](broken-reference/) the JFrog Platform connection details on GitLab
3. Optionally set the URL of your Artifactory Docker Registry as the value of the **JF\_DOCKER\_REGISTRY** variable
4. [Add](broken-reference/) the **setup-jfrog** pipeline script in your GitLab pipeline

#### Storing the JFrog Platform Connection Details

Store the connection details of your JFrog Platform as [GitLab CI/CD variables](https://docs.gitlab.com/ee/ci/variables/) by using one of the following variables combinations:

1. **JF\_URL** - Anonymous access (no authentication)
2. **JF\_URL** + **JF\_USER** + **JF\_PASSWORD** - Basic authentication
3. **JF\_URL** + **JF\_ACCESS\_TOKEN** - Authentication with JFrog Access Token. NOTE: When pulling and pushing docker images from/to Artifactory, the **JF\_USER** variable is also required, in addition to the **JF\_ACCESS\_TOKEN** variables

#### Adding the setup-jfrog Script in Your Pipeline

**Including the Script**

The templates included in this repository already have the **setup-jfrog** script included as follows:

```yaml
include:
  - remote: 'https://releases.jfrog.io/artifactory/jfrog-cli/gitlab/v2/.setup-jfrog-unix.yml'
```

For Windows agents, use:

```yaml
include:
  - remote: 'https://releases.jfrog.io/artifactory/jfrog-cli/gitlab/v2/.setup-jfrog-windows.yml'
```

You also have the option of downloading the matching script from [releases.jfrog.io](https://releases.jfrog.io/artifactory/jfrog-cli/gitlab/v2/), adding it to your project, and including it in your pipeline as follows:

```yaml
  - local: '.setup-jfrog-unix.yml'
```

You can also include it from one of your projects as follows:

```yaml
  - project: 'my-group/my-project'
    file: '/script/.setup-jfrog-unix.yml'
```

**Referencing the Script**

Once the script is included in your pipeline, you'll need to reference it from any `script` or `before_script` sections in the pipeline as shown below:

```yaml
job:
  script:
    - !reference [.setup_jfrog, script]
```

At the end of your `script`, or as part of `after_script`, you should add the cleanup reference:

```yaml
job:
  after_script:
    - !reference [.cleanup_jfrog, script]
```

**Downloading the setup-jfrog script and JFrog CLI from Artifactory**

If your GitLab environment is air-gapped, you would want your pipeline to avoid downloading the **setup-jfrog** script and also JFrog CLI from `https://releases.jfrog.io/artifactory`. Here's how you do this:

As shown in the above [Including the Script](broken-reference/) and [Referencing the Script](broken-reference/) sections, you have the option of copying the **setup-jfrog** script into your pipeline, and thus avoiding its download. Since the **setup-jfrog** script downloads JFrog CLI from `https://releases.jfrog.io/artifactory`, you should also configure the script to download JFrog CLI from a remote repository in your JFrog Artifactory instance. Follow these steps to have JFrog CLI downloaded from your Artifactory instance:

1. Create a remote generic repository in Artifactory pointing to `https://releases.jfrog.io/artifactory/`
2. Add the **JF\_RELEASES\_REPO** variable to GitLab with the name of the repository you created

#### Additional Optional Variables

Configurations can be done via Project Settings > CI/CD > Variables:

| Variable                   | Usage                                                                                                                                                                                                            |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| JF\_DOCKER\_REGISTRY       | Docker registry in Artifactory. For more info, see [Getting Started with Artifactory as a Docker Registry](https://www.jfrog.com/confluence/display/JFROG/Getting+Started+with+Artifactory+as+a+Docker+Registry) |
| JFROG\_CLI\_BUILD\_PROJECT | JFrog project key to be used by commands which expect build name and build number. Determines the project of the published build.                                                                                |
| JFROG\_CLI\_VERSION        | Use a specific JFrog CLI version instead of the latest version. The minimal version allowed is: 2.17.0                                                                                                           |

See more environment variables in the JFrog CLI [documentation](https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory#CLIforJFrogArtifactory-EnvironmentVariables).

#### Prerequisites

* For Linux / Mac: **cURL**

#### Limitations

1. If the `JF_DOCKER_REGISTRY` and `JF_ACCESS_TOKEN` variables are set, then the `JF_USER` variable is required.
2. Build info collection is unavailable when:
   * Working with a docker registry without JFrog CLI.
   * Running separate jobs on temporary agents or docker containers.

#### Behind the Scenes

The **setup-jfrog** scripts are maintained in the [jfrog-cli repository](https://github.com/jfrog/jfrog-cli/blob/v2/build/gitlab/v2/). Each yaml includes two hidden jobs with scripts named `.setup_jfrog` and `.cleanup_jfrog`, which can be referenced by the pipeline after the script is included.

### Templates

#### Build and Upload to JFrog Artifactory

* .NET
* Go
* Gradle
* Maven
* npm
* NuGet
* Pip
* Pipenv
* Yarn Berry

```
$ jf mvn clean install
18:26:42 [🔵Info] Running Mvn...
18:26:58 [🔵Info] These files were uploaded:

📦 libs-snapshot-local
└── 📁 org
    └── 📁 jfrog
        └── 📁 test
            ├── 📁 multi2
            │   └── 📁 3.7-SNAPSHOT
            │       ├── 📄 multi2-3.7-20230227.162644-67.jar
            │       └── 📄 multi2-3.7-20230227.162644-67.pom
            ├── 📁 multi1
            │   └── 📁 3.7-SNAPSHOT
            │       ├── 📄 multi1-3.7-20230227.162644-69.pom
            │       ├── 📄 multi1-3.7-20230227.162644-69-tests.jar
            │       ├── 📄 multi1-3.7-20230227.162644-69-sources.jar
            │       └── 📄 multi1-3.7-20230227.162644-69.jar
            ├── 📁 multi
            │   └── 📁 3.7-SNAPSHOT
            │       └── 📄 multi-3.7-20230227.162644-67.pom
            └── 📁 multi3
                └── 📁 3.7-SNAPSHOT
                    ├── 📄 multi3-3.7-20230227.162644-67.pom
                    └── 📄 multi3-3.7-20230227.162644-67.war
```

#### Security Audit with JFrog Xray

* .NET
* Go
* Gradle
* Maven
* npm
* NuGet
* Pip
* Pipenv
* Yarn Berry
