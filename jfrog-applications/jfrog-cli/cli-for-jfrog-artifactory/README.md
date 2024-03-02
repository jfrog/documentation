# CLI for JFrog Artifactory

## Overview

This page describes how to use JFrog CLI with JFrog Artifactory.

Read more about JFrog CLI [here](../README.md).

## Environment Variables

The Artifactory upload command makes use of the following environment variable:

|                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Variable Name**                                  | **Description**                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| **JFROG\_CLI\_MIN\_CHECKSUM\_DEPLOY\_SIZE\_KB**    | <p>[Default: 10]<br><br>Minimum file size in KB for which JFrog CLI performs checksum deploy optimization.</p>                                                                                                                                                                                                                                                                                                                                                                 |
| **JFROG\_CLI\_RELEASES\_REPO**                     | <p>Configured Artifactory repository name to download the jar needed by the mvn/gradle command.<br>This environment variable's value format should be <code>&#x3C;server ID configured by the 'jf c add' command>/&#x3C;repo name></code>.<br>The repository should proxy https://releases.jfrog.io.<br>This environment variable is used by the 'jf mvn' and 'jf gradle' commands, and also by the 'jf audit' command, when used for maven or gradle projects.</p> |
| **JFROG\_CLI\_DEPENDENCIES\_DIR**                  | <p>[Default: $JFROG_CLI_HOME_DIR/dependencies]<br><br>Defines the directory to which JFrog CLI's internal dependencies are downloaded.</p>                                                                                                                                                                                                                                                                                                                                     |
| **JFROG\_CLI\_REPORT\_USAGE**                      | <p>[Default: true]<br><br>Set to false to block JFrog CLI from sending usage statistics to Artifactory.</p>                                                                                                                                                                                                                                                                                                                                                                    |
| **JFROG\_CLI\_SERVER\_ID**                         | Server ID configured using the config command, unless sent as a command argument or option.                                                                                                                                                                                                                                                                                                                                                                                    |
| **JFROG\_CLI\_BUILD\_NAME**                        | Build name to be used by commands which expect a build name, unless sent as a command argument or option.                                                                                                                                                                                                                                                                                                                                                                      |
| **JFROG\_CLI\_BUILD\_NUMBER**                      | Build number to be used by commands which expect a build number, unless sent as a command argument or option.                                                                                                                                                                                                                                                                                                                                                                  |
| **JFROG\_CLI\_BUILD\_PROJECT**                     | JFrog project key to be used by commands that expect build name and build number. Determines the project of the published build.                                                                                                                                                                                                                                                                                                                                              |
| **JFROG\_CLI\_BUILD\_URL**                         | Sets the CI server build URL in the build-info. The "jf rt build-publish" command uses the value of this environment variable unless the --build-url command option is sent.                                                                                                                                                                                                                                                                                                  |
| **JFROG\_CLI\_ENV\_EXCLUDE**                       | <p>[Default: *password*;*secret*;*key*;*token*]<br><br>List of case insensitive patterns in the form of "value1;value2;...". Environment variables match those patterns will be excluded. This environment variable is used by the "jf rt build-publish" command, in case the --env-exclude command option is not sent.</p>                                                                                                                                                    |
| **JFROG\_CLI\_TRANSITIVE\_DOWNLOAD\_EXPERIMENTAL** | <p>[Default: false]<br><br>Used by the "jf rt download" command. Set to true to download artifacts also from remote repositories. This feature is experimental and available on Artifactory version 7.17.0 or higher.</p>                                                                                                                                                                                                                                                     |
| **JFROG\_CLI\_UPLOAD\_EMPTY\_ARCHIVE** | <p>[Default: false]<br><br>Used by the "jf rt upload" command. Set to true if you'd like to upload an empty archive when '--archive' is set but all files were excluded by exclusions pattern.</p>                                                                                                                                                                                                                                                     |

***

**Note**

> Read about additional environment variables at the [Welcome to JFrog CLI](../../README.md) page.

***

## Authentication

When used with Artifactory, JFrog CLI offers several means of authentication: JFrog CLI does not support accessing Artifactory without authentication.

### Authenticating with Username and Password / API Key

To authenticate yourself using your JFrog login credentials, either configure your credentials once using the **jf c add** command or provide the following option to each command.

| Command option | Description                                                           |
| -------------- | --------------------------------------------------------------------- |
| --url          | JFrog Artifactory API endpoint URL. It usually ends with /artifactory |
| --user         | JFrog username                                                        |
| --password     | JFrog password or API key                                             |

For enhanced security, when JFrog CLI is configured to use a username and password / API key, it automatically generates an access token to authenticate with Artifactory. The generated access token is valid for one hour only. JFrog CLI automatically refreshed the token before it expires. The **jfrog c add** command allows disabling this functionality. This feature is currently not supported by commands which use external tools or package managers or work with JFrog Distribution.

### Authenticating with an Access Token

To authenticate yourself using an Artifactory Access Token, either configure your Access Token once using the **jf c add** command or provide the following option to each command.

| Command option | Description                                                           |
| -------------- | --------------------------------------------------------------------- |
| --url          | JFrog Artifactory API endpoint URL. It usually ends with /artifactory |
| --access-token | JFrog access token                                                    |

