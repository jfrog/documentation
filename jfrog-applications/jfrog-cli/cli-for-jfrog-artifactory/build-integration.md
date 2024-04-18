# Build Integration

## Overview

JFrog CLI integrates with any development ecosystem allowing you to collect build-info and then publish it to Artifactory. By publishing build-info to Artifactory, JFrog CLI empowers Artifactory to provide visibility into artifacts deployed, dependencies used and extensive information on the build environment to allow fully traceable builds. Read more about build-info and build integration with Artifactory [here](https://jfrog.com/help/r/jfrog-integrations-documentation/Build-Integration).

Many of JFrog CLI's commands accept two optional command options: **--build-name** and **--build-number**. When these options are added, JFrog CLI collects and records the build info locally for these commands.\
When running multiple commands using the same build and build number, JFrog CLI aggregates the collected build info into one build.\
The recorded build-info can be later published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#publishing-build-info) command.

## Collecting Build-Info

Build-info is collected by adding the `--build-name` and `--build-number` options to different CLI commands. The CLI commands can be run several times and cumulatively collect build-info for the specified build name and number until it is published to Artifactory. For example, running the `jf rt download` command several times with the same build name and number will accumulate each downloaded file in the corresponding build-info.

### Collecting Dependencies

Dependencies are collected by adding the `--build-name` and `--build-number` options to the `jf rt download` command.

For example, the following command downloads the `cool-froggy.zip` file found in repository `my-local-repo`, but it also specifies this file as a dependency in build `my-build-name` with build number 18:

```
jf rt dl my-local-repo/cool-froggy.zip --build-name=my-build-name --build-number=18
```

### Collecting Build Artifacts

Build artifacts are collected by adding the `--build-name` and `--build-number` options to the `jf rt upload` command.

For example, the following command specifies that file `froggy.tgz` uploaded to repository `my-local-repo` is a build artifact of build `my-build-name` with build number 18:

```
jf rt u froggy.tgz my-local-repo --build-name=my-build-name --build-number=18
```

### Collecting Environment Variables

This command is used to collect environment variables and attach them to a build.

Environment variables are collected using the `build-collect-env` (`bce`) command.

#### Usage

```jf rt bce <build name> <build number>```

#### Commands Params

The following table lists the command arguments and flags:

|                   |                                         |
|-------------------|-----------------------------------------|
| Command name      | rt build-collect-env                    |
| Abbreviation      | rt bce                                  |
| Command arguments | The command accepts two arguments.      |
| Build name        | Build name.                             |
| Build number      | Build number.                           |
| Command options   |                                         |
| --project         | <p>[Optional]<br>JFrog project key.</p> |

#### Examples
##### Example 1

The following command collects all currently known environment variables, and attaches them to the build-info for build `my-build-name` with build number 18:

```
jf rt bce my-build-name 18
```

##### Example 2

Collect environment variables for build name: frogger-build and build number: 17

```
jf rt bce frogger-build 17
```

### Collecting Information from Git

The `build-add-git` (bag) command collects the Git revision and URL from the local .git directory and adds it to the build-info. It can also collect the list of tracked project issues (for example, issues stored in JIRA or other bug tracking systems) and add them to the build-info. The issues are collected by reading the git commit messages from the local git log. Each commit message is matched against a pre-configured regular expression, which retrieves the issue ID and issue summary. The information required for collecting the issues is retrieved from a yaml configuration file provided to the command.

#### Usage

```jf rt bag [command options] <build name> <build number> [Path To .git]```

#### Commands Params

The following table lists the command arguments and flags:

|                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name      | rt build-add-git                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| Abbreviation      | rt bag                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| Command arguments | The command accepts three arguments.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| Build name        | Build name.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| Build number      | Build number.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| .git path         | Optional - Path to a directory containing the .git directory. If not specific, the .git directory is assumed to be in the current directory or in one of the parent directories.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| Command options   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| --config          | <p>[Optional]<br>Path to a yaml configuration file, used for collecting tracked project issues and adding them to the build-info.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| --server-id       | <p>[Optional]<br>Server ID configured using the <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration">'jf config' command</a>. This is the server to which the build-info will be later published, using the <code>jf rt build-publish</code> command. This option, if provided, overrides the serverID value in this command's yaml configuration. If both values are not provided, the default server, configured by the <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration">'jf config' command</a>, is used.</p> |
| --project         | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |

#### Configuration file properties

|                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Property name     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| Version           | The schema version is intended for internal use. Do not change!                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| serverID          | <p>Artifactory server ID configured by the <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration">'jf config' command</a>. The command uses this server for fetching details about previous published builds. The <strong>--server-id</strong> command option, if provided, overrides the <strong>serverID</strong> value.<br>If both the <strong>serverID</strong> property and the <strong>--server-id</strong> command options are not provided, the default server, configured by the <a href="https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration">'jf config' command</a> is used.</p> |
| trackerName       | The name (type) of the issue tracking system. For example, JIRA. This property can take any value.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| regexp            | <p>A regular expression used for matching the git commit messages. The expression should include two capturing groups - for the issue key (ID) and the issue summary. In the example above, the regular expression matches the commit messages as displayed in the following example:<br><br>HAP-1007 - This is a sample issue</p>                                                                                                                                                                                                                                                                                                                                                                                              |
| keyGroupIndex     | <p>The capturing group index in the regular expression used for retrieving the issue key. In the example above, setting the index to "1" retrieves <strong>HAP-1007</strong> from this commit message:<br><br>HAP-1007 - This is a sample issue</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| summaryGroupIndex | <p>The capturing group index in the regular expression for retrieving the issue summary. In the example above, setting the index to "2" retrieves the sample issue from this commit message:<br><br>HAP-1007 - This is a sample issue</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| trackerUrl        | The issue tracking URL. This value is used for constructing a direct link to the issues in the Artifactory build UI.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| aggregate         | Set to true, if you wish all builds to include issues from previous builds.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| aggregationStatus | If aggregate is set to true, this property indicates how far in time should the issues be aggregated. In the above example, issues will be aggregated from previous builds, until a build with a RELEASE status is found. Build statuses are set when a build is promoted using the **jf rt build-promote** command.                                                                                                                                                                                                                                                                                                                                                                                                            |

#### Example

```
jf rt bag frogger-build 17 checkout-dir
```

This is the configuration file structure.

```yaml
version: 1
issues: 
  # The serverID yaml property is optional. The --server-id command option, if provided, overrides the serverID value.
  # If both the serverID property and the --server-id command options are not provided,
  # the default server, configured by the "jfrog config add" command is used.
  serverID: my-artifactory-server

  trackerName: JIRA
  regexp: (.+-\[0-9\]+)\\s-\\s(.+)
  keyGroupIndex: 1
  summaryGroupIndex: 2
  trackerUrl: https://my-jira.com/issues
  aggregate: true
  aggregationStatus: RELEASED
```

### Adding Files as Build Dependencies

The download command, as well as other commands which download dependencies from Artifactory accept the **--build-name** and **--build-number** command options. Adding these options records the downloaded files as build dependencies. In some cases however, it is necessary to add a file, which has been downloaded by another tool, to a build. Use the **build-add-dependencies** command to this.

By default, the command collects the files from the local file system. If you'd like the files to be collected from Artifactory however, add the **--from-rt** option to the command.

#### Usage

```jf rt bad [command options] <build name> <build number> <pattern>```
```jf rt bad --spec=<File Spec path> [command options] <build name> <build number>```

#### Commands Params

|                   |                                                                                                                                                                                                                                                                                                                                                                     |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name      | rt build-add-dependencies                                                                                                                                                                                                                                                                                                                                           |
| Abbreviation      | rt bad                                                                                                                                                                                                                                                                                                                                                              |
| Command arguments | The command takes three arguments.                                                                                                                                                                                                                                                                                                                                  |
| Build name        | The build name to add the dependencies to                                                                                                                                                                                                                                                                                                                           |
| Build number      | The build number to add the dependencies to                                                                                                                                                                                                                                                                                                                         |
| Pattern           | Specifies the local file system path to dependencies which should be added to the build info. You can specify multiple dependencies by using wildcards or a regular expression as designated by the --regexp command option. If you have specified that you are using regular expressions, then the first one used in the argument must be enclosed in parenthesis. |
| Command options   | <p>When using the * or ; characters in the command options or arguments, make sure to wrap the whole options or arguments string in quotes (") to make sure the * or ; characters are not interpreted as literals.</p>                                                                                                                                              |
| --from-rt         | <p>[Default: false]<br>Set to true to search the files in Artifactory, rather than on the local file system. The --regexp option is not supported when --from-rt is set to true.</p>                                                                                                                                                                                |
| --server-id       | <p>[Optional]<br>Server ID configured using the config command.</p>                                                                                                                                                                                                                                                                                                 |
| --spec            | <p>[Optional]<br>Path to a File Spec.</p>                                                                                                                                                                                                                                                                                                                           |
| --spec-vars       | <p>[Optional]<br>List of variables in the form of "key1=value1;key2=value2;..." to be replaced in the File Spec. In the File Spec, the variables should be used as follows: ${key1}.</p>                                                                                                                                                                            |
| --recursive       | <p>[Default: true]<br>When false, artifacts inside sub-folders in Artifactory will not be affected.</p>                                                                                                                                                                                                                                                             |
| --regexp          | <p>[Optional: false]<br>[Default: false] Set to true to use a regular expression instead of wildcards expression to collect files to be added to the build info.This option is not supported when --from-rt is set to true.</p>                                                                                                                                     |
| --dry-run         | <p>[Default: false]<br>Set to true to only get a summery of the dependencies that will be added to the build info.</p>                                                                                                                                                                                                                                              |
| --module          | <p>[Optional]<br>Optional module name in the build-info for adding the dependency.</p>                                                                                                                                                                                                                                                                              |
| --exclusions      | A list of Semicolon-separated exclude patterns. Allows using wildcards or a regular expression according to the value of the 'regexp' option.                                                                                                                                                                                                                       |

#### Examples
##### Example 1

Add all files located under the **path/to/build/dependencies/dir** directory as dependencies of a build. The build name is **my-build-name** and the build number is **7**. The build-info is only updated locally. To publish the build-info to Artifactory use the **jf rt build-publish** command.

```
jf rt bad my-build-name 7 "path/to/build/dependencies/dir/"
```

##### Example 2

Add all files located in the **m-local-repo** Artifactory repository, under the **dependencies** folder, as dependencies of a build. The build name is **my-build-name** and the build number is **7**. The build-info is only updated locally. To publish the build-info to Artifactory use the **jf rt build-publish** command.

```
jf rt bad my-build-name 7 "my-local-repo/dependencies/" --from-rt
```

##### Example 3

Add all files located under the **path/to/build/dependencies/dir** directory as dependencies of a build. The build name is **my-build-name**, the build number is **7** and module is m1. The build-info is only updated locally. To publish the build-info to Artifactory use the **jf rt build-publish** command.

```
jf rt bad my-build-name 7 "path/to/build/dependencies/dir/" --module m1
```

## Publishing Build-Info

This command is used to publish build info to Artifactory. To publish the accumulated build-info for a build to Artifactory, use the **build-publish** command. 
For example, the following command publishes all the build-info collected for build **my-build-name** with build number 18:

### Usage

```jf rt bp [command options] <build name> <build number>```


### Commands Params

|                   |                                                                                                                                                                                  |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name      | rt build-publish                                                                                                                                                                 |
| Abbreviation      | rt bp                                                                                                                                                                            |
| Command arguments | The command accepts two arguments.                                                                                                                                               |
| Build name        | Build name to be published.                                                                                                                                                      |
| Build number      | Build number to be published.                                                                                                                                                    |
| Command options   |                                                                                                                                                                                  |
| --server-id       | <p>[Optional]<br>Server ID configured using the config command. If not specified, the default configured Artifactory server is used.</p>                                         |
| --project         | <p>[Optional]<br>JFrog project key.</p>                                                                                                                                          |
| --build-url       | <p>[Optional]<br>Can be used for setting the CI server build URL in the build-info.</p>                                                                                          |
| --env-include     | <p>[Default: *]<br>List of patterns in the form of "value1;value2;..." Only environment variables that match those patterns will be included in the build info.</p>              |
| --env-exclude     | <p>[Default: *password*;*secret*;*key*]<br>List of case insensitive patterns in the form of "value1;value2;..." environment variables match those patterns will be excluded.</p> |
| --dry-run         | <p>[Default: false]<br>Set to true to disable communication with Artifactory.</p>                                                                                                |
| --insecure-tls    | <p>[Default: false]<br>Set to true to skip TLS certificates verification.</p>                                                                                                    | |

### Example

Publishes to Artifactory all the build-info collected for build **my-build-name** with build number 18

```
jf rt bp my-build-name 18
```

## Aggregating Published Builds

The build-info, which is collected and published to Artifactory by the **jf rt build-publish** command, can include multiple modules. Each module in the build-info represents a package, which is the result of a single build step, or in other words, a JFrog CLI command execution. For example, the following command adds a module named **m1** to a build named **my-build** with **1** as the build number:

```
jf rt upload "a/*.zip" generic-local --build-name my-build --build-number 1 --module m1
```

The following command, adds a second module, named **m2** to the same build:

```
jf rt upload "b/*.zip" generic-local --build-name my-build --build-number 1 --module m2
```

You now publish the generated build-info to Artifactory using the following command:

```
jf rt build-publish my-build 1
```

Now that you have your build-info published to Artifactory, you can perform actions on the entire build. For example, you can download, copy, move or delete all or some of the artifacts of a build. Here's how you do this.

```
jf rt download "*" --build my-build/1
```

In some cases though, your build is composed of multiple build steps, which are running on multiple different machines or spread across different time periods. How do you aggregate those build steps, or in other words, aggregate those command executions, into one build-info?

The way to do this, is to create a separate build-info for every section of the build, and publish it independently to Artifactory. Once all the build-info instances are published, you can create a new build-info, which references all the previously published build-info instances. The new build-info can be viewed as a "master" build-info, which references other build-info instances.

So the next question is - how can this reference between the two build-instances be created?

The way to do this is by using the **build-append** command. Running this command on an unpublished build-info, adds a reference to a different build-info, which has already been published to Artifactory. This reference is represented by a new module in the new build-info. The ID of this module will have the following format: \*\*\<referenced build name>/\<referenced build number>.

Now, when downloading the artifacts of the "master" build, you'll actually be downloading the artifacts of all of its referenced builds. The examples below demonstrates this,

### Usage

```jf rt ba <build name> <build number> <build name to append> <build number to append>```

### Commands Params

|                        |                                                           |
|------------------------|-----------------------------------------------------------|
| Command name           | rt build-append                                           |
| Abbreviation           | rt ba                                                     |
| Command arguments      | The command accepts four arguments.                       |
| Build name             | The current (not yet published) build name.               |
| Build number           | The current (not yet published) build number,             |
| build name to append   | The published build name to append to the current build   |
| build number to append | The published build number to append to the current build |
| Command options        | This command has no options.                              |


**Requirements**

Artifactory version 7.25.4 and above.

### Example

This script illustrates the process of creating two build-info instances, publishing both to Artifactory, and subsequently generating a third build-info that consolidates the published instances before publishing it to Artifactory.

```yaml
# Create and publish build a/1
jf rt upload "a/*.zip" generic-local --build-name a --build-number 1
jf rt build-publish a 1

# Create and publish build b/1
jf rt upload "b/*.zip" generic-local --build-name b --build-number 1
jf rt build-publish b 1

# Append builds a/1 and b/1 to build aggregating-build/10
jf rt build-append aggregating-build 10 a 1
jf rt build-append aggregating-build 10 b 1

# Publish build aggregating-build/10
jf rt build-publish aggregating-build 10

# Download the artifacts of aggregating-build/10, which is the same as downloading the of a/1 and b/1
jf rt download --build aggregating-build/10
```

## Promoting a Build

This command is used to [promote build](https://jfrog.com/knowledge-base/how-does-build-promotion-work/) in Artifactory.

### Usage

```jf rt bpr [command options] <build name> <build number> <target repository>```

### Commands Params

|                        |                                                                                                                                             |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | rt build-promote                                                                                                                            |
| Abbreviation           | rt bpr                                                                                                                                      |
| Command arguments      | The command accepts three arguments.                                                                                                        |
| Build name             | Build name to be promoted.                                                                                                                  |
| Build number           | Build number to be promoted.                                                                                                                |
| Target repository      | Build promotion target repository.                                                                                                          |
| Command options        |                                                                                                                                             |
| --server-id            | <p>[Optional]<br>Server ID configured using the config command. If not specified, the default configured Artifactory server is used.</p>    |
| --project              | <p>[Optional]<br>JFrog project key.</p>                                                                                                     |
| --status               | <p>[Optional]<br>Build promotion status.</p>                                                                                                |
| --comment              | <p>[Optional]<br>Build promotion comment.</p>                                                                                               |
| --source-repo          | <p>[Optional]<br>Build promotion source repository.</p>                                                                                     |
| --include-dependencies | <p>[Default: false]<br>If set to true, the build dependencies are also promoted.</p>                                                        |
| --copy                 | <p>[Default: false]<br>If set true, the build artifacts and dependencies are copied to the target repository, otherwise they are moved.</p> |
| --props                | <p>[Optional]<br>List of properties in the form of "key1=value1;key2=value2,...". to attach to the build artifacts.</p>                     |
| --dry-run              | <p>[Default: false]<br>If true, promotion is only simulated. The build is not promoted.</p>                                                 |
| --insecure-tls         | <p>[Default: false]<br>Set to true to skip TLS certificates verification.</p>                                                               | |

### Example

This example involves moving the artifacts associated with the published build-info, identified by the build name 'my-build-name' and build number '18', from their existing Artifactory repository to a new Artifactory repository called 'target-repository'.

```
jf rt bpr my-build-name 18 target-repository
```

## Cleaning up the Build

Build-info is accumulated by the CLI according to the commands you apply until you publish the build-info to Artifactory. If, for any reason, you wish to "reset" the build-info and cleanup (i.e. delete) any information accumulated so far, you can use the `build-clean` (`bc`) command.

### Usage

```jf rt bc <build name> <build number>```

### Commands Params

The following table lists the command arguments and flags:

|                   |                                    |
|-------------------|------------------------------------|
| Command name      | rt build-clean                     |
| Abbreviation      | rt bc                              |
| Command arguments | The command accepts two arguments. |
| Build name        | Build name.                        |
| Build number      | Build number.                      |
| Command options   | The command has no options.        |


### Example

The following command cleans up any build-info collected for build `my-build-name` with build number 18:

```
jf rt bc my-build-name 18
```

## Discarding Old Builds from Artifactory

This command is used to discard builds previously published to Artifactory using the [build-publish](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#publishing-build-info) command.

### Usage

```jf rt bdi [command options] <build name>```

### Commands Params

The following table lists the command arguments and flags:

|                    |                                                                                                                                          |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| Command name       | rt build-discard                                                                                                                         |
| Abbreviation       | rt bdi                                                                                                                                   |
| Command arguments  | The command accepts one argument.                                                                                                        |
| Build name         | Build name.                                                                                                                              |
| Command options    |                                                                                                                                          |
| --server-id        | <p>[Optional]<br>Server ID configured using the config command. If not specified, the default configured Artifactory server is used.</p> |
| --max-days         | <p>[Optional]<br>The maximum number of days to keep builds in Artifactory.</p>                                                           |
| --max-builds       | <p>[Optional]<br>The maximum number of builds to store in Artifactory.</p>                                                               |
| --exclude-builds   | <p>[Optional]<br>List of build numbers in the form of "value1,value2,...", that should not be removed from Artifactory.</p>              |
| --delete-artifacts | <p>[Default: false]<br>If set to true, automatically removes build artifacts stored in Artifactory.</p>                                  |
| --async            | <p>[Default: false]<br>If set to true, build discard will run asynchronously and will not wait for response.</p>                         |
                                                                                                                     |
### Examples
#### Example 1

Discard the oldest build numbers of build **my-build-name** from Artifactory, leaving only the 10 most recent builds.

```
jf rt bdi my-build-name --max-builds= 10
```

#### Example 2

Discard the oldest build numbers of build **my-build-name** from Artifactory, leaving only builds published during the last 7 days.

```
jf rt bdi my-build-name --max-days=7
```

#### Example 3

Discard the oldest build numbers of build **my-build-name** from Artifactory, leaving only builds published during the last 7 days. **b20** and **b21** will not be discarded.

```
jf rt bdi my-build-name --max-days 7 --exclude-builds "b20,b21"
```
