# JetBrains CLI Configuration
	
The JFrog CLI JetBrains command provides automated configuration for all JetBrains IDEs to use JFrog Artifactory as the plugin repository source, which enables organizations to centrally manage, curate, and distribute JetBrains plugins through their private Artifactory repositories.

## Command Overview

| Command | Purpose | Description |
|---------|---------|-------------|
| `jf jetbrains config` | Configure JetBrains IDEs | Updates all detected JetBrains IDEs to use Artifactory plugin repository |

## jf jetbrains config

Configure all installed JetBrains IDEs to use JFrog Artifactory as the plugin repository source.

### Syntax

```bash
jf jetbrains config --repo=<REPO_KEY> [command options]
```

### Command Arguments

| Argument | Description |
|----------|-------------|
| `--repo` | **[Mandatory]** JetBrains plugins repository key in Artifactory |

### Command Options

| Option | Description | Default |
|--------|-------------|---------|
| `--artifactory-url` | **[Optional]** Artifactory server URL. If not provided, the default server configuration is used. | Uses configured default server. |

### Examples

#### Example 1: Basic Configuration
Configure all JetBrains IDEs to use an Artifactory plugins repository:

```bash
jf jetbrains config --repo=jetbrains-plugins --artifactory-url=https://mycompany.jfrog.io/
```

#### Example 2: Using Default Server Configuration
If you have a default JFrog CLI server configuration:

```bash
jf jetbrains config --repo=jetbrains-plugins
```

#### Example 3: Platform-Specific Usage

**macOS:**
```bash
jf jetbrains config --repo=jetbrains-plugins --artifactory-url=https://mycompany.jfrog.io/
```

**Windows:**
```bash
jf jetbrains config --repo=jetbrains-plugins --artifactory-url=https://mycompany.jfrog.io/
```

**Linux:**
```bash
jf jetbrains config --repo=jetbrains-plugins --artifactory-url=https://mycompany.jfrog.io/
```

### Supported JetBrains IDEs

The command automatically detects and configures all installed JetBrains IDEs:

| IDE | Product Code | Description |
|-----|--------------|-------------|
| **IntelliJ IDEA** | `IntelliJIdea` | Java/Kotlin/Scala development |
| **PyCharm** | `PyCharm` | Python development |
| **WebStorm** | `WebStorm` | JavaScript/TypeScript development |
| **PhpStorm** | `PhpStorm` | PHP development |
| **RubyMine** | `RubyMine` | Ruby development |
| **CLion** | `CLion` | C/C++ development |
| **DataGrip** | `DataGrip` | Database tools |
| **GoLand** | `GoLand` | Go development |
| **Rider** | `Rider` | .NET development |
| **Android Studio** | `AndroidStudio` | Android development |
| **AppCode** | `AppCode` | iOS/macOS development |
| **RustRover** | `RustRover` | Rust development |
| **Aqua** | `Aqua` | Test automation |

### How It Works

1. **Auto-Detection**: Scans system for installed JetBrains IDEs
2. **Configuration Discovery**: Locates configuration directories for each IDE
3. **Backup Creation**: Creates timestamped backups of `idea.properties` files
4. **Configuration Update**: Adds or updates `idea.plugins.host` property
5. **Verification**: Confirms changes were applied successfully

### Configuration Details

The command modifies the `idea.properties` file for each detected IDE:

**Before Configuration:**
```properties
# Empty or existing properties
```

**After Configuration:**
```properties
# JFrog Artifactory plugins repository
idea.plugins.host=https://mycompany.jfrog.io/artifactory/api/jetbrainsplugins/jetbrains-plugins
```

### Repository URL Format

The command constructs the plugin repository URL in the following format:

```
https://[artifactory-url]/artifactory/api/jetbrainsplugins/[repo-key]
```

### File Locations

The command modifies `idea.properties` files in the following locations:

#### macOS
```
~/Library/Application Support/JetBrains/[IDE][VERSION]/idea.properties
```

**Examples:**
- `~/Library/Application Support/JetBrains/IntelliJIdea2023.3/idea.properties`
- `~/Library/Application Support/JetBrains/PyCharm2023.3/idea.properties`
- `~/Library/Application Support/JetBrains/WebStorm2023.3/idea.properties`

#### Windows
```
%APPDATA%\JetBrains\[IDE][VERSION]\idea.properties
```

