# CLI for IDE Setup

![](../../../.gitbook/assets/cli-ide-header.png)


JFrog CLI provides powerful commands to configure IDEs to use JFrog Artifactory as the repository source for extensions, plugins, and packages, streamlining the setup process for development teams by allowing them to easily configure their IDEs to use private or curated repositories hosted in Artifactory.

## Overview

The CLI IDE setup commands automate the configuration of popular IDEs to point to JFrog Artifactory repositories instead of public repositories, which enables organizations to:

- **Centralize IDE Extensions Management**: Host and manage IDE extensions and plugins in Artifactory
- **Ensure Security and Compliance**: Use curated, scanned, and approved extensions from private repositories  
- **Enable Offline Development**: Work with locally cached extensions without internet dependency
- **Standardize Development Environment**: Ensure all team members use the same approved set of extensions

## Supported IDEs

| IDE | Configuration Type | Commands | Platform Support |
|-----|-------------------|----------|------------------|
| **Visual Studio Code** | Extensions Gallery | `jf vscode config`<br/>`jf vscode install` | Windows, macOS, Linux |
| **JetBrains IDEs** | Plugin Repository | `jf jetbrains config` | Windows, macOS, Linux |

### JetBrains IDEs Support

The JetBrains configuration command supports all JetBrains IDEs:
- IntelliJ IDEA
- PyCharm  
- WebStorm
- PhpStorm
- RubyMine
- CLion
- DataGrip
- GoLand
- Rider
- Android Studio
- AppCode
- RustRover
- Aqua

## How It Works

### VSCode Configuration
The VSCode commands modify the IDE's configuration to point to JFrog Artifactory for extension downloads:

1. **Configuration**: Updates `product.json` to change the extensions gallery URL
2. **Installation**: Downloads and installs extensions directly from Artifactory repositories

### JetBrains Configuration  
The JetBrains command configures all detected JetBrains IDEs:

1. **Auto-Detection**: Scans for installed JetBrains IDEs on the system
2. **Configuration**: Updates `idea.properties` files to use Artifactory plugin repositories
3. **Backup**: Creates automatic backups before making changes

## Prerequisites

### Repository Setup
Before using the CLI commands, ensure you have:

1. **JFrog Artifactory** with appropriate repository types:
   - **VSCode**: Virtual or Remote repository of type `Generic` configured for VSCode extensions
   - **JetBrains**: Virtual or Remote repository of type `Generic` configured for JetBrains plugins

2. **Repository Configuration** in Artifactory:
   - VSCode extensions repository with appropriate layout
   - JetBrains plugins repository with JetBrains plugin repository layout

3. **Access Permissions**: Appropriate permissions to read from the configured repositories

### System Requirements

#### Platform-Specific Requirements

**macOS & Linux:**
- `sudo` access may be required for system-installed IDEs
- Commands will prompt for elevated privileges when needed

**Windows:**  
- "Run as Administrator" may be required for system-installed IDEs
- PowerShell execution policy must allow script execution

## Quick Start

### VSCode Setup
```bash
# Configure VSCode to use Artifactory extensions repository
jf vscode config --repo=vscode-extensions --artifactory-url=https://mycompany.jfrog.io/

# Install an extension from the configured repository
jf vscode install --publisher=microsoft --extension-name=vscode-python --repo=vscode-extensions
```

### JetBrains Setup  
```bash
# Configure all JetBrains IDEs to use Artifactory plugins repository
jf jetbrains config --repo=jetbrains-plugins --artifactory-url=https://mycompany.jfrog.io/
```

## Advanced Configuration

### Using Server Configurations
If you have JFrog CLI server configurations set up, you can omit the `--artifactory-url` parameter:

```bash
# Configure with default server
jf vscode config --repo=vscode-extensions

# Configure with specific server
jf c use my-artifactory-server
jf jetbrains config --repo=jetbrains-plugins
```

### Repository URL Formats

The CLI automatically constructs the appropriate repository URLs:

**VSCode Extensions:**
```
https://[artifactory-url]/artifactory/api/vscodeextensions/[repo-key]/_apis/public/gallery
```

**JetBrains Plugins:**
```
https://[artifactory-url]/artifactory/api/jetbrainsplugins/[repo-key]
```

## Backup and Recovery

### Automatic Backups
Both commands create automatic backups before making changes:

- **VSCode**: Backups stored in `~/.jfrog/backup/ide/vscode/`
- **JetBrains**: Backups stored in `~/.jfrog/backup/ide/jetbrains/`

### Manual Recovery
If you need to restore original settings:

1. Locate the backup files in the backup directories
2. Copy the backup files back to their original locations
3. Restart the affected IDE

## Troubleshooting

### Common Issues

**Permission Denied:**
- Run the command with elevated privileges (`sudo` on macOS/Linux, "Run as Administrator" on Windows)

**IDE Not Detected:**
- Ensure the IDE is installed in standard locations
- Check if the IDE is running and close it before configuration

**Repository Not Accessible:**
- Verify repository exists and is accessible
- Check network connectivity to Artifactory
- Ensure proper authentication is configured

### Manual Configuration
If automatic configuration fails, each command provides manual setup instructions in the error output.

## Security Considerations

### File Modifications
The CLI commands modify system-level IDE configuration files. Always:
- Review backup files before making changes
- Test changes in a development environment first
- Ensure you have proper backups of IDE configurations

### Repository Access
- Use access tokens instead of username/password when possible
- Ensure repository permissions are properly configured
- Consider using repository scanning and policies for enhanced security

## Examples and Use Cases

For detailed examples and specific use cases, see:
- [VSCode CLI Configuration](vscode-cli.md)
- [JetBrains CLI Configuration](jetbrains-cli.md) 