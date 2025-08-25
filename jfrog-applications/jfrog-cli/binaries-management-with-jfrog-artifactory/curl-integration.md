# cURL Integration

## Overview

This command executes a `cURL` command using the Artifactory details you have configured. The command expects the `cURL` client to be included in your `PATH`.

Note: This command supports only Artifactory REST APIs, which are accessible under `https://<JFrog base URL>/artifactory/api/`.

## Commands Params

|                          |                                                                                                                                                                                                                                                                                                  |   |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | - |
| Command name             | rt curl                                                                                                                                                                                                                                                                                          |   |
| Abbreviation             | rt cl                                                                                                                                                                                                                                                                                            |   |
| **Command options:**     |                                                                                                                                                                                                                                                                                                  |   |
| `--server-id`            | <p>[Optional]<br>Server ID configured using the <strong>jf c add</strong> command. If not specified, the default configured server is used.</p>                                                                                                                                                  |   |
| **Command arguments:**   |                                                                                                                                                                                                                                                                                                  |   |
| cUrl arguments and flags | <p>The same list of arguments and flags passed to cUrl, except for the following changes:<br><br>1. The full Artifactory URL should not be passed. Instead, the REST endpoint URI should be sent.<br>2. The login credentials should not be passed. Instead, the --server-id should be used.</p> |   |

Currently only servers configured with username and password / API key are supported.

### Examples

#### Example 1: Send a Request to the Default Server

Execute the `cURL` client to send a GET request to the `/api/build` endpoint on the default configured Artifactory server.

```
jf rt curl -XGET /api/build
```

#### Example 2: Send a Request to a Specific Server

Execute the `cURL` client to send a GET request to the `/api/build` endpoint on the Artifactory server configured with the ID `my-rt-server`.

```
jf rt curl -XGET /api/build --server-id my-rt-server
```
