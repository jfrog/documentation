# Scan Azure Repos Pull Request

After you create a new pull request, Frogbot will automatically scan it.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the pull request. Vulnerabilities that aren't new, and existed in the code before the pull request was created, will not be included in the report. In order to include all the vulnerabilities in the report, including older ones that weren't added by this PR, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot Azure Repos scan workflow is:

1. The developer opens a pull request.
2. Frogbot scans the pull request and adds a comment with the scan results.
3. Frogbot can be triggered again following new commits, by adding a comment with the `rescan` text.
