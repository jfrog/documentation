# cURL Integration
## Overview

Execute a cURL command, using the configured Artifactory details. The command expects the cUrl client to be included in the PATH.

> **Note** - This command supports only Artifactory REST APIs, which are accessible under https://\<JFrog base URL>/artifactory/api/

## Commands Params

|                          |                                                                                                                                                                                                                                                                                                  |   |
|--------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|
| Command name             | rt curl                                                                                                                                                                                                                                                                                          |   |
| Abbreviation             | rt cl                                                                                                                                                                                                                                                                                            |   |
| **Command options:**          |                                                                                                                                                                                                                                                                                                  |   |
| `--server-id` | <p>[Optional]<br>Server ID configured using the <strong>jf c add</strong> command. If not specified, the default configured server is used.</p>                                                                                                                                                  |   |
| **Command arguments:**        |                                                                                                                                                                                                                                                                                                  |   |
| cUrl arguments and flags | <p>The same list of arguments and flags passed to cUrl, except for the following changes:<br><br>1. The full Artifactory URL should not be passed. Instead, the REST endpoint URI should be sent.<br>2. The login credentials should not be passed. Instead, the --server-id should be used.</p> |   |

Currently only servers configured with username and password / API key are supported.

## Examples
### Example 1

Execute the cUrl client, to send a GET request to the /api/build endpoint to the default Artifactory server

```
jf rt curl -XGET /api/build
```

### Example 2

Execute the cUrl client, to send a GET request to the /api/build endpoint to the configured my-rt-server server ID.

```
jf rt curl -XGET /api/build --server-id my-rt-server
```
