# Proxy Support

JFrog CLI supports using an HTTP/S proxy. All you need to do is set HTTP\_PROXY or HTTPS\_PROXY environment variable with the proxy URL.

HTTP\_PROXY, HTTPS\_PROXY and NO\_PROXY are the industry standards for proxy usages.

| Variable Name | Description                                                                                                                                                                                                                                  |
| ------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| HTTP\_PROXY   | Determines a URL to an HTTP proxy.                                                                                                                                                                                                           |
| HTTPS\_PROXY  | Determines a URL to an HTTPS proxy.                                                                                                                                                                                                          |
| NO\_PROXY     | Use this variable to bypass the proxy to IP addresses, subnets or domains. This may contain a comma-separated list of hostnames or IPs without protocols and ports. A typical usage may be to set this variable to Artifactory’s IP address. |
