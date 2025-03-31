# JFrog Platform Configuration

## Web Login to the JFrog Platform

You can use the `jf login` command to authenticate with the JFrog Platform through the web browser.
This command is solely interactive, meaning it does not receive any options and cannot be used in a CI server.

![](../../../.gitbook/assets/login-page.png)
![](../../../.gitbook/assets/login-successful.png)

## Creating Access Tokens

This command allows creating [Access Tokens](https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-tokens) for users in the JFrog Platform. By default, a user-scoped token will be created. Administrators may provide the scope explicitly with '--scope', or implicitly with '--groups', '--grant-admin'.

### Commands Params

|                   |                                                                                                                                                                                                                                                                                                                                                                    |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name      | access-token-create                                                                                                                                                                                                                                                                                                                                                |
| Abbreviation      | atc                                                                                                                                                                                                                                                                                                                                                                |
| **Command arguments:** |                                                                                                                                                                                                                                                                                                                                                                    |
| username          | The username for which this token is created. If not specified, the token will be created for the current user.                                                                                                                                                                                                                                                    |
| **Command options:**   |                                                                                                                                                                                                                                                                                                                                                                    |
| `--audience` | <p>[Optional]</p><p>A space-separated list of the other instances or services that should accept this token identified by their Service-IDs.</p>                                                                                                                                                                                                                   |
| `--description` | <p>[Optional]</p><p>Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.</p>                                                                                                                                                                                                                                         |
| `--expiry` | <p>[Optional]</p><p>The amount of time, in seconds, it would take for the token to expire. Must be non-negative. If not provided, the platform default will be used. To specify a token that never expires, set to zero. Non-admin may only set a value that is equal or lower than the platform default that was set by an administrator (1 year by default).</p> |
| `--grant-admin` | <p>[Default: false]</p><p>Set to true to provide admin privileges to the access token. This is only available for administrators.</p>                                                                                                                                                                                                                              |
| `--groups` | <p>[Optional]</p><p>A list of comma-separated(,) groups for the access token to be associated with. This is only available for administrators.</p>                                                                                                                                                                                                                 |
| `--project` | <p>[Optional]</p><p>The project for which this token is created. Enter the project name on which you want to apply this token.</p>                                                                                                                                                                                                                                 |
| `--reference` | <p>[Default: false]</p><p>Generate a Reference Token (alias to Access Token) in addition to the full token (available from Artifactory 7.38.10).</p>                                                                                                                                                                                                               |
| `--refreshable` | <p>[Default: false]</p><p>Set to true if you'd like the token to be refreshable. A refresh token will also be returned in order to be used to generate a new token once it expires.</p>                                                                                                                                                                            |
| `--scope` | <p>[Optional]</p><p>The scope of access that the token provides. This is only available for administrators.</p>                                                                                                                                                                                                                                                    |

### Examples

#### Example 1

