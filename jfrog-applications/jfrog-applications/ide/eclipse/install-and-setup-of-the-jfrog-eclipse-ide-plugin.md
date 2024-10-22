# Install and Setup of the JFrog Eclipse IDE Plugin

To install and work with the plugin:

1. Install the JFrog plugin
2. If JFrog Xray is behind an HTTP proxy, configure the proxy settings as described [here](https://help.eclipse.org/kepler/index.jsp?topic=%2Forg.eclipse.platform.doc.user%2Freference%2Fref-net-preferences.htm). This is supported since version 1.1.0 of the JFrog Eclipse Plugin.
3. Configure the plugin to connect to JFrog Xray
4. Scan and view the results
5. Filter Xray Scanned Results

### Download the JFrog Eclipse IDE Plugin

| **Version** | **Download link**                                                                                                              | **Compatibility**   |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------ | ------------------- |
| 1.1.2       | [Download](https://github.com/jfrog/jfrog-eclipse-plugin/releases/download/1.1.2/com.jfrog.ide.eclipse.releng.update-site.zip) | Eclipse 4.13 - 4.20 |
| 1.1.1       | [Download](https://github.com/jfrog/jfrog-eclipse-plugin/releases/download/1.1.1/com.jfrog.ide.eclipse.releng.update-site.zip) | Eclipse 4.10 - 4.19 |

**Install the JFrog Eclipse IDE Plugin**

1. Download the plugin zip.
2. Go to **Help | Install New Software**,click **Add** and then click **Archive**.
3. Choose the plugin zip file you downloaded and click **Add**.
4. Click **Next**.

### Configure the JFrog Eclipse IDE Plugin

This section describes how to configure the JFrog Eclipse IDE Plugin and reviews connecting to JFrog Xray and Scanning Gradel projects with the Plugin. It reviews the following:

* Connect to JFrog Xray \[315]
* Scan Gradle Projects with the JFrog Eclipse IDE Plugin \[316]

**Connect to JFrog Xray**

Once the plugin is successfully installed, connect the plugin to your instance of JFrog Xray.

1. Go to **Eclipse (Preferences)**, click **JFrog Xray**.
2. Set your JFrog Xray URL and login credentials.
3. Test your connection to Xray using the **Test Connection** button.
