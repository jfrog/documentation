---
description: The SAST configuration file is in beta.
---

# Customizing Your Scan

## Configuration File

In your project's root directory, create a file named `jfrog-sast-beta.yaml` with the following schema:

```yaml
# Define the scans for the project
# If you need more than one scan, add another scan element
scans:
    # Specify the programming language for the scan
    # Supported languages: python, js, java
  - language: python
    # Relative paths to the project's source code
    source_roots:
      - "src/module1"
    # Exclude directories from the scan
    # Use relative glob-pattern paths of files and folders
    exclude_patterns:
      - "**/*test*/**/*"
    # Exclude a specific scan rule
    exclude_rules:
      - "sast-rule-1"
    # Exclude a specific scan tag
    exclude_tags:
      - "cwe-123"
```

## Ignore Findings

The SAST scanner allows ignoring a vulnerability finding simply by placing a `JFS-ignore` annotation directly in the code. Place the `JFS-ignore` annotation as a comment anywhere in the finding's data-flow to make it ignored.

The following example shows how to ignore a log-injection issue:

{% code lineNumbers="true" %}
```javascript
source = input()
# JFS-ignore
print(source)
```
{% endcode %}

> **Suggestion:** Place the `JFS-ignore` annotation above the finding's execution line
