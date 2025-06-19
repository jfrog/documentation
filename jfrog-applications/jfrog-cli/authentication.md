# Authentication

JFrog CLI is a command-line tool that enhances the automation and management of JFrog services, including Artifactory, Xray, and other components within the JFrog ecosystem. Authentication is a vital component of using JFrog CLI, ensuring secure interactions with the JFrog services.

When working with JFrog Xray, you have two primary authentication options: username/password pairs and access tokens. Each method allows you to secure access to your JFrog instance and interact with the API effectively.

&#x20;

#### Prerequisites

Before proceeding with authentication using JFrog CLI, ensure that you meet the following prerequisites:

1. **JFrog CLI Installed:** Make sure that you have the JFrog CLI installed on your system. You can download and install it from the[ JFrog CLI Download & Install](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/install).
2. **JFrog Account:** An active JFrog account with appropriate permissions to access the service. Ensure you have the necessary login credentials (username and password) or an access token.
3. **Token Validity (if using access tokens):** If you choose to authenticate with an access token, ensure that it is in a valid JWT format and has not expired. Review your token’s scope and permissions to confirm it grants the required access to Xray.

### Authentication using CLI

When using JFrog CLI with Xray, authentication is mandatory. JFrog CLI does not support access to Xray without valid authentication credentials. You can authenticate using either a username and password or an access token. Below are detailed instructions for both methods.

* Using password & Username
* Access token

#### Authenticating with Username and Password

To authenticate using your Xray login credentials:

* **Configuration Options**

You can configure your credentials permanently using the `jf c add` command. Alternatively, you can provide your credentials dynamically for each command.

**Configure Once Using `jf c add`**

Run the following command:

Follow the prompts to enter the necessary information:

`jf c`

* **Enter a unique server identifier:** Your chosen name for this configuration (e.g., xray\_server).
* J**Frog Platform URL:** The base URL for your JFrog instance (e.g., \<https://yourjfroginstance.jfrog.io).>
* **JFrog username:** Your username.
* **JFrog password:** Your password.

\


* **Using Command Options**

For each command, you can specify the following options:

| Command Option | Description                                                      |
| -------------- | ---------------------------------------------------------------- |
| --url          | JFrog Xray API endpoint URL. Typically ends with /xray.          |
| --user         | Your JFrog username.                                             |
| --password     | Your JFrog password.                                             |

**Example Command**

```
jf rt
ping
--url
"<https://yourjfroginstance.jfrog.io/xray>"
--user
"your_username"
--password
"your_password"
```



***

#### Authenticating with an Access Token

To authenticate using an Xray Access Token:

**Configuration Options**

Similar to username/password authentication, you can configure your access token using the jf c add command, or you can include it directly with each command.

**Configure Once Using `jf c add`**

When prompted, enter your access token instead of a password.

**Using Command Options**

You can specify the following options for authentication:

| Command Option   | Description                                                      |
| ---------------- | ---------------------------------------------------------------- |
| --url            | JFrog Xray API endpoint URL. Typically ends with /xray.          |
| --access-token   | Your JFrog access token. Ensure it is a valid JWT format token.  |

**Example Command**

```
jf rt
ping
--url
"<https://yourjfroginstance.jfrog.io/xray>"
--access-token
"your_access_token"
```



***

**Note**

* **Security:** Ensure that your credentials and access tokens are kept secure and not hardcoded in scripts wherever possible. Consider using environment variables or secure vaults for sensitive information.
* **Token Expiration:** Access tokens may have an expiration time. Be aware of this and renew your token as needed to maintain access.





***
