# Connect VS Code to the JFrog Platform

Once the JFrog Extension is installed in VS Code, click on the JFrog tab:

![jfrogTab](../../.gitbook/assets/vscode/jfrogTab.png)

This will open the Sign in page:

![SighInPage](../../.gitbook/assets/vscode/sighInPage.png)

Fill in your connection details and click on the `Sign In` button to start using the extension

**Note**: If you would like to use custom URLs for Artifactory or Xray, click on `Advanced`.

You can also choose other option to authenticate with your JFrog Platform instance:

1. [SSO](connect-vs-code-to-the-jfrog-platform.md#connect-using-sso)
2. [JFrog CLI's Connection Details](connect-vs-code-to-the-jfrog-platform.md#connect-using-jfrog-cli-connection-details)
3. [Using Environment Variables](connect-vs-code-to-the-jfrog-platform.md#connect-using-environment-variables).

## Connect Using SSO

To sign in using SSO, follow these steps:

1. On the sign-in page, click the `Continue with SSO` button:

![SighInSsoButton](../../.gitbook/assets/vscode/ssoButton.png)

2. After entering your JFrog platform URL, click on `Sign in With SSO`.
3. It will take a few seconds for the browser to redirect you to the SSO sign in page.
4. You should now be signed in in at vscode.

## Connect Using JFrog CLI Connection Details

If JFrog CLI is installed on your machine and is configured with your JFrog Platform connection details, then you should see the message popup in the Sigh in page:

![LoginPageJfrogCli](../../.gitbook/assets/vscode/sighInPageJFrogCli.png)

## Connect Using Environment Variables

You may set the connection details using the following environment variables. VS Code will read them after it is launched.

* `JFROG_IDE_URL` - JFrog URL
* `JFROG_IDE_USERNAME` - JFrog username
* `JFROG_IDE_PASSWORD` - JFrog password
* `JFROG_IDE_ACCESS_TOKEN` - JFrog access token
* `JFROG_IDE_STORE_CONNECTION` - Set the value of this environment variable to **true**, if you'd like VS Code to store the connection details after reading them from the environment variables.

Once the above environment variables are configured, you can expect to see a message popup in the Sigh in page:

![LoginPageEnvVar](../../.gitbook/assets/vscode/sighInPageEnvVar.png)

**Note**: For security reasons, it is recommended to unset the environment variables after launching VS Code.
