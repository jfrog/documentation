# Extension Settings

The extension is licensed under Apache License 2.0.

To open the extension settings, use the following VS Code menu command:

* On Windows/Linux - File > Preferences > Settings > Extensions > JFrog
* On macOS - Code > Preferences > Settings > Extensions > JFrog

#### Exclude Paths from Scan

By default, paths containing the words `test`, `venv` and `node_modules` are excluded from Xray scan. The exclude pattern can be configured in the [Extension Settings](extension-settings.md).

#### Proxy Configuration

If your JFrog environment is behind an HTTP/S proxy, follow these steps to configure the proxy server:

1. Go to Preferences --> Settings --> Application --> Proxy
2. Set the proxy URL under 'Proxy'.
3. Make sure 'Proxy Support' is 'override' or 'on'.

* Alternatively, you can use the HTTP\_PROXY and HTTPS\_PROXY environment variables.

### External Resource Repository

By default, the JFrog extension downloads the necessary tools needed from https://releases.jfrog.io. If the machine that runs JFrog extension has no access to it, you can create a remote repository in Artifactory which proxy https://releases.jfrog.io and set the JFrog extension setting:

![externalResourcesRepository](../../.gitbook/assets/vscode/externalResourcesRepository.png)

 or set the following enviable variable

```
JFROG_IDE_RELEASES_REPO=jfrog-releases-repository
```


To set up a remote repository that acts as a proxy for [https://releases.jfrog.io](https://releases.jfrog.io), follow these steps:

1. Log in using credentials with administrative privileges.

2. Create a Remote Repository:
   - Navigate to the Remote Repository creation section.
   - Configure the repository with the following properties:

   Basic Configuration:
   - Package Type: Generic
   - Repository Key: jfrog-releases-repository
   - URL: [https://releases.jfrog.io](https://releases.jfrog.io)

   Advanced Configuration:
   - Uncheck the 'Store Artifacts Locally' option.

These settings will establish the remote repository as a proxy for the specified URL.
Remember to set `jfrog-releases-repository` as the releases repository using either an environment variable or in the External Resource Repository within the extension settings.

#### Proxy Authorization

If your proxy server requires credentials, follow these steps:

1. Follow 1-3 steps under [Proxy configuration](extension-settings#proxy-configuration).

**Basic authorization**

1. Encode with base64: `[Username]:[Password]`.
2. Under 'Proxy Authorization' click on 'Edit in settings.json'.
3. Add to settings.json:

* `"http.proxyAuthorization": "Basic [Encoded credentials]"`.

**Access token authorization**

1. Under 'Proxy Authorization' click on 'Edit in settings.json'.
2. Add to settings.json:

* `"http.proxyAuthorization": "Bearer [Access token]"`.

**Example**

* `Username: foo`
* `Password: bar`

settings.json:

```
{
    "http.proxyAuthorization": "Basic Zm9vOmJhcg=="
}
```

### Troubleshooting

Change the log level to 'debug', 'info', 'warn', or 'err' in the [Extension Settings](extension-settings.md).