**Examples:**
- `%APPDATA%\JetBrains\IntelliJIdea2023.3\idea.properties`
- `%APPDATA%\JetBrains\PyCharm2023.3\idea.properties`
- `%APPDATA%\JetBrains\WebStorm2023.3\idea.properties`

#### Linux
```
~/.config/JetBrains/[IDE][VERSION]/idea.properties
```

**Examples:**
- `~/.config/JetBrains/IntelliJIdea2023.3/idea.properties`
- `~/.config/JetBrains/PyCharm2023.3/idea.properties`
- `~/.config/JetBrains/WebStorm2023.3/idea.properties`

**Legacy Location (older installations):**
```
~/.JetBrains/[IDE][VERSION]/idea.properties
```

## Prerequisites

### Repository Configuration
Ensure your Artifactory repository is properly configured:

1. **Repository Type**: Generic repository
2. **Layout**: Default or custom layout supporting JetBrains plugins
3. **Repository Key**: Unique identifier for the repository
4. **Permissions**: Read access for users/groups that will download plugins

### JetBrains IDE Installation
- At least one JetBrains IDE must be installed on the system
- IDEs should be installed in standard locations for auto-detection
- IDEs should be closed during configuration for best results

## Platform-Specific Requirements

### macOS
- **Permissions**: Generally no special permissions required
- **File Access**: Ensure the user has read/write access to `~/Library/Application Support/JetBrains/`
- **IDE Locations**: IDEs typically installed in `/Applications/`

### Windows
- **Permissions**: Usually no elevated permissions required
- **File Access**: Ensure access to `%APPDATA%\JetBrains\`
- **IDE Locations**: IDEs typically installed in `%LOCALAPPDATA%\JetBrains\Toolbox\apps\` or `%PROGRAMFILES%\JetBrains\`

### Linux
- **Permissions**: Generally no special permissions required
- **File Access**: Ensure access to `~/.config/JetBrains/` or `~/.JetBrains/`
- **IDE Locations**: Various locations including `/opt/`, `~/Applications/`, or snap packages

## Command Output

### Successful Execution
```
$ jf jetbrains config --repo=jetbrains-plugins --artifactory-url=https://mycompany.jfrog.io/

[Info] Configuring JetBrains IDEs plugin repository...
[Info] Validating repository...
[Info] Repository validation successful
[Info] Found 3 JetBrains IDE installation(s):
[Info]   IntelliJ IDEA 2023.3
[Info]   PyCharm 2023.3
[Info]   WebStorm 2023.3
[Info] Configuring IntelliJ IDEA 2023.3...
[Info] 	Backup created at: ~/.jfrog/backup/ide/jetbrains/idea.properties.backup.20240105-143022
[Info] 	Added idea.plugins.host property
[Info] IntelliJ IDEA 2023.3 configured successfully
[Info] Configuring PyCharm 2023.3...
[Info] 	Backup created at: ~/.jfrog/backup/ide/jetbrains/idea.properties.backup.20240105-143023
[Info] 	Added idea.plugins.host property
[Info] PyCharm 2023.3 configured successfully
[Info] Configuring WebStorm 2023.3...
[Info] 	Backup created at: ~/.jfrog/backup/ide/jetbrains/idea.properties.backup.20240105-143024
[Info] 	Updated existing idea.plugins.host property
[Info] WebStorm 2023.3 configured successfully
[Info] Successfully configured 3 out of 3 JetBrains IDE(s)
[Info] Repository URL: https://mycompany.jfrog.io/artifactory/api/jetbrainsplugins/jetbrains-plugins
[Info] Please restart your JetBrains IDEs to apply changes
```

### No IDEs Found
```
$ jf jetbrains config --repo=jetbrains-plugins

[Info] Configuring JetBrains IDEs plugin repository...
[Error] No JetBrains IDEs found

Manual JetBrains IDE Setup Instructions:
=======================================

1. Close all JetBrains IDEs

2. Locate your IDE configuration directory:
   ~/Library/Application Support/JetBrains/[IDE][VERSION]/idea.properties

   Examples:
   • IntelliJ IDEA: IntelliJIdea2023.3/idea.properties
   • PyCharm: PyCharm2023.3/idea.properties
   • WebStorm: WebStorm2023.3/idea.properties

3. Open or create the idea.properties file in a text editor

4. Add or modify the following line:
   idea.plugins.host=https://mycompany.jfrog.io/artifactory/api/jetbrainsplugins/jetbrains-plugins

