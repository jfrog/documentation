# VSCode CLI Configuration

The JFrog CLI VSCode commands provide seamless integration between Visual Studio Code and JFrog Artifactory, allowing you to configure VSCode to use Artifactory repositories for extension management and install extensions directly from your private repositories.

## Commands Overview

The VSCode CLI functionality includes two main commands:

| Command | Purpose | Description |
|---------|---------|-------------|
| `jf vscode config` | Configure VSCode | Updates VSCode to use Artifactory extensions repository |
| `jf vscode install` | Install Extensions | Downloads and installs extensions from Artifactory |

## jf vscode config

Configure Visual Studio Code to use JFrog Artifactory as the extensions repository source.

### Syntax

```bash
jf vscode config --repo=<REPO_KEY> [command options]
```

### Command Arguments

| Argument | Description |
|----------|-------------|
| `--repo` | **[Mandatory]** VSCode extensions repository key in Artifactory |

### Command Options

| Option | Description | Default |
|--------|-------------|---------|
| `--artifactory-url` | **[Optional]** Artifactory server URL. If not provided, uses default server configuration | Uses configured default server |

### Examples

#### Example 1: Basic Configuration
Configure VSCode to use an Artifactory extensions repository:

```bash
jf vscode config --repo=vscode-extensions --artifactory-url=https://mycompany.jfrog.io/
```

#### Example 2: Using Default Server Configuration
If you have a default JFrog CLI server configuration:

```bash
jf vscode config --repo=vscode-extensions
```

#### Example 3: Platform-Specific Usage

**macOS (with system-installed VSCode):**
```bash
sudo jf vscode config --repo=vscode-extensions --artifactory-url=https://mycompany.jfrog.io/
```

**Windows (Run as Administrator):**
```bash
jf vscode config --repo=vscode-extensions --artifactory-url=https://mycompany.jfrog.io/
```

**Linux:**
```bash
sudo jf vscode config --repo=vscode-extensions --artifactory-url=https://mycompany.jfrog.io/
```

### How It Works

1. **Detection**: Automatically detects VSCode installation location
2. **Backup**: Creates backup of original `product.json` file
3. **Modification**: Updates `extensionsGallery.serviceUrl` in `product.json`
4. **Verification**: Confirms the changes were applied successfully

### Configuration Details

The command modifies the `product.json` file to change the extensions gallery URL:

**Before:**
```json
{
  "extensionsGallery": {
    "serviceUrl": "https://marketplace.visualstudio.com/_apis/public/gallery"
  }
}
```

**After:**
```json
{
  "extensionsGallery": {
    "serviceUrl": "https://mycompany.jfrog.io/artifactory/api/vscodeextensions/vscode-extensions/_apis/public/gallery"
  }
}
```

### File Locations

The command modifies the following files based on your platform:

| Platform | File Location |
|----------|---------------|
| **macOS** | `/Applications/Visual Studio Code.app/Contents/Resources/app/product.json` |
| **Windows** | `%LOCALAPPDATA%\Programs\Microsoft VS Code\resources\app\product.json` |
| **Linux** | `/usr/share/code/resources/app/product.json` |

## jf vscode install

Install VSCode extensions directly from JFrog Artifactory repositories.

### Syntax

```bash
jf vscode install --publisher=<PUBLISHER> --extension-name=<EXTENSION_NAME> --repo=<REPO_KEY> [command options]
```

### Command Arguments

| Argument | Description |
|----------|-------------|
| `--publisher` | **[Mandatory]** Extension publisher name |
| `--extension-name` | **[Mandatory]** Extension name |
| `--repo` | **[Mandatory]** VSCode extensions repository key in Artifactory |

### Command Options

| Option | Description | Default |
|--------|-------------|---------|
| `--version` | **[Optional]** Specific extension version to install | Latest version |
| `--artifactory-url` | **[Optional]** Artifactory server URL. If not provided, uses default server configuration | Uses configured default server |

### Examples

#### Example 1: Install Latest Extension
Install the latest version of the Python extension:

```bash
jf vscode install --publisher=ms-python --extension-name=python --repo=vscode-extensions
```

#### Example 2: Install Specific Version
Install a specific version of an extension:

```bash
jf vscode install --publisher=ms-python --extension-name=python --version=2023.12.0 --repo=vscode-extensions
```

#### Example 3: Install with Explicit Server URL
```bash
jf vscode install --publisher=microsoft --extension-name=vscode-typescript --repo=vscode-extensions --artifactory-url=https://mycompany.jfrog.io/
```

### How It Works

1. **Validation**: Verifies extension exists in the Artifactory repository
2. **Download**: Downloads extension package (.vsix) from Artifactory
3. **Installation**: Uses VSCode CLI to install the extension
4. **Cleanup**: Removes temporary files after installation

### Extension URL Format

The command constructs download URLs in the following format:

**Latest Version:**
```
https://[artifactory-url]/artifactory/api/vscodeextensions/[repo-key]/_apis/public/gallery/publishers/[publisher]/extensions/[extension-name]/latest/vspackage
```

**Specific Version:**
```
https://[artifactory-url]/artifactory/api/vscodeextensions/[repo-key]/_apis/public/gallery/publishers/[publisher]/extensions/[extension-name]/[version]/vspackage
```

## Prerequisites

### VSCode CLI Requirements
The `jf vscode install` command requires the VSCode CLI (`code` command) to be available in your system PATH.

