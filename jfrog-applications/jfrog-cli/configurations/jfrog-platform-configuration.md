# JFrog Platform Configuration

## Web Sign-in to the JFrog Platform

Use the `jf login` command to authenticate with the JFrog Platform through a web browser. This command is solely interactive; it does not receive any options and cannot be used in a CI server.

![](../../../.gitbook/assets/login-page.png) ![](../../../.gitbook/assets/login-successful.png)

## Create Access Tokens

Use this command to create [Access Tokens](https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-tokens) in the JFrog Platform. By default, a user-scoped token is created. Administrators can provide the scope explicitly with `--scope`, or implicitly with `--groups` or `--grant-admin`.

### Commands Arguments

|                        |                                                                                                                                                                                                                                                                                                                             |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name           | access-token-create                                                                                                                                                                                                                                                                                                         |
| Abbreviation           | atc                                                                                                                                                                                                                                                                                                                         |
| **Command arguments:** |                                                                                                                                                                                                                                                                                                                             |
| username               | The username for whom you are creating the token. If not specified, the token is created for the current user.                                                                                                                                                                                                              |
| **Command options:**   |                                                                                                                                                                                                                                                                                                                             |
| `--audience`           | <p>[Optional]</p><p>Specifies a space-separated list of the other instances or services, identified by their Service-IDs, that should accept this token.</p>                                                                                                                                                                |
| `--description`        | <p>[Optional]</p><p>Provides a free text description for the token. This is useful for filtering and managing tokens and is limited to 1024 characters.</p>                                                                                                                                                                 |
| `--expiry`             | <p>[Optional]</p><p>The amount of time, in seconds, until the token expires. This value must be non-negative. If not provided, the platform default is used. To specify a token that never expires, set the value to zero. Non-administrators may only set a value that is equal to or lower than the platform default.</p> |
| `--grant-admin`        | <p>[Default: false]</p><p>Set to <code>true</code> to provide admin privileges to the access token. This option is only available for administrators.</p>                                                                                                                                                                   |
| `--groups`             | <p>[Optional]</p><p>A comma-separated list of groups to associate with the access token. This option is only available for administrators.</p>                                                                                                                                                                              |
| `--project`            | <p>[Optional]</p><p>The project for which this token is created. Enter the project name on which you want to apply this token.</p>                                                                                                                                                                                          |
| `--reference`          | <p>[Default: false]</p><p>Generate a Reference Token (an alias for an Access Token) in addition to the full token. This is available from Artifactory 7.38.10.</p>                                                                                                                                                          |
| `--refreshable`        | <p>[Default: false]</p><p>Set to <code>true</code> to make the token refreshable. A refresh token is also returned to be used to generate a new token after expiration.</p>                                                                                                                                                 |
| `--scope`              | <p>[Optional]</p><p>The scope of access that the token provides. This option is only available for administrators.</p>                                                                                                                                                                                                      |

### Examples

#### Example 1

Create an access token for the current user on the default server configured by the `jf c add` command.

```
jf atc
```

#### Example 2

Create an access token for the user with the username `toad`.

```
jf atc toad
```

## Add and Edit Configured Servers

The `config add` or `config edit` commands add and edit JFrog Platform server configurations, which are stored in JFrog CLI's configuration storage. Other commands can use these configured servers. You can override the configured server details for any command by passing in alternative values for the URL and sign-in credentials. The configured values are saved in a file under the JFrog CLI home directory.

