# Enrich your SBOM JSONs & XMLs

The sbom enrichment command takes an exported SBOM file in XML/JSON format and enriches your file with package vulnerabilities found by XRAY.

This _**jf sbom enrich \<file\_path>**_ command enriches a file that is found on file\_path.

***

**Note**

> This command requires:

* Version 3.101.3 or above of Xray
* Version 2.60.0 or above of JFrog CLI

***

#### Commands Params

|                       |                                                                                                                                         |
| --------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| **Command name**      | sbom-enrich                                                                                                                             |
| **Abbreviation**      | se                                                                                                                                      |
| **Command options**   |                                                                                                                                         |
| `--server-id`         | <p>[Optional]<br>Server ID configured using the <em>jf c add</em> command. If not specified, the default configured server is used.</p> |
| **Command arguments** |                                                                                                                                         |
| `file_path`           | the sbom file path.                                                                                                                     |

#### Example 1

Enriches an XML file

```
jf se "path/to/file.xml"
```

#### Example 2

Enriches a JSON file

```
jf se "path/to/files/file.json"
```
