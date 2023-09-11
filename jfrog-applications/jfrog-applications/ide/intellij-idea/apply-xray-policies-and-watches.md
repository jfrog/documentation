# Apply Xray Policies and Watches

You can configure the JFrog Plugin to use the security policies you create in Xray. Policies enable you to create a set of rules, in which each rule defines security criteria, with a corresponding set of automatic actions according to your needs. Policies are enforced when applying them to Watches.

If you'd like to use a JFrog Project that is associated with the policy, follow these steps:

1. Create a [JFrog Project](https://jfrog.com/help/r/jfrog-platform-administration-documentation/introduction-to-projects), or obtain the relevant JFrog Project key.
2. Create a [Policy](https://jfrog.com/help/r/jfrog-security-documentation/creating-xray-policies-and-rules) on JFrog Xray.
3. Create a [Watch](https://jfrog.com/help/r/jfrog-security-documentation/configuring-xray-watches) on JFrog Xray and assign your Policy and Project as resources to it.
4. Configure your Project key in the plugin settings: under **Settings (Preferences)** | **Other Settings**, click **JFrog Global Configuration** and go to the **Settings** tab.

If however your policies are referenced through Xray Watches, follow these steps instead:

1. Create one or more [Watches](https://jfrog.com/help/r/jfrog-security-documentation/configuring-xray-watches) on JFrog Xray.
2. Configure your Watches in the plugin settings: under **Settings (Preferences)** | **Other Settings**, click **JFrog Global Configuration** and go to the **Settings** tab.