Create an access token for the user in the default server configured by the [jf c add](#adding-and-editing-configured-servers) command:

```
jf atc
```

#### Example 2

Create an access token for the user with the **toad** username:

```
jf atc toad
```

## Adding and Editing Configured Servers

The **config add** and **config edit** commands are used to add and edit JFrog Platform server configuration, stored in JFrog CLI's configuration storage. These configured servers can be used by the other commands. The configured servers' details can be overridden per command by passing in alternative values for the URL and login credentials. The values configured are saved in a file under the JFrog CLI home directory.

Starting from version 2.75.0, `jf c add` supports authentication using OIDC tokens. This is particularly useful for CI environments. When using OIDC, the command must be run in non-interactive mode (e.g. using `--interactive=false`).

|                        |                                                                                                                                                                                                                                                                                                                                                                                                    |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command Name           | config add / config edit                                                                                                                                                                                                                                                                                                                                                                           |
| Abbreviation           | c add / c edit                                                                                                                                                                                                                                                                                                                                                                                     |
| **Command options:**        |                                                                                                                                                                                                                                                                                                                                                                                                    |
| `--access-token` | <p>[Optional]</p><p>Access token.</p>                                                                                                                                                                                                                                                                                                                                                              |
| `--artifactory-url` | <p>[Optional]</p><p>JFrog Artifactory URL. (example: https://acme.jfrog.io/artifactory)</p>                                                                                                                                                                                                                                                                                                        |
| `--basic-auth-only` | <p>[Default: false]</p><p>Used for Artifactory authentication. Set to true to disable replacing username and password/API key with automatically created access token that's refreshed hourly. Username and password/API key will still be used with commands which use external tools or the JFrog Distribution service. Can only be passed along with username and password/API key options.</p> |
| `--client-cert-key-path` | <p>[Optional]</p><p>Private key file for the client certificate in PEM format.</p>                                                                                                                                                                                                                                                                                                                 |
| `--client-cert-path` | <p>[Optional]</p><p>Client certificate file in PEM format.</p>                                                                                                                                                                                                                                                                                                                                     |
| `--dist-url` | <p>[Optional]</p><p>Distribution URL. (example: https://acme.jfrog.io/distribution)</p>                                                                                                                                                                                                                                                                                                            |
| `--enc-password` | <p>[Default: true]<br>If true, the configured password will be encrypted using Artifactory's <a href="https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetUserEncryptedPassword">encryption API</a> before being stored. If false, the configured password will not be encrypted.</p>                                                                          |
| `--insecure-tls` | <p>[Default: false]</p><p>Set to true to skip TLS certificates verification, while encrypting the Artifactory password during the config process.</p>                                                                                                                                                                                                                                              |
| `--interactive` | <p>[Default: true, unless $CI is true]</p><p>Set to false if you do not want the config command to be interactive.</p>                                                                                                                                                                                                                                                                             |
| `--mission-control-url` | <p>[Optional]</p><p>JFrog Mission Control URL. (example: https://acme.jfrog.io/ms)</p>                                                                                                                                                                                                                                                                                                             |
| `--password` | <p>[Optional]</p><p>JFrog Platform password.</p>                                                                                                                                                                                                                                                                                                                                                   |
| `--ssh-key-path` | <p>[Optional]</p><p>For authentication with Artifactory. SSH key file path.</p>                                                                                                                                                                                                                                                                                                                    |
| `--url` | <p>[Optional]</p><p>JFrog Platform URL. (example: https://acme.jfrog.io)</p>                                                                                                                                                                                                                                                                                                                       |
| `--user` | <p>[Optional]</p><p>JFrog Platform username.</p>                                                                                                                                                                                                                                                                                                                                                   |
| `--xray-url` | \[Optional] Xray URL. (example: https://acme.jfrog.io/xray)                                                                                                                                                                                                                                                                                                                                        |
| `--overwrite` | <p>[Available for <em>config add</em> only]<br>[Default: false]<br>Overwrites the instance configuration if an instance with the same ID already exists.</p>                                                                                                                                                                                                                                       |
| `--oidc-provider-name` | <p>[Optional]</p><p>OIDC provider name for CI authentication.</p> |
| `--oidc-provider-type` | <p>[Default: GitHub]</p><p>OIDC provider type (e.g., GitHub, Azure).</p> |
| `--oidc-token-id` | <p>[Optional]</p><p>The OIDC token ID to use for authentication.</p> |
| `--oidc-audience` | <p>[Optional]</p><p>Audience for the OIDC token.</p> |
| `--application-key` | <p>[Optional]</p><p>JFrog Application Key to associate with the authentication.</p> |
| **Command arguments:**      |                                                                                                                                                                                                                                                                                                                                                                                                    |
| server ID              | A unique ID for the server configuration.                                                                                                                                                                                                                                                                                                                                                          |

## Removing Configured Servers

The _config remove_ command is used to remove JFrog Platform server configuration, stored in JFrog CLI's configuration storage.

|                   |                                                                                      |
|-------------------|--------------------------------------------------------------------------------------|
| Command name      | config remove                                                                        |
| Abbreviation      | c rm                                                                                 |
| **Command options:**   |                                                                                      |
| `--quiet` | <p>[Default: $CI]</p><p>Set to true to skip the delete confirmation message.</p>     |
| **Command arguments:** |                                                                                      |
| server ID         | The server ID to remove. If no argument is sent, all configured servers are removed. |


## Showing the Configured Servers

The _config show_ command shows the stored configuration. You may show a specific server's configuration by sending its ID as an argument to the command.

|                   |                                                                                         |
|-------------------|-----------------------------------------------------------------------------------------|
| Command name      | config show                                                                             |
| Abbreviation      | c s                                                                                     |
| **Command arguments:** |                                                                                         |
| server ID         | The ID of the server to show. If no argument is sent, all configured servers are shown. |

## Setting a Server as Default

The _config use_ command sets a configured server as default. The following commands will use this server.

|                   |                                         |
|-------------------|-----------------------------------------|
| Command name      | config use                              |
| **Command arguments:** |                                         |
| server ID         | The ID of the server to set as default. |

## Exporting and Importing Configuration

The _config export_ command generates a token, which stores the server configuration. This token can be used by the _config import_ command, to import the configuration stored in the token, and save it in JFrog CLI's configuration storage.

### Export

|                   |                                |
|-------------------|--------------------------------|
| Command name      | config export                  |
| Abbreviation      | c ex                           |
| **Command arguments:** |                                |
| server ID         | The ID of the server to export |

### Import

|                   |                     |
|-------------------|---------------------|
| Command name      | config import       |
| Abbreviation      | c im                |
| **Command arguments:** |                     |
| server token      | The token to import |

## Sensitive Data Encryption

### File-Based Encryption

Starting from version 1.37.0, JFrog CLI introduces support for encrypting sensitive data stored in its configuration using an encryption key stored in a file. Follow these steps to enable encryption:

1. Generate a random 32-character master key. Ensure that the key size is exactly 32 characters. For example: _f84hc22dQfhe9f8ydFwfsdn48!wejh8A_
2. Create a file named **security.yaml** under **~/.jfrog/security**.

   > If you've customized the default JFrog CLI home directory by setting the JFROG_CLI_HOME_DIR environment variable, create the **security/security.yaml** file under the configured home directory.

3. Add the generated master key to the **security.yaml** file:

   ```yaml
   version: 1
   masterKey: "your master key"
   ```

4. Ensure that the **security.yaml** file has only read permissions for the user running JFrog CLI.

The configuration will be encrypted the next time JFrog CLI accesses the config. If you have existing configurations stored before creating the file, you'll need to reconfigure the servers stored in the config.

> **Warning:** When upgrading JFrog CLI from a version prior to 1.37.0 to version 1.37.0 or above, automatic changes are made to the content of the **~/.jfrog** directory to support the new functionality introduced. Before making these changes, the content of the **~/.jfrog** directory is backed up inside the **~/.jfrog/backup** directory. After enabling sensitive data encryption, it is recommended to remove the **backup** directory to ensure no sensitive data is left unencrypted.

### Environment Variable-Based Encryption

Starting from version 2.36.0, JFrog CLI also supports encrypting sensitive data in its configuration using an encryption key stored in an environment variable. To enable encryption, follow these steps:

1. Generate a random 32-character master key. Ensure that the key size is exactly 32 characters. For example: _f84hc22dQfhe9f8ydFwfsdn48!wejh8A_
2. Store the key in an environment variable named **JFROG_CLI_ENCRYPTION_KEY**.

The configuration will be encrypted the next time JFrog CLI attempts to access the config. If you have configurations already stored before setting the environment variable, you'll need to reconfigure the servers stored in the config.


---



### Exchanging OIDC Token for Access Token

The `exchange-oidc-token` (alias: `eot`) command is used to exchange an OIDC token (such as those provided by GitHub Actions or other CI systems) for a JFrog Platform access token and associated username. This is useful in automation workflows where credentials must be derived securely via an identity provider.

🔒 Important Notes about OIDC Authentication:

OIDC tokens are short-lived and **do not support refresh**.
OIDC access tokens are not renewable. They are designed for one-time use during CI pipelines and do not have an automatic refresh mechanism like other types of tokens. As a result, the authentication will remain valid only for the duration of that pipeline or until the token expires.


This flow is built specifically for CI/CD pipelines where the identity provider
(e.g. GitHub Actions) issues a fresh OIDC tokenID for every pipeline run.
This ensures maximum security and minimizes long-lived secrets.

|                   |                                                                                  |
|-------------------|----------------------------------------------------------------------------------|
| Command name      | exchange-oidc-token                                                              |
| Abbreviation      | eot                                                                              |
| **Command options:** |                                                                                  |
| `--platformUrl`       | <p>[Mandatory]</p><p>The URL of the JFrog Platform instance.</p>            |
| `--oidc-token-id`     | <p>[Mandatory]</p><p>The token ID from the OIDC provider.</p>               |
| `--oidc-provider-name`| <p>[Mandatory]</p><p>The name of the OIDC provider.</p>                     |
| `--oidc-audience`     | <p>[Optional]</p><p>The audience for the OIDC token.</p>                    |
| `--oidc-provider-type`| <p>[Optional, default: GitHub]</p><p>The type of provider (e.g. GitHub).</p>|
| `--application-key`   | <p>[Optional]</p><p>JFrog Application key for attribution.</p>              |
| `--project`           | <p>[Optional]</p><p>Project key (if applicable).</p>                         |
| `--repository`        | <p>[Optional]</p><p>Source code repository name.</p>                         |

### Example

```bash
jf eot --platformUrl=https://platform.jfrog.io \
       --oidc-token-id=$JFROG_CLI_OIDC_EXCHANGE_TOKEN_ID \
       --oidc-provider-name=my-intergraion-name
```

### Sample Output
```bash
{ AccessToken: **** Username: **** }
```