|                          |                                                                                                                                                                                                                                                                                                                           |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command Name             | config add / config edit                                                                                                                                                                                                                                                                                                  |
| Abbreviation             | c add / c edit                                                                                                                                                                                                                                                                                                            |
| **Command options:**     |                                                                                                                                                                                                                                                                                                                           |
| `--access-token`         | <p>[Optional]</p><p>Access token.</p>                                                                                                                                                                                                                                                                                     |
| `--artifactory-url`      | \[Optional] Provides the JFrog Artifactory URL (for example, acme.jfrog.io/artifactory).                                                                                                                                                                                                                                  |
| `--basic-auth-only`      | \[Default: false] For Artifactory authentication. Set to `true` to disable replacing the username and password/API key with an automatically created access token.                                                                                                                                                        |
| `--client-cert-key-path` | <p>[Optional]</p><p>Private key file for the client certificate in PEM format.</p>                                                                                                                                                                                                                                        |
| `--client-cert-path`     | <p>[Optional]</p><p>Client certificate file in PEM format.</p>                                                                                                                                                                                                                                                            |
| `--dist-url`             | <p>[Optional]</p><p>Distribution URL. (example: https://acme.jfrog.io/distribution)</p>                                                                                                                                                                                                                                   |
| `--enc-password`         | <p>[Default: true]<br>If true, the configured password will be encrypted using Artifactory's <a href="https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetUserEncryptedPassword">encryption API</a> before being stored. If false, the configured password will not be encrypted.</p> |
| `--insecure-tls`         | <p>[Default: false]</p><p>Set to true to skip TLS certificates verification, while encrypting the Artifactory password during the config process.</p>                                                                                                                                                                     |
| `--interactive`          | <p>[Default: true, unless $CI is true]</p><p>Set to false if you do not want the config command to be interactive.</p>                                                                                                                                                                                                    |
| `--mission-control-url`  | <p>[Optional]</p><p>JFrog Mission Control URL. (example: https://acme.jfrog.io/ms)</p>                                                                                                                                                                                                                                    |
| `--password`             | <p>[Optional]</p><p>JFrog Platform password.</p>                                                                                                                                                                                                                                                                          |
| `--ssh-key-path`         | <p>[Optional]</p><p>For authentication with Artifactory. SSH key file path.</p>                                                                                                                                                                                                                                           |
| `--url`                  | <p>[Optional]</p><p>JFrog Platform URL. (example: https://acme.jfrog.io)</p>                                                                                                                                                                                                                                              |
| `--user`                 | <p>[Optional]</p><p>JFrog Platform username.</p>                                                                                                                                                                                                                                                                          |
| `--xray-url`             | \[Optional] Xray URL. (example: https://acme.jfrog.io/xray)                                                                                                                                                                                                                                                               |
| `--overwrite`            | <p>[Available for <em>config add</em> only]<br>[Default: false]<br>Overwrites the instance configuration if an instance with the same ID already exists.</p>                                                                                                                                                              |
| **Command arguments:**   |                                                                                                                                                                                                                                                                                                                           |
| server ID                | A unique ID for the server configuration.                                                                                                                                                                                                                                                                                 |

## Remove Configured Servers

The `config remove` command removes a JFrog Platform server configuration from JFrog CLI's configuration storage.

|                        |                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------- |
| Command name           | config remove                                                                            |
| Abbreviation           | c rm                                                                                     |
| **Command options:**   |                                                                                          |
| `--quiet`              | <p>[Default: $CI]</p><p>Set to true to skip the delete confirmation message.</p>         |
| **Command arguments:** |                                                                                          |
| server ID              | The server ID to remove. If no argument is provided, all configured servers are removed. |

## Show Configured Servers

The `config show` command shows the stored configuration. To show a specific server's configuration, provide its ID as an argument.

|                        |                                                                                             |
| ---------------------- | ------------------------------------------------------------------------------------------- |
| Command name           | config show                                                                                 |
| Abbreviation           | c s                                                                                         |
| **Command arguments:** |                                                                                             |
| server ID              | The ID of the server to show. If no argument is provided, all configured servers are shown. |

## Set a Server as Default

The `config use` command sets a configured server as the default for subsequent commands.

|                        |                                         |
| ---------------------- | --------------------------------------- |
| Command name           | config use                              |
| **Command arguments:** |                                         |
| server ID              | The ID of the server to set as default. |

## Export and Import Configuration

The `config export` command generates a token that stores a server configuration. The `config import` command uses this token to import the configuration and save it to JFrog CLI's configuration storage.

### Export

|                        |                                |
| ---------------------- | ------------------------------ |
| Command name           | config export                  |
| Abbreviation           | c ex                           |
| **Command arguments:** |                                |
| server ID              | The ID of the server to export |

### Import

|                        |                     |
| ---------------------- | ------------------- |
| Command name           | config import       |
| Abbreviation           | c im                |
| **Command arguments:** |                     |
| server token           | The token to import |

## Sensitive Data Encryption

### File-Based Encryption

Starting from version 1.37.0, JFrog CLI supports encrypting sensitive configuration data using an encryption key stored in a file. To enable encryption:

1. Generate a random 32-character master key. The key must be exactly 32 characters. For example: `f84hc22dQfhe9f8ydFwfsdn48!wejh8A`
2.  Create a file named `security.yaml` under `~/.jfrog/security`. If you customized the JFrog CLI home directory using the `JFROG_CLI_HOME_DIR` environment variable, create the file in the configured home directory.

    > If you've customized the default JFrog CLI home directory by setting the JFROG\_CLI\_HOME\_DIR environment variable, create the **security/security.yaml** file under the configured home directory.
3.  Add the generated master key to the **security.yaml** file:

    ```yaml
    version: 1
    masterKey: "your master key"
    ```
4. Ensure that the **security.yaml** file has only read permissions for the user running JFrog CLI.

The configuration is encrypted the next time JFrog CLI accesses it. If you have existing configurations, you must reconfigure the servers.

> **Warning**: When upgrading JFrog CLI from a version prior to 1.37.0, the `~/.jfrog` directory is backed up to `~/.jfrog/backup`. After enabling encryption, it is recommended to remove the backup directory to ensure no sensitive data is left unencrypted.

### Environment Variable-Based Encryption

Starting from version 2.36.0, JFrog CLI also supports encryption using a key stored in an environment variable. To enable this method:

1. Generate a random 32-character master key. Ensure that the key size is exactly 32 characters. For example: _**f84hc22dQfhe9f8ydFwfsdn48!wejh8A**_
2. Store the key in an environment variable named `JFROG_CLI_ENCRYPTION_KEY`.

The configuration is encrypted the next time JFrog CLI accesses it. If you have existing configurations, you must reconfigure the servers.
