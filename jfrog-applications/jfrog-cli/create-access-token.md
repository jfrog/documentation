# Creating Access Tokens

## Overview

This page describes the usage of the `jf access-token-create` command.  
The command allows creating [Access Tokens](https://jfrog.com/help/r/jfrog-rest-apis/access-tokens) for users in Artifactory via the Access service [REST API](https://jfrog.com/help/r/jfrog-rest-apis/access-tokens).

**Note**

> This command will require authentication with a valid access token. 
It is possible to provide the token with each command, as shown in the example below, or to configure your server once by adding the `--access-token` option in the `jf c add / edit` command. You can find the `jf c add / edit` command
in the [JFrog Platform Configuration](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/configurations/jfrog-platform-configuration) page.
***

#### Commands Params

|                   |                                                                                                                                                                                                                                                                                                                                                                                          |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name      |  access-token-create                                                                                                                                                                                                                                                                                                                                                                   |
| Abbreviation      |  atc                                                                                                                                                                                                                                                                                                                                                                                   |
| Command options   |                                                                                                                                                                                                                                                                                                                                                                                          |
| --access-token    | <p>[Optional]<br><br>JFrog access token.</p>           |
| --audience        | <p>[Optional]<br><br>A space-separated list of the other instances or services that should accept this token identified by their Service-IDs.</p>           |
| --description     | <p>[Optional]<br><br>Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.</p> 
| --expiry          | <p>[Optional]<br><br>The amount of time, in seconds, it would take for the token to expire. Must be non-negative. If not provided, the platform default (1 year) will be used. To specify a token that never expires, set to zero. Non-admin may only set a value that is equal to or lower than the platform default that was set by an administrator.</p>                                |
| --grant-admin     | <p>[Default: false]<br><br>Set to true to provide admin privileges to the access token. This is only available for administrators.</p>                                                                                                               |
| --groups          | <p>[Optional]<br><br>A list of comma-separated groups for the access token to be associated with. This is only available for administrators.</p> |
| --reference       | <p>[Default: false]<br><br>Set to true to generate a Reference Token (alias to Access Token) in addition to the full token (available from Artifactory 7.38.10)</p> 
| --refreshable     | <p>[Default: false]<br><br>Set to true if you'd like the token to be refreshable. A refresh token will also be returned in order to be used to generate a new token once it expires.</p>                                                                  |
| --scope           | <p>[Optional]<br><br>The scope of access that the token provides. This is only available for administrators.</p>                                                                                                                                                             |
| Command arguments |                                                                                                                                                                                                                                                                                                                                                                                          |
| username          | Optional - The user name for which this token is created. If not specified, the configured user is used.                                                                                                                                                                                                                                                                                 |

#### Example

Create an access token for username **commander-will-riker**.

```
jf atc commander-will-riker --scope=applied-permissions/groups:readers,riker-group --reference=true --expiry=31556952 --description="Token for Riker-group. Owner topdog@riker.com"
```