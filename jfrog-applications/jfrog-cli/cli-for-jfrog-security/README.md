# CLI for JFrog Security

### Authentication

When used with Xray, JFrog CLI offers several means of authentication: JFrog CLI does not support accessing Xray without authentication.

#### Authenticating with Username and Password

To authenticate yourself using your Xray login credentials, either configure your credentials once using the\_jf c add\_command or provide the following option to each command.

<table>
   <thead>
      <tr>
         <th width="424.5"></th>
         <th></th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td>Command option</td>
         <td>Description</td>
      </tr>
      <tr>
         <td>--url</td>
         <td>JFrog Xray API endpoint URL. It usually ends with /xray</td>
      </tr>
      <tr>
         <td>--user</td>
         <td>JFrog username</td>
      </tr>
      <tr>
         <td>--password</td>
         <td>JFrog password</td>
      </tr>
   </tbody>
</table>

#### Authenticating with an Access Token

To authenticate yourself using an Xray Access Token, either configure your Access Token once using the \_jf c add\_command or provide the following option to each command.

|                |                                                         |
|----------------|---------------------------------------------------------|
| Command option | Description                                             |
| --url          | JFrog Xray API endpoint URL. It usually ends with /xray |
| --access-token | JFrog access token                                      |