5. Save the file and restart your IDE

Repository URL: https://mycompany.jfrog.io/artifactory/api/jetbrainsplugins/jetbrains-plugins

Supported IDEs: IntelliJ IDEA, PyCharm, WebStorm, PhpStorm, RubyMine, CLion, DataGrip, GoLand, Rider, Android Studio, AppCode, RustRover, Aqua
```

## Backup and Recovery

### Automatic Backups
The command automatically creates timestamped backups before making changes:

**Backup Location:**
```
~/.jfrog/backup/ide/jetbrains/idea.properties.backup.[TIMESTAMP]
```

**Backup Filename Format:**
```
idea.properties.backup.YYYYMMDD-HHMMSS
```

**Example:**
```
~/.jfrog/backup/ide/jetbrains/idea.properties.backup.20240105-143022
```

### Manual Recovery
To restore original settings:

1. **Locate Backup**: Find the appropriate backup file in `~/.jfrog/backup/ide/jetbrains/`
2. **Identify Target**: Determine which IDE configuration to restore
3. **Copy Backup**: Copy the backup file to the original location
4. **Restart IDE**: Restart the affected IDE to apply changes

**Recovery Example:**
```bash
# Restore IntelliJ IDEA configuration
cp ~/.jfrog/backup/ide/jetbrains/idea.properties.backup.20240105-143022 \
   ~/Library/Application\ Support/JetBrains/IntelliJIdea2023.3/idea.properties
```

### Backup Management
Clean up old backups periodically:

```bash
# Remove backups older than 30 days
find ~/.jfrog/backup/ide/jetbrains/ -name "*.backup.*" -mtime +30 -delete
```

## Troubleshooting

### Common Issues and Solutions

#### No JetBrains IDEs Detected

**Symptoms:**
```
Error: no JetBrains IDEs found
```

**Possible Causes:**
- IDEs not installed in standard locations
- IDEs installed via snap packages (Linux)
- Non-standard installation directories

**Solutions:**
1. **Verify Installation**: Ensure at least one JetBrains IDE is installed
2. **Check Locations**: Verify IDEs are in standard installation directories
3. **Manual Configuration**: Use the manual setup instructions provided in the error output
4. **Custom Installations**: For non-standard installations, manually create configuration directories

#### Configuration Directory Not Accessible

**Symptoms:**
```
Error: failed to create config directory
```

**Solutions:**
1. **Check Permissions**: Ensure user has write access to configuration directories
2. **Create Directories**: Manually create missing configuration directories
3. **File System Issues**: Check for file system problems or disk space

#### Repository Validation Fails

**Symptoms:**
```
Error: repository validation failed
```

**Solutions:**
1. **Check Repository**: Verify repository exists in Artifactory
2. **Verify URL**: Ensure Artifactory URL is correct and accessible
3. **Authentication**: Ensure proper authentication is configured
4. **Network**: Check network connectivity to Artifactory

#### Backup Creation Fails

**Symptoms:**
```
Error: failed to create backup
```

**Solutions:**
1. **Disk Space**: Ensure sufficient disk space for backup files
2. **Permissions**: Check write permissions to backup directory
3. **Directory Structure**: Ensure backup directory structure exists

### Manual Configuration

If automatic configuration fails, you can manually configure each IDE:

#### Step-by-Step Manual Setup

1. **Close IDE**: Ensure the IDE is completely closed
2. **Locate Configuration**: Find the IDE's configuration directory
3. **Create/Edit Properties File**: Open or create `idea.properties`
4. **Add Plugin Host**: Add the repository URL
5. **Save and Restart**: Save the file and restart the IDE

#### Manual Configuration Example

**File Location (macOS):**
```
~/Library/Application Support/JetBrains/IntelliJIdea2023.3/idea.properties
```

**File Content:**
```properties
# JFrog Artifactory plugins repository
idea.plugins.host=https://mycompany.jfrog.io/artifactory/api/jetbrainsplugins/jetbrains-plugins

# Other existing properties...
```

#### Verification
After manual configuration:

1. **Start IDE**: Launch the configured IDE
2. **Check Plugins**: Go to Settings → Plugins
3. **Verify Repository**: Ensure plugins are loaded from Artifactory
4. **Test Installation**: Try installing a plugin from the repository

## Advanced Usage

### Batch Configuration
Configure multiple systems using scripts:

```bash
#!/bin/bash
# Configure JetBrains IDEs across multiple machines

