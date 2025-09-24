# Build Integration

### Overview

JFrog CLI integrates with any development ecosystem, allowing you to collect build-info and then publish it to Artifactory. By publishing build-info to Artifactory, JFrog CLI empowers Artifactory to provide visibility into deployed artifacts, used dependencies, and extensive information on the build environment to allow fully traceable builds. For more information about build-info and build integration with Artifactory, see the [JFrog Integrations Documentation](https://app.gitbook.com/u/wlkoInwXRRMoHDanNwbVXlSfBo52).

Many JFrog CLI commands accept the optional `--build-name` and `--build-number` options. When you add these options, JFrog CLI collects and records the build info locally for these commands. When running multiple commands with the same build name and build number, JFrog CLI aggregates the collected build info into one build. The recorded build-info can be published to Artifactory later using the `build-publish` command.

> **New:** You can now optionally collect **environment variables** and **Git metadata** _at publish time_ with `jf rt bp` using `--collect-env` and `--collect-git-info`, reducing the need for separate `bce`/`bag` steps (details below). Legacy commands remain supported.

***

### Collecting Build-Info

Build-info is collected by adding the `--build-name` and `--build-number` options to different CLI commands. The CLI commands can be run multiple times to cumulatively collect build-info for the specified build name and number until it is published to Artifactory. For example, running the `jf rt download` command several times with the same build name and number will accumulate each downloaded file in the corresponding build-info.

***

### Collecting Dependencies

Dependencies are collected by adding the `--build-name` and `--build-number` options to the `jf rt download` command.

**Example**\
The following command downloads the `cool-froggy.zip` file from the `my-local-repo` repository and also specifies this file as a dependency in build `my-build-name` with build number `18`:

```bash
jf rt dl my-local-repo/cool-froggy.zip --build-name=my-build-name --build-number=18
```

***

### Collecting Build Artifacts

Build artifacts are collected by adding the `--build-name` and `--build-number` options to the `jf rt upload` command.

**Example**\
The following command specifies that the file `froggy.tgz` uploaded to the `my-local-repo` repository is a build artifact of build `my-build-name` with build number `18`:

```bash
jf rt u froggy.tgz my-local-repo --build-name=my-build-name --build-number=18
```

***

### Collecting Environment Variables

#### Option A (recommended): Collect at publish time (new)

Use `jf rt bp` with:

* `--collect-env` – include environment variables in the build-info.
* Combine with `--env-include` / `--env-exclude` to control what’s sent.

**Example**

```bash
jf rt bp my-build-name 18 \
  --collect-env \
  --env-exclude "*password*;*secret*;*key*;*token*;*auth*"
```

#### Option B (legacy): `build-collect-env (bce)`

Use this command to collect environment variables and attach them to a build. Environment variables are collected using the `build-collect-env (bce)` command.

**Usage**

```bash
jf rt bce <build name> <build number>
```

#### Commands Params

| **Command name**       | **Abbreviation** |
| ---------------------- | ---------------- |
| `rt build-collect-env` | `rt bce`         |

**Command arguments:** (The command accepts two arguments.)

| Argument     | Description   |
| ------------ | ------------- |
| Build name   | Build name.   |
| Build number | Build number. |

**Command options:**

| Option      | Description                     |
| ----------- | ------------------------------- |
| `--project` | _(Optional)_ JFrog project key. |

#### Examples

**Example 1**\
Collect all currently known environment variables and attach them to the build-info for build `my-build-name` with build number `18`:

```bash
jf rt bce my-build-name 18
```

**Example 2**\
Collect environment variables for build name `frogger-build` and build number `17`:

```bash
jf rt bce frogger-build 17
```

> **Recommendation:** Prefer `jf rt bp ... --collect-env` to capture env at publish time and reduce pipeline steps. `bce` remains supported for backward compatibility.

***

### Collecting Information from Git

#### Option A (recommended): Collect at publish time (new)

Use `jf rt bp` with:

* `--collect-git-info` – read Git revision and remote URL.
* `--dot-git-path` – path to `.git` if not at project root.
* `--git-config-file-path` – path to a Git config file for advanced layouts.

**Example**

```bash
jf rt bp my-build-name 18 \
  --collect-git-info \
  --dot-git-path ./.git
```

#### Option B (legacy): `build-add-git (bag)`

The `build-add-git (bag)` command collects the Git revision and URL from the local `.git` directory and adds it to the build-info. The command also collects the list of tracked project issues (for example, issues stored in JIRA or other bug tracking systems) and adds them to the build-info. The issues are collected by reading the git commit messages from the local git log. Each commit message is matched against a pre-configured regular expression, which retrieves the issue ID and issue summary. The information required for collecting the issues is retrieved from a YAML configuration file provided to the command.

**Usage**

```bash
jf rt bag [command options] <build name> <build number> [Path To .git]
```

#### Commands Params

| **Command name**   | **Abbreviation** |
| ------------------ | ---------------- |
| `rt build-add-git` | `rt bag`         |

**Command arguments:** (The command accepts three arguments.)

| Argument     | Description                                                                                                                                                                             |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Build name   | Build name.                                                                                                                                                                             |
| Build number | Build number.                                                                                                                                                                           |
| `.git` path  | _(Optional)_ Path to a directory containing the `.git` directory. If not specified, the `.git` directory is assumed to be in the current directory or in one of the parent directories. |

**Command options:**

| Option        | Description                                                                                                                                                                                                                                                                                                                                                                       |
| ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--config`    | _(Optional)_ Path to a yaml configuration file, used for collecting tracked project issues and adding them to the build-info.                                                                                                                                                                                                                                                     |
| `--server-id` | _(Optional)_ Server ID configured using the `jf config` command. This is the server to which the build-info will be later published, using the `jf rt build-publish` command. This option, if provided, overrides the `serverID` value in this command's yaml configuration. If both values are not provided, the default server, configured by the `jf config` command, is used. |
| `--project`   | _(Optional)_ JFrog project key.                                                                                                                                                                                                                                                                                                                                                   |

#### Configuration file properties

| **Property name**   | **Description**                                                                                                                                                                                                                                                                                                                                                                           |
| ------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Version`           | The schema version is intended for internal use. Do not change!                                                                                                                                                                                                                                                                                                                           |
| `serverID`          | Artifactory server ID configured by the `jf config` command. The command uses this server for fetching details about previous published builds. The `--server-id` command option, if provided, overrides the `serverID` value. If both the `serverID` property and the `--server-id` command options are not provided, the default server, configured by the `jf config` command is used. |
| `trackerName`       | The name (type) of the issue tracking system. For example, JIRA. This property can take any value.                                                                                                                                                                                                                                                                                        |
| `regexp`            | A regular expression used for matching the git commit messages. The expression should include two capturing groups - for the issue key (ID) and the issue summary. Example message: `HAP-1007 - This is a sample issue`                                                                                                                                                                   |
| `keyGroupIndex`     | The capturing group index in the regular expression used for retrieving the issue key.                                                                                                                                                                                                                                                                                                    |
| `summaryGroupIndex` | The capturing group index in the regular expression for retrieving the issue summary.                                                                                                                                                                                                                                                                                                     |
| `trackerUrl`        | The issue tracking URL. Used for constructing links to issues in the Artifactory build UI.                                                                                                                                                                                                                                                                                                |
| `aggregate`         | Set to `true` to include issues from previous builds.                                                                                                                                                                                                                                                                                                                                     |
| `aggregationStatus` | If `aggregate` is `true`, indicates how far in time to aggregate (e.g., until a build with a `RELEASE` status is found).                                                                                                                                                                                                                                                                  |

**Example**

```bash
jf rt bag frogger-build 17 checkout-dir
```

**Configuration file structure (example)**

```yaml
version: 1
issues: 
  # The serverID yaml property is optional. The --server-id command option, if provided, overrides the serverID value.
  # If both the serverID property and the --server-id command options are not provided,
  # the default server, configured by the "jfrog config add" command is used.
  serverID: my-artifactory-server

  trackerName: JIRA
  regexp: (.+-[0-9]+)\s-\s(.+)
  keyGroupIndex: 1
  summaryGroupIndex: 2
  trackerUrl: https://my-jira.com/issues
  aggregate: true
  aggregationStatus: RELEASED
```

> **Recommendation:** Prefer `jf rt bp ... --collect-git-info [--dot-git-path ...]` to gather Git metadata at publish time. Use `bag` only when you need commit-message issue extraction via YAML.

***

### Adding Files as Build Dependencies

The download command and other commands that download dependencies from Artifactory accept the `--build-name` and `--build-number` options to record the downloaded files as build dependencies. In cases where it is necessary to add a file downloaded by another tool as a dependency, use the `build-add-dependencies (bad)` command.

By default, the command collects files from the local file system. To collect files from Artifactory instead, add the `--from-rt` option.

**Usage**

```bash
jf rt bad [command options] <build name> <build number> <pattern>
jf rt bad --spec=<File Spec path> [command options] <build name> <build number>
```

#### Commands Params

| **Command name**            | **Abbreviation** |
| --------------------------- | ---------------- |
| `rt build-add-dependencies` | `rt bad`         |

**Command arguments:** (The command takes three arguments.)

| Argument     | Description                                                                                                                                                                                                                                                                                                                                                           |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Build name   | The build name to add the dependencies to                                                                                                                                                                                                                                                                                                                             |
| Build number | The build number to add the dependencies to                                                                                                                                                                                                                                                                                                                           |
| Pattern      | Specifies the local file system path to dependencies which should be added to the build info. You can specify multiple dependencies by using wildcards or a regular expression as designated by the `--regexp` command option. If you have specified that you are using regular expressions, then the first one used in the argument must be enclosed in parenthesis. |

**Command options:**

> When using the `*` or `;` characters in the command options or arguments, wrap the whole string in quotes (`"`) so they are not interpreted as literals.

| Option         | Description                                                                                                                                                                             |
| -------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--from-rt`    | _(Default: false)_ Set to true to search the files in Artifactory, rather than on the local file system. The `--regexp` option is not supported when `--from-rt` is set to true.        |
| `--server-id`  | _(Optional)_ Server ID configured using the `jf config` command.                                                                                                                        |
| `--spec`       | _(Optional)_ Path to a File Spec.                                                                                                                                                       |
| `--spec-vars`  | _(Optional)_ Semicolon-separated variables in the form of `"key1=value1;key2=value2;..."` to be replaced in the File Spec. In the File Spec, the variables should be used as `${key1}`. |
| `--recursive`  | _(Default: true)_ When false, artifacts inside sub-folders in Artifactory will not be affected.                                                                                         |
| `--regexp`     | _(Optional: false; Default: false)_ Use a regular expression instead of wildcards to collect files (not supported with `--from-rt`).                                                    |
| `--dry-run`    | _(Default: false)_ Only show a summary of dependencies that will be added.                                                                                                              |
| `--module`     | _(Optional)_ Optional module name in the build-info for adding the dependency.                                                                                                          |
| `--exclusions` | A list of semicolon-separated exclude patterns. Allows wildcards or a regular expression according to the value of the `--regexp` option.                                               |

#### Examples

**Example 1**\
Add all files located under the `path/to/build/dependencies/dir` directory as dependencies of a build. The build name is `my-build-name` and the build number is `7`. The build-info is only updated locally. To publish the build-info to Artifactory use the `jf rt build-publish` command.

```bash
jf rt bad my-build-name 7 "path/to/build/dependencies/dir/"
```

**Example 2**\
Add all files located in the `m-local-repo` Artifactory repository, under the `dependencies` folder, as dependencies of a build.

```bash
jf rt bad my-build-name 7 "my-local-repo/dependencies/" --from-rt
```

**Example 3**\
Add all files located under the directory to module `m1`.

```bash
jf rt bad my-build-name 7 "path/to/build/dependencies/dir/" --module m1
```

***

### Publishing Build-Info (updated)

Use the `build-publish (bp)` command to publish the accumulated build-info to Artifactory.

**Usage**

```bash
jf rt bp [command options] <build name> <build number>
```

#### Commands Params

| **Command name**   | **Abbreviation** |
| ------------------ | ---------------- |
| `rt build-publish` | `rt bp`          |

**Command arguments:** (The command accepts two arguments.)

| Argument     | Description                   |
| ------------ | ----------------------------- |
| Build name   | Build name to be published.   |
| Build number | Build number to be published. |

**Command options (updated list):**

| Option                   | Description                                                                                                                                                                          |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `--server-id`            | _(Optional)_ Server ID configured using the `jf config` command. If not specified, the default configured Artifactory server is used.                                                |
| `--project`              | _(Optional)_ JFrog project key.                                                                                                                                                      |
| `--build-url`            | _(Optional)_ CI server build URL to include in the build-info.                                                                                                                       |
| `--collect-env`          | _(Default: false)_ **New.** Collect environment variables at publish time.                                                                                                           |
| `--env-include`          | _(Default: `*`)_ Semicolon-separated patterns in the form `"value1;value2;..."`. Only environment variables that match those patterns are included.                                  |
| `--env-exclude`          | _(Default: `*passwd*;*password*;*secret*;*key*;*token*;*auth*`)_ Case-insensitive semicolon-separated patterns to exclude sensitive variables. _(Updated default set for security.)_ |
| `--collect-git-info`     | _(Default: false)_ **New.** Collect Git revision and remote URL at publish time.                                                                                                     |
| `--dot-git-path`         | _(Optional)_ **New.** Path to a `.git` directory (if not at project root); CLI otherwise searches upward from the current directory.                                                 |
| `--git-config-file-path` | _(Optional)_ **New.** Path to a Git config file (for advanced repository layouts).                                                                                                   |
| `--dry-run`              | _(Default: false)_ Set to true to disable communication with Artifactory.                                                                                                            |
| `--insecure-tls`         | _(Default: false)_ Skip TLS certificates verification.                                                                                                                               |
| `--overwrite`            | _(Default: false)_ Overwrite all existing occurrences of build infos with the provided name and number. Build artifacts will not be deleted.                                         |

#### Examples

**Publish only**

```bash
jf rt bp my-build-name 18
```

**Publish + collect env + git (one step)**

```bash
jf rt bp my-build-name 18 \
  --collect-env \
  --env-exclude "*password*;*secret*;*key*;*token*;*auth*" \
  --collect-git-info \
  --dot-git-path ./.git
```

> **Security note:** Always use `--env-exclude` to prevent leaking secrets. The default covers common names; extend it for your organization’s variables.

***

### Aggregating Published Builds

In complex builds where steps run across multiple machines, you can publish separate build-info instances and then aggregate them into a "master" build-info record. The `build-append (ba)` command adds a reference from a new build-info to a previously published one.

```bash
jf rt upload "a/*.zip" generic-local --build-name my-build --build-number 1 --module m1
jf rt upload "b/*.zip" generic-local --build-name my-build --build-number 1 --module m2
jf rt build-publish my-build 1
jf rt download "*" --build my-build/1
```

If your build is composed of multiple build steps across different machines/time periods, create a separate build-info for each section and publish it independently. Once all are published, create a new build-info that **references** all previously published build-info instances using `build-append`. The reference is represented by a new module in the new build-info, with ID format: `**<referenced build name>/<referenced build number>`.

**Usage**

```bash
jf rt ba <build name> <build number> <build name to append> <build number to append>
```

#### Commands Params

| **Command name**  | **Abbreviation** |
| ----------------- | ---------------- |
| `rt build-append` | `rt ba`          |

**Command arguments:** (The command accepts four arguments.)

| Argument               | Description                                                |
| ---------------------- | ---------------------------------------------------------- |
| Build name             | The current (not yet published) build name.                |
| Build number           | The current (not yet published) build number.              |
| build name to append   | The published build name to append to the current build.   |
| build number to append | The published build number to append to the current build. |

**Command options:**\
This command has no options.

**Requirements**\
Artifactory version 7.25.4 and above.

**Example**

```bash
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

# Download the artifacts of aggregating-build/10 (includes a/1 and b/1)
jf rt download --build aggregating-build/10
```

***

### Promoting a Build

This command is used to promote build in Artifactory. Use the `build-promote (bpr)` command to promote a build in Artifactory, which typically involves moving or copying build artifacts to a target repository.

**Usage**

```bash
jf rt bpr [command options] <build name> <build number> <target repository>
```

#### Commands Params

| **Command name**   | **Abbreviation** |
| ------------------ | ---------------- |
| `rt build-promote` | `rt bpr`         |

**Command arguments:** (The command accepts three arguments.)

| Argument          | Description                        |
| ----------------- | ---------------------------------- |
| Build name        | Build name to be promoted.         |
| Build number      | Build number to be promoted.       |
| Target repository | Build promotion target repository. |

**Command options:**

| Option                   | Description                                                                                                                           |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| `--server-id`            | _(Optional)_ Server ID configured using the `jf config` command. If not specified, the default configured Artifactory server is used. |
| `--project`              | _(Optional)_ JFrog project key.                                                                                                       |
| `--status`               | _(Optional)_ Build promotion status.                                                                                                  |
| `--comment`              | _(Optional)_ Build promotion comment.                                                                                                 |
| `--source-repo`          | _(Optional)_ Build promotion source repository.                                                                                       |
| `--include-dependencies` | _(Default: false)_ If set to true, the build dependencies are also promoted.                                                          |
| `--copy`                 | _(Default: false)_ If set true, the build artifacts and dependencies are copied to the target repository, otherwise they are moved.   |
| `--props`                | _(Optional)_ Semicolon-separated properties in the form `"key1=value1;key2=value2,..."` to attach to the build artifacts.             |
| `--dry-run`              | _(Default: false)_ If true, promotion is only simulated. The build is not promoted.                                                   |
| `--insecure-tls`         | _(Default: false)_ Skip TLS certificates verification.                                                                                |

**Example**\
Move the artifacts associated with the published build-info `my-build-name/18` to `target-repository`:

```bash
jf rt bpr my-build-name 18 target-repository
```

***

### Cleaning up the Build

To reset the locally accumulated build-info and delete any information collected so far, use the `build-clean (bc)` command.

**Usage**

```bash
jf rt bc <build name> <build number>
```

#### Commands Params

| **Command name** | **Abbreviation** |
| ---------------- | ---------------- |
| `rt build-clean` | `rt bc`          |

**Command arguments:** (The command accepts two arguments.)

| Argument     | Description   |
| ------------ | ------------- |
| Build name   | Build name.   |
| Build number | Build number. |

**Command options:**\
This command has no options.

**Example**

```bash
jf rt bc my-build-name 18
```

***

### Discarding Old Builds from Artifactory

Use the `build-discard (bdi)` command to remove old builds previously published to Artifactory.

**Usage**

```bash
jf rt bdi [command options] <build name>
```

#### Commands Params

| **Command name**   | **Abbreviation** |
| ------------------ | ---------------- |
| `rt build-discard` | `rt bdi`         |

**Command arguments:** (The command accepts one argument.)

| Argument   | Description |
| ---------- | ----------- |
| Build name | Build name. |

**Command options:**

| Option               | Description                                                                                                                           |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `--server-id`        | _(Optional)_ Server ID configured using the `jf config` command. If not specified, the default configured Artifactory server is used. |
| `--max-days`         | _(Optional)_ The maximum number of days to keep builds in Artifactory.                                                                |
| `--max-builds`       | _(Optional)_ The maximum number of builds to store in Artifactory.                                                                    |
| `--exclude-builds`   | _(Optional)_ Comma-separated build numbers in the form `"build1,build2,..."` that should not be removed.                              |
| `--delete-artifacts` | _(Default: false)_ If set to true, automatically removes build artifacts stored in Artifactory.                                       |
| `--async`            | _(Default: false)_ If set to true, build discard will run asynchronously and will not wait for response.                              |

#### Examples

**Example 1**\
Discard the oldest build numbers of build `my-build-name`, leaving only the 10 most recent:

```bash
jf rt bdi my-build-name --max-builds=10
```

**Example 2**\
Discard the oldest build numbers, leaving only builds published during the last 7 days:

```bash
jf rt bdi my-build-name --max-days=7
```

**Example 3**\
Discard by age but keep specific builds:

```bash
jf rt bdi my-build-name --max-days 7 --exclude-builds "b20,b21"
```

***

#### Troubleshooting

* **Build succeeded but nothing was deployed**
  * _Gradle (Artifactory plugin):_ specify publications (`artifactory.publish.publications=release` or via CLI/DSL).
  * _Maven:_ run a deploy phase/goal (`deploy`) or ensure your setup deploys during `install`.
  * _npm/yarn:_ `install` only resolves; publish with `jf npm publish` (or `jf rt u`), then `jf rt bp`.
  * _Docker/Podman:_ ensure `jf docker login` (or remove `--skip-login`), and that the repo permissions are correct.
* **Secrets in env:** use `--env-exclude` with `--collect-env`; extend default patterns to match your org’s variables.



***
