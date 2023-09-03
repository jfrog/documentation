# Behind the Scenes

#### Go Projects

Behind the scenes, the JFrog VS Code Extension scans all the project dependencies, both direct and indirect (transitive), even if they are not declared in the project's go.mod. It builds the Go dependencies tree by running `go mod graph` and intersecting the results with `go list -f '{{with .Module}}{{.Path}} {{.Version}}{{end}}' all` command. Therefore, please make sure to have Go CLI in your system PATH.

#### Maven Projects

The JFrog VS Code Extension builds the Maven dependencies tree by running `mvn dependency:tree`. View licenses and top issue severities directly from the pom.xml.

Important notes:

1. To have your project dependencies scanned by JFrog Xray, make sure Maven is installed, and that the mvn command is in your system PATH.
2. For projects which include the [Maven Dependency Plugin](https://maven.apache.org/plugins/maven-dependency-plugin/examples/resolving-conflicts-using-the-dependency-tree.html) as a build plugin, with include or exclude configurations, the scanning functionality is disabled. For example:

```
      <plugins>
        <plugin>
          <groupId>org.apache.maven.plugins</groupId>
          <artifactId>maven-dependency-plugin</artifactId>
          <configuration>
            <includes>org.apache.*</includes>
          </configuration>
        </plugin>
      </plugins>
```

#### Npm Projects

Behind the scenes, the extension builds the npm dependencies tree by running `npm list`. View licenses and top issue severities directly from the package.json.

Important: To have your project dependencies scanned by JFrog Xray, make sure the npm CLI is installed on your local machine and that it is in your system PATH. In addition, the project dependencies must be installed using `npm install`.

#### Yarn v1 Projects

Behind the scenes, the extension builds the Yarn dependencies tree by running `yarn list`. View licenses and top issue severities directly from the yarn.lock.

Important:

* To have your project dependencies scanned by JFrog Xray, make sure the Yarn CLI is installed on your local machine and that it is in your system PATH.
* Yarn v2 is not yet supported.

#### Pypi Projects

Behind the scenes, the extension builds the Pypi dependencies tree by running `pipdeptree` on your Python virtual environment. It also uses the Python interpreter path configured by the [Python extension](https://marketplace.visualstudio.com/items?itemName=ms-python.python). View licenses and top issue severities directly from your requirements.txt files. The scan your Pypi dependencies, make sure the following requirements are met:

1. The [Python extension for VS Code](https://code.visualstudio.com/docs/python/python-tutorial#\_install-visual-studio-code-and-the-python-extension) is installed.
2. Depending on your project, Please make sure Python 2 or 3 are included in your system PATH.
3. Create and activate a virtual env as instructed in [VS-Code documentation](https://code.visualstudio.com/docs/python/environments#\_global-virtual-and-conda-environments). Make sure that Virtualenv Python interpreter is selected as instructed [here](https://code.visualstudio.com/docs/python/environments#\_select-and-activate-an-environment).
4. Open a new terminal and activate your Virtualenv:
   *   On macOS and Linux:

       ```
       source <venv-dir>/bin/activate

       # For example:
       source .env/bin/activate
       ```
   *   On Windows:

       ```
       .\<venv-dir>\Scripts\activate

       # For example:
       .\env\Scripts\activate
       ```
5. In the same terminal, install your python project and dependencies according to your project specifications.

#### .NET Projects

For .NET projects which use NuGet packages as dependencies, the extension displays the NuGet dependencies tree, together with the information for each dependency. Behind the scenes, the extension builds the NuGet dependencies tree using the [NuGet deps tree](https://github.com/jfrog/nuget-deps-tree) npm package.

Important:

* Does your project define its NuGet dependencies using a _packages.config_ file? If so, then please make sure the `nuget` CLI is installed on your local machine and that it is in your system PATH. The extension uses the `nuget` CLI to find the location of the NuGet packages on the local file-system.
* The project must be restored using `nuget restore` or `dotnet restore` prior to scanning. After this action, you should click on the Refresh  button, for the tree view to be refreshed and updated.
