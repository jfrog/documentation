# cUrl Integration
### Running cUrl

Execute a cUrl command, using the configured Xray details. The command expects the cUrl client to be included in the PATH.

#### Commands Params

|                          |                                                                                                                                                                                                                                                                                           |
|--------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name             | xr curl                                                                                                                                                                                                                                                                                   |
| Abbreviation             | xr cl                                                                                                                                                                                                                                                                                     |
| **Command options:**          |                                                                                                                                                                                                                                                                                           |
| `--server-id` | <p>[Optional]<br>Server ID configured using the <em>jf c add</em> command. If not specified, the default configured server is used.</p>                                                                                                                                                   |
| **Command arguments:**        |                                                                                                                                                                                                                                                                                           |
| cUrl arguments and flags | <p>The same list of arguments and flags passed to cUrl, except for the following changes:<br><br>1. The full Xray URL should not be passed. Instead, the REST endpoint URI should be sent.<br>2. The login credentials should not be passed. Instead, the --server-id should be used.</p> |

#### Example 1

Execute the cUrl client, to send a GET request to the /api/system/version endpoint to the default configured Xray server.

```
jf xr curl -XGET /api/v1/system/version
```

#### Example 2

Execute the cUrl client, to send a GET request to the /api/v1/system/version endpoint to the configured my-xr-server server ID.

```
jf rt curl -XGET /api/v1/system/version --server-id my-xr-server
```
