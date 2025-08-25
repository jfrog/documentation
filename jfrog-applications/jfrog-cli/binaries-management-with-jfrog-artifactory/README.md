# Binaries Management with JFrog Artifactory

JFrog CLI is a compact and smart client that provides a simple interface to automate access to JFrog products. It simplifies your automation scripts, making them more readable and easier to maintain. JFrog CLI works with JFrog Artifactory to make your scripts more efficient and reliable in several ways.

**Advanced upload and download capabilities**

With JFrog CLI, you can upload and download artifacts concurrently using a configurable number of threads to help your automated builds run faster. For large artifacts, you can define a number of chunks to split files for parallel download.

JFrog CLI optimizes both upload and download operations by skipping artifacts that already exist in their target location. Before uploading an artifact, JFrog CLI queries Artifactory with the artifact's checksum. If the artifact already exists in Artifactory's storage, the CLI skips sending the file, and Artifactory only updates its database if necessary. Similarly, when you download an artifact, it will be skipped if it already exists in the target download path. Checksum optimization also allows you to pause long upload and download operations and resume them later from where they left off.

JFrog CLI supports uploading files to Artifactory using wildcard patterns, regular expressions, and ANT patterns, giving you an easy way to collect all the files you wish to upload. You can also download files using wildcard patterns.

**Support for Popular Package Managers and Build Tools**

JFrog CLI offers comprehensive support for popular package managers and build tools. It seamlessly integrates with package managers like npm, Maven, NuGet, Docker, and more, allowing you to easily manage and publish packages.

**Support for Build-Info**

Build-Info is a comprehensive metadata Software Bill of Materials (SBOM) that captures detailed information about the components used in a build. It serves as a vital source of information, containing version history, artifacts, project modules, dependencies, and other crucial data collected during the build process. By storing this metadata in Artifactory, you gain traceability and analysis capabilities to improve the quality and security of your builds. The Build-Info encompasses project module details, artifacts, dependencies, environment variables, and more. It is collected and output in a JSON format, facilitating easy access to information about the build and its components. JFrog CLI can create and store the Build-Info in Artifactory.

For more information, see [here](../).
