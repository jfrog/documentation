# Xray Policies and Watches

You can configure the JFrog VS-Code extension to use the security policies you create in Xray. Policies enable you to create a set of rules, in which each rule defines security criteria, with a corresponding set of automatic actions according to your needs. Policies are enforced when applying them to Watches.

If you'd like to use a JFrog Project that is associated with the policy, follow these steps:

1. Create a [JFrog Project](https://www.jfrog.com/confluence/display/JFROG/Projects), or obtain the relevant JFrog Project key.
2. Create a [Policy](https://www.jfrog.com/confluence/display/JFROG/Creating+Xray+Policies+and+Rules) on JFrog Xray.
3. Create a [Watch](https://www.jfrog.com/confluence/display/JFROG/Configuring+Xray+Watches) on JFrog Xray and assign your Policy and Project as resources to it.
4. Configure your Project key in the [Extension Settings](extension-settings.md).

If however your policies are referenced through an Xray Watch or Watches, follow these steps instead:

1. Create one or more [Watches](https://www.jfrog.com/confluence/display/JFROG/Configuring+Xray+Watches) on JFrog Xray.
2. Configure your Watches in the [Extension Settings](extension-settings.md).
