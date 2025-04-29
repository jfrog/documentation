# Authentication

When used with Xray, JFrog CLI offers several means of authentication: JFrog CLI does not support accessing Xray without authentication.

#### Authenticating with Username and Password

To authenticate yourself using your Xray login credentials, either configure your credentials once using the _**jf c add**_ command or provide the following option to each command.

<table><thead><tr><th width="424.5">Command Option</th><th>Description</th></tr></thead><tbody><tr><td>--url</td><td>JFrog Xray API endpoint URL. It usually ends with /xray</td></tr><tr><td>--user</td><td>JFrog username</td></tr><tr><td>--password</td><td>JFrog password</td></tr></tbody></table>

#### Authenticating with an Access Token

To authenticate yourself using an Xray Access Token, either configure your Access Token once using the _**jf c add**_ command or provide the following option to each command.

<table><thead><tr><th width="218.5">Command Option</th><th>Description</th></tr></thead><tbody><tr><td>--url</td><td>JFrog Xray API endpoint URL. It usually ends with /xray</td></tr><tr><td>--access-token</td><td>JFrog access token</td></tr></tbody></table>

***
