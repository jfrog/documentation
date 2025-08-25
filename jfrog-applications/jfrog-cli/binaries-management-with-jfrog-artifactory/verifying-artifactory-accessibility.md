# Verifying Artifactory's Accessibility

## Overview

Use this command to verify that Artifactory is accessible by sending an application ping to Artifactory.

## Commands Params

|                        |                                                                                                                                                             |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name           | `rt ping`                                                                                                                                                   |
| Abbreviation           | `rt p`                                                                                                                                                      |
|                        |                                                                                                                                                             |
| **Command options:**   |                                                                                                                                                             |
| `--url`                | <p>[Optional]<br>JFrog Artifactory URL. (example: https://acme.jfrog.io/artifactory)</p>                                                                    |
| `--server-id`          | <p>[Optional]<br>Server ID configured using the <strong>jf c add</strong> command. If not specified, the default configured Artifactory server is used.</p> |
| `--insecure-tls`       | <p>[Default: false]<br>Set to true to skip TLS certificates verification.</p>                                                                               |
| **Command arguments:** | The command accepts no arguments.                                                                                                                           |

## Examples

### Example 1: Ping the default server

Pings the configured default Artifactory server.

```
jf rt ping
```

### Example 2: Ping a specific server by ID

Pings the configured Artifactory server with the ID `rt-server-1`.

```
jf rt ping --server-id=rt-server-1
```

### Example 3: Ping a specific server by URL

Pings the Artifactory server accessible at the specified URL.

```
jf rt ping --url=https://my-rt-server.com/artifactory
```
