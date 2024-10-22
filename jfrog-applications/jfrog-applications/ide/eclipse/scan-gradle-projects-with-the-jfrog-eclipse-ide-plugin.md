# Scan Gradle Projects with the JFrog Eclipse IDE Plugin

Behind the scenes, the JFrog plugin executes a Gradle script, which creates the dependencies tree of the project. The plugin reads the Gradle configuration defined in Eclipse. This configuration is added to Eclipse by the [Buildship](https://marketplace.eclipse.org/content/buildship-gradle-integration) plugin You can access this configuration by going in **Preferences** | **Gradle** | **Gradle distribution**

{% hint style="info" %}


If the Gradle configuration is not set, then Gradle Wrapper will be used. If the project does not include the Gradle Wrapper configuration, Gradle will be automatically downloaded.
{% endhint %}

<div align="left">

<figure><img src="../../../.gitbook/assets/0 (1).jpeg" alt="" width="375"><figcaption></figcaption></figure>

</div>
