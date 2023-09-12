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