**Installation:**
1. Open VSCode
2. Open Command Palette (`Ctrl+Shift+P` or `Cmd+Shift+P`)
3. Type "Shell Command: Install 'code' command in PATH"
4. Select and execute the command

### Repository Configuration
Ensure your Artifactory repository is properly configured:

1. **Repository Type**: Generic repository
2. **Layout**: Default or custom layout supporting VSCode extensions
3. **Permissions**: Read access for users/groups that will install extensions

## Platform-Specific Requirements

### macOS
- **System Installation**: Requires `sudo` for modifications
- **User Installation**: No special permissions required if VSCode is installed in `~/Applications/`
- **File Permissions**: Ensure the user has read/write access to VSCode installation directory

### Windows  
- **System Installation**: Requires "Run as Administrator"
- **User Installation**: May require elevated permissions depending on installation location
- **PowerShell**: Ensure PowerShell execution policy allows script execution

### Linux
- **System Installation**: Requires `sudo` for modifications in `/usr/share/code/`
- **Snap Installation**: May require special handling for snap-packaged VSCode
- **File Permissions**: Ensure proper permissions on configuration files

## Troubleshooting

### Common Issues and Solutions

#### Permission Denied Errors

**Symptoms:**
```
Error: insufficient permissions to modify VSCode configuration
```

**Solutions:**
- **macOS/Linux**: Run with `sudo`
- **Windows**: Run Command Prompt/PowerShell as Administrator
- **Alternative**: Install VSCode in user directory instead of system directory

#### VSCode Not Detected

**Symptoms:**
```
Error: VSCode installation not found in standard locations
```

**Solutions:**
1. Verify VSCode is installed
2. Check if VSCode is installed in non-standard location
3. Use manual configuration instructions provided in error output

#### Extension Installation Fails

**Symptoms:**
```
Error: failed to install extension via VSCode CLI
```

**Solutions:**
1. Ensure VSCode CLI (`code` command) is installed and accessible
2. Verify extension exists in the Artifactory repository
3. Check network connectivity to Artifactory
4. Ensure proper authentication is configured

#### Repository Access Issues

**Symptoms:**
```
Error: repository validation failed
```

**Solutions:**
1. Verify repository exists and is accessible
2. Check repository permissions
3. Ensure proper authentication (access token, username/password)
4. Verify network connectivity to Artifactory

### Manual Configuration

If automatic configuration fails, you can manually configure VSCode:

1. **Locate product.json file** (see file locations table above)
2. **Create backup** of the original file
3. **Edit the file** to change `extensionsGallery.serviceUrl`
4. **Restart VSCode** to apply changes

**Manual Configuration Example:**
```json
{
  "extensionsGallery": {
    "serviceUrl": "https://your-artifactory.com/artifactory/api/vscodeextensions/your-repo/_apis/public/gallery",
    "cacheUrl": "https://vscode.blob.core.windows.net/gallery/index",
    "itemUrl": "https://marketplace.visualstudio.com/items"
  }
}
```

## Advanced Usage

### Batch Extension Installation
Install multiple extensions using a script:

```bash
#!/bin/bash
REPO="vscode-extensions"
ARTIFACTORY_URL="https://mycompany.jfrog.io/"

# Array of extensions to install
extensions=(
    "ms-python:python"
    "ms-vscode:vscode-typescript"
    "ms-dotnettools:csharp"
    "golang:go"
)

for ext in "${extensions[@]}"; do
    IFS=':' read -r publisher name <<< "$ext"
    echo "Installing $publisher.$name..."
    jf vscode install --publisher="$publisher" --extension-name="$name" --repo="$REPO" --artifactory-url="$ARTIFACTORY_URL"
done
```

### CI/CD Integration
Configure VSCode in CI/CD environments:

```yaml
# GitHub Actions example
- name: Configure VSCode for Artifactory
  run: |
    sudo jf vscode config --repo=vscode-extensions --artifactory-url=${{ secrets.ARTIFACTORY_URL }}
  
- name: Install Required Extensions
  run: |
    jf vscode install --publisher=ms-python --extension-name=python --repo=vscode-extensions
```

### Validation Scripts
Verify VSCode configuration:

```bash
#!/bin/bash
# Check if VSCode is configured to use Artifactory
PRODUCT_JSON="/Applications/Visual Studio Code.app/Contents/Resources/app/product.json"
EXPECTED_URL="your-artifactory.com"

if grep -q "$EXPECTED_URL" "$PRODUCT_JSON"; then
    echo "✓ VSCode is configured to use Artifactory"
else
    echo "✗ VSCode is not configured for Artifactory"
    exit 1
fi
```

## Best Practices

### Security
1. **Use Access Tokens**: Prefer access tokens over username/password
2. **Repository Permissions**: Configure least-privilege access to repositories
3. **Network Security**: Use HTTPS for all Artifactory communications
4. **Backup Configurations**: Always maintain backups of IDE configurations

### Team Workflow
1. **Standardization**: Use the same repository keys across the team
2. **Documentation**: Document required extensions and versions
3. **Automation**: Include IDE configuration in onboarding scripts
4. **Testing**: Test configuration changes in development environment first

### Repository Management
1. **Curation**: Only include approved and scanned extensions
2. **Versioning**: Maintain specific versions for production environments
3. **Cleanup**: Regularly clean up old extension versions
4. **Monitoring**: Monitor repository usage and access patterns 