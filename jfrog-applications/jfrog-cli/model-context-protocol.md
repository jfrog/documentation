# Managing JFrog Model Context Protocol (MCP)

## Overview

The JFrog Model Context Protocol (MCP) server enables integration between JFrog tools and AI models/agents. The `jf mcp` command allows you to manage the MCP server, including starting it with specific configurations and updating it to the latest version.

## Commands

### Start MCP Server

Starts the MCP server with configured options.

```
jf mcp start [options]
```

#### Options

| Flag | Environment Variable | Description |
|------|---------------------|-------------|
| `--toolsets` | `JFROG_MCP_TOOLSETS` | Comma-separated list of toolsets to enable (e.g., "artifactory,xray"). Default: "read" |
| `--tools-access` | `JFROG_MCP_TOOL_ACCESS` | Semicolon-separated list of tool access rights. Default: "all-toolsets" |
| `--mcp-server-version` | - | Specific version of the MCP server to use. Default: "[RELEASE]" (latest) |

#### Examples

```
# Start the MCP server with default settings
jf mcp start

# Start the MCP server with specific toolsets
jf mcp start --toolsets=artifactory,xray

# Start the MCP server with specific tools access
jf mcp start --tools-access="read;write"
```

### Update MCP Server

Updates the MCP server binary to the latest version.

```
jf mcp update
```

#### Example

```
jf mcp update
```

## Implementation Details

### Binary Management

- **Storage Location**: The MCP server binary is stored in `~/.jfrog/cli-mcp/cli-mcp-server` (or platform equivalent).
- **Download**: The command automatically downloads the binary if not present or when updating.
- **Version Checking**: When updating, the command checks if you already have the latest version.

### Access Control

The MCP server can be configured to provide different levels of access to JFrog tools:

- Use `--toolsets` to specify which JFrog tools are accessible via MCP
- Use `--tools-access` to specify access rights (read, write, etc.)

### Common Use Cases

1. **AI Tool Integration**: Start the MCP server to provide JFrog platform access to AI agents.
2. **Regular Updates**: Use the update command periodically to ensure you have the latest MCP server features.
3. **Controlled Access**: Specify which JFrog tools and access levels are available to AI agents.