### Authenticating with RSA Keys

***

**Note**

> Currently, authentication with RSA keys is not supported when working with external package managers and build tools (Maven, Gradle, Npm, Docker, Go and NuGet) or with the cUrl integration.

***

From version 4.4, Artifactory supports SSH authentication using RSA public and private keys. To authenticate yourself to Artifactory using RSA keys, execute the following instructions:

* Enable SSH authentication as described in [Configuring SSH](https://jfrog.com/help/r/jfrog-platform-administration-Documentation/Managing-Ssh-Keys).
*   Configure your Artifactory URL to have the following format: `ssh://[host]:[port]` There are two ways to do this:

    * For each command, use the `--url` command option.
    * Specify the Artifactory URL in the correct format using the **jfrog c add** command.

    ***

    **Warning**\
    \
    **Don't include your Artifactory context URL**

    > Make sure that the \[host] component of the URL only includes the hostname or the IP, but not your Artifactory context URL.

    ***
* Configure the path to your SSH key file. There are two ways to do this:
  * For each command, use the `--ssh-key-path` command option.
  * Specify the path using the **jfrog c add** command.

### Authenticating using Client Certificates (mTLS)

From Artifactory release 7.38.4, you can authenticate users using a client certificate ([mTLS](https://en.wikipedia.org/wiki/Mutual\_authentication#mTLS)). To do so will require a reverse proxy and some setup on the front reverse proxy (Nginx). Read about how to set this up [here](https://jfrog.com/help/r/jfrog-artifactory-documentation/Http-Settings).

To authenticate with the proxy using a client certificate, either configure your certificate once using the **jf c add** command or use the --`client-cert-path` and`--client-cert-ket-path` command options with each command.

***

**Note**

> Authentication using client certificates (mTLS) is not supported by commands which integrate with package managers.

***

Not Using a Public CA (Certificate Authority)?

This section is relevant for you if you're not using a public CA (Certificate Authority) to issue the SSL certificate used to connect to your Artifactory domain. You may not be using a public CA either because you're using self-signed certificates or you're running your own PKI services in-house (often by using a Microsoft CA).

In this case, you'll need to make those certificates available for JFrog CLI, by placing them inside the **security/certs** directory, which is under JFrog CLI's home directory. By default, the home directory is **\~/.jfrog**, but it can be also set using the **JFROG\_CLI\_HOME\_DIR** environment variable.

**Note**

1. The supported certificate format is PEM.
2. Some commands support the **--insecure-tls** option, which skips the TLS certificates verification.
3. Before version 1.37.0, JFrog CLI expected the certificates to be located directly under the **security** directory. JFrog CLI will automatically move the certificates to the new directory when installing version 1.37.0 or above. Downgrading back to an older version requires replacing the configuration directory manually. You'll find a backup if the old configuration under **.jfrog/backup**

## General Commands

The following sections describe the commands available in the JFrog CLI for use with Artifactory.

### Verifying Artifactory is Accessible

This command can be used to verify that Artifactory is accessible by sending an applicative ping to Artifactory.

#### Commands Params

|                   |                                                                                                                                                                 |
| ----------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name      | rt ping                                                                                                                                                         |
| Abbreviation      | rt p                                                                                                                                                            |
|                   |                                                                                                                                                                 |
| Command options   |                                                                                                                                                                 |
| --url             | <p>[Optional]<br><br>Artifactory URL.</p>                                                                                                                       |
| --server-id       | <p>[Optional]<br><br>Server ID configured using the <strong>jf c add</strong> command. If not specified, the default configured Artifactory server is used.</p> |
| --insecure-tls    | <p>[Default: false]<br><br>Set to true to skip TLS certificates verification.</p>                                                                               |
| Command arguments | The command accepts no arguments.                                                                                                                               |

#### Example 1

Ping the configured default Artifactory server.

```
jf rt ping
```

#### Example 2

Ping the configured Artifactory server with ID **rt-server-1**.

```
jf rt ping --server-id=rt-server-1
```

#### Example 3

Ping the Artifactory server. accessible through the specified URL.

```
jf rt ping --url=https://my-rt-server.com/artifactory
```

### Creating Access Tokens

This command allows creating [Access Tokens](https://jfrog.com/help/r/jfrog-platform-administration-Documentation/Access-Tokens) for users in Artifactory

#### Commands Params

|                   |                                                                                                                                                                                                                                                                                                                                                                                          |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name      | rt access-token-create                                                                                                                                                                                                                                                                                                                                                                   |
| Abbreviation      | rt atc                                                                                                                                                                                                                                                                                                                                                                                   |
| Command options   |                                                                                                                                                                                                                                                                                                                                                                                          |
| --groups          | <p>[Default: *]<br><br>A list of comma-separated groups for the access token to be associated with. Specify * to indicate that this is a 'user-scoped token', i.e., the token provides the same access privileges that the current subject has, and is therefore evaluated dynamically. A non-admin user can only provide a scope that is a subset of the groups to which he belongs</p> |
| --grant-admin     | <p>[Default: false]<br><br>Set to true to provides admin privileges to the access token. This is only available for administrators.</p>                                                                                                                                                                                                                                                  |
| --expiry          | <p>[Default: 3600]<br><br>The time in seconds for which the token will be valid. To specify a token that never expires, set to zero. Non-admin can only set a value that is equal to or less than the default 3600.</p>                                                                                                                                                                  |
| --refreshable     | <p>[Default: false]<br><br>Set to true if you'd like the the token to be refreshable. A refresh token will also be returned in order to be used to generate a new token once it expires.</p>                                                                                                                                                                                             |
| --audience        | <p>[Optional]<br><br>A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs, as obtained by the 'jf rt curl api/system/serviceid' command.</p>                                                                                                                                                   |
| Command arguments |                                                                                                                                                                                                                                                                                                                                                                                          |
| username          | Optional - The user name for which this token is created. If not specified, the configured user is used.                                                                                                                                                                                                                                                                                 |

#### Example

Create an access token for the user with the **commander-will-riker** username.

```
jf rt atc commander-will-riker
```

### Cleaning Up Unreferenced Files from a Git LFS Repository

This command is used to clean up files from a Git LFS repository. This deletes all files from a Git LFS repository, which are no longer referenced in a corresponding Git repository.

#### Commands Params

|                   |                                                                                                                                              |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name      | rt git-lfs-clean                                                                                                                             |
| Abbreviation      | rt glc                                                                                                                                       |
| Command options   |                                                                                                                                              |
| --refs            | \[Default: refs/remotes/\*] List of Git references in the form of "ref1,ref2,..." which should be preserved.                                 |
| --repo            | \[Optional] Local Git LFS repository in Artifactory which should be cleaned. If omitted, the repository is detected from the Git repository. |
| --quiet           | \[Default: false] Set to true to skip the delete confirmation message.                                                                       |
| --dry-run         | \[Default: false] If true, cleanup is only simulated. No files are actually deleted.                                                         |
| Command arguments | If no arguments are passed in, the command assumes the .git repository is located at current directory.                                      |
| path to .git      | Path to the directory which includes the .git directory.                                                                                     |

#### Example 1

Cleans up Git LFS files from Artifactory, using the configuration in the .git directory located at the current directory.

```
 jf rt glc
```

#### Example 2

Cleans up Git LFS files from Artifactory, using the configuration in the .git directory located inside the path/to/git/config directory.

```
jf rt glc path/to/git/config
```

### Running cUrl

Execute a cUrl command, using the configured Artifactory details. The command expects the cUrl client to be included in the PATH.

> **Note** - This command supports only Artifactory REST APIs, which are accessible under https://\<JFrog base URL>/artifactory/api/

#### Commands Params

|                          |                                                                                                                                                                                                                                                                                                  |   |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | - |
| Command name             | rt curl                                                                                                                                                                                                                                                                                          |   |
| Abbreviation             | rt cl                                                                                                                                                                                                                                                                                            |   |
| Command options          |                                                                                                                                                                                                                                                                                                  |   |
| --server-id              | <p>[Optional]<br><br>Server ID configured using the <strong>jf c add</strong> command. If not specified, the default configured server is used.</p>                                                                                                                                              |   |
| Command arguments        |                                                                                                                                                                                                                                                                                                  |   |
| cUrl arguments and flags | <p>The same list of arguments and flags passed to cUrl, except for the following changes:<br><br>1. The full Artifactory URL should not be passed. Instead, the REST endpoint URI should be sent.<br>2. The login credentials should not be passed. Instead, the --server-id should be used.</p> |   |

Currently only servers configured with username and password / API key are supported.

#### Example 1

Execute the cUrl client, to send a GET request to the /api/build endpoint to the default Artifactory server

```
jf rt curl -XGET /api/build
```

#### Example 2

Execute the cUrl client, to send a GET request to the /api/build endpoint to the configured my-rt-server server ID.

```
jf rt curl -XGET /api/build --server-id my-rt-server
```

## Downloading the Maven and Gradle Extractor JARs

For integrating with Maven and Gradle, JFrog CLI uses the build-info-extractor jars files. These jar files are downloaded by JFrog CLI from jcenter the first time they are needed.

If you're using JFrog CLI on a machine which has no access to the internet, you can configure JFrog CLI to download these jar files from an Artifactory instance. Here's how to configure Artifactory and JFrog CLI to download the jars files.

1. Create a remote Maven repository in Artifactory and name it **extractors**. When creating the repository, configure it to proxy [https://releases.jfrog.io/artifactory/oss-release-local](https://releases.jfrog.io/artifactory/oss-release-local)
2. Make sure that this Artifactory server is known to JFrog CLI, using the [jfrog c show](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration#showing-the-configured-servers) command. If not, configure it using the [jfrog c add](https://jfrog.com/help/r/jfrog-cli/Adding-and-Editing-Configured-Servers) command.
3. Set the **JFROG\_CLI\_EXTRACTORS\_REMOTE** environment variable with the server ID of the Artifactory server you configured, followed by a slash, and then the name of the repository you created. For example _**my-rt-server/extractors**_