REPO="jetbrains-plugins"
ARTIFACTORY_URL="https://mycompany.jfrog.io/"

echo "Configuring JetBrains IDEs..."
jf jetbrains config --repo="$REPO" --artifactory-url="$ARTIFACTORY_URL"

if [ $? -eq 0 ]; then
    echo "✓ JetBrains IDEs configured successfully"
    echo "Please restart your IDEs to apply changes"
else
    echo "✗ Configuration failed"
    exit 1
fi
```

### CI/CD Integration
Include in deployment or onboarding scripts:

```yaml
# GitHub Actions example
- name: Configure JetBrains IDEs
  run: |
    jf jetbrains config --repo=jetbrains-plugins --artifactory-url=${{ secrets.ARTIFACTORY_URL }}
  
- name: Verify Configuration
  run: |
    # Check if configuration was applied
    if grep -q "idea.plugins.host.*artifactory" ~/.config/JetBrains/*/idea.properties; then
      echo "✓ JetBrains IDEs configured for Artifactory"
    else
      echo "✗ Configuration verification failed"
      exit 1
    fi
```

### Configuration Validation
Verify configuration across multiple IDEs:

```bash
#!/bin/bash
# Validate JetBrains IDE configuration

EXPECTED_HOST="mycompany.jfrog.io"
CONFIG_BASE="$HOME/.config/JetBrains"

echo "Validating JetBrains IDE configurations..."

for ide_dir in "$CONFIG_BASE"/*; do
    if [ -d "$ide_dir" ]; then
        ide_name=$(basename "$ide_dir")
        properties_file="$ide_dir/idea.properties"
        
        if [ -f "$properties_file" ]; then
            if grep -q "idea.plugins.host.*$EXPECTED_HOST" "$properties_file"; then
                echo "✓ $ide_name: Configured for Artifactory"
            else
                echo "✗ $ide_name: Not configured or incorrect URL"
            fi
        else
            echo "? $ide_name: No properties file found"
        fi
    fi
done
```

### Team Onboarding Script
Automate developer environment setup:

```bash
#!/bin/bash
# Developer onboarding script

echo "=== JetBrains IDE Setup ==="
echo "Configuring IDEs to use company plugin repository..."

# Configure JetBrains IDEs
jf jetbrains config --repo=jetbrains-plugins

if [ $? -eq 0 ]; then
    echo ""
    echo "✓ JetBrains IDEs configured successfully!"
    echo ""
    echo "Next steps:"
    echo "1. Restart all JetBrains IDEs"
    echo "2. Go to Settings → Plugins in your IDE"
    echo "3. Install approved company plugins"
    echo "4. Contact IT if you encounter issues"
else
    echo ""
    echo "✗ Configuration failed!"
    echo "Please contact IT support for manual setup"
fi
```

## Best Practices

### Repository Management
1. **Curated Plugins**: Only include approved and tested plugins
2. **Version Control**: Maintain specific plugin versions for stability
3. **Security Scanning**: Scan plugins for vulnerabilities before inclusion
4. **Documentation**: Document approved plugins and their purposes

### Team Workflow
1. **Standardization**: Use consistent repository keys across the organization
2. **Communication**: Inform team when repository configuration changes
3. **Testing**: Test plugin updates in development environment first
4. **Rollback Plan**: Maintain ability to rollback to previous configurations

### Security
1. **Access Control**: Implement proper repository permissions
2. **Authentication**: Use access tokens for CLI authentication
3. **Network Security**: Use HTTPS for all communications
4. **Audit Trail**: Monitor repository access and plugin downloads

### Maintenance
1. **Regular Updates**: Keep plugin repositories updated
2. **Cleanup**: Remove old plugin versions periodically
3. **Monitoring**: Monitor repository usage and performance
4. **Backup Management**: Regularly clean up old configuration backups

### Plugin Installation Workflow
After configuration, developers can install plugins through their IDE:

1. **Open IDE**: Launch any configured JetBrains IDE
2. **Access Plugins**: Go to Settings → Plugins (or Preferences → Plugins on macOS)
3. **Browse Repository**: Browse available plugins from the Artifactory repository
4. **Install Plugins**: Install approved plugins directly from the IDE
5. **Restart if Required**: Restart IDE if prompted after plugin installation 