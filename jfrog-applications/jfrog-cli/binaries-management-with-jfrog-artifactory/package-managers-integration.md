# Package Managers Integration

## Running Maven Builds

JFrog CLI integrates with Maven, which you can use to resolve dependencies and deploy build artifacts from and to Artifactory, while collecting build-info and storing it in Artifactory.

### Setting maven repositories

Before you use the `jf mvn` command, you must first configure the project with the Artifactory server and repositories to be used for building and publishing. Use the `jf mvn-config` command once to add this configuration to the project. You should run the command from inside the root directory of the project. The command stores the configuration in the `.jfrog` directory at the project's root.

|                             |                                                                                                                                                                                                                                                                   |
| --------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name                | mvn-config                                                                                                                                                                                                                                                        |
| Abbreviation                | mvnc                                                                                                                                                                                                                                                              |
| **Command options:**        |                                                                                                                                                                                                                                                                   |
| `--global`                  | <p>[Optional]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p>                                                                                        |
| `--server-id-resolve`       | <p>[Optional]<br>Server ID for resolution. The server should configured using the 'jf rt c' command.</p>                                                                                                                                                          |
| `--server-id-deploy`        | <p>[Optional]<br>Server ID for deployment. The server should be configured using the 'jf rt c' command.</p>                                                                                                                                                       |
| `--repo-resolve-releases`   | <p>[Optional]<br>Resolution repository for release dependencies.</p>                                                                                                                                                                                              |
| `--repo-resolve-snapshots`  | <p>[Optional]<br>Resolution repository for snapshot dependencies.</p>                                                                                                                                                                                             |
| `--repo-deploy-releases`    | <p>[Optional]<br>Deployment repository for release artifacts.</p>                                                                                                                                                                                                 |
| `--repo-deploy-snapshots`   | <p>[Optional]<br>Deployment repository for snapshot artifacts.</p>                                                                                                                                                                                                |
| `--include-patterns`        | <p>[Optional]<br>Filter deployed artifacts by setting a wildcard pattern that specifies which artifacts to include. You may provide multiple comma-separated(,) patterns followed by a white-space. For example<br><br>artifact-<em>.jar, artifact-</em>.pom</p>  |
| `--exclude-patterns`        | <p>[Optional]<br>Filter deployed artifacts by setting a wildcard pattern that specifies which artifacts to exclude. You may provide multiple comma-separated(,) followed by a white-space. For example<br><br>artifact-<em>-test.jar, artifact-</em>-test.pom</p> |
| `--disable-snapshots`       | <p>[Default:false]<br>Set to true to disable snapshot resolution.</p>                                                                                                                                                                                             |
| `--snapshots-update-policy` | <p>[Optional]<br>Set snapshot update policy. Defaults to daily.</p>                                                                                                                                                                                               |
| **Command arguments:**      | The command accepts no arguments                                                                                                                                                                                                                                  |

#### Examples

Before using `jf mvn-config`, you must first configure your Artifactory server with JFrog CLI using the `jf c add` command. For instance:

```
jf c add my-artifactory-server --url=your-artifactory-url.jfrog.io --user=your-user --password=your-password
```

Replace `my-artifactory-server`, `your-artifactory-url.jfrog.io`, `your-user`, and `your-password` with your actual Artifactory instance details.

Once your Artifactory server is configured, you can set your Maven repositories from your project's root directory.

Example 1: Setting Project-Specific Repositories This example defines the resolution and deployment repositories for the current project.

```
jf mvn-config \
    --server-id-resolve=my-artifactory-server \
    --repo-resolve-releases=maven-virtual-releases \
    --repo-resolve-snapshots=maven-virtual-snapshots \
    --server-id-deploy=my-artifactory-server \
    --repo-deploy-releases=maven-releases-local \
    --repo-deploy-snapshots=maven-snapshots-local
```

* my-artifactory-server: The server ID you configured using `jf c add`.
* maven-virtual-releases: Your Artifactory repository for resolving release dependencies.
* maven-virtual-snapshots: Your Artifactory repository for resolving snapshot dependencies.
* maven-releases-local: Your local Artifactory repository for deploying release artifacts.
* maven-snapshots-local: Your local Artifactory repository for deploying snapshot artifacts.

### Running maven

The **mvn** command triggers the maven client, while resolving dependencies and deploying artifacts from and to Artifactory.

> **Note**: Before running the **mvn** command on a project for the first time, the project should be configured with the **jf mvn-config** command.

