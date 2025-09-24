# JFrog BO*Bs CLIck ;-p

![](../../.gitbook/assets/jfrog-cli-header.png)

JFrog CLI is a compact and smart client that provides a simple interface to automate access to JFrog products, simplifying your automation scripts and making them more readable and easier to maintain.

JFrog CLI works with JFrog Artifactory, Xray, and Distribution (through their respective REST APIs), making your scripts more efficient and reliable in several ways.



**Advanced upload and download capabilities**

JFrog CLI allows you to upload and download artifacts concurrently by a configurable number of threads that help your automated builds run faster. For big artifacts, you can define a number of chunks to split files for parallel download.

To optimize both uploads and downloads, JFrog CLI avoids transferring artifacts that already exist in the target location. Before uploading, the CLI checks the artifact's checksum with Artifactory. If the artifact is already present in Artifactory’s storage, the CLI skips the upload, and Artifactory may just update its database to reflect the new upload. Similarly, when downloading an artifact, if it already exists in the specified download path, it will be skipped. This checksum optimization also allows you to pause long upload and download operations and resume them later from where you left off.

JFrog CLI simplifies file uploads by supporting wildcard patterns, regular expressions, and ANT patterns, allowing you to easily select all the files you want to upload. You can also use wildcard patterns for downloading files.



**Support for popular package managers and build tools**

JFrog CLI offers comprehensive support for popular package managers and build tools. It seamlessly integrates with package managers like npm, Maven, NuGet, Docker, and more, allowing you to easily manage and publish packages.

**Source code and binaries scanning**

JFrog CLI empowers you with robust scanning capabilities to ensure the security and compliance of your source code and software artifacts, including containers. It integrates with JFrog Xray, enabling you to scan and analyze your projects and packages, including containers, for vulnerabilities, license compliance, and quality issues. With JFrog CLI, you can proactively identify and mitigate potential risks, ensuring the integrity and safety of your software supply chain.

**Support for Build-Info**

Build-Info is a comprehensive metadata Software Bill of Materials (SBOM) that captures detailed information about the components used in a build. It serves as a vital source of information, containing version history, artifacts, project modules, dependencies, and other crucial data collected during the build process. By storing this metadata in Artifactory, developers gain traceability and analysis capabilities to improve the quality and security of their builds. The Build-Info encompasses project module details, artifacts, dependencies, environment variables, and more. It is collected and outputted in a JSON format, facilitating easy access to information about the build and its components. JFrog CLI can create a Build-Info and store the Build-Info in Artifactory.

#### System Requirements

JFrog CLI runs on any modern OS that fully supports the [Go programming language](https://golang.org/).

#### Contributing to JFrog CLI Documentation

Your input is valuable in making the JFrog CLI documentation better. You can help enhance and improve it by recommending changes and additions. To contribute, follow these steps:

Go to the documentation project on GitHub:

[GitHub - jfrog/documentation](https://github.com/jfrog/documentation)  and create a pull request with your proposed changes and additions.

Your contributions will be reviewed, and if accepted, they will be merged into the documentation to benefit the entire JFrog CLI community.

\
