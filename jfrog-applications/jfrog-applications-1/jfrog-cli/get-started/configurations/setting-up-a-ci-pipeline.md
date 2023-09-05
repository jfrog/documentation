# Setting up a CI Pipeline

The **ci-setup** command allows setting up a basic CI pipeline with the JFrog Platform, while automatically configuring the JFrog Platform to serve the pipeline. It is an interactive command, which prompts you with a series for questions, such as your source control details, your build tool, build command and your CI provider. The command then uses this information to do following:

* Create the repositories in JFrog Artifactory, to be used by the pipeline to resolve dependencies.
* Configure JFrog Xray to scan the build.
* Generate a basic CI pipeline, which builds and scans your code.

You can use the generated CI pipeline as a working starting point and then expand it as needed.

The command currently supports the following package managers:

* Maven
* Gradle
* npm.

and the following CI providers:

* JFrog Pipelines
* Jenkins
* GitHub Actions.

Usage:

```
jf ci-setup
```