> **Note**: If the machine running JFrog CLI has no access to the internet, make sure to read the [Downloading the Maven and Gradle Extractor JARs](package-managers-integration.md#downloading-the-maven-and-gradle-extractor-jars) section.

#### Commands Params

The following table lists the command arguments and flags:

|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | mvn                                                                                                                                                                                                                                      |
| Abbreviation           | mvn                                                                                                                                                                                                                                      |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--threads`            | <p>[Default: 3]<br>Number of threads for uploading build artifacts.</p>                                                                                                                                                                  |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--insecure-tls`       | <p>[Default: false]<br>Set to true to skip TLS certificates verification.</p>                                                                                                                                                            |
| `--scan`               | <p>[Default: false]<br>Set if you'd like all files to be scanned by Xray on the local file system prior to the upload, and skip the upload if any of the files are found vulnerable.</p>                                                 |
| `--format`             | <p>[Default: table]<br>Should be used with the --scan option. Defines the scan output format. Accepts table or json as values.</p>                                                                                                       |
| **Command arguments:** | The command accepts the same arguments and options as the mvn client.                                                                                                                                                                    |

#### Deploying Maven Artifacts

The deployment to Artifacts is triggered both by the deployment and install phases. To disable artifacts deployment, add **-Dartifactory.publish.artifacts=false** to the list of goals and options. For example: "**jf mvn clean install -Dartifactory.publish.artifacts=false**"

#### Example

Run clean and install with maven.

```
jf mvn clean install -f /path/to/pom.xml
```

## Running Gradle Builds

JFrog CLI integrates with Gradle, which you can use to resolve dependencies and deploy build artifacts from and to Artifactory, while collecting build-info and storing it in Artifactory.

### Setting Gradle Repositories

Before using the `jf gradle` command, you must configure the project with the Artifactory server and repositories for building and publishing. The `jf gradle-config` command is used once per project to add this configuration. Run the command from the root directory of the project. The configuration is stored by the command in the `.jfrog` directory at the project's root.

|                           |                                                                                                                                                                            |
| ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name              | gradle-config                                                                                                                                                              |
| Abbreviation              | gradlec                                                                                                                                                                    |
| **Command options:**      |                                                                                                                                                                            |
| `--global`                | <p>[Optional]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p> |
| `--server-id-resolve`     | <p>[Optional]<br>Server ID for resolution. The server should configured using the 'jf c add' command.</p>                                                                  |
| `--server-id-deploy`      | <p>[Optional]<br>Server ID for deployment. The server should be configured using the 'jf c add' command.</p>                                                               |
| `--repo-resolve`          | <p>[Optional]<br>Repository for dependencies resolution.</p>                                                                                                               |
| `--repo-deploy`           | <p>[Optional]<br>Repository for artifacts deployment.</p>                                                                                                                  |
| `--uses-plugin`           | <p>[Default: false]<br>Set to true if the Gradle Artifactory Plugin is already applied in the build script.</p>                                                            |
| `--use-wrapper`           | <p>[Default: false]<br>Set to true if you'd like to use the Gradle wrapper.</p>                                                                                            |
| `--deploy-maven-desc`     | <p>[Default: true]<br>Set to false if you do not wish to deploy Maven descriptors.</p>                                                                                     |
| `--deploy-ivy-desc`       | <p>[Default: true]<br>Set to false if you do not wish to deploy Ivy descriptors.</p>                                                                                       |
| `--ivy-desc-pattern`      | <p>[Default: '[organization]/[module]/ivy-[revision].xml'<br><br>Set the deployed Ivy descriptor pattern.</p>                                                              |
| `--ivy-artifacts-pattern` | <p>[Default: '[organization]/[module]/[revision]/[artifact]-<a href="-[classifier]/">revision</a>.[ext]'<br><br>Set the deployed Ivy artifacts pattern.</p>                |
| **Command arguments:**    | The command accepts no arguments                                                                                                                                           |

### Running gradle

The `jf gradle` command runs the Gradle client, resolving dependencies and deploying artifacts from and to Artifactory using the settings defined by `gradle-config`.

**Note:** Before running the `jf gradle` command on a project for the first time, you must configure the project with the `jf gradle-config` command.

**Note:** If the machine running JFrog CLI does not have internet access, be sure to read the Downloading the Maven and Gradle Extractor JARs section.

#### Commands Params

The following table lists the command arguments and flags:

|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | gradle                                                                                                                                                                                                                                   |
| Abbreviation           | gradle                                                                                                                                                                                                                                   |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--threads`            | <p>[Default: 3]<br>Number of threads for uploading build artifacts.</p>                                                                                                                                                                  |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--scan`               | <p>[Default: false]<br>Set if you'd like all files to be scanned by Xray on the local file system prior to the upload, and skip the upload if any of the files are found vulnerable.</p>                                                 |
| `--format`             | <p>[Default: table]<br>Should be used with the --scan option. Defines the scan output format. Accepts table or json as values.</p>                                                                                                       |
| **Command arguments:** | The command accepts the same arguments and options as the gradle client.                                                                                                                                                                 |

#### Example

Build the project using the `artifactoryPublish` task, while resolving dependencies from and deploying artifacts to Artifactory.

```
jf gradle clean artifactoryPublish -b path/to/build.gradle
```

## Downloading the Maven and Gradle Extractor JARs

For integration with Maven and Gradle, JFrog CLI uses the `build-info-extractor` JAR files. These JAR files are downloaded by JFrog CLI from jcenter the first time they are needed.

If you are using JFrog CLI on a machine that has no access to the internet, you can configure it to download these JAR files from an Artifactory instance instead.

To configure JFrog CLI to download the extractor JARs from Artifactory:

1. Create a remote Maven repository in Artifactory and name it `extractors`. When creating the repository, configure it to proxy the following URL: `https://releases.jfrog.io/artifactory/oss-release-local`
2. Verify that this Artifactory server is configured in JFrog CLI by using the `jf c show` command. If it is not, configure it using the `jf c add` command.
3. Set the `JFROG_CLI_EXTRACTORS_REMOTE` environment variable. The value should be the Server ID of the configured Artifactory server, followed by a slash, and then the name of the repository you created. For example: `my-rt-server/extractors`.

## Running Builds with MSBuild

JFrog CLI includes integration with MSBuild and Artifactory, allowing you to resolve dependencies and deploy build artifacts from and to Artifactory, while collecting build-info and storing it in Artifactory. This is done by having JFrog CLI in your search path and adding JFrog CLI commands to the MSBuild `csproj` file.

For detailed instructions, please refer to our [MSBuild Project Example](https://github.com/eyalbe4/project-examples/tree/master/msbuild-example) on GitHub.

## Managing Docker Images

JFrog CLI provides full support for pulling and publishing docker images from and to Artifactory using the docker client running on the same machine. This allows you to collect build-info for your docker build and then publish it to Artifactory. You can also promote the pushed docker images from one repository to another in Artifactory.

To build and push your docker images to Artifactory, follow these steps:

1. Make sure Artifactory can be used as docker registry. Please refer to [Getting Started with Docker and Artifactory](https://jfrog.com/help/r/jfrog-artifactory-documentation/Getting-Started-With-Artifactory-As-A-docker-Registry) in the JFrog Artifactory User Guide.
2. Make sure that the installed docker client has version **17.07.0-ce (2017-08-29)** or above. To verify this, run **docker -v**\*\*
3. To ensure that the docker client and your Artifactory docker registry are correctly configured to work together, run the following code snippet.

```
docker pull hello-world
docker tag hello-world:latest &lt;artifactoryDockerRegistry&gt;/hello-world:latest
docker login &lt;artifactoryDockerRegistry&gt;
docker push &lt;artifactoryDockerRegistry&gt;/hello-world:latest
```

If everything is configured correctly, pushing any image including the hello-world image should be successfully uploaded to Artifactory.

> **Note**: When running the docker-pull and docker-push commands, the CLI will first attempt to log in to the docker registry. In case of a login failure, the command will not be executed.

### Examples

Check out our [docker project examples on GitHub](https://github.com/jfrog/project-examples/tree/master/docker-oci-examples).

### Pulling Docker Images Using the Docker Client

Running **jf docker pull** command allows pulling docker images from Artifactory, while collecting the build-info and storing it locally, so that it can be later published to Artifactory, using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

#### Commands Params

The following table lists the command arguments and flags:

|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | docker pull                                                                                                                                                                                                                              |
| Abbreviation           | dpl                                                                                                                                                                                                                                      |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--server-id`          | <p>[Optional]<br>Server ID configured using the 'jf config' command. If not specified, the default configured Artifactory server is used.</p>                                                                                            |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`             | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--skip-login`         | <p>[Default: false]<br>Set to true if you'd like the command to skip performing docker login.</p>                                                                                                                                        |
| **Command arguments:** | The same arguments and options supported by the docker client/                                                                                                                                                                           |

#### Example

The subsequent command utilizes the docker client to pull the 'my-docker-registry.io/my-docker-image:latest' image from Artifactory. This operation logs the image layers as dependencies of the local build-info identified by the build name 'my-build-name' and build number '7'. This local build-info can subsequently be released to Artifactory using the command 'jf rt bp my-build-name 7'.

```
jf docker pull my-docker-registry.io/my-docker-image:latest --build-name=my-build-name --build-number=7
```

You can then publish the build-info collected by the **jf docker pull** command to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

### Pushing Docker Images Using the Docker Client

After building your image using the docker client, the **jf docker push** command pushes the image layers to Artifactory, while collecting the build-info and storing it locally, so that it can be later published to Artifactory, using the **jf rt build-publish** command.

#### Commands Params

The following table lists the command arguments and flags:

|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | docker push                                                                                                                                                                                                                              |
| Abbreviation           | dp                                                                                                                                                                                                                                       |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--server-id`          | <p>[Optional]<br>Server ID configured using the 'jf config' command. If not specified, the default configured Artifactory server is used.</p>                                                                                            |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`             | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--skip-login`         | <p>[Default: false]<br>Set to true if you'd like the command to skip performing docker login.</p>                                                                                                                                        |
| `--threads`            | <p>[Default: 3]<br>Number of working threads.</p>                                                                                                                                                                                        |
| `--detailed-summary`   | <p>[Default: false]<br>Set true to include a list of the affected files as part of the command output summary.</p>                                                                                                                       |
| **Command arguments:** | The same arguments and options supported by the docker client/                                                                                                                                                                           |

#### Example

The subsequent command utilizes the docker client to push the 'my-docker-registry.io/my-docker-image:latest' image to Artifactory. This operation logs the image layers as artifacts of the local build-info identified by the build name 'my-build-name' and build number '7'. This local build-info can subsequently be released to Artifactory using the command 'jf rt bp my-build-name 7'.

```
jf docker push my-docker-registry.io/my-docker-image:latest --build-name=my-build-name --build-number=7
```

You can then publish the build-info collected by the **docker-push** command to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

### Pulling Docker Images Using Podman

[Podman](https://podman.io/) is a daemon-less container engine for developing, managing, and running OCI Containers. Running the **podman-pull** command allows pulling docker images from Artifactory using podman, while collecting the build-info and storing it locally, so that it can be later published to Artifactory, using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

#### Commands Params

The following table lists the command arguments and flags:

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | rt podman-pull                                                                                                                                                                                                                           |
| Abbreviation         | rt ppl                                                                                                                                                                                                                                   |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--server-id`        | <p>[Optional]<br>Server ID configured using the 'jf config' command. If not specified, the default configured Artifactory server is used.</p>                                                                                            |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--skip-login`       | <p>[Default: false]<br>Set to true if you'd like the command to skip performing docker login.</p>                                                                                                                                        |
| Command argument     |                                                                                                                                                                                                                                          |
| Image tag            | The docker image tag to pull.                                                                                                                                                                                                            |
| Source repository    | Source repository in Artifactory.                                                                                                                                                                                                        |

#### Example

In this example, podman is employed to pull the local image 'my-docker-registry.io/my-docker-image:latest' from the docker-local Artifactory repository. During this process, it registers the image layers as dependencies within a build-info identified by the build name 'my-build-name' and build number '7'. This build-info is initially established locally and must be subsequently published to Artifactory using the command 'jf rt build-publish my-build-name 7'.

```
jf rt podman-pull my-docker-registry.io/my-docker-image:latest docker-local --build-name my-build-name --build-number 7
```

You can then publish the build-info collected by the **podman-pull** command to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

### Pushing Docker Images Using Podman

[Podman](https://podman.io/) is a daemon-less container engine for developing, managing, and running OCI Containers. After building your image, the **podman-push** command pushes the image layers to Artifactory, while collecting the build-info and storing it locally, so that it can be later published to Artifactory, using the **build-publish** command.

#### Commands Params

The following table lists the command arguments and flags:

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | rt podman-push                                                                                                                                                                                                                           |
| Abbreviation         | rt pp                                                                                                                                                                                                                                    |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--server-id`        | <p>[Optional]<br>Server ID configured using the 'jf config' command. If not specified, the default configured Artifactory server is used.</p>                                                                                            |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--skip-login`       | <p>[Default: false]<br>Set to true if you'd like the command to skip performing docker login.</p>                                                                                                                                        |
| `--threads`          | <p>[Default: 3]<br>Number of working threads.</p>                                                                                                                                                                                        |
| `--detailed-summary` | <p>[Default: false]<br>Set to true to include a list of the affected files as part of the command output summary.</p>                                                                                                                    |
| Command argument     |                                                                                                                                                                                                                                          |
| Image tag            | The docker image tag to push.                                                                                                                                                                                                            |
| Target repository    | Target repository in Artifactory.                                                                                                                                                                                                        |

#### Example

In this illustration, podman is employed to push the local image 'my-docker-registry.io/my-docker-image:latest' to the docker-local Artifactory repository. During this process, it registers the image layers as artifacts within a build-info identified by the build name 'my-build-name' and build number '7'. This build-info is initially established locally and must be subsequently published to Artifactory using the command 'jf rt build-publish my-build-name 7'.

```
jf rt podman-push my-docker-registry.io/my-docker-image:latest docker-local --build-name=my-build-name --build-number=7
```

You can then publish the build-info collected by the **podman-push** command to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

### Pushing Docker Images Using Kaniko

JFrog CLI allows pushing containers to Artifactory using [Kaniko](https://github.com/GoogleContainerTools/kaniko#kaniko---build-images-in-kubernetes), while collecting build-info and storing it in Artifactory.\
For detailed instructions, please refer to our [Kaniko project example on GitHub](https://github.com/jfrog/project-examples/tree/master/docker-oci-examples/kaniko-example).

### Pushing Docker Images Using buildx

JFrog CLI allows pushing containers to Artifactory using [buildx](https://github.com/GoogleContainerTools/kaniko#kaniko---build-images-in-kubernetes), while collecting build-info and storing it in Artifactory.\
For detailed instructions, please refer to our [buildx project example on GitHub](https://github.com/jfrog/project-examples/tree/master/docker-oci-examples/fat-manifest-example).

### Pushing Docker Images Using the OpenShift CLI

JFrog CLI allows pushing containers to Artifactory using the [OpenShift CLI](https://docs.openshift.com/container-platform/4.2/cli_reference/openshift_cli/getting-started-cli.html), while collecting build-info and storing it in Artifactory.\
For detailed instructions, please refer to our [OpenShift build project example on GitHub](https://github.com/jfrog/project-examples/tree/master/docker-oci-examples/openshift-examples/openshift-build-example).

### Adding Published Docker Images to the Build-Info

The **build-docker-create** command allows adding a docker image, which is already published to Artifactory, into the build-info. This build-info can be later published to Artifactory, using the **build-publish** command.

#### Commands Params

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | rt build-docker-create                                                                                                                                                                                                                   |
| Abbreviation         | rt bdc                                                                                                                                                                                                                                   |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--image-file`       | <p>Path to a file which includes one line in the following format: IMAGE-TAG@sha256:MANIFEST-SHA256. For example:<br><br>cat image-file-details<br>superfrog-docker.jfrog.io/hello-frog@sha256:30f04e684493fb5ccc030969df6de0</p>        |
| `--server-id`        | <p>[Optional]<br>Server ID configured using the 'jf config' command. If not specified, the default configured Artifactory server is used.</p>                                                                                            |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--threads`          | <p>[Default: 3]<br>Number of working threads.</p>                                                                                                                                                                                        |
| Command argument     |                                                                                                                                                                                                                                          |
| Target repository    | The name of the repository to which the image was pushed.                                                                                                                                                                                |

> **Note:**\
> If your Docker image has **multiple tags** pointing to the same digest, you can provide them in a **comma-separated format** in the `--image-file`.\
> All listed tags will be processed and added to the build-info individually.

#### Example

In this example, a Docker image that has already been deployed to Artifactory is incorporated into a locally created, unpublished build-info identified by the build name `myBuild` and build number '1'. This local build-info can subsequently be published to Artifactory using the command 'jf rt bp myBuild 1'.

```
jf rt bdc docker-local --image-file image-file-details --build-name myBuild --build-number 1
```

You can then publish the build-info collected by the **podman-push** command to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

### Promoting Docker Images

Promotion is the action of moving or copying a group of artifacts from one repository to another, to support the artifacts' lifecycle. When it comes to docker images, there are two ways to promote a docker image which was pushed to Artifactory:

1. Create build-info for the docker image, and then promote the build using the **jf rt build-promote** command.
2. Use the **jf rt docker-promote** command as described below.

#### Commands Params

The following table lists the command arguments and flags:

|                         |                                                                                                                                               |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name            | rt docker-promote                                                                                                                             |
| Abbreviation            | rt dpr                                                                                                                                        |
| **Command options:**    |                                                                                                                                               |
| `--server-id`           | <p>[Optional]<br>Server ID configured using the 'jf config' command. If not specified, the default configured Artifactory server is used.</p> |
| `--copy`                | <p>[Default: false]<br>If set true, the Docker image is copied to the target repository, otherwise it is moved.</p>                           |
| `--source-tag`          | <p>[Optional]<br>The tag name to promote.</p>                                                                                                 |
| `--target-docker-image` | <p>[Optional]<br>Docker target image name.</p>                                                                                                |
| `--target-tag`          | <p>[Optional]<br>The target tag to assign the image after promotion.</p>                                                                      |
| Command argument        |                                                                                                                                               |
| source docker image     | The docker image name to promote.                                                                                                             |
| source repository       | Source repository in Artifactory.                                                                                                             |
| target repository       | Target repository in Artifactory.                                                                                                             |

#### Examples

Promote the **hello-world** docker image from the **docker-dev-local** repository to the **docker-staging-local** repository.

```
jf rt docker-promote hello-world docker-dev-local docker-staging-local
```

**Note:** The `jf rt docker-promote` command currently requires the source and target repositories to be different. It does not support promoting a Docker image to the same repository while assigning it a different target image name. If you need to perform this type of promotion, consider using the Artifactory REST API directly.

## Building Npm Packages Using the Npm Client

JFrog CLI provides full support for building npm packages using the npm client. This allows you to resolve npm dependencies, and publish your npm packages from and to Artifactory, while collecting build-info and storing it in Artifactory.

Follow these guidelines when building npm packages:

* You can download npm packages from any npm repository type - local, remote or virtual, but you can only publish to a local or virtual Artifactory repository, containing local repositories. To publish to a virtual repository, you first need to set a default local repository. For more details, please refer to [Deploying to a Virtual Repository.](https://jfrog.com/help/r/jfrog-artifactory-documentation/deploy-to-a-virtual-repository)

When building npm packages, it is important to understand how the jf npm publish command handles publishing scripts. The behavior differs based on whether the --run-native flag is used:

* **Default Behavior (Without the --run-native flag):** JFrog CLI runs the pack command in the background, which is followed by an upload action not based on the npm client's native publish command. Therefore, if your npm package includes prepublish or postpublish scripts, you must rename them to prepack and postpack respectively to ensure they are executed.
* **Behavior with the --run-native flag:** When this flag is used, the command utilizes the native npm client's own publish lifecycle. In this mode, standard npm script names such as prepublish, publish, and postpublish are handled directly by npm itself, and no renaming is necessary.



**Prerequisites**

* Npm client version 5.4.0 and above.
* Artifactory version 5.5.2 and above.

## Setting npm repositories

Before using the jf npm install, jf npm ci, and jf npm publish commands, the project needs to be pre-configured with the Artifactory server and repositories for building and publishing. The configuration method depends on your workflow:

* **Standard JFrog CLI Configuration:** The jf npm-config command should be used once to add the configuration to the project. This command should be run from the project's root directory and stores the configuration in the .jfrog directory.
* **Native Client Configuration (--run-native):** When the --run-native flag is used, JFrog CLI bypasses the configuration in the .jfrog directory. Instead, it uses the user's own .npmrc file for all configurations, including authentication tokens and other settings.



| **Command-name**                                            | **Abbreviation** | **Description**                                                                              |
| ----------------------------------------------------------- | ---------------- | -------------------------------------------------------------------------------------------- |
| <p><code>npm-config</code><br><br>Short form</p><p>npmc</p> | \`npmc\`         | Configures Artifactory server and repository details for npm builds within a project.        |
| `--server-id-resolve`                                       | <p><br></p>      | \[Optional] Artifactory server ID for dependency resolution (configured using \`jf c add\`). |
| `--server-id-deploy`                                        | <p><br></p>      | \[Optional] Artifactory server ID for artifact deployment (configured using \`jf c add\`).   |
| `--repo-resolve`                                            | <p><br></p>      | \[Optional] Repository for resolving dependencies.                                           |
| `--repo-deploy`                                             | <p><br></p>      | \[Optional] Repository for deploying artifacts.                                              |
| **Command arguments**                                       | <p><br></p>      | Accepts no arguments.                                                                        |

\


\


### Installing Npm Packages

The jf npm install and jf npm ci commands execute npm's install and ci commands respectively, to fetch the npm dependencies from the npm repositories.

Commands Params

The following table lists the command arguments and flags:

Command-name

npm (covers install and ci subcommands typically, for example jf npm install, jf npm ci)

Abbreviation

(No specific abbreviation listed for jf npm, but for the underlying commands like npm i)

\


|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | npm                                                                                                                                                                                                                                      |
| Abbreviation           |                                                                                                                                                                                                                                          |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`             | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--threads`            | <p>[Default: 3]<br>Number of working threads for build-info collection.</p>                                                                                                                                                              |
| **Command arguments:** | The command accepts the same arguments and options as the npm client.                                                                                                                                                                    |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`             | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| **Command arguments:** | The command accepts the same arguments and options as the npm client.                                                                                                                                                                    |

**Command options:**

* \--run-native \[Optional] \[Default: false] Set to true to use the native npm client and the user's existing .npmrc configuration file. When this flag is active, JFrog CLI will not create its own temporary .npmrc file. All configurations, including authentication, must be handled by the user's .npmrc file.

**Note**

The "deployment view" and "details summary" features are not supported by the jf npm install and jf npm ci commands. This limitation applies regardless of whether the --run-native flag is used.

####

#### Examples

**Example 1**

The following example installs the dependencies and records them locally as part of build my-build-name/1. The build-info can later be published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command. The dependencies are resolved from the Artifactory server and repository configured by npm-config command.

```
jf npm install --build-name=my-build-name --build-number=1
```

**Example 2**

The following example installs the dependencies. The dependencies are resolved from the Artifactory server and repository configured by **npm-config** command.

```
jf npm install
```

**Example 3**

The following example installs the dependencies using the npm-ci command. The dependencies are resolved from the Artifactory server and repository configured by npm-config command.

```
jf npm ci
```

**Example 4**

The following example installs dependencies using the native npm client, based on the .npmrc configuration.

```
jf npm install --run-native




Eg: jf npm install --run-native --build-name=my-native-build --build-number=1

```



### Publishing the Npm Packages into Artifactory

The npm-publish command packs and deploys the npm package to the designated npm repository.

Before running the npm-publish command on a project for the first time, the project should be configured using the jf npm-config command. This configuration includes the Artifactory server and repository to which the package should deploy.

When using the --run-native flag, the jf npm-config command and the resulting .jfrog directory configuration are bypassed. Instead, JFrog CLI uses the native npm client , which relies exclusively on the user's .npmrc file for all configurations. Therefore, you must ensure your .npmrc file is correctly configured for publishing to the desired Artifactory repository, including all necessary repository URLs and authentication details.

> **Warning**:&#x20;
>
> If your npm package includes the prepublish or postpublish scripts and you are not using the --run-native flag, please refer to the guidelines above (rename to prepack and postpack).
>
> When using --run-native, standard npm script names are respected by the npm client.

#### Commands Params

The following table lists the command arguments and flags:\


|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | `npm publish`                                                                                                                                                                                                                            |
| Abbreviation         |                                                                                                                                                                                                                                          |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--detailed-summary` | <p>[Default: false]<br>Set true to include a list of the affected files as part of the command output summary.</p>                                                                                                                       |
| `--scan`             | <p>[Default: false]<br>Set if you'd like all files to be scanned by Xray on the local file system prior to the upload, and skip the upload if any of the files are found vulnerable.</p>                                                 |
| `--format`           | <p>[Default: table]<br>Should be used with the --scan option. Defines the scan output format. Accepts table or JSON as values.</p>                                                                                                       |
| Command argument     | The command accepts the same arguments and options that the **npm pack** command expects.                                                                                                                                                |
| `--run-native`       | \[Optional] \[Default: false] Set to true to use the native npm client for publishing. This allows leveraging all features and configurations specified in the user's .npmrc file.                                                       |

**Note:**&#x20;

* Require a valid .npmrc file with appropriate configurations and authentication tokens.
* Performance: Using this flag may result in performance degradation compared to the default JFrog CLI publish mechanism (which uses multi-threading).
* Unsupported Features with this flag: "Deployment view" and "details summary" are not supported when this flag is used.

Command argument

The command accepts the same arguments and options that the npm pack command expects (when not using --run-native) or npm publish command expects (when using --run-native).

\


#### Example 1

To pack and publish the npm package and also record it locally as part of build my-build-name/1, run the following command. The build-info can later be published to Artifactory using the build-publish command. The package is published to the Artifactory server and repository configured by npm-config command.

```
jf npm publish --build-name=my-build-name --build-number=1
```

#### Example 2

Publishing an npm package using the native npm client and user's .npmrc.

```
jf npm publish --run-native
```

(Ensure your package.json and .npmrc are configured for publishing)

## Building Npm Packages Using the Yarn Client

JFrog CLI provides full support for building npm packages using the yarn client. This allows you to resolve npm dependencies, while collecting build-info and storing it in Artifactory. You can download npm packages from any npm repository type - local, remote or virtual. Publishing the packages to a local npm repository is supported through the **jf rt upload** command.

**Note:** "Yarn versions from 2.4.0 up to, but not including, Yarn 4.x are supported. Yarn 4.x is currently not supported by JFrog CLI."

### Setting npm repositories

Before using the **jf yarn** command, the project needs to be pre-configured with the Artifactory server and repositories, to be used for building the project. The **yarn-config** command should be used once to add the configuration to the project. The command should run while inside the root directory of the project. The configuration is stored by the command in the **.jfrog** directory at the root directory of the project.

|                        |                                                                                                                                                                            |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | yarn-config                                                                                                                                                                |
| Abbreviation           | yarnc                                                                                                                                                                      |
| **Command options:**   |                                                                                                                                                                            |
| `--global`             | <p>[Optional]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p> |
| `--server-id-resolve`  | <p>[Optional]<br>Artifactory server ID for resolution. The server should configured using the 'jf c add' command.</p>                                                      |
| `--repo-resolve`       | <p>[Optional]<br>Repository for dependencies resolution.</p>                                                                                                               |
| **Command arguments:** | The command accepts no arguments                                                                                                                                           |

### Installing Npm Packages

The **jf yarn** command executes the yarn client, to fetch the npm dependencies from the npm repositories.

> **Note**: Before running the command on a project for the first time, the project should be configured using the **jf yarn-config** command.

#### Commands Params

The following table lists the command arguments and flags:

|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | yarn                                                                                                                                                                                                                                     |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`             | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--threads`            | <p>[Default: 3]<br>Number of working threads for build-info collection.</p>                                                                                                                                                              |
| **Command arguments:** | The command accepts the same arguments and options as the yarn client.                                                                                                                                                                   |

#### Examples

**Example 1**

The following example installs the dependencies and records them locally as part of build **my-build-name/1**. The build-info can later be published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command. The dependencies are resolved from the Artifactory server and repository configured by \*\*yarn-config command.

```
jf yarn install --build-name=my-build-name --build-number=1
```

**Example 2**

The following example installs the dependencies. The dependencies are resolved from the Artifactory server and repository configured by **jf yarn-config** command.

```
jf yarn install
```

## Building Go Packages

### General

JFrog CLI provides full support for building Go packages using the Go client. This allows resolving Go dependencies from and publish your Go packages to Artifactory, while collecting build-info and storing it in Artifactory.

### Requirements

JFrog CLI client version 1.20.0 and above.

Artifactory version 6.1.0 and above.

Go client version 1.11.0 and above.

### Example project

To help you get started, you can use [this sample project on GitHub](https://github.com/jfrog/project-examples/tree/master/golang-example).

### Setting Go repositories

Before you can use JFrog CLI to build your Go projects with Artifactory, you first need to set the resolutions and deployment repositories for the project.

Here's how you set the repositories.

1. 'cd' into to the root of the Go project.
2. Run the **jf go-config** command.

#### Commands Params

|                       |                                                                                                                                                                                 |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name          | go-config                                                                                                                                                                       |
| Abbreviation          |                                                                                                                                                                                 |
| **Command options:**  |                                                                                                                                                                                 |
| `--global`            | <p>[Default false]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p> |
| `--server-id-resolve` | <p>[Optional]<br>Artifactory server ID for resolution. The server should configured using the 'jf c add' command.</p>                                                           |
| `--server-id-deploy`  | <p>[Optional]<br>Artifactory server ID for deployment. The server should be configured using the 'jf c add' command.</p>                                                        |
| `--repo-resolve`      | <p>[Optional]<br>Repository for dependencies resolution.</p>                                                                                                                    |
| `--repo-deploy`       | <p>[Optional]<br>Repository for artifacts deployment.</p>                                                                                                                       |

#### Examples

**Example 1**

Set repositories for this go project.

```
jf go-config
```

**Example 2**

Set repositories for all go projects on this machine.

```
jf go-config --global
```

### Running Go commands

The **go** command triggers the go client.

> **Note**: Before running the **go** command on a project for the first time, the project should be configured using the **jf go-config** command.

#### Commands Params

The following table lists the command arguments and flags:

|                        |                                                                                                                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | go                                                                                                                                                                                                                                       |
| Abbreviation           | go                                                                                                                                                                                                                                       |
| **Command options:**   |                                                                                                                                                                                                                                          |
| `--build-name`         | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`       | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`            | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--no-fallback`        | <p>[Default false]<br>Set to avoid downloading packages from the VCS, if they are missing in Artifactory.</p>                                                                                                                            |
| `--module`             | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| **Command arguments:** |                                                                                                                                                                                                                                          |
| Go command             | The command accepts the same arguments and options as the go client.                                                                                                                                                                     |

#### Examples

**Example 1**

The following example runs Go build command. The dependencies resolved from Artifactory via the go-virtual repository.

> **Note**: Before using this example, please make sure to set repositories for the Go project using the go-config command.

```
jf go build
```

**Example 2**

The following example runs Go build command, while recording the build-info locally under build name **my-build** and build number **1**. The build-info can later be published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info)command.

> **Note**: Before using this example, please make sure to set repositories for the Go project using the go-config command.

```
jf rt go build --build-name=my-build --build-number=1
```

### Publishing Go Packages to Artifactory

The **jf go-publish** command packs and deploys the Go package to the designated Go repository in Artifactory.

> **Note**: Before running the **jf go-publish** command on a project for the first time, the project should be configured using the **jf go-config** command.

#### Commands Params

The following table lists the command arguments and flags:

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | go-publish                                                                                                                                                                                                                               |
| Abbreviation         | gp                                                                                                                                                                                                                                       |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| `--detailed-summary` | <p>[Default: false]<br>Set true to include a list of the affected files as part of the command output summary.</p>                                                                                                                       |
| Command argument     |                                                                                                                                                                                                                                          |
| Version              | The version of the Go project that is being published                                                                                                                                                                                    |

#### Examples

**Example 1**

To pack and publish the Go package, run the following command. Before running this command on a project for the first time, the project should be configured using the **jf go-config** command.

```
jf gp v1.2.3 
```

**Example 2**

To pack and publish the Go package and also record the build-info as part of build **my-build-name/1** , run the following command. The build-info can later be published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command. Before running this command on a project for the first time, the project should be configured using the **jf go-config** command.

```
jf gp v1.2.3 --build-name my-build-name --build-number 1
```

## Building Python Packages

### Pip, Pipenv and Twine

JFrog CLI provides full support for building Python packages using the **pip** and **pipenv** package managers, and deploying distributions using **twine**. This allows resolving python dependencies from Artifactory, using for **pip** and **pipenv**, while recording the downloaded packages. After installing and packaging the project, the distributions and wheels can be deployed to Artifactory using **twine**, while recording the uploaded packages. The downloaded packages are stored as dependencies in the build-info stored in Artifactory, while the uploaded ones are stored as artifacts.

#### Example projects

To help you get started, you can use [the sample projects on GitHub](https://github.com/jfrog/project-examples/tree/master/python-example).

#### Setting Python repository

Before you can use JFrog CLI to build your Python projects with Artifactory, you first need to set the repository for the project.

Here's how you set the repositories.

1. 'cd' into the root of the Python project.
2. Run the **jf pip-config** or **jf pipenv-config** commands, depending on whether you're using the **pip** or **pipenv** clients.

**Commands Params**

|                       |                                                                                                                                                                                 |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name          | pip-config / pipenv-config                                                                                                                                                      |
| Abbreviation          | pipc / pipec                                                                                                                                                                    |
| **Command options:**  |                                                                                                                                                                                 |
| `--global`            | <p>[Default false]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p> |
| `--server-id-resolve` | <p>[Optional]<br>Artifactory server ID for resolution. The server should configured using the 'jf c add' command.</p>                                                           |
| `--repo-resolve`      | <p>[Optional]<br>Repository for dependencies resolution.</p>                                                                                                                    |
| `--server-id-deploy`  | <p>[Optional]<br>Artifactory server ID for deployment. The server should configured using the 'jf c add' command.</p>                                                           |
| `--repo-deploy`       | <p>[Optional]<br>Repository for artifacts deployment.</p>                                                                                                                       |

**Examples**

**Example 1**

Set repositories for this Python project when using the pip client (for pipenv: `jf pipec`).

```
jf pipc
```

**Example 2**

Set repositories for all Python projects using the pip client on this machine (for pipenv: `jf pipec --global`).

```
jf pipc --global
```

#### Installing Python packages

The **jf pip install** and **jf pipenv install** commands use the **pip** and **pipenv** clients respectively, to install the project dependencies from Artifactory. The **jf pip install** and **jf pipenv install** commands can also record these packages as build dependencies as part of the build-info published to Artifactory.

> **Note**: Before running the **pip install** and **pipenv install** commands on a project for the first time, the project should be configured using the **jf pip-config** or **jf pipenv-config** commands respectively.

**Recording all dependencies**

JFrog CLI records the installed packages as build-info dependencies. The recorded dependencies are packages installed during the **jf pip install** and **jf pipenv install** command execution. When running the command inside a Python environment, which already has some of the packages installed, the installed packages will not be included as part of the build-info, because they were not originally installed by JFrog CLI. A warning message will be added to the log in this case.

**How to include all packages in the build-info?**

The details of all the installed packages are always cached by the **jf pip install** and **jf pipenv install** command in the **.jfrog/projects/deps.cache.json** file, located under the root of the project. JFrog CLI uses this cache for including previously installed packages in the build-info.\
If the Python environment had some packages installed prior to the first execution of the `install` command, those previously installed packages will be missing from the cache and therefore will not be included in the build-info.

Running the `install` command with both the **no-cache-dir** and **force-reinstall** pip options, should re-download and install these packages, and they will therefore be included in the build-info and added to the cache. It is also recommended to run the command from inside a [virtual environment](https://packaging.python.org/guides/installing-using-pip-and-virtual-environments/).

**Commands Params**

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | pip / pipenv                                                                                                                                                                                                                             |
| Abbreviation         |                                                                                                                                                                                                                                          |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| Command argument     | The command accepts the same arguments and options as the pip / pipenv clients.                                                                                                                                                          |

**Examples**

**Example 1**

The following command triggers pip install, while recording the build dependencies as part of build name **my-build** and build number **1** .

```
jf pip install . --build-name my-build --build-number 1
```

**Example 2**

The following command triggers pipenv install, while recording the build dependencies as part of build name **my-build** and build number **1** .

```
jf pipenv install . --build-name my-build --build-number 1
```

#### Publishing Python packages using Twine

The **jf twine upload** command uses the **twine**, to publish the project distributions to Artifactory. The **jf twine upload** command can also record these packages as build artifacts as part of the build-info published to Artifactory.

> **Note**: Before running the **twine upload** command on a project for the first time, the project should be configured using the **jf pip-config** or **jf pipenv-config** commands, with deployer configuration.

**Commands Params**

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | twine                                                                                                                                                                                                                                    |
| Abbreviation         |                                                                                                                                                                                                                                          |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| Command argument     | The command accepts the arguments and options supported by twine client, except for repository configuration and authentication options.                                                                                                 |

**Examples**

**Example 1**

The following command triggers twine upload, while recording the build artifacts as part of build name **my-build** and build number **1** .

```
jf twine upload "dist/*" --build-name my-build --build-number 1
```

## Configure Poetry Resolution for Artifactory (JFrog CLI)

### Introduction

Use this procedure to configure which Artifactory PyPI repository Poetry resolves dependencies from when you use the JFrog-wrapped Poetry commands (`jf poetry` …). This setup is not used by native Poetry execution; for native workflows, configure credentials in Poetry via your repository’s Set Me Up instructions.

### CLI Structure

```
jf poetry-config --server-id-resolve <SERVER_ID> --repo-resolve <PYPI_VIRTUAL_OR_REMOTE> [--global]
# shorthand:
jf poc --server-id-resolve <SERVER_ID> --repo-resolve <PYPI_VIRTUAL_OR_REMOTE> [--global]
```

* `<SERVER_ID>`: The configured JFrog CLI server ID (create with jf c add).
* `<PYPI_VIRTUAL_OR_REMOTE>`: The Artifactory PyPI virtual/remote repository used for dependency resolution.
* `--global`: Apply on this machine for all projects (project config can override).\


**Example**

`# Project-scoped configuration`

`jf poc --server-id-resolve art --repo-resolve pypi-virtual`

\


`# Global configuration`

`jf poc --global --server-id-resolve art --repo-resolve pypi-virtual`

### Additional information

* Native Poetry execution ignores this configuration; use Set Me Up to configure Poetry credentials and repo there.\
  \


***

## Install Python Dependencies with Poetry (JFrog CLI)

### Introduction

Use this procedure to install project dependencies from Artifactory using Poetry. You can run:

* **Wrapped mode** (jf poetry install): uses the `jf poetry-config` resolution settings from the previous procedure.\

* **Native mode** (poetry install with JFrog CLI pass-through): does not use `jf poetry-config`; configure Poetry via Set Me Up.\
  \
  \


**Tip**: You can capture build-info during install by adding `--build-name` and `--build-number` flags to the wrapped command.\


**CLI command**&#x20;

```
jf poetry install [--build-name <BUILD_NAME>] [--build-number <BUILD_NUMBER>] [--project <PROJECT_KEY>] [--module <MODULE_ID>] [<NATIVE_POETRY_ARGS>...
```

* `<BUILD_NAME> / <BUILD_NUMBER>`: Identifiers used to collect and later publish build-info.
* `<PROJECT_KEY>`: JFrog Project key (optional).
* `<MODULE_ID>`: Optional module name override in build-info.
* `<NATIVE_POETRY_ARGS>`: Any arguments supported by Poetry; JFrog CLI passes them through in native execution mode.\
  \


**Full example**

`# Wrapped install using configured resolution; collect build-info`

`jf poetry install --build-name my-python-app --build-number 1.2.0`

\


`# Native Poetry (credentials/repo set via Set Me Up)`

`poetry install`

### Additional information

* Wrapped vs. native: users already on wrapped can continue, but we encourage moving to native over time.\

* After install, continue to Publish Python Package with Poetry (JFrog CLI) if you need to deploy artifacts.\
  \


***

## Publish Python Package with Poetry (JFrog CLI)

### Introduction

Use this procedure to build and deploy your Poetry package to an Artifactory PyPI local repository, while collecting build-info.



**CLI command**

```
jf poetry publish --repo <PYPI_LOCAL> [--build-name <BUILD_NAME>] [--build-number <BUILD_NUMBER>] [--project <PROJECT_KEY>] [--module <MODULE_ID>] [<NATIVE_POETRY_ARGS>...]
```



* `<PYPI_LOCAL>`: Target local PyPI repository in Artifactory for deployment.\

* `<BUILD_NAME> / <BUILD_NUMBER>`: Build-info identifiers used to associate published artifacts to the build.\

* `<PROJECT_KEY> / <MODULE_ID>`: Optional metadata for build-info organization.\

* `<NATIVE_POETRY_ARGS>`: Any Poetry flags passed through by the CLI.\
  \


**Full example**

```
# Build and publish with build-info

jf poetry publish --repo pypi-local \
  --build-name my-python-app --build-number 1.2.0
```

### Additional information

* If you prefer a fully native flow: `poetry build` then `poetry publish -r <REPO_ALIAS>` (credentials configured via Set Me Up), and publish build-info separately from your pipeline as needed.\


**Next step**: finalize the record in Artifactory:\
\
`jf build-publish --build-name my-python-app --build-number 1.2.0`

\


***

### Universal package note (applies to all Poetry procedures)

JFrog CLI is a pass-through to Poetry for execution and output. Output format/behavior belongs to Poetry; if Poetry changes its output in a newer version, JFrog CLI would simply honor that without any custom adjustment.



## Building NuGet Packages

JFrog CLI provides full support for restoring NuGet packages using the NuGet client or the .NET Core CLI. This allows you to resolve NuGet dependencies from and publish your NuGet packages to Artifactory, while collecting build-info and storing it in Artifactory.

NuGet dependencies resolution is supported by the `jf nuget` command, which uses the NuGet client or the `jf dotnet` command, which uses the .NET Core CLI.

To publish your NuGet packages to Artifactory, use the [jf rt upload](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#uploading-files) command.

### Setting NuGet repositories

Before using the**nuget** or **dotnet** commands, the project needs to be pre-configured with the Artifactory server and repository, to be used for building the project.

Before using the nuget or dotnet commands, the **nuget-config** or **dotnet-config** commands should be used respectively. These commands configure the project with the details of the Artifactory server and repository, to be used for the build. The **nuget-config** or **dotnet-config** commands should be executed while inside the root directory of the project. The configuration is stored by the command in the **.jfrog** directory at the root directory of the project. You then have the option of storing the .jfrog directory with the project sources, or creating this configuration after the sources are checked out.

The following table lists the commands' options:

|                        |                                                                                                                                                                            |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | nuget-config / dotnet-config                                                                                                                                               |
| Abbreviation           | nugetc / dotnetc                                                                                                                                                           |
| **Command options:**   |                                                                                                                                                                            |
| `--global`             | <p>[Optional]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p> |
| `--server-id-resolve`  | <p>[Optional]<br>Artifactory server ID for resolution. The server should configured using the 'jf c add' command.</p>                                                      |
| `--repo-resolve`       | <p>[Optional]<br>Repository for dependencies resolution.</p>                                                                                                               |
| --nuget-v2             | <p>[Default: false]<br>Set to true if you'd like to use the NuGet V2 protocol when restoring packages from Artifactory (instead of NuGet V3).</p>                          |
| **Command arguments:** | The command accepts no arguments                                                                                                                                           |

### Running Nuget and Dotnet commands

The **nuget** command runs the **NuGet client** and the **dotnet** command runs the \*\*.NET Core CLI.

> Before running the nuget command on a project for the first time, the project should be configured using the nuget-config command.

> Before running the dotnet command on a project for the first time, the project should be configured using the dotnet-config command.

#### Commands Params

The following table lists the commands arguments and options:

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | nuget / dotnet                                                                                                                                                                                                                           |
| Abbreviation         |                                                                                                                                                                                                                                          |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                  |
| `--module`           | <p>[Optional]<br>Optional module name for the build-info.</p>                                                                                                                                                                            |
| Command argument     | The command accepts the same arguments and options as the NuGet client / .NET Core CLI.                                                                                                                                                  |

#### Examples

**Example 1**

Run nuget restore for the solution at the current directory, while resolving the NuGet dependencies from the pre-configured Artifactory repository. Use the NuGet client for this command

```
jf nuget restore
```

**Example 2**

Run dotnet restore for the solution at the current directory, while resolving the NuGet dependencies from the pre-configured Artifactory repository. Use the .NET Core CLI for this command

```
jf dotnet restore
```

**Example 3**

Run dotnet restore for the solution at the current directory, while resolving the NuGet dependencies from the pre-configured Artifactory repository.

```
jf dotnet restore --build-name=my-build-name --build-number=1
```

In addition, record the build-info as part of build **my-build-name/1**. The build-info can later be published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#publishing-build-info) command.

## Packaging and Publishing Terraform Modules

JFrog CLI supports packaging Terraform modules and publishing them to a Terraform repository in Artifactory using the **jf terraform publish** command.

We recommend using [this example project on GitHub](https://github.com/jfrog/project-examples/tree/master/terraform-example) for an easy start up.

Before using the **jf terraform publish** command for the first time, you first need to configure the Terraform repository for your Terraform project. To do this, follow these steps:

1. 'cd' into the root directory for your Terraform project.
2. Run the interactive **jf terraform-config** command and set deployment repository name.

### terraform-config

The **jf terraform-config** command will store the repository name inside the **.jfrog** directory located in the current directory. You can also add the **--global** command option, if you prefer the repository configuration applies to all projects on the machine. In that case, the configuration will be saved in JFrog CLI's home directory.

#### Commands Params

The following table lists the command options:

|                        |                                                                                                                                                                            |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | terraform-config                                                                                                                                                           |
| Abbreviation           | tfc                                                                                                                                                                        |
| **Command options:**   |                                                                                                                                                                            |
| `--global`             | <p>[Optional]<br>Set to true, if you'd like the configuration to be global (for all projects on the machine). Specific projects can override the global configuration.</p> |
| `--server-id-deploy`   | <p>[Optional]<br>Artifactory server ID for deployment. The server should configured using the 'jf c add' command.</p>                                                      |
| `--repo-deploy`        | <p>[Optional]<br>Repository for artifacts deployment.</p>                                                                                                                  |
| **Command arguments:** | The command accepts no arguments                                                                                                                                           |

#### Examples

**Example 1**

Configuring the Terraform repository for a project, while inside the root directory of the project

```
jf tfc
```

**Example 2**

Configuring the Terraform repository for all projects on the machine

```
jf tfc --global
```

### terraform publish

The **terraform publish** command creates a terraform package for the module in the current directory, and publishes it to the configured Terraform repository in Artifactory.

#### Commands Params

The following table lists the commands arguments and options:

|                      |                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name         | terraform publish                                                                                                                                                                                                                        |
| Abbreviation         | tf p                                                                                                                                                                                                                                     |
| **Command options:** |                                                                                                                                                                                                                                          |
| `--namespace`        | <p>[Mandatory]<br>Terraform module namespace</p>                                                                                                                                                                                         |
| `--provider`         | <p>[Mandatory]<br>Terraform module provider</p>                                                                                                                                                                                          |
| `--tag`              | <p>[Mandatory]<br>Terraform module tag</p>                                                                                                                                                                                               |
| `--exclusions`       | <p>[Optional]<br>A list of semicolon-separated(;) exclude patterns wildcards. Paths inside the module matching one of the patterns are excluded from the deployed package.</p>                                                           |
| `--build-name`       | <p>[Optional]<br>Build name. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p>   |
| `--build-number`     | <p>[Optional]<br>Build number. For more details, please refer to <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory#build-integration">Build Integration</a>.</p> |
| `--project`          |                                                                                                                                                                                                                                          |
| Command argument     | The command accepts no arguments                                                                                                                                                                                                         |

#### Examples

**Example 1**

The command creates a package for the Terraform module in the current directory, and publishes it to the Terraform repository (configured by the **jf tfc command**) with the provides namespace, provider and tag.

```
jf tf p --namespace example --provider aws --tag v0.0.1
```

**Example 2**

The command creates a package for the Terraform module in the current directory, and publishes it to the Terraform repository (configured by the **jf tfc command**) with the provides namespace, provider and tag. The published package will not include the module paths which include either **test** or **ignore** .

```
jf tf p --namespace example --provider aws --tag v0.0.1 --exclusions "\*test\*;\*ignore\*"
```

**Example 3**

The command creates a package for the Terraform module in the current directory, and publishes it to the Terraform repository (configured by the **jf tfc** command) with the provides namespace, provider and tag. The published module will be recorded as an artifact of a build named **my-build** with build number **1**. The **jf rt bp** command publishes the build to Artifactory.

```
jf tf p --namespace example --provider aws --tag v0.0.1 --build-name my-build --build-number 1
jf rt bp my-build 1
```